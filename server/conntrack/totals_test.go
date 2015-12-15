package conntrack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTotals(t *testing.T) {
	Convey("Given a line to parse", t, func() {
		line := `  eth0: 4031605052 2955324    0    0    0     0          0         0 42854733  469771    0    0    0     0       0          0`

		Convey("When totals are parsed", func() {
			totals, err := parseTotals(line)

			So(err, ShouldBeNil)
			So(totals, ShouldNotBeNil)
			So(totals.Receive.Bytes, ShouldEqual, 4031605052)
			So(totals.Receive.Packets, ShouldEqual, 2955324)
			So(totals.Send.Bytes, ShouldEqual, 42854733)
			So(totals.Send.Packets, ShouldEqual, 469771)
		})
	})
}
