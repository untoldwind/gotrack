package conntrack

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"testing"
)

func TestProcProvider(t *testing.T) {
	Convey("Given a file to parse", t, func() {
		config := &config.ProviderConfig{
			Type:     "proc",
			ProcFile: "_procfile_example",
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
	})
}
