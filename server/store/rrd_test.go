package store

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/gotrack/server/conntrack"
	"testing"
	"time"
)

func TestProcProvider(t *testing.T) {
	Convey("Given a rrd for 5 min", t, func() {
		var now int64

		rrd_now = func() int64 {
			return now
		}
		rrd := newRRD(300, 1)

		So(rrd.start, ShouldEqual, 0)
		So(rrd.end, ShouldEqual, 300)
		So(rrd.entries, ShouldHaveLength, 300)

		Convey("When rrd is filled", func() {
			for i := uint64(0); i < 300; i++ {
				now = int64(i)
				rrd.addTotals(
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
				So(span.Entries, ShouldHaveLength, 300)

				for i := 0; i < 300; i++ {
					if span.Entries[i].BytesIn != uint64(100*i) ||
						span.Entries[i].PacketsIn != uint64(10*i) ||
						span.Entries[i].BytesOut != uint64(200*i) ||
						span.Entries[i].PacketsOut != uint64(20*i) {
						t.Fail()
					}
				}
			})

			Convey("When 100 more entries are added", func() {
				for i := uint64(300); i < 400; i++ {
					now = int64(i)
					rrd.addTotals(
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

				So(rrd.start, ShouldEqual, 100)
				So(rrd.end, ShouldEqual, 400)

				Convey("When span is retrieved", func() {
					span := rrd.getSpan()

					So(span.Start, ShouldResemble, time.Unix(100, 0))
					So(span.End, ShouldResemble, time.Unix(400, 0))
					So(span.Entries, ShouldHaveLength, 300)

					for i := 0; i < 300; i++ {
						if span.Entries[i].BytesIn != uint64(100*(i+100)) ||
							span.Entries[i].PacketsIn != uint64(10*(i+100)) ||
							span.Entries[i].BytesOut != uint64(200*(i+100)) ||
							span.Entries[i].PacketsOut != uint64(20*(i+100)) {
							t.Fail()
						}
					}
				})
			})
		})
	})
}
