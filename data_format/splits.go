package data_format

type Splits []*Split

type Split struct {
	Key   string `json:"key"`
	Index int    `json:"index"`
}
