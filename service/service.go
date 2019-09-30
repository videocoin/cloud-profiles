package service

import (
	ds "github.com/videocoin/cloud-profiles/datastore"
	"github.com/videocoin/cloud-profiles/manager"
	"github.com/videocoin/cloud-profiles/rpc"
)

type Service struct {
	cfg *Config
	rpc *rpc.RpcServer
}

func NewService(cfg *Config) (*Service, error) {
	ds, err := ds.NewDatastore(cfg.DBURI)
	if err != nil {
		return nil, err
	}

	manager := manager.NewManager(
		&manager.ManagerOpts{
			Ds:     ds,
			Logger: cfg.Logger.WithField("system", "manager"),
		})

	rpcConfig := &rpc.RpcServerOpts{
		Logger:  cfg.Logger,
		Addr:    cfg.RPCAddr,
		Ds:      ds,
		Manager: manager,
	}

	rpc, err := rpc.NewRpcServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg: cfg,
		rpc: rpc,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.rpc.Start()
	return nil
}

func (s *Service) Stop() error {
	return nil
}
