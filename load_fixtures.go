package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

		// profile := new(profile.Profile)
		// var data map[string]interface{}
		// err = json.Unmarshal(content, profile)
		// if err != nil {
		// 	panic(err)
		// }

		// p := &v1.Profile{
		// 	Name:        profile.Name,
		// 	Description: profile.Description,
		// 	IsEnabled:   profile.IsEnabled,
		// }

		// cc, _ := profile.SerializeComponents()
		// p.Components = cc

		// _, err = ds.Profile.Create(context.Background(), p)
		// if err != nil {
		// 	panic(err)
		// }

		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, data); err != nil {
			panic(err)
		}

		var objmap map[string]*json.RawMessage
		if err := json.Unmarshal(buffer.Bytes(), &objmap); err != nil {
		}

		profile := new(v1.Profile)
		if err := json.Unmarshal(*objmap["id"], &profile.Id); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(*objmap["name"], &profile.Name); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(*objmap["description"], &profile.Description); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(*objmap["is_enabled"], &profile.IsEnabled); err != nil {
			panic(err)
		}

		profile.Components = *objmap["components"]

		_, err = ds.Profile.Create(context.Background(), profile)
		if err != nil {
			panic(err)
		}
	}
}
