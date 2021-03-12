package parse

type Node struct {
	First       bool `json:"first"`
	Last        bool `json:"last"`
	Eq          int  `json:"eq"`
	Parent      bool `json:"parent"`
	Children    bool `json:"children"`
	PrevSibling bool `json:"prev_sibling"`
	NextSibling bool `json:"next_sibling"`
}
