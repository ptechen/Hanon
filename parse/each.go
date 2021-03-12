package parse

type Each struct {
	All    *SelectParams        `json:"all"`
	One    *SelectParams        `json:"one"`
	Fields *HashMapSelectParams `json:"fields"`
}

func (p *Each) each(ds DocumentSelection) (res interface{}, err error) {
	if p.All != nil {
		return p.all(ds), nil
	} else if p.One != nil {
		return p.one(ds), nil
	} else {
		return p.fields(ds)
	}
}

func (p *Each) all(ds DocumentSelection) interface{} {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		ds := NewDocumentSelection(nil, node)
		v := ds.parse(p.All)
		if v == nil {
			continue
		}
		array = append(array, v)
	}
	return array
}

func (p *Each) one(ds DocumentSelection) interface{} {
	for _, node := range ds.Selection.Nodes {
		ds := NewDocumentSelection(nil, node)
		v := ds.parse(p.One)
		if v == nil {
			continue
		}
		return v
	}
	return nil // todo 添加默认值
}

func (p *Each) fields(ds DocumentSelection) (res interface{}, err error) {
	array := make([]interface{}, 0, len(ds.Selection.Nodes))
	for _, node := range ds.Selection.Nodes {
		ds := NewDocumentSelection(nil, node)
		curMap, err := p.Fields.ParsingHtml("", ds)
		if err != nil {
			return res, err
		}
		array = append(array, curMap)
	}
	return array, err
}
