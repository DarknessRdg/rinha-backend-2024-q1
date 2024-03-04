package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyCents_LowerThan(t *testing.T) {
	suite := []struct {
		name   string
		money  MoneyCents
		value  int
		expect bool
	}{
		{
			name:   "When money is lower than value, Then return false",
			money:  MoneyCents(99),
			value:  100,
			expect: true,
		},
		{
			name: "When money is equal value, Then return false",
			money: MoneyCents(100),
			value: 100,
			expect: false,
		},
		{
			name: "When money is greater tha value, Then return true",
			money: MoneyCents(100),
			value: 99,
			expect: false,
		},
	}
	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			got := test.money.LowerThan(test.value)
			assert.Equal(t, test.expect, got, test.name)
		})
	}
}
