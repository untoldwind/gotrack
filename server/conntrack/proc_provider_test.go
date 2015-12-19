package conntrack

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"testing"
)

func TestProcProvider(t *testing.T) {
	Convey("Given a provider with test files", t, func() {
		config := &config.ContrackConfig{
			Type:          "proc",
			ConntrackFile: "_procfile_example",
			DevFile:       "_dev_example",
			WanInterface:  "eth1",
		}
		logger := logging.NewSimpleLoggerNull()
		provider, err := NewProvider(config, logger)

		So(err, ShouldBeNil)
		So(provider, ShouldNotBeNil)

		Convey("When records are read", func() {
			records, err := provider.Records()

			So(err, ShouldBeNil)
			So(records, ShouldHaveLength, 4)
			So(records[0].Protocol, ShouldEqual, "tcp")
			So(records[1].Protocol, ShouldEqual, "tcp")
			So(records[2].Protocol, ShouldEqual, "udp")
			So(records[3].Protocol, ShouldEqual, "udp")
		})

		Convey("When totals are read", func() {
			totals, err := provider.Totals()

			So(err, ShouldBeNil)
			So(totals, ShouldNotBeNil)
			So(totals.Receive.Bytes, ShouldEqual, 44153805)
			So(totals.Receive.Packets, ShouldEqual, 483380)
			So(totals.Send.Bytes, ShouldEqual, 4096298788)
			So(totals.Send.Packets, ShouldEqual, 2956550)
		})
	})
}
