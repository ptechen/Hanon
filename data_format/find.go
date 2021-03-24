package data_format

import (
	"context"
	"regexp"
)

type Find []string

func (p Find) find(ctx context.Context, params string) string {
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
