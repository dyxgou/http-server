package main

import (
	"context"
	"log/slog"
	"net"
)

const defaultListenAdrr = ":3000"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	Ln  net.Listener
	Ctx context.Context
}

func CreateServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAdrr
	}
	return &Server{
		Config: cfg,
		Ctx:    context.Background(),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	defer ln.Close()

	if err != nil {
		return err
	}
	s.Ln = ln

	s.acceptConns()
	return nil
}

func (s *Server) acceptConns() {
	ctx, cancel := context.WithCancel(s.Ctx)
	defer cancel()

	for {
		conn, err := s.Ln.Accept()

		if err != nil {
			slog.Error("accepting conn err :", err)
			continue
		}
	}
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn) {
	select {
	case <-ctx.Done():
		return
	default:
	}
}
