package tests

import (
	"Hanon/parse"
	"github.com/ptechen/config"
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
	params := &parse.HashMapSelectParams{}
	config.New().YAML("../test_pages/nodes.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface {}{"eq":"first", "first":"first", "last":"last", "last1":""}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}