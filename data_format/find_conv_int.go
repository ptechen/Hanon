package data_format

import (
	"context"
	"regexp"
	"strconv"
)

type FindConvInt []string

func (p FindConvInt) findConvInt(ctx context.Context, params string) interface{} {
	for _, s := range p {
		matched, _ := regexp.MatchString(s, params)
		if matched {
			reg := regexp.MustCompile(s)
			res := reg.FindStringSubmatch(params)
			n, _ := string2int(res[1])
			return *n
		}
	}
	return nil
}

func string2int(params string) (Int, error) {
	n, err := strconv.Atoi(params)
	return &n, err
}