package parse

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type DocumentSelection struct {
	Selection *goquery.Selection
	Node      *html.Node
}

func NewDocumentSelection(selection *goquery.Selection, node *html.Node) (res *DocumentSelection) {
	return &DocumentSelection{
		Selection: selection,
		Node:      node,
	}
}

func (p DocumentSelection) parse(params *SelectParams) interface{} {
	return p.parseExecOrder(params)
}

func (p DocumentSelection) parseExecOrder(params *SelectParams) interface{} {
	if params.ExecOrder == nil {
		params.ExecOrder = DefaultExecOrder
	}
	for _, val := range params.ExecOrder {
		switch val {
		case "each":
			if params.Each != nil {
				params.Each.each(p)
			}
		}
	}
	return
}

func (p DocumentSelection) html(params *SelectParams) (interface{}, error){
	if p.Selection != nil {
		return p.Selection.Html()
	} else if p.Node != nil {
		return goquery.Selection{}
	}
}