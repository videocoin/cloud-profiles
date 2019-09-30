package rpc

import (
	"context"
	"net"

	protoempty "github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	"github.com/videocoin/cloud-api/rpc"
	"github.com/videocoin/cloud-pkg/grpcutil"
	ds "github.com/videocoin/cloud-profiles/datastore"
	"github.com/videocoin/cloud-profiles/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcServerOpts struct {
	Addr    string
	Ds      *ds.Datastore
	Manager *manager.Manager
	Logger  *logrus.Entry
}

type RpcServer struct {
	addr    string
	grpc    *grpc.Server
	listen  net.Listener
	ds      *ds.Datastore
	manager *manager.Manager
	logger  *logrus.Entry
}

func NewRpcServer(opts *RpcServerOpts) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	rpcServer := &RpcServer{
		addr:    opts.Addr,
		grpc:    grpcServer,
		listen:  listen,
		ds:      opts.Ds,
		manager: opts.Manager,
		logger:  opts.Logger.WithField("system", "rpc"),
	}

	v1.RegisterProfilesServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *RpcServer) Start() error {
	s.logger.Infof("starting rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}

func (s *RpcServer) Health(ctx context.Context, req *protoempty.Empty) (*rpc.HealthStatus, error) {
	return &rpc.HealthStatus{Status: "OK"}, nil
}
