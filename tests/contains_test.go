package tests

import (
	"Hanon/parse"
	"github.com/ptechen/config"
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
	params := &parse.HashMapSelectParams{}
	config.New().YAML("../test_pages/contains.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface{}{"contains": "北京北京1上海2杭州3", "each_contains": []interface{}{"北京", "杭州"}}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
