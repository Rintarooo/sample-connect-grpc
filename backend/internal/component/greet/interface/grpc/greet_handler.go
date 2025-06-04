package grpc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	greetv1 "github.com/Rintarooo/sample-connect-grpc/internal/generated/grpc/greet/v1"
	"github.com/Rintarooo/sample-connect-grpc/internal/generated/grpc/greet/v1/greetv1connect"
)

type greetingServiceHandler struct{}

func NewGreetingServiceHandler() greetv1connect.GreetingServiceHandler {
	return &greetingServiceHandler{}
}

func (h *greetingServiceHandler) GetGreeting(
	ctx context.Context,
	req *connect.Request[greetv1.GetGreetingRequest],
) (*connect.Response[greetv1.GetGreetingResponse], error) {
	name := req.Msg.Name
	if name == "" {
		name = "World"
	}

	greeting := &greetv1.Greeting{
		Id:      fmt.Sprintf("greeting-%s", name),
		Message: fmt.Sprintf("Hello, %s!", name),
	}

	response := &greetv1.GetGreetingResponse{
		Greeting: greeting,
	}

	return connect.NewResponse(response), nil
}
