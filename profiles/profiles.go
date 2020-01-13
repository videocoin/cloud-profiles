package profiles

import (
	"strings"

	v1 "github.com/videocoin/cloud-api/profiles/v1"
	ds "github.com/videocoin/cloud-profiles/datastore"
)

type Profile struct {
	*ds.Profile
}

func (p *Profile) Render(input, output string) string {
	built := []string{"ffmpeg"}

	if p.Name == "test" {
		input = "/tmp/in.mp4"
	}

	for _, c := range p.Spec.Components {
		if c.Type == v1.ComponentTypeDemuxer {
			built = append(built, c.Render())
		}
	}

	built = append(built, "-i "+input)

	for _, c := range p.Spec.Components {
		if c.Type != v1.ComponentTypeDemuxer {
			built = append(built, c.Render())
		}
	}

	output += "/index.m3u8"

	built = append(built, output)
	return strings.Join(built, " ")
}
