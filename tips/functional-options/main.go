package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	s, err := NewServer("127.0.0.1", 5432, Protocol("http"), Timeout(10*time.Second))
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	fmt.Printf("server:%+v\n", s)
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type option func(*Server)

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func Protocol(protocol string) option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func Timeout(duration time.Duration) option {
	return func(s *Server) {
		s.Timeout = duration
	}
}

func MaxConns(maxConns int) option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func Tls(t *tls.Config) option {
	return func(s *Server) {
		s.TLS = t
	}
}
