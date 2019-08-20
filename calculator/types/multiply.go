package types

import "github.com/looking-4-justice/calculator/interfaces"

type Multiply BinaryOperator

func (bo Multiply) Evaluate() float64 {
	return bo.Left.Evaluate() * bo.Right.Evaluate()
}

func NewMultiply(left interfaces.Expression, right interfaces.Expression) Multiply {
	return Multiply{Left: left, Right: right, Operator: '*'}
}