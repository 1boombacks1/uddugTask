package main

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

const (
	Hour  = time.Hour
	Day   = 24 * Hour
	Week  = 7 * Day
	Month = 30 * Day
)

var (
	NotValidIntervalError = errors.New("not valid interval, it should be: hour, day, week, month")
)

func FormatTransactionByInterval(transactions []Transaction, interval time.Duration) ([]Transaction, error) {
	if !isValidInterval(interval) {
		return nil, NotValidIntervalError
	}

	groupedTxs := make(map[int64]Transaction, 1)

	for _, currTx := range transactions {
		var roundedTimestamp time.Time
		if interval == Month {
			roundedTimestamp = getRoundedDateByMonth(currTx.Timestamp)
		} else {
			roundedTimestamp = currTx.Timestamp.UTC().Truncate(interval)
		}

		groupTx, ok := groupedTxs[roundedTimestamp.Unix()]
		if !ok || currTx.Timestamp.After(groupTx.Timestamp) {
			groupedTxs[roundedTimestamp.Unix()] = currTx
		}
	}

	result := make([]Transaction, 0, len(groupedTxs))
	for ts, tx := range groupedTxs {
		formatedTx := Transaction{tx.Value, time.Unix(ts, 0).UTC()}
		result = append(result, formatedTx)
	}

	return result, nil
}

func getRoundedDateByMonth(ts time.Time) time.Time {
	return time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func isValidInterval(interval time.Duration) bool {
	return interval == Hour || interval == Day || interval == Week || interval == Month
}

func main() {
	txs := []Transaction{
		{Value: 4456, Timestamp: time.Date(2021, 03, 18, 0, 10, 48, 0, time.UTC)},
		{Value: 4231, Timestamp: time.Date(2021, 03, 17, 23, 10, 48, 0, time.UTC)},
		{Value: 5212, Timestamp: time.Date(2021, 03, 17, 22, 10, 48, 0, time.UTC)},
		{Value: 4321, Timestamp: time.Date(2021, 03, 16, 10, 10, 48, 0, time.UTC)},
		{Value: 4567, Timestamp: time.Date(2021, 03, 16, 5, 10, 48, 0, time.UTC)},
	}

	formatedTxs, err := FormatTransactionByInterval(txs, Day)
	if err != nil {
		fmt.Print(err)
		return
	}
	for i, tx := range formatedTxs {
		fmt.Printf("%d. Value - %d\n   Date - %v\n", i+1, tx.Value, tx.Timestamp)
	}
}
