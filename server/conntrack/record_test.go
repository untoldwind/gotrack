package conntrack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRecord(t *testing.T) {
	Convey("Given a tcp line with accounting to parse", t, func() {
		line := `tcp      6 12 TIME_WAIT src=192.168.3.133 dst=84.53.142.43 sport=60968 dport=80 packets=15 bytes=1172 src=84.53.142.43 dst=192.168.2.106 sport=80 dport=60968 packets=21 bytes=23026 [ASSURED] mark=0 use=2`

		Convey("When the line is parsed", func() {
			record, err := parseRecord(line)

			So(err, ShouldBeNil)
			So(record, ShouldNotBeNil)
			So(record.Protocol, ShouldEqual, "tcp")
			So(record.Ttl, ShouldEqual, 12)
			So(record.State, ShouldEqual, "TIME_WAIT")
			So(record.Send.Src, ShouldEqual, "192.168.3.133")
			So(record.Send.Dst, ShouldEqual, "84.53.142.43")
			So(record.Send.SrcPort, ShouldEqual, 60968)
			So(record.Send.DstPort, ShouldEqual, 80)
			So(record.Send.Packets, ShouldEqual, 15)
			So(record.Send.Bytes, ShouldEqual, 1172)
			So(record.Receive.Src, ShouldEqual, "84.53.142.43")
			So(record.Receive.Dst, ShouldEqual, "192.168.2.106")
			So(record.Receive.SrcPort, ShouldEqual, 80)
			So(record.Receive.DstPort, ShouldEqual, 60968)
			So(record.Receive.Packets, ShouldEqual, 21)
			So(record.Receive.Bytes, ShouldEqual, 23026)
		})
	})

	Convey("Given a tcp line without accounting to parse", t, func() {
		line := `tcp      6 12 TIME_WAIT src=192.168.3.133 dst=84.53.142.43 sport=60968 dport=80 src=84.53.142.43 dst=192.168.2.106 sport=80 dport=60968 [ASSURED] mark=0 use=2`

		Convey("When the line is parsed", func() {
			record, err := parseRecord(line)

			So(err, ShouldBeNil)
			So(record, ShouldNotBeNil)
			So(record.Protocol, ShouldEqual, "tcp")
			So(record.Ttl, ShouldEqual, 12)
			So(record.State, ShouldEqual, "TIME_WAIT")
			So(record.Send.Src, ShouldEqual, "192.168.3.133")
			So(record.Send.Dst, ShouldEqual, "84.53.142.43")
			So(record.Send.SrcPort, ShouldEqual, 60968)
			So(record.Send.DstPort, ShouldEqual, 80)
			So(record.Send.Packets, ShouldEqual, 0)
			So(record.Send.Bytes, ShouldEqual, 0)
			So(record.Receive.Src, ShouldEqual, "84.53.142.43")
			So(record.Receive.Dst, ShouldEqual, "192.168.2.106")
			So(record.Receive.SrcPort, ShouldEqual, 80)
			So(record.Receive.DstPort, ShouldEqual, 60968)
			So(record.Receive.Packets, ShouldEqual, 0)
			So(record.Receive.Bytes, ShouldEqual, 0)
		})
	})

	Convey("Given an udp line with accounting to parse", t, func() {
		line := `udp      17 15 src=192.168.3.133 dst=192.168.3.1 sport=64761 dport=53 packets=1 bytes=66 src=192.168.3.1 dst=192.168.3.133 sport=53 dport=64761 packets=1 bytes=157 mark=0 use=2`

		Convey("When the line is parsed", func() {
			record, err := parseRecord(line)

			So(err, ShouldBeNil)
			So(record, ShouldNotBeNil)
			So(record.Protocol, ShouldEqual, "udp")
			So(record.Ttl, ShouldEqual, 15)
			So(record.State, ShouldEqual, "")
			So(record.Send.Src, ShouldEqual, "192.168.3.133")
			So(record.Send.Dst, ShouldEqual, "192.168.3.1")
			So(record.Send.SrcPort, ShouldEqual, 64761)
			So(record.Send.DstPort, ShouldEqual, 53)
			So(record.Send.Packets, ShouldEqual, 1)
			So(record.Send.Bytes, ShouldEqual, 66)
			So(record.Receive.Src, ShouldEqual, "192.168.3.1")
			So(record.Receive.Dst, ShouldEqual, "192.168.3.133")
			So(record.Receive.SrcPort, ShouldEqual, 53)
			So(record.Receive.DstPort, ShouldEqual, 64761)
			So(record.Receive.Packets, ShouldEqual, 1)
			So(record.Receive.Bytes, ShouldEqual, 157)
		})
	})
}
