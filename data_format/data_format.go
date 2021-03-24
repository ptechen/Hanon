package data_format

import "context"

var DataFormatExecOrder = []string{"deletes", "replaces", "splits", "find", "find_iter"}

type DataFormat struct {
	ExecOrder []string `json:"exec_order" yaml:"exec_order"`
	Splits    Splits   `json:"splits" yaml:"splits"`
	Deletes   Deletes  `json:"deletes" yaml:"deletes"`
	Replaces  Replaces `json:"replaces" yaml:"replaces"`
	Find      Find     `json:"find" yaml:"find"`
	FindIter  FindIter `json:"find_iter" yaml:"find_iter"`
}

func (p *DataFormat) DataFormat(ctx context.Context, params interface{}) interface{} {
	if p.ExecOrder == nil || len(p.ExecOrder) == 0 {
		p.ExecOrder = DataFormatExecOrder
	}
	for _, s := range p.ExecOrder {
		params = p.dataFormat(ctx, s, params)
	}
	return params
}

func (p *DataFormat) dataFormat(ctx context.Context, pat, params interface{}) interface{} {
	switch pat {
	case "deletes":
		if p.Deletes != nil {
			return p.Deletes.deletes(ctx, params)
		}
	case "replaces":
		if p.Replaces != nil {
			return p.Replaces.replaces(ctx, params)
		}
	case "splits":
		if p.Splits != nil {
			param, ok := params.(string)
			if ok {
				return p.Splits.splits(ctx, param)
			}
		}
	case "find":
		if p.Find != nil {
			param, ok := params.(string)
			if ok {
				return p.Find.find(ctx, param)
			}
		}
	case "find_iter":
		if p.FindIter != nil {
			param, ok := params.(string)
			if ok {
				return p.FindIter.findIter(ctx, param)
			}
		}
	}
	return params
}
