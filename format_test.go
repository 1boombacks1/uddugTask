package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInterval(t *testing.T) {
	type testCase struct {
		transactions []Transaction
		interval     time.Duration

		expected []Transaction
	}

	t.Log("Testing for correct input interval and transactions")
	t.Run("check all interval types", func(t *testing.T) {
		testCases := []testCase{
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: Day,
				expected: []Transaction{
					{Value: 4456, Timestamp: time.Date(2021, 03, 18, 0, 0, 0, 0, time.UTC)},
					{Value: 4231, Timestamp: time.Date(2021, 03, 16, 0, 0, 0, 0, time.UTC)},
				},
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: Hour,
				expected: []Transaction{
					{Value: 5212, Timestamp: time.Date(2021, 03, 16, 12, 0, 0, 0, time.UTC)},
					{Value: 4321, Timestamp: time.Date(2021, 03, 16, 10, 0, 0, 0, time.UTC)},
					{Value: 4567, Timestamp: time.Date(2021, 03, 16, 5, 0, 0, 0, time.UTC)},
					{Value: 4456, Timestamp: time.Date(2021, 03, 18, 0, 0, 0, 0, time.UTC)},
					{Value: 4231, Timestamp: time.Date(2021, 03, 16, 22, 0, 0, 0, time.UTC)},
				},
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: Week,
				expected: []Transaction{
					{4456, time.Date(2021, 03, 15, 0, 0, 0, 0, time.UTC)},
				},
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: Month,
				expected: []Transaction{
					{4456, time.Date(2021, 03, 01, 0, 0, 0, 0, time.UTC)},
				},
			},
		}

		for i, test := range testCases {
			t.Logf("Test #%d", i+1)
			actual, err := FormatTransactionByInterval(test.transactions, test.interval)

			assert.Nil(t, err)
			assert.ElementsMatch(t, actual, test.expected)
		}
	})

	t.Log("Testing for incorrect input interval")
	t.Run("check incorrect interval input", func(t *testing.T) {
		testCases := []testCase{
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: time.Second,
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: 6*Day + 59*time.Minute,
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: Week + time.Hour,
			},
			{
				transactions: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1615932648, 0)},
					{Value: 5212, Timestamp: time.Unix(1615899048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: time.Date(2021, 3, 4, 1, 0, 0, 0, time.UTC).Sub(time.Date(2021, 3, 4, 5, 0, 0, 0, time.UTC)),
			},
		}

		for i, test := range testCases {
			t.Logf("Test #%d", i+1)
			actual, err := FormatTransactionByInterval(test.transactions, test.interval)

			assert.ErrorIs(t, err, NotValidIntervalError)
			assert.Nil(t, actual)
		}
	})
}
