package svc

import (
	"imservice/sdk/config"
	"imservice/sdk/conn"
	"imservice/sdk/types"
)

type ServiceContext struct {
	Config       config.Config
	client       *conn.Client
	EventHandler types.EventHandler
}

func NewServiceContext(
	config config.Config,
) *ServiceContext {
	return &ServiceContext{
		Config: config,
	}
}

func (s *ServiceContext) SetEventHandler(eventHandler types.EventHandler) {
	s.EventHandler = eventHandler
}

func (s *ServiceContext) Client() *conn.Client {
	if s.client == nil {
		s.client = conn.NewClient(s.Config.Client, s.EventHandler)
	}
	return s.client
}
