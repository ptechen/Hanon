package tests

import (
	"context"
	"github.com/ptechen/config"
	"github.com/ptechen/hanon/parse_html"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestTextAttrHtml(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/text_attr_html.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params := &parse_html.HashMapSelectParams{}
	config.New().YAML("../test_pages/text_attr_html.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface {}{"attr":"/attr", "html":"\n    <a href=\"/attr\">test</a>\n", "text":"text"}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
