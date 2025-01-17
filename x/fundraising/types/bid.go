package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (b Bid) GetBidder() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(b.Bidder)
	if err != nil {
		panic(err)
	}
	return addr
}

func (b *Bid) SetMatched(status bool) {
	b.IsMatched = status
}

// ConvertToSellingAmount converts to selling amount depending on the bid coin denom.
// Note that we take as little coins as possible to prevent from overflowing the remaining selling coin.
func (b Bid) ConvertToSellingAmount(denom string) (amount sdk.Int) {
	if b.Coin.Denom == denom {
		return b.Coin.Amount.ToDec().QuoTruncate(b.Price).TruncateInt() // BidAmount / BidPrice
	}
	return b.Coin.Amount
}

// ConvertToPayingAmount converts to paying amount depending on the bid coin denom.
// Note that we take as many coins as possible by ceiling numbers from bidder.
func (b Bid) ConvertToPayingAmount(denom string) (amount sdk.Int) {
	if b.Coin.Denom == denom {
		return b.Coin.Amount
	}
	return b.Coin.Amount.ToDec().Mul(b.Price).Ceil().TruncateInt() // BidAmount * BidPrice
}
