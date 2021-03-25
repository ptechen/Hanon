package tests

import (
	"context"
	"github.com/ptechen/config"
	"github.com/ptechen/hanon/parse_html"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMatchParseHtml(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/match_html.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params:= parse_html.MatchHtmlMany{}
	config.New().YAML("../test_pages/match_html.yml", &params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.RegexesMatchParseHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface {}{"test":"test", "version":"2"}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}