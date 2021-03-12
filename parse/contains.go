package parse

type Contains struct {
	Contains    *TextHtml `json:"contains"`
	NotContains *TextHtml `json:"not_contains"`
}

type TextHtml struct {
	Html []string `json:"html"`
	Text []string `json:"text"`
}
