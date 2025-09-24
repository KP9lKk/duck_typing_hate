package grpcserver

import (
	"context"
	"fmt"
	"net"

	pbgrpc "google.golang.org/grpc"
)

type Server struct {
	App     *pbgrpc.Server
	notify  chan error
	address string
}

func New(address string) *Server {
	s := &Server{
		App:     pbgrpc.NewServer(),
		notify:  make(chan error, 1),
		address: address,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		var lc net.ListenConfig

		ln, err := lc.Listen(context.Background(), "tcp", net.JoinHostPort("", s.address))
		if err != nil {
			s.notify <- fmt.Errorf("failed to listen: %w", err)
			close(s.notify)
			return
		}
		s.notify <- s.App.Serve(ln)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) ShutDown() error {
	s.App.GracefulStop()
	return nil
}
