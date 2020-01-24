package v1

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (p *Pipeline) Render(input, output string) string {
	built := make([]string, 0)

	if len(input) > 0 {
		built = append(built, "ffmpeg")
		for _, c := range p.Components {
			if c.Type == ComponentTypeDemuxer {
				built = append(built, c.Render())
			}
		}

		built = append(built, "-i "+input)
	}

	for _, c := range p.Components {
		if c.Type != ComponentTypeDemuxer {
			built = append(built, c.Render())
		}
	}

	built = append(built, output)
	result := strings.Join(built, " ")
	return strings.Join(strings.Fields(result), " ")
}

func (c *Component) Render() string {
	var built []string
	for _, p := range c.Params {
		built = append(built, p.Render())
	}

	return strings.Join(built, " ")
}

func (p *Param) Render() string {
	return p.Key + " " + p.Value
}

func (ct *ComponentType) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(ComponentType_name[int32(*ct)])
	return b, err
}

func (ct *ComponentType) UnmarshalJSON(b []byte) error {
	ctRaw, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	value := ComponentType(ComponentType_value[ctRaw])
	*ct = value

	return nil
}
