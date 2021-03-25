package parse_html

import "context"

type Each struct {
	All    *SelectParams        `json:"all" yaml:"all"`
	One    *SelectParams        `json:"one" yaml:"one"`
	Fields *HashMapSelectParams `json:"fields" yaml:"fields"`
}

func (p *Each) each(ctx context.Context, ds *DocumentSelection) interface{} {
	if p.All != nil {
		return p.all(ctx, ds)
	} else if p.One != nil {
		return p.one(ctx, ds)
	} else if p.Fields != nil {
		return p.fields(ctx, ds)
	}
	return nil
}

func (p *Each) all(ctx context.Context, ds *DocumentSelection) interface{} {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(node)
		v, err := curDs.parse(ctx, p.All)
		if err != nil {
			continue
		}
		if v == nil || v == ""{
			continue
		}
		array = append(array, v)
	}
	return array
}

func (p *Each) one(ctx context.Context, ds *DocumentSelection) interface{} {
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(node)
		v, err := curDs.parse(ctx, p.One)
		if err != nil {
			continue
		}
		if v == nil || v == ""{
			continue
		}
		return v
	}
	return p.One.defaultVal(ctx)
}

func (p *Each) fields(ctx context.Context, ds *DocumentSelection) interface{} {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(node)
		curMap, err := p.Fields.parsingHtml(ctx, curDs)
		if err != nil {
			continue
		}
		array = append(array, curMap)
	}
	return array
}
