package parse

import (
	"context"
	"strings"
)

type Contains struct {
	Contains    *TextClassAttrHtml `yaml:"contains"`
	NotContains *TextClassAttrHtml `yaml:"not_contains"`
}

type TextClassAttrHtml struct {
	Class    []string `yaml:"class"`
	Attr     []string `yaml:"attr"`
	Html     []string `yaml:"html"`
	Text     []string `yaml:"text"`
}

func (p *Contains) call(ctx context.Context, ds *DocumentSelection) bool {
	if p.contains(ctx, ds) && p.notContains(ctx, ds) {
		return true
	}
	return false
}

func (p *Contains) contains(ctx context.Context, ds *DocumentSelection) bool {
	if p.Contains == nil {
		return true
	}
	return p.Contains.containsCall(ctx, ds)
}

func (p *Contains) notContains(ctx context.Context, ds *DocumentSelection) bool {
	if p.NotContains == nil {
		return true
	}
	return p.NotContains.notContainsCall(ctx, ds)
}

func (p *TextClassAttrHtml) containsCall(ctx context.Context, ds *DocumentSelection) bool {
	if p.Class != nil {
		if !p.class(ctx, ds) {
			return false
		}
	}

	if p.Attr != nil {
		if !p.attr(ctx, ds) {
			return false
		}
	}

	if p.Text != nil {
		if !p.text(ctx, ds) {
			return false
		}
	}

	if p.Html != nil {
		if !p.html(ctx, ds) {
			return false
		}
	}

	return true
}

func (p *TextClassAttrHtml) notContainsCall(ctx context.Context, ds *DocumentSelection) bool {
	if p.Class != nil {
		if p.class(ctx, ds) {
			return false
		}
	}

	if p.Attr != nil {
		if p.attr(ctx, ds) {
			return false
		}
	}

	if p.Text != nil {
		if p.text(ctx, ds) {
			return false
		}
	}

	if p.Html != nil {
		if p.html(ctx, ds) {
			return false
		}
	}

	return true
}

func (p *TextClassAttrHtml) html(ctx context.Context, ds *DocumentSelection) bool {
	html, _ := ds.Selection.Html()
	for _, pat := range p.Html {
		if !strings.Contains(html, pat) {
			return false
		}
	}
	return true
}

func (p *TextClassAttrHtml) text(ctx context.Context, ds *DocumentSelection) bool {
	text := ds.Selection.Text()
	for _, pat := range p.Text {
		if !strings.Contains(text, pat) {
			return false
		}
	}
	return true
}

func (p *TextClassAttrHtml) attr(ctx context.Context, ds *DocumentSelection) bool {
	for _, pat := range p.Text {
		_, b := ds.Selection.Attr(pat)
		if !b {
			return false
		}
	}
	return true
}

func (p *TextClassAttrHtml) class(ctx context.Context, ds *DocumentSelection) bool {
	for _, pat := range p.Class {
		if !ds.Selection.HasClass(pat) {
			return false
		}
	}
	return true
}
