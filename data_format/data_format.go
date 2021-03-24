package data_format

type DataFormat struct {
	Splits   Splits   `json:"splits" yaml:"splits"`
	Deletes  Deletes  `json:"deletes" yaml:"deletes"`
	Replaces Replaces `json:"replaces" yaml:"replaces"`
	Find     Find     `json:"find" yaml:"find"`
	FindIter FindIter `json:"find_iter" yaml:"find_iter"`
}
