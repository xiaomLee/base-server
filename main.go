package main

import (
	"gate-ws/router"
	"github.com/judwhite/go-svc/svc"
	"net/http"
)

type GatewayWs struct {
	server *http.Server
}

func (s *GatewayWs) Init(env svc.Environment) error {

	return nil
}

func (s *GatewayWs) Start() error {
	s.server = &http.Server{
		Addr: ":8087",
		Handler: router.NewEngine(),
	}
	return nil
}

func (s *GatewayWs) Stop() error {

	return nil
}

func main() {
	if err := svc.Run(&GatewayWs{}); err != nil {
		println(err.Error())
	}
}