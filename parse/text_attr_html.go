package parse

import (
	"context"
)

type TextAttrHtml struct {
	Text bool   `yaml:"text"`
	Html bool   `yaml:"html"`
	Attr string `yaml:"attr"`
}

func newTextAttrHtml() *TextAttrHtml {
	return &TextAttrHtml{
		Text: true,
	}
}

func (p *TextAttrHtml) call(ctx context.Context, ds *DocumentSelection) (string, error) {
	if p.Text == true {
		return p.text(ctx, ds), nil
	} else if p.Html == true {
		return p.html(ctx, ds)
	} else if p.Attr != "" {
		v, _ := p.attr(ctx, ds)
		return v, nil
	}
	return "", nil
}

func (p *TextAttrHtml) text(ctx context.Context, ds *DocumentSelection) string {
	val := ds.Selection.Text()
	return val
}

func (p *TextAttrHtml) html(ctx context.Context, ds *DocumentSelection) (string, error) {
	return ds.Selection.Html()
}

func (p *TextAttrHtml) attr(ctx context.Context, ds *DocumentSelection) (string, bool) {
	return ds.Selection.Attr(p.Attr)
}
