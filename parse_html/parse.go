package parse_html

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/ptechen/hanon/data_format"
	"strings"
	"sync"
)

var DefaultExecOrder = []string{"selects", "each", "select_params", "nodes", "contains"}

type SelectParams struct {
	ExecOrder      ExecOrder               `json:"exec_order" yaml:"exec_order"`
	Selects        Selects                 `json:"selects" yaml:"selects"`
	Each           *Each                   `json:"each" yaml:"each"`
	SelectParams   *SelectParams           `json:"select_params" yaml:"select_params"`
	Nodes          *Node                   `json:"nodes" yaml:"nodes"`
	Contains       *Contains               `json:"contains" yaml:"contains"`
	TextAttrHtml   *TextAttrHtml           `json:"text_attr_html" yaml:"text_attr_html"`
	DataFormat     *data_format.DataFormat `json:"data_format" yaml:"data_format"`
	DefaultValType string                  `json:"default_val_type" yaml:"default_val_type"`
}

type (
	HashMapSelectParams map[string]*SelectParams
	ExecOrder           []string
	Selects             []string
)

// ParsingHtml is 解析html的入口
func (params *HashMapSelectParams) ParsingHtml(ctx context.Context, html string) (res map[string]interface{}, err error) {
	html = strings.ReplaceAll(html," ", "")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	ds := &DocumentSelection{
		Selection: doc.Selection,
	}
	return params.parsingHtml(ctx, ds)
}

func (params *HashMapSelectParams) parsingHtml(ctx context.Context, ds *DocumentSelection) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	for key, selectParams := range *params {
		wg.Add(1)
		go func(key string, selectParams *SelectParams) {
			defer wg.Done()
			val, _ := ds.parse(ctx, selectParams)
			lock.Lock()
			res[key] = val
			lock.Unlock()
		}(key, selectParams)
	}
	wg.Wait()
	return
}

func (p *SelectParams) defaultVal(ctx context.Context) interface{} {
	switch p.DefaultValType {
	case "str", "string":
		return ""
	case "[]", "slice", "array":
		return []interface{}{}
	case "map":
		return map[interface{}]interface{}{}
	default:
		return nil
	}
}
