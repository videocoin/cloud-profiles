package rpc

import (
	"context"

	protoempty "github.com/gogo/protobuf/types"
	"github.com/jinzhu/copier"
	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	"github.com/videocoin/cloud-api/rpc"
	"github.com/videocoin/cloud-profiles/datastore"
	"github.com/videocoin/cloud-profiles/profiles"
)

func (s *Server) Get(ctx context.Context, req *v1.ProfileRequest) (*v1.GetProfileResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("id", req.ID)
	logger := s.logger.WithField("id", req.ID)

	profile, err := s.manager.GetProfileByID(ctx, req.ID)
	if err != nil {
		if err == datastore.ErrProfileNotFound {
			return nil, rpc.ErrRpcNotFound
		}

		logFailedTo(logger, "get profile", err)
		return nil, rpc.ErrRpcInternal
	}

	response := &v1.GetProfileResponse{}
	if err := copier.Copy(&response, &profile); err != nil {
		return nil, err
	}
	response.MachineType = profile.Spec.MachineType
	response.Cost = profile.Spec.Cost
	response.Components = profile.Spec.Components

	return response, nil
}

func (s *Server) List(ctx context.Context, req *protoempty.Empty) (*v1.ProfileListResponse, error) {
	profiles, err := s.manager.ListEnabledProfiles(ctx)
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

func (s *Server) Render(ctx context.Context, req *v1.RenderRequest) (*v1.RenderResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("id", req.ID)
	logger := s.logger.WithField("id", req.ID)

	if req.Input == "" || req.Output == "" {
		return nil, rpc.ErrRpcBadRequest
	}

	profile, err := s.manager.GetProfileByID(ctx, req.ID)
	if err != nil {
		if err == datastore.ErrProfileNotFound {
			return nil, rpc.ErrRpcNotFound
		}

		logFailedTo(logger, "get profile", err)
		return nil, rpc.ErrRpcInternal
	}

	if len(req.Components) > 0 {
		profile.Spec.Components = req.Components
	}

	p := profiles.Profile{Profile: profile}

	return &v1.RenderResponse{
		Render: p.Render(req.Input, req.Output),
	}, nil
}
