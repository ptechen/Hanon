package parse

import "context"

type Each struct {
	All    *SelectParams        `json:"all"`
	One    *SelectParams        `json:"one"`
	Fields *HashMapSelectParams `json:"fields"`
}

func (p *Each) each(ctx context.Context, ds *DocumentSelection) (interface{}, error) {
	if p.All != nil {
		return p.all(ctx, ds)
	} else if p.One != nil {
		return p.one(ctx, ds)
	} else if p.Fields != nil {
		return p.fields(ctx, ds)
	}
	return nil, nil
}

func (p *Each) all(ctx context.Context, ds *DocumentSelection) (interface{}, error) {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(ctx, node)
		v, err := curDs.parse(ctx, p.All)
		if err == nil {
			return nil, err
		}
		if v == nil {
			continue
		}
		array = append(array, v)
	}
	return array, nil
}

func (p *Each) one(ctx context.Context, ds *DocumentSelection) (interface{}, error) {
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(ctx, node)
		v, err := curDs.parse(ctx, p.One)
		if err == nil {
			return nil, err
		}
		if v == nil {
			continue
		}
		return v, nil
	}
	return nil, nil // todo 添加默认值
}

func (p *Each) fields(ctx context.Context, ds *DocumentSelection) (res interface{}, err error) {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		curDs := NewDocumentSelectionByNode(ctx, node)
		curMap, err := p.Fields.ParsingHtml(ctx, "", curDs)
		if err != nil {
			return res, err
		}
		array = append(array, curMap)
	}
	return array, err
}
