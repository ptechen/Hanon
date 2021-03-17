package parse

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type DocumentSelection struct {
	Selection *goquery.Selection
}

func NewDocumentSelection(ctx context.Context, selection *goquery.Selection) (res *DocumentSelection) {
	return &DocumentSelection{
		Selection: selection,
	}
}

func NewDocumentSelectionByNode(ctx context.Context, node *html.Node) (res *DocumentSelection) {
	return &DocumentSelection{
		Selection: &goquery.Selection{Nodes: []*html.Node{node}},
	}
}

func (p *DocumentSelection) parse(ctx context.Context, params *SelectParams) (interface{}, error) {
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
			return params.Each.each(ctx, p)

		case "select_params":
			if params.SelectParams == nil {
				continue
			}
			return p.parse(ctx, params)
		default:
			p.parseKey(ctx, params, val)
		}
	}
	return p.content(ctx, params)
}

func (p *DocumentSelection) parseKey(ctx context.Context, params *SelectParams, key string)  {
	switch key {
	case "selects":
		if params.Selects == nil {
			return
		}
		p.selects(ctx, params.Selects)
	case "nodes":
		if params.Nodes == nil {
			return
		}
		params.Nodes.run(ctx, p)
	//case "has":
	//	if params.Has == nil {
	//		return
	//	}
	//	params.Has.
	}
}

func (p *DocumentSelection) content(ctx context.Context, params *SelectParams) (string, error) {
	if params.TextAttrHtml == nil {
		params.TextAttrHtml = newTextAttrHtml()
	}
	return params.TextAttrHtml.call(ctx, p)
}

func (p *DocumentSelection) selects(ctx context.Context, params []string)  {
	for _, param := range params {
		p.Selection = p.Selection.Find(param)
	}
}

func (p *DocumentSelection) html(ctx context.Context, params *SelectParams) (interface{}, error) {
	return p.Selection.Html()
}

func (p *DocumentSelection) text(ctx context.Context, params *SelectParams) string {
	return p.Selection.Text()
}