package dhcp

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"testing"
)

func TestDnsmasqProvider(t *testing.T) {
	Convey("Given a provider with test files", t, func() {
		config := &config.DhcpConfig{
			Type:        "dnsmasq",
			DnsmasqFile: "_dnsmasq_lease_example",
		}
		logger := logging.NewSimpleLoggerNull()
		provider, err := NewProvider(config, logger)

		So(err, ShouldBeNil)

		Convey("When leases are queried", func() {
			leases, err := provider.Leases()

			So(err, ShouldBeNil)
			So(leases, ShouldHaveLength, 2)

			So(leases[0].ExpiresAt, ShouldEqual, 1451742692)
			So(leases[0].Name, ShouldEqual, "Prospero")
			So(leases[0].MacAddress, ShouldEqual, "0c:2d:e9:a6:6a:7f")
			So(leases[0].IpAddress, ShouldEqual, "192.168.3.151")

			So(leases[1].ExpiresAt, ShouldEqual, 1451764075)
			So(leases[1].Name, ShouldEqual, "Clavain")
			So(leases[1].MacAddress, ShouldEqual, "6c:20:08:94:22:a4")
			So(leases[1].IpAddress, ShouldEqual, "192.168.3.157")
		})
	})
}
