package profiles

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/videocoin/cloud-profiles/components"
)

type Profile struct {
	Components components.Components `json:"components"`
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
	var built []string

	built = append(built, "-i "+input)

	for _, c := range p.Components {
		built = append(built, c.Render())
	}

	built = append(built, output)
	return strings.Join(built, " ")
}
