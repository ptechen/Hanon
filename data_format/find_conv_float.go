package data_format

import (
	"context"
	"regexp"
	"strconv"
)

type FindConvFloat []string


func (p FindConvFloat) findConvFloat(ctx context.Context, params string) interface{} {
	for _, s := range p {
		matched, _ := regexp.MatchString(s, params)
		if matched {
			reg := regexp.MustCompile(s)
			res := reg.FindStringSubmatch(params)
			n, _ := string2float(res[1])
			return *n
		}
	}
	return nil
}

func string2float(params string) (Float64, error) {
	n, err := strconv.ParseFloat(params, 64)
	return &n, err
}