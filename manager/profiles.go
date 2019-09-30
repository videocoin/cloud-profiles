package manager

import (
	"context"

	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	tracer "github.com/videocoin/cloud-pkg/tracer"
)

func (m *Manager) GetProfileByID(ctx context.Context, id string) (*v1.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.GetProfileByID")
	defer span.Finish()

	profile, err := m.ds.Profile.Get(ctx, id)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profile, nil
}

func (m *Manager) ListProfiles(ctx context.Context) ([]*v1.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "manager.List")
	defer span.Finish()

	profiles, err := m.ds.Profile.List(ctx)
	if err != nil {
		tracer.SpanLogError(span, err)
		return nil, err
	}

	return profiles, nil
}
