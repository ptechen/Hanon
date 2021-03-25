package data_format

import (
	"context"
	"regexp"
)

type FindIter []string

func (p FindIter) findIter(ctx context.Context, params interface{}) interface{} {
	switch params.(type) {
	case string:
		return p.string(ctx, params.(string))
	case map[string]interface{}:
		return p.hashmap(ctx, params.(map[string]interface{}))
	}
	return []string{}
}

func (p FindIter) string(ctx context.Context, params string) []string {
	for _, s := range p {
		matched, _ := regexp.MatchString(s, params)
		if matched {
			reg := regexp.MustCompile(s)
			res := reg.FindStringSubmatch(params)
			return res
		}
	}
	return []string{}
}

func (p FindIter) hashmap(ctx context.Context, params map[string]interface{}) interface{} {
	for key, val := range params {
		params[key] = p.findIter(ctx, val)
	}
	return params
}

func (p FindIter) slice(ctx context.Context, params []interface{}) interface{} {
	for idx, val := range params {
		params[idx] = p.findIter(ctx, val)
	}
	return params
}