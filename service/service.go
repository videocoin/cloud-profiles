package service

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/videocoin/cloud-profiles/datastore"
	"github.com/videocoin/cloud-profiles/manager"
	"github.com/videocoin/cloud-profiles/rpc"
)

type Service struct {
	cfg        *Config
	rpc        *rpc.Server
	managerRPC *rpc.ManagerServer
}

func NewService(cfg *Config) (*Service, error) {
	ds, err := datastore.NewDatastore(cfg.DBURI)
	if err != nil {
		return nil, err
	}

	manager := manager.NewManager(
		&manager.Opts{
			Ds:     ds,
			Logger: cfg.Logger.WithField("system", "manager"),
		})

	rpcConfig := &rpc.ServerOpts{
		Logger:  cfg.Logger,
		Addr:    cfg.RPCAddr,
		Ds:      ds,
		Manager: manager,
	}

	rpcServer, err := rpc.NewServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	managerRPCConfig := &rpc.ManagerServerOpts{
		Logger:  cfg.Logger,
		Addr:    cfg.ManagerRPCAddr,
		Ds:      ds,
		Manager: manager,
	}
	managerRPC, err := rpc.NewManagerServer(managerRPCConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg:        cfg,
		rpc:        rpcServer,
		managerRPC: managerRPC,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.rpc.Start()  //nolint
	go s.managerRPC.Start()  //nolint
	return nil
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) LoadFixtures(presetsRoot string) error {
	var presetsFiles []string

	err := filepath.Walk(presetsRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			presetsFiles = append(presetsFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	ds, err := datastore.NewDatastore(s.cfg.DBURI)
	if err != nil {
		return err
	}

	m := &runtime.JSONPb{OrigName: true, EmitDefaults: true, EnumsAsInts: false}
	ctx := context.Background()
	profileIds := []string{}

	for _, file := range presetsFiles {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		profile := new(datastore.Profile)
		err = m.Unmarshal(data, &profile)
		if err != nil {
			return err
		}

		_, err = ds.Profile.Get(ctx, profile.ID)
		if err != nil {
			if err == datastore.ErrProfileNotFound {
				_, createErr := ds.Profile.Create(ctx, profile)
				if createErr != nil {
					return createErr
				}

				profileIds = append(profileIds, profile.ID)
				continue
			}

			return err
		}

		err = ds.Profile.Update(ctx, profile)
		if err != nil {
			return err
		}

		profileIds = append(profileIds, profile.ID)
	}

	if len(profileIds) > 0 {
		err = ds.Profile.DeleteAllExceptIds(ctx, profileIds)
		if err != nil {
			return err
		}
	}

	return nil
}
