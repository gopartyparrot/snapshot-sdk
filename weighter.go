package sdk

import (
	"math/big"

	"github.com/gagliardetto/solana-go"
)

type Weighter struct {
	ByOwner map[solana.PublicKey]*big.Int //key: owner, value: owner weight
	Sum     *big.Int                      //total weight
}

func NewWeighter() Weighter {
	return Weighter{
		ByOwner: make(map[solana.PublicKey]*big.Int),
		Sum:     big.NewInt(0),
	}
}

func (m *Weighter) Add(account solana.PublicKey, weight *big.Int) {
	if weight.Cmp(big.NewInt(0)) <= 0 {
		return
	}
	if _weight := m.ByOwner[account]; _weight == nil {
		m.ByOwner[account] = weight
	} else {
		m.ByOwner[account] = _weight.Add(_weight, weight)
	}
	m.Sum = m.Sum.Add(m.Sum, weight)
}
