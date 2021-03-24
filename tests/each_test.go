package tests

import (
	"Hanon/parse"
	"github.com/ptechen/config"
	"golang.org/x/net/context"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestEach(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("../test_pages/each.html")
	if err != nil {
		t.Error(err)
	}
	dataStr := string(dataBytes)
	params := &parse.HashMapSelectParams{}
	config.New().YAML("../test_pages/each.yml", params)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := params.ParsingHtml(ctx, dataStr)
	if err != nil {
		t.Error(err)
	}
	v := map[string]interface{}{"each": []interface{}{"上海"},
		"each_all": []interface{}{"北京", "上海", "杭州"},
		"each_fields": []interface{}{
			map[string]interface{}{"city_field": "北京", "code_field": "1"},
			map[string]interface{}{"city_field": "上海", "code_field": "2"},
			map[string]interface{}{"city_field": "杭州", "code_field": "3"}},
		"each_one": "北京", "each_one_contains": "上海"}
	if !reflect.DeepEqual(res, v) {
		t.Errorf("left: %#v, right: %#v", res, v)
	}
}
