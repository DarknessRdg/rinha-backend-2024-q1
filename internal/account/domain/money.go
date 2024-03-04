package domain

type MoneyCents int

func (m MoneyCents) LowerThan(value int) bool {
	return int(m) < value
}
