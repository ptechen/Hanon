package tests

import (
	"Hanon/parse"
	"fmt"
	"github.com/ptechen/config"
	"golang.org/x/net/context"
	"io/ioutil"
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
	res, err := params.ParsingHtml(ctx, dataStr, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
