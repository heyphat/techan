package techan

import (
	"strings"
	"time"

	"github.com/sdcoffey/big"
)

// OrderSide is a simple enumeration representing the side of an Order (buy or sell)
type OrderSide int

// BUY and SELL enumerations
const (
	BUY OrderSide = iota
	SELL
)

// Order represents a trade execution (buy or sell) with associated metadata.
type Order struct {
	Side          OrderSide
	Security      string
	Price         big.Decimal
	Amount        big.Decimal
	ExecutionTime time.Time
}

// String returns a `BUY` or `SELL` string.
func (os OrderSide) String() string {
	if os == BUY {
		return "BUY"
	}
	return "SELL"
}

// OrderSideFromString return tachan OrderSide from generic string.
func OrderSideFromString(s string) OrderSide {
	if strings.ToUpper(s) == "BUY" {
		return BUY
	}
	return SELL
}
