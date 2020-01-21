package manager

import (
	"context"

	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/profiles/manager/v1"
	tracer "github.com/videocoin/cloud-pkg/tracer"
	ds "github.com/videocoin/cloud-profiles/datastore"
)

func (m *Manager) Create(ctx context.Context, req *v1.ProfileCreateRequest) (*ds.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.Create")
	defer span.Finish()

	deposit := "10000000000000000000"
	if req.Deposit != "" {
		deposit = req.Deposit
	}

	reward := "10000000000000000"
	if req.Reward != "" {
		reward = req.Reward
	}

	profile := &ds.Profile{
		Name:        req.Name,
		Description: req.Description,
		IsEnabled:   false,
		Rel:         req.Rel,
		Deposit:     deposit,
		Reward:      reward,
	}

	if req.Spec != nil {
		profile.Spec = *req.Spec
	}

	profile, err := m.ds.Profile.Create(ctx, profile)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profile, nil
}

func (m *Manager) GetProfileByID(ctx context.Context, id string) (*ds.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.GetProfileByID")
	defer span.Finish()

	profile, err := m.ds.Profile.Get(ctx, id)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profile, nil
}

func (m *Manager) ListEnabledProfiles(ctx context.Context) ([]*ds.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.ListEnabledProfiles")
	defer span.Finish()

	profiles, err := m.ds.Profile.ListEnabled(ctx)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profiles, nil
}

func (m *Manager) ListAllProfiles(ctx context.Context) ([]*ds.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.ListAllProfiles")
	defer span.Finish()

	profiles, err := m.ds.Profile.List(ctx)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profiles, nil
}
