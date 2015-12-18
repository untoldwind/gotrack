package store

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/gotrack/server/conntrack"
	"testing"
	"time"
)

func TestProcProvider(t *testing.T) {
	Convey("Given a rrd for 5 min", t, func() {
		rrd := newRRD(time.Unix(0, 0), 300, 1)

		So(rrd.start, ShouldEqual, 0)
		So(rrd.end, ShouldEqual, 300)
		So(rrd.entries, ShouldHaveLength, 300)

		Convey("When rates a queried", func() {
			rates := rrd.getRate(5)

			So(rates.BytesIn, ShouldEqual, 0)
			So(rates.PacketsIn, ShouldEqual, 0)
			So(rates.BytesOut, ShouldEqual, 0)
			So(rates.PacketsOut, ShouldEqual, 0)
		})

		Convey("When rrd is filled", func() {
			for i := uint64(0); i < 300; i++ {
				rrd.addTotals(
					time.Unix(int64(i), 0),
					&conntrack.Transfer{
						Bytes:   100 * i,
						Packets: 10 * i,
					},
					&conntrack.Transfer{
						Bytes:   200 * i,
						Packets: 20 * i,
					},
				)
			}

			So(rrd.start, ShouldEqual, 0)
			So(rrd.end, ShouldEqual, 300)
			So(rrd.last, ShouldEqual, 299)

			for i := 0; i < 300; i++ {
				if rrd.entries[i].BytesIn != uint64(100*i) ||
					rrd.entries[i].PacketsIn != uint64(10*i) ||
					rrd.entries[i].BytesOut != uint64(200*i) ||
					rrd.entries[i].PacketsOut != uint64(20*i) {
					t.Fail()
				}
			}

			Convey("When span is retrieved", func() {
				span := rrd.getSpan()

				So(span.Start, ShouldResemble, time.Unix(0, 0))
				So(span.End, ShouldResemble, time.Unix(300, 0))
				So(span.Deltas, ShouldHaveLength, 299)

				for i := 0; i < 299; i++ {
					if span.Deltas[i].BytesIn != 100 ||
						span.Deltas[i].PacketsIn != 10 ||
						span.Deltas[i].BytesOut != 200 ||
						span.Deltas[i].PacketsOut != 20 {
						t.Fail()
					}
				}
			})

			Convey("When rates a queried", func() {
				rates := rrd.getRate(5)

				So(rates.BytesIn, ShouldEqual, 100)
				So(rates.PacketsIn, ShouldEqual, 10)
				So(rates.BytesOut, ShouldEqual, 200)
				So(rates.PacketsOut, ShouldEqual, 20)
			})

			Convey("When 100 more entries are added", func() {
				for i := uint64(300); i < 400; i++ {
					rrd.addTotals(
						time.Unix(int64(i), 0),
						&conntrack.Transfer{
							Bytes:   1000 * i,
							Packets: 100 * i,
						},
						&conntrack.Transfer{
							Bytes:   2000 * i,
							Packets: 200 * i,
						},
					)
				}

				So(rrd.start, ShouldEqual, 100)
				So(rrd.end, ShouldEqual, 400)

				Convey("When span is retrieved", func() {
					span := rrd.getSpan()

					So(span.Start, ShouldResemble, time.Unix(100, 0))
					So(span.End, ShouldResemble, time.Unix(400, 0))
					So(span.Deltas, ShouldHaveLength, 299)

					for i := 0; i < 199; i++ {
						if span.Deltas[i].BytesIn != 100 ||
							span.Deltas[i].PacketsIn != 10 ||
							span.Deltas[i].BytesOut != 200 ||
							span.Deltas[i].PacketsOut != 20 {
							t.Fail()
						}
					}
					for i := 2000; i < 299; i++ {
						if span.Deltas[i].BytesIn != 1000 ||
							span.Deltas[i].PacketsIn != 100 ||
							span.Deltas[i].BytesOut != 2000 ||
							span.Deltas[i].PacketsOut != 200 {
							t.Fail()
						}
					}
				})

				Convey("When rates a queried", func() {
					rates := rrd.getRate(5)

					So(rates.BytesIn, ShouldEqual, 1000)
					So(rates.PacketsIn, ShouldEqual, 100)
					So(rates.BytesOut, ShouldEqual, 2000)
					So(rates.PacketsOut, ShouldEqual, 200)
				})
			})
		})
	})
}
