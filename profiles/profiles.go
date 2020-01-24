package profiles

import (
	"strings"

	ds "github.com/videocoin/cloud-profiles/datastore"
)

type Profile struct {
	*ds.Profile
}

func (p *Profile) Pipelines() []*v1.Pipeline {
	return p.Spec.Pipelines
}

func (p *Profile) Render(input, output string) string {
	built := []string{"ffmpeg"}

	if p.Name == "test" {
		input = "/tmp/in.mp4"
	}

	for _, p := range p.Spec.Pipelines {
		for _, c := range p.Components {
			if c.Type == v1.ComponentTypeDemuxer {
				built = append(built, c.Render())
				break
			}
		}
	}

	built = append(built, "-i "+input)

	for _, p := range p.Spec.Pipelines {
		built = append(built, p.Render("", output))
	}

	return strings.Join(built, " ")
}
