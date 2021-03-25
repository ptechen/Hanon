package data_format

import (
	"context"
	"strings"
)

type Splits []*Split
type Index int

type Split struct {
	Key   string `json:"key" yaml:"key"`
	Index *Index `json:"index" yaml:"index"`
}

func (s Splits) splits(ctx context.Context, params string) interface{} {
	for _, split := range s {
		if split.Index != nil {
			params = split.splitIndex(ctx, params)
		} else {
			return split.split(ctx, params)
		}
	}
	return params
}

func (s Split) splitIndex(ctx context.Context, params string) string {
	data := strings.Split(params, s.Key)
	if int(*s.Index) < len(data) {
		return data[int(*s.Index)]
	} else {
		return ""
	}
}

func (s Split) split(ctx context.Context, params string) []string {
	return strings.Split(params, s.Key)
}
