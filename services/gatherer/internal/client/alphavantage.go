package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
)

type TimeSeries struct {
	Date  string  `json:"date"`
	Close float64 `json:"close"`
}

func FetchDailyTimeSeries(symbol string) ([]TimeSeries, error) {
	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&apikey=%s",
		symbol, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	timeSeries := raw["Time Series (Daily)"].(map[string]interface{})
	dates := make([]string, 0, len(timeSeries))
	for date := range timeSeries {
		dates = append(dates, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	result := []TimeSeries{}
	for _, date := range dates[:5] {
		entry := timeSeries[date].(map[string]interface{})
		closeVal := entry["4. close"].(string)
		var close float64
		fmt.Sscanf(closeVal, "%f", &close)
		result = append(result, TimeSeries{Date: date, Close: close})
	}

	return result, nil
}
