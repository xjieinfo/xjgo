package main

import (
	"github.com/xjieinfo/xjgo/examples/rpcdemo"
	"math"
)

// MathUtil 用于数学计算
type RemoteMathUtil struct{}

// CaculateCircleArea 计算圆的面积
func (m *RemoteMathUtil) CalculateCircleArea(req float64, resp *float64) error {
	*resp = math.Pi * req * req
	return nil
}

func (m *RemoteMathUtil) Sum(req rpcdemo.Req1, resp *int) error {
	*resp = req.I1 + req.I2 + req.I3
	return nil
}

func (m *RemoteMathUtil) Sum2(req rpcdemo.Req1, resp *rpcdemo.Resp1) error {
	*resp = rpcdemo.Resp1{
		Flag: true,
		Sum:  req.I1 + req.I2 + req.I3,
	}
	return nil
}

func (m *RemoteMathUtil) Add(req int, resp *rpcdemo.Resp1) error {
	*resp = rpcdemo.Resp1{
		Flag: true,
		Sum:  req + req,
	}
	return nil
}
