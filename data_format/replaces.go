package data_format

import (
	"context"
	"strings"
)

type Replaces []*Replace

type Replace struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func (p Replaces) replaces(ctx context.Context, params string) string {
	for _, replace := range p {
		params = replace.replace(ctx, params)
	}
	return params
}

func (p *Replace)replace(ctx context.Context, params string ) string {
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