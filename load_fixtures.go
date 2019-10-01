package main

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	ds "github.com/videocoin/cloud-profiles/datastore"
)

const presetsRoot = "presets/"

func main() {
	dbURI := os.Getenv("DBM_MSQLURI")
	ds, err := ds.NewDatastore(dbURI)
	if err != nil {
		panic(err)
	}

	var presetsFiles []string
	err = filepath.Walk(presetsRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			presetsFiles = append(presetsFiles, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range presetsFiles {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		profile := new(v1.Profile)

		m := &runtime.JSONPb{OrigName: true, EmitDefaults: true, EnumsAsInts: false}
		if err := m.Unmarshal(data, &profile); err != nil {
			panic(err)
		}

		_, err = ds.Profile.Create(context.Background(), profile)
		if err != nil {
			panic(err)
		}
	}
}
