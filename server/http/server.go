package http

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"github.com/untoldwind/routing"
	"net"
	"net/http"
)

type Server struct {
	config   *config.ServerConfig
	listener net.Listener
	logger   logging.Logger
}

func NewServer(config *config.ServerConfig, parent logging.Logger) *Server {
	return &Server{
		config: config,
		logger: parent.WithContext(map[string]interface{}{"package": "http"}),
	}
}

func (s *Server) Start() error {
	ip := net.ParseIP(s.config.BindAddress)

	if ip == nil {
		return errors.Errorf("Failed to parse IP: %v", s.config.BindAddress)
	}
	bindAddr := &net.TCPAddr{IP: ip, Port: s.config.HttpPort}

	var err error
	s.listener, err = net.Listen(bindAddr.Network(), bindAddr.String())
	if err != nil {
		return err
	}

	go http.Serve(s.listener, s.routeHandler())

	s.logger.Infof("Started http server on %s", bindAddr.String())
	return nil
}

func (s *Server) Stop() {
	s.logger.Info("Stopping http server ...")
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) routeHandler() http.Handler {
	return routing.NewRouteHandler(
		routing.PrefixSeq("/v1",
			InternalRoutes(s.logger),
		),
		SendError(s.logger, NotFound()),
	)
}
