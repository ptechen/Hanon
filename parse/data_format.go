package parse

import "Hanon/data_format"

type DataFormat struct {
	Splits   *data_format.Splits   `json:"splits"`
	Deletes  *data_format.Deletes  `json:"deletes"`
	Replaces *data_format.Replaces `json:"replaces"`
	Find     *data_format.Find     `json:"find"`
	FindIter *data_format.FindIter `json:"find_iter"`
}
