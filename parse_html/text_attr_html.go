package parse_html

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

func (p *TextAttrHtml) call(ctx context.Context, ds *DocumentSelection, contains *Contains) (string, error) {
	if p.Text == true {
		return p.text(ctx, ds, contains), nil
	} else if p.Html == true {
		return p.html(ctx, ds, contains)
	} else if p.Attr != "" {
		return p.attr(ctx, ds, contains), nil
	}
	return "", nil
}

func (p *TextAttrHtml) text(ctx context.Context, ds *DocumentSelection, contains *Contains) string {
	if contains == nil{
		val := ds.Selection.Text()
		return val
	} else if len(ds.Selection.Nodes) == 1 {
		ok := contains.call(ctx, *ds)
		if ok {
			val := ds.Selection.Text()
			return val
		}
	} else if len(ds.Selection.Nodes) > 1 {
		res := ""
		for _, node := range ds.Selection.Nodes {
			curDs := NewDocumentSelectionByNode(node)
			ok := contains.call(ctx, *curDs)
			if ok {
				val := curDs.Selection.Text()
				res += val
			}
		}
		return res
	}
	return ""
}

func (p *TextAttrHtml) html(ctx context.Context, ds *DocumentSelection, contains *Contains) (string, error) {
	if contains == nil{
		val, err := ds.Selection.Html()
		return val, err
	} else if len(ds.Selection.Nodes) == 1{
		ok := contains.call(ctx, *ds)
		if ok {
			val, err := ds.Selection.Html()
			return val, err
		}
	} else if len(ds.Selection.Nodes) > 1 {
		res := ""
		for _, node := range ds.Selection.Nodes {
			curDs := NewDocumentSelectionByNode(node)
			ok := contains.call(ctx, *curDs)
			if ok {
				val, err := curDs.Selection.Html()
				if err != nil {
					return "", err
				}
				res += val
			}
		}
		return res, nil
	}
	return "", nil
}

func (p *TextAttrHtml) attr(ctx context.Context, ds *DocumentSelection, contains *Contains) string {
	if contains == nil{
		val, _ := ds.Selection.Attr(p.Attr)
		return val
	} else if len(ds.Selection.Nodes) == 1{
		ok := contains.call(ctx, *ds)
		if ok {
			val,_ := ds.Selection.Attr(p.Attr)
			return val
		}
	} else if len(ds.Selection.Nodes) > 1 {
		res := ""
		for _, node := range ds.Selection.Nodes {
			curDs := NewDocumentSelectionByNode(node)
			ok := contains.call(ctx, *curDs)
			if ok {
				val, _ := curDs.Selection.Attr(p.Attr)
				res += val
			}
		}
		return res
	}
	return ""
}
