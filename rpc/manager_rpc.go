package rpc

import (
	"context"

	protoempty "github.com/gogo/protobuf/types"
	"github.com/jinzhu/copier"
	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/profiles/manager/v1"
	"github.com/videocoin/cloud-api/rpc"
	"github.com/videocoin/cloud-profiles/datastore"
)

func (s *ManagerServer) Create(ctx context.Context, req *v1.ProfileCreateRequest) (*v1.ProfileResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("name", req.Name)
	logger := s.logger

	resp := &v1.ProfileResponse{}

	profile, err := s.manager.Create(ctx, req)
	if err != nil {
		logFailedTo(logger, "create profile", err)
		return nil, rpc.ErrRpcInternal
	}

	if err := copier.Copy(&resp, &profile); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ManagerServer) Get(ctx context.Context, req *v1.ProfileRequest) (*v1.ProfileResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("id", req.Id)
	logger := s.logger.WithField("id", req.Id)

	profile, err := s.manager.GetProfileByID(ctx, req.Id)
	if err != nil {
		if err == datastore.ErrProfileNotFound {
			return nil, rpc.ErrRpcNotFound
		}

		logFailedTo(logger, "get profile", err)
		return nil, rpc.ErrRpcInternal
	}

	response := &v1.ProfileResponse{}
	if err := copier.Copy(&response, &profile); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ManagerServer) List(ctx context.Context, req *protoempty.Empty) (*v1.ProfileListResponse, error) {
	profiles, err := s.manager.ListAllProfiles(ctx)
	if err != nil {
		logFailedTo(s.logger, "profiles list", err)
		return nil, rpc.ErrRpcInternal
	}

	response := &v1.ProfileListResponse{}
	if err := copier.Copy(&response.Items, &profiles); err != nil {
		return nil, err
	}

	return response, nil
}
