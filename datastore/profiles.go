package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/opentracing/opentracing-go"
	v1 "github.com/videocoin/cloud-api/profiles/v1"

	"github.com/jinzhu/gorm"
)

var (
	ErrProfileNotFound = errors.New("profile is not found")
)

type ProfileDatastore struct {
	db *gorm.DB
}

func NewProfileDatastore(db *gorm.DB) (*ProfileDatastore, error) {
	db.AutoMigrate(&v1.Profile{})
	return &ProfileDatastore{db: db}, nil
}

func (ds *ProfileDatastore) Create(ctx context.Context, profile *v1.Profile) (*v1.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Create")
	defer span.Finish()

	tx := ds.db.Begin()

	if err := tx.Create(profile).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return profile, nil
}

func (ds *ProfileDatastore) Delete(ctx context.Context, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()

	span.SetTag("id", id)

	profile := &v1.Profile{
		Id: id,
	}

	if err := ds.db.Delete(profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrProfileNotFound
		}

		return fmt.Errorf("failed to get profile by id %s: %s", id, err.Error())
	}

	return nil
}

func (ds *ProfileDatastore) Get(ctx context.Context, id string) (*v1.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Get")
	defer span.Finish()

	span.SetTag("id", id)

	profile := &v1.Profile{}

	if err := ds.db.Where("id = ?", id).First(profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrProfileNotFound
		}

		return nil, fmt.Errorf("failed to get profile by id %s: %s", id, err.Error())
	}

	return profile, nil
}

func (ds *ProfileDatastore) List(ctx context.Context) ([]*v1.Profile, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "List")
	defer span.Finish()

	profiles := []*v1.Profile{}

	if err := ds.db.Where("is_enabled = ?", true).Find(&profiles).Error; err != nil {
		return nil, fmt.Errorf("failed to list profiles: %s", err)
	}

	return profiles, nil
}
