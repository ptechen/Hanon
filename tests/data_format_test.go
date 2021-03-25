package tests

import (
	"context"
	"github.com/ptechen/config"
	"github.com/ptechen/hanon/parse_html"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestDataFormat(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/data_format.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params := &parse_html.HashMapSelectParams{}
	config.New().YAML("../test_pages/data_format.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface{}{"deletes": "ttttt", "find": "123", "find_iter": []string{"123", "123"},
		"replaces": "1234ttttt", "splits": []string{"testt", "tttt"}, "splits_index": "testt"}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
