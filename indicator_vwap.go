package techan

import (
	"time"

	"github.com/sdcoffey/big"
)

type vwapIndicator struct {
	series   *TimeSeries
	timezone *time.Location
}

func NewVWAPIndicator(ts *TimeSeries, tz *time.Location) Indicator {
	return vwapIndicator{
		series:   ts,
		timezone: tz,
	}
}

func (vwap vwapIndicator) Calculate(index int) big.Decimal {
	ctp, cvp := big.ZERO, big.ZERO
	vi, tpi := NewVolumeIndicator(vwap.series), NewTypicalPriceIndicator(vwap.series)
	start := vwap.series.Candles[index].Period.Start.Truncate(time.Hour * 24)
	for vwap.series.Candles[index].Period.Start.After(start) ||
		vwap.series.Candles[index].Period.Start.Equal(start) {
		ctp = ctp.Add(tpi.Calculate(index).Mul(vwap.series.Candles[index].Volume))
		cvp = cvp.Add(vi.Calculate(index))
		index -= 1
		if index < 0 {
			break
		}
	}
	if cvp.EQ(big.ZERO) {
		return big.ZERO
	}
	return ctp.Div(cvp)
}
