package data_format

type DataFormat struct {
	Splits   *Splits   `json:"splits"`
	Deletes  *Deletes  `json:"deletes"`
	Replaces *Replaces `json:"replaces"`
	Find     *Find     `json:"find"`
	FindIter *FindIter `json:"find_iter"`
}
