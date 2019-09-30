package components

import (
	"encoding/json"
	"strings"
)

type Type string

const (
	TypeFilter     = "filter"
	TypeTranscoder = "transcoder"
)

type IComponent interface {
	GetType() Type
	GetName() string
	GetDescription() string
	GetParams() Params
	Render() string
	Serialize() string
}

type Components []Component

type Component struct {
	Type        Type   `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      Params `json:"params"`
}

func Deserialize(d []byte) (*Component, error) {
	c := &Component{}
	err := json.Unmarshal(d, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Component) Render() string {
	switch c.Type {
	case TypeFilter:
	case TypeTranscoder:
	}

	var built []string
	for _, p := range c.Params {
		built = append(built, p.Render())
	}

	return strings.Join(built, " ")
}

func (c *Component) GetType() Type {
	return c.Type
}

func (c *Component) GetName() string {
	return c.Name
}

func (c *Component) GetDescription() string {
	return c.Description
}

func (c *Component) GetParams() Params {
	return c.Params
}
