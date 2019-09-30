package components

type Range struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}

type Param struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Range       Range  `json:"range"`
}

type Params []Param

func (p Param) Render() string {
	return p.Key + " " + p.Value
}
