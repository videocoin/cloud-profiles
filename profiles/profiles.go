package profiles

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	ds "github.com/videocoin/cloud-profiles/datastore"
)

type Profile struct {
	*ds.Profile
}

func ProfileFromContent(content []byte) (*Profile, error) {
	profile := new(Profile)
	err := json.Unmarshal(content, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func ProfileFromFile(filepath string) (*Profile, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return ProfileFromContent(content)
}

func (p *Profile) Render(input, output string) string {
	built := []string{"ffmpeg"}

	if p.Name == "test" {
		input = "/tmp/in.mp4"
	}

	built = append(built, "-i "+input)

	for _, c := range p.Spec.Components {
		built = append(built, c.Render())
	}

	output += "/index.m3u8"

	built = append(built, output)
	return strings.Join(built, " ")
}
