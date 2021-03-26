package data_format

import "context"

var DataFormatExecOrder = []string{"deletes", "replaces", "splits", "find", "find_iter", "find_conv_int", "find_conv_float"}

type DataFormat struct {
	ExecOrder     []string      `json:"exec_order" yaml:"exec_order"`
	Splits        Splits        `json:"splits" yaml:"splits"`
	Deletes       Deletes       `json:"deletes" yaml:"deletes"`
	Replaces      Replaces      `json:"replaces" yaml:"replaces"`
	Find          Find          `json:"find" yaml:"find"`
	FindIter      FindIter      `json:"find_iter" yaml:"find_iter"`
	FindConvInt   FindConvInt   `json:"find_conv_int" yaml:"find_conv_int"`
	FindConvFloat FindConvFloat `json:"find_conv_float" yaml:"find_conv_float"`
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
	case "find_conv_int":
		if p.FindConvInt != nil {
			param, ok := params.(string)
			if ok {
				return p.FindConvInt.findConvInt(ctx, param)
			}
		}
	case "find_conv_float":
		if p.FindConvFloat != nil {
			param, ok := params.(string)
			if ok {
				return p.FindConvFloat.findConvFloat(ctx, param)
			}
		}
	}
	return params
}
