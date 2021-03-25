package parse_html

import (
	"context"
	"errors"
	"regexp"
)

type MatchParseHtml struct {
	/// Regex match html
	Regex string `json:"regex" yaml:"regex"`
	/// Custom error message, return error message directly if the regular expression matches successfully
	Err string `json:"err" yaml:"err"`
	/// Parse the configuration of html
	Fields *HashMapSelectParams `json:"fields" yaml:"fields"`
	/// Add version, you can not add
	Version string `json:"version" yaml:"version"`
}

type MatchHtmlMany []*MatchParseHtml

type RegexesMatchParseHtml struct {
	RegexesMatchParseHtml []*MatchParseHtml `json:"regexes_match_parse_html" yaml:"regexes_match_parse_html"`
}

// RegexesMatchParseHtml is 正则匹配解析 html 入口
func (p *MatchHtmlMany) RegexesMatchParseHtml(ctx context.Context, html string) (map[string]interface{}, error) {
	for i, matchHtml := range *p {
		reg := regexp.MustCompile(matchHtml.Regex)
		result := reg.FindAllStringSubmatch(html, -1)
		if result == nil {
			if i == len(*p)-1 {
				return nil, errors.New("all rules were failed")
			}
			continue
		}
		if matchHtml.Err != "" {
			return nil, errors.New(matchHtml.Err)
		}
		data, err := matchHtml.Fields.ParsingHtml(ctx, html)
		if err != nil {
			return nil, err
		}
		data["version"] = matchHtml.Version
		return data, err
	}
	return nil, nil
}
