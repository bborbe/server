package server

import (
	"errors"
	"net"
	"net/http"
	"strconv"
	"github.com/bborbe/log"
	"os"
	"os/signal"
	"syscall"
)
var logger = log.DefaultLogger

type Server interface {
	Start() error
	Stop() error
	Wait()
	Run()
}

type server struct {
	addr     string
	listener net.Listener
	handler  http.Handler
	done     chan bool
}

func NewServerPort(port int, handler http.Handler) *server {
	addr := "0.0.0.0:" + strconv.Itoa(port)
	return NewServer(addr, handler)
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


func  (s *server) Run() {
	defer logger.Close()
	{
		err := s.Start()
		if err != nil {
			logger.Errorf("start server failed, %v", err)
			return
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		err := s.Stop()
		if err != nil {
			logger.Errorf("stop server failed, %v", err)
			return
		}
		logger.Debug("server finished")
		logger.Close()
		os.Exit(0)
	}()

	s.Wait()
}
