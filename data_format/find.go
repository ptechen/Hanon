package data_format

import (
	"context"
	"regexp"
)

type Find []string

func (p Find) find(ctx context.Context, params interface{}) interface{} {
	switch params.(type) {
	case string:
		return p.string(ctx, params.(string))
	case []interface{}:
		return p.slice(ctx, params.([]interface{}))
	case map[string]interface{}:
		return p.hashmap(ctx, params.(map[string]interface{}))
	}
	return ""
}

func (p Find) string(ctx context.Context, params string) string {
	for _, s := range p {
		matched, _ := regexp.MatchString(s, params)
		if matched {
			reg := regexp.MustCompile(s)
			res := reg.FindStringSubmatch(params)
			return res[1]
		}
	}
	return ""
}

func (p Find) hashmap(ctx context.Context, params map[string]interface{}) interface{} {
	for key, val := range params {
		params[key] = p.find(ctx, val)
	}
	return params
}

func (p Find) slice(ctx context.Context, params []interface{}) interface{} {
	for idx, val := range params {
		params[idx] = p.find(ctx, val)
	}
	return params
}

