package server

import (
	"artemzinchenko.com/auth/internal/api/task"
	"artemzinchenko.com/auth/internal/service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"log"
)

const (
	jwks = "https://dev-artem-zinchenko.eu.auth0.com/.well-known/jwks.json"
	aud  = "https://task-api.local"
	iss  = "https://dev-artem-zinchenko.eu.auth0.com/"
)

func NewGRPCServer() *grpc.Server {
	//create new JWT manager
	jwtManager := service.NewAuth0JWTManager(jwks, aud, iss)

	//sets authorisation interceptor
	so := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(jwtManager.Validate)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(jwtManager.Validate)),
	}
	//create new gGRP
	gsrv := grpc.NewServer(so...)

	//register Task API server
	if t, err := task.NewApiServer(); err == nil {
		t.RegisterService(gsrv)
	} else {
		log.Fatalf("while registering TaskApiServer %v", err)
	}
	return gsrv
}

const (
	jwks = "https://dev-artem-zinchenko.eu.auth0.com/.well-known/jwks.json"
	aud  = "https://plantcatalog.api"
	iss  = "https://dev-artem-zinchenko.eu.auth0.com/"
)

func NewGRPCServer() *grpc.Server {
	//create new JWT manager
	jwtManager := service.NewAuth0JWTManager(jwks, aud, iss)

	//sets authorisation interceptor
	so := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(jwtManager.Validate)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(jwtManager.Validate)),
	}
	//create new gGRP
	gsrv := grpc.NewServer(so...)

	//register Task API server
	if t, err := task.NewApiServer(); err == nil {
		t.RegisterService(gsrv)
	} else {
		log.Fatalf("while registering TaskApiServer %v", err)
	}
	return gsrv
}
