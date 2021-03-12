package parse

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

var DefaultExecOrder = []string{"selects", "each", "select_params", "nodes", "has", "contains"}

type SelectParams struct {
	ExecOrder      ExecOrder    `json:"exec_order"`
	Selects        Selects      `json:"selects"`
	Each           *Each         `json:"each"`
	SelectParams   *SelectParams `json:"select_params"`
	Nodes          *Node         `json:"nodes"`
	Has            *Has          `json:"has"`
	Contains       *Contains     `json:"contains"`
	TextAttrHtml   *TextAttrHtml `json:"text_attr_html"`
	DataFormat     *DataFormat   `json:"data_format"`
	DefaultValType string        `json:"default_val_type"`

	doc *goquery.Document
}

type (
	HashMapSelectParams map[string]*SelectParams
	ExecOrder           []string
	Selects             []string
)

// ParsingHtml is 解析html的入口
func (params *HashMapSelectParams) ParsingHtml(html string, ds *DocumentSelection) (res map[string]interface{}, err error) {
	if html != "" && ds == nil {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			return nil, err
		}
		ds = &DocumentSelection{
			Selection: doc.Selection,
			Node:      nil,
		}
	}
	res = params.InitResMap()
	wg := sync.WaitGroup{}
	for key, selectParams := range *params {
		res[key] = nil
		wg.Add(1)
		go func(key string, selectParams *SelectParams) {
			defer wg.Done()
			res[key] = ds.parse(selectParams)
		}(key, selectParams)
	}
	wg.Wait()
	return
}

// initResMap
func (params *HashMapSelectParams) InitResMap() (res map[string]interface{}) {
	res = map[string]interface{}{}
	for key, _ := range *params {
		res[key] = nil
	}
	return
}
