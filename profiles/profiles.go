package profiles

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	v1 "github.com/videocoin/cloud-api/profiles/v1"
)

type Profile struct {
	*v1.Profile
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

	built = append(built, "-i "+input)

	for _, c := range p.Spec.Components {
		built = append(built, c.Render())
	}

	built = append(built, output)
	return strings.Join(built, " ")
}
