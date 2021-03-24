package data_format

import (
	"context"
	"strings"
)

type Replaces []*Replace

type Replace struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

func (p Replaces) replaces(ctx context.Context, params interface{}) interface{} {
	for _, replace := range p {
		params = replace.replace(ctx, params)
	}
	return params
}

func (p *Replace) replace(ctx context.Context, params interface{}) interface{} {
	switch params.(type) {
	case string:
		return p.string(ctx, params.(string))
	case map[string]interface{}:
		return p.hashmap(ctx, params.(map[string]interface{}))
	case []interface{}:
		return p.slice(ctx, params.([]interface{}))
	}
	return params
}

func (p *Replace) string(ctx context.Context, params string) interface{} {
	if p != nil {
		if p.From == "\\n" {
			params = strings.ReplaceAll(params, "\n", p.To)
		} else if p.From == "\\t" {
			params = strings.ReplaceAll(params, "\t", p.To)
		} else {
			params = strings.ReplaceAll(params, p.From, p.To)
		}
	}
	return params
}

func (p *Replace) hashmap(ctx context.Context, params map[string]interface{}) interface{} {
	for key, val := range params {
		params[key] = p.replace(ctx, val)
	}
	return params
}

func (p Replace) slice(ctx context.Context, params []interface{}) interface{} {
	for i, param := range params {
		params[i] = p.replace(ctx, param)
	}
	return params
}
