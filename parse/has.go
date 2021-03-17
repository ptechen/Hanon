package parse

import "context"

type Has struct{
	Class string `json:"class"`
	Attr string `json:"attr"`
}

func (p *Has) class(ctx context.Context, ds *DocumentSelection)  {

}