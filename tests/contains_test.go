package tests

import (
	"github.com/ptechen/config"
	"github.com/ptechen/hanon/parse_html"
	"golang.org/x/net/context"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/contains.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params := &parse_html.HashMapSelectParams{}
	config.New().YAML("../test_pages/contains.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface{}{"contains": []interface{}{"test4"}}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
