package entity

type Config struct {
	Input  *Input  `json:"input"`
	Action *Action `json:"action"`
}

type Input struct {
	Type string `json:"type"`
	File string `json:"file"`
}

type Action struct {
	Type  string      `json:"type"`
	Param interface{} `json:"param"`
}
