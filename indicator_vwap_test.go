package techan

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestVWAPIndicator(t *testing.T) {
	start := 1691884500
	s := [][]float64{
		[]float64{10, 12, 12, 8},
		[]float64{11, 14, 14, 9},
		[]float64{10, 20, 24, 10},
		[]float64{9, 10, 11, 9},
		[]float64{11, 14, 14, 9},
		[]float64{9, 10, 11, 9},
		[]float64{10, 12, 12, 10},
		[]float64{9, 10, 11, 8},
		[]float64{6, 5, 8, 1},
		[]float64{15, 12, 18, 9},
		[]float64{35, 25, 50, 20},
	}
	ts := NewTimeSeries()
	for i, ochl := range s {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(start+i*60), 0), time.Minute))
		candle.OpenPrice = big.NewDecimal(ochl[0])
		candle.ClosePrice = big.NewDecimal(ochl[1])
		candle.MaxPrice = big.NewDecimal(ochl[2])
		candle.MinPrice = big.NewDecimal(ochl[3])
		candle.Volume = big.NewDecimal(float64(10))
		ts.AddCandle(candle)
	}
	loc, _ := time.LoadLocation("UTC")
	vwap := NewVWAPIndicator(ts, loc)
	results := []string{
		"10.6667",
		"11.5000",
		"13.6667",
		"12.7500",
		"12.6667",
		"10.0000",
		"10.6667",
		"10.3333",
		"8.9167",
		"9.7333",
		"13.3889"}
	for i, r := range results {
		//fmt.Println(vwap.Calculate(i).FormattedString(4))
		assert.EqualValues(t, r, vwap.Calculate(i).FormattedString(4))
	}
}
