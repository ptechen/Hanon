package data_format

import (
	"context"
	"strings"
)

type Deletes = []string

func (p Deletes) deletes(ctx context.Context, params string) string {
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