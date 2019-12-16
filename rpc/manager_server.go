package rpc

import (
	"net"

	"github.com/sirupsen/logrus"
	v1 "github.com/videocoin/cloud-api/profiles/manager/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	ds "github.com/videocoin/cloud-profiles/datastore"
	"github.com/videocoin/cloud-profiles/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type ManagerRpcServerOpts struct {
	Addr    string
	Ds      *ds.Datastore
	Manager *manager.Manager
	Logger  *logrus.Entry
}

type ManagerRpcServer struct {
	addr    string
	grpc    *grpc.Server
	listen  net.Listener
	ds      *ds.Datastore
	manager *manager.Manager
	logger  *logrus.Entry
}

func NewManagerRpcServer(opts *ManagerRpcServerOpts) (*ManagerRpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)
	healthService := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthService)
	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	rpcServer := &ManagerRpcServer{
		addr:    opts.Addr,
		grpc:    grpcServer,
		listen:  listen,
		ds:      opts.Ds,
		manager: opts.Manager,
		logger:  opts.Logger.WithField("system", "managerrpc"),
	}

	v1.RegisterProfileManagerServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *ManagerRpcServer) Start() error {
	s.logger.Infof("starting manager rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}
