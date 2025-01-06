package math

type Math struct {
	A        int
	B        int
	Operator string
}

func NewMath(a int, b int, operator string) *Math {
	return &Math{A: a, B: b, Operator: operator}
}

func (m *Math) Calculate() int {
	switch m.Operator {
	case "sum":
		return m.Sum()
	case "sub":
		return m.Sub()
	case "times":
		return m.Times()
	default:
		return 0
	}
}

func (m *Math) Sum() int {
	return Sum(m.A, m.B)
}

func (m *Math) Sub() int {
	return Sub(m.A, m.B)
}

func (m *Math) Times() int {
	return Times(m.A, m.B)
}
