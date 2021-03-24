package data_format

import (
	"context"
	"strings"
)

type Deletes []string

func (p Deletes) deletes(ctx context.Context, params interface{}) interface{} {
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

func (p Deletes) string(ctx context.Context, params string) string {
	for _, s := range p {
		switch s {
		case "\\n":
			params = strings.ReplaceAll(params, "\n", "")
		case "\\t":
			params = strings.ReplaceAll(params, "\t", "")
		default:
			params = strings.ReplaceAll(params, s, "")
		}
	}
	return params
}

func (p Deletes) hashmap(ctx context.Context, params map[string]interface{}) map[string]interface{} {
	for key, val := range params {
		params[key] =  p.deletes(ctx, val)
	}
	return params
}

func (p Deletes) slice(ctx context.Context, params []interface{}) []interface{} {
	for i, param := range params {
		params[i] = p.deletes(ctx, param)
	}
	return params
}