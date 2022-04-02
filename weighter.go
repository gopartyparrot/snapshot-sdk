package sdk

import (
	"github.com/gagliardetto/solana-go"
	"github.com/shopspring/decimal"
)

type Weighter struct {
	ByOwner map[solana.PublicKey]decimal.Decimal //key: owner, value: owner weight
	Sum     decimal.Decimal                      //total weight
}

func NewWeighter() Weighter {
	return Weighter{
		ByOwner: make(map[solana.PublicKey]decimal.Decimal),
		Sum:     decimal.New(0, 0),
	}
}

func (m *Weighter) Add(account solana.PublicKey, weight decimal.Decimal) {
	if weight.Cmp(decimal.New(0, 0)) <= 0 {
		return
	}
	if w, ok := m.ByOwner[account]; !ok {
		m.ByOwner[account] = weight
	} else {
		m.ByOwner[account] = w.Add(weight)
	}
	m.Sum = m.Sum.Add(weight)
}
