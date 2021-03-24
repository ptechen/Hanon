package data_format

import (
	"context"
	"strings"
)

type Splits []*Split

type Split struct {
	Key    string `json:"key" yaml:"key"`
	Index  int    `json:"index" yaml:"index"`
	Enable bool   `json:"enable" yaml:"enable"`
}

func (s Splits) splits(ctx context.Context, params string) interface{} {
	for _, split := range s {
		if split.Enable {
			params = split.splitEnableTrue(ctx, params)
		} else {
			return split.splitEnableTrue(ctx, params)
		}
	}
	return nil
}

func (s Split) splitEnableTrue(ctx context.Context, params string) string {
	data := strings.Split(params, s.Key)
	if s.Index < len(data) {
		return data[s.Index]
	} else {
		return ""
	}
}

func (s Split) splitEnableFalse(ctx context.Context, params string) []string {
	return strings.Split(params, s.Key)
}
