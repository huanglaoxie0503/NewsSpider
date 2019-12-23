package rpcDemo

import "errors"

// Service.Method

type DemoService struct{}

type Params struct {
	A, B int
}

// 做一个除法
func (DemoService) Div(p Params, result *float64) error {
	if p.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(p.A) / float64(p.B)

	return nil
}

// {"method":"DemoService.Div", "params": [{"A":5, "B": 10}], "id":1}
