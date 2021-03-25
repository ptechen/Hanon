package tests

import (
	"github.com/ptechen/config"
	"github.com/ptechen/hanon/parse_html"
	"golang.org/x/net/context"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNodes(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/nodes.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params := &parse_html.HashMapSelectParams{}
	config.New().YAML("../test_pages/nodes.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface{}{"children": "123", "eq": "first1", "first": "first", "last": "123",
		"next_sibling": "first1last", "parent": "\n    first\n    first1\n    last\n", "prev_sibling": "first1"}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
