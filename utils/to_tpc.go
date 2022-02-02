package utils

import "math/big"

// ToTpc number of TPC to Wei
func ToTpc(tpc uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(tpc), big.NewInt(1e18))
}
