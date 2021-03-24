package parse

import (
	"context"
	"fmt"
)

type Node struct {
	First       bool `yaml:"first"`
	Last        bool `yaml:"last"`
	Parent      bool `yaml:"parent"`
	Children    bool `yaml:"children"`
	PrevSibling bool `yaml:"prev_sibling"`
	NextSibling bool `yaml:"next_sibling"`
	Eq          int  `yaml:"eq"`
}

func (p *Node) call(ctx context.Context, params *DocumentSelection)  {
	if p.First == true {
		p.first(ctx, params)
	} else if p.Last == true {
		p.last(ctx, params)
	} else if p.Parent == true {
		p.parent(ctx, params)
	} else if p.Children == true {
		p.children(ctx, params)
	} else if p.PrevSibling != true {
		p.prevSibling(ctx, params)
	} else if p.NextSibling == true {
		p.nextSibling(ctx, params)
	} else {
		p.eq(ctx, params)
	}
}

func (p *Node) first(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.First()
	vsl := params.Selection.Text()
	fmt.Println(vsl)
}

func (p *Node) last(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Last()
}

func (p *Node) parent(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Parent()
}

func (p *Node) children(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Children()
}

func (p *Node) prevSibling(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Prev()
}

func (p *Node) nextSibling(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Next()
}

func (p *Node) eq(ctx context.Context, params *DocumentSelection) {
	params.Selection = params.Selection.Eq(p.Eq)
}
