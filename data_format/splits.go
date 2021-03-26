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

func (p Splits) splits(ctx context.Context, params interface{}) interface{} {
	for _, split := range p {
		params = split.split(ctx, params)
	}
	return params
}

func (p Split) split(ctx context.Context, params interface{}) interface{} {
	switch params.(type) {
	case string:
		return p.string(ctx, params.(string))
	case []interface{}:
		return p.slice(ctx, params.([]interface{}))
	case map[string]interface{}:
		return p.hashmap(ctx, params.(map[string]interface{}))
	}
	return params
}

func (p Split) string(ctx context.Context, params string) interface{} {
	if p.Index != nil {
		return p.splitIndex(ctx, params)
	} else {
		return strings.Split(params, p.Key)
	}
}

func (p Split) hashmap(ctx context.Context, params map[string]interface{}) interface{} {
	for key, val := range params {
		params[key] = p.split(ctx, val)
	}
	return params
}

func (p Split) slice(ctx context.Context, params []interface{}) interface{} {
	for idx, val := range params {
		params[idx] = p.split(ctx, val)
	}
	return params
}

func (p Split) splitIndex(ctx context.Context, params string) string {
	data := strings.Split(params, p.Key)
	if int(*p.Index) < len(data) {
		return data[int(*p.Index)]
	} else {
		return ""
	}
}


