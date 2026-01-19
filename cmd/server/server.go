package main

import (
	"context"
	"fmt"

	"github.com/youssefhmidi/gRPC-example/services"
)

type Server struct {
	services.UnimplementedMsgServerServer
}

var _ services.MsgServerServer = &Server{}

func (s *Server) Send(ctx context.Context, msg *services.ObjMessage) (*services.Response, error) {
	resString := msg.GetMsg()
	msgString := msg.String()

	resp := fmt.Sprintf("msg : %s,\n all: {%s}\n", resString, msgString)

	return &services.Response{Res: resp}, nil
}
