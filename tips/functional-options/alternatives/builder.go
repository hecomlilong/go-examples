package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	serverbuilder := NewServerBuilder("127.0.0.1", 5432).Protocol("tcp").MaxConns(100)
	fmt.Printf("server:%+v\n", serverbuilder.server)
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type ServerBuilder struct {
	server *Server
}

func NewServerBuilder(addr string, port int) *ServerBuilder {
	var res = new(ServerBuilder)
	res.server = &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	return res
}

func (p *ServerBuilder) Protocol(protocol string) *ServerBuilder {
	p.server.Protocol = protocol
	return p
}

func (p *ServerBuilder) Timeout(timeout time.Duration) *ServerBuilder {
	p.server.Timeout = timeout
	return p
}
func (p *ServerBuilder) MaxConns(conns int) *ServerBuilder {
	p.server.MaxConns = conns
	return p
}
func (p *ServerBuilder) Tls(t *tls.Config) *ServerBuilder {
	p.server.TLS = t
	return p
}
