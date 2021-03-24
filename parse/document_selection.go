package parse

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type DocumentSelection struct {
	Selection *goquery.Selection `json:"selection"`
}

func NewDocumentSelectionByNode(ctx context.Context, node *html.Node) (res *DocumentSelection) {
	return &DocumentSelection{Selection: &goquery.Selection{Nodes: []*html.Node{node}}}
}

func NewDocumentSelectionBySelection(ctx context.Context, selection *goquery.Selection) (res *DocumentSelection) {
	return &DocumentSelection{Selection: selection}
}

func (p DocumentSelection) parse(ctx context.Context, params *SelectParams) (interface{}, error) {
	return p.parseExecOrder(ctx, params)
}

func (p *DocumentSelection) parseExecOrder(ctx context.Context, params *SelectParams) (interface{}, error) {
	if params.ExecOrder == nil {
		params.ExecOrder = DefaultExecOrder
	}
	for _, val := range params.ExecOrder {
		switch val {
		case "each":
			if params.Each == nil {
				continue
			}
			return params.Each.each(ctx, p), nil

		case "select_params":
			if params.SelectParams == nil {
				continue
			}
			return p.parse(ctx, params)
		default:
			err := p.parseKey(ctx, params, val)
			if err != nil {
				return nil, err
			}
		}
	}
	return p.content(ctx, params)
}

func (p *DocumentSelection) parseKey(ctx context.Context, params *SelectParams, key string) error {
	switch key {
	case "selects":
		if params.Selects == nil {
			return nil
		}
		p.selects(ctx, params.Selects)
	case "nodes":
		if params.Nodes == nil {
			return nil
		}
		params.Nodes.call(ctx, p)

	case "contains":
		if params.Contains == nil {
			return nil
		}
		ok := params.Contains.call(ctx, *p)
		if !ok {
			return errors.New("class or attr or html or text is not exist")
		}
	}
	return nil
}

func (p *DocumentSelection) content(ctx context.Context, params *SelectParams) (interface{}, error) {
	if params.TextAttrHtml == nil {
		params.TextAttrHtml = newTextAttrHtml()
	}
	data, err := params.TextAttrHtml.call(ctx, p)
	if err != nil {
		return "", err
	}
	if params.DataFormat != nil {
		return params.DataFormat.DataFormat(ctx, data), err
	}
	return data, err
}

func (p *DocumentSelection) selects(ctx context.Context, params []string) {
	for _, param := range params {
		p.Selection = p.Selection.Find(param)
	}
}
