package server

import (
	"errors"
	"net"
	"net/http"
)

type Server interface {
	Start() error
	Stop() error
	Wait()
}

type server struct {
	addr     string
	listener net.Listener
	handler  http.Handler
	done     chan bool
}

func NewServer(addr string, handler http.Handler) *server {
	srv := new(server)
	srv.addr = addr
	srv.handler = handler
	return srv
}

func (s *server) Start() error {
	if s.listener != nil {
		return errors.New("already started")
	}
	l, e := net.Listen("tcp", s.addr)
	if e != nil {
		return e
	}
	s.listener = l
	srv := &http.Server{Addr: s.addr, Handler: s.handler}
	go func() {
		srv.Serve(l)
		s.done <- true
	}()
	return nil
}

func (s *server) Wait() {
	<-s.done
}

func (s *server) Stop() error {
	if s.listener == nil {
		return errors.New("already stopped")
	}
	l := s.listener
	s.listener = nil
	return l.Close()
}
