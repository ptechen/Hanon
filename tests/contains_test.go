package tests

import (
	"Hanon/parse"
	"fmt"
	"github.com/ptechen/config"
	"golang.org/x/net/context"
	"io/ioutil"
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
	fmt.Println(res)
}
