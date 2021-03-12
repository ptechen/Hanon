package data_format

type Replaces []*Replace

type Replace struct {
	From string `json:"from"`
	To   string `json:"to"`
}
