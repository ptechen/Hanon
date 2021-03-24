package data_format

type Splits []*Split

type Split struct {
	Key   string `yaml:"key"`
	Index int    `yaml:"index"`
}
