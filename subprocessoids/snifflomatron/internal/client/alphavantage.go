package client

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Printf("Using Alpha Vantage API key: %s", apiKey)
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s",
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

	if note, ok := raw["Note"]; ok {
		return nil, fmt.Errorf("rate limited: %v", note)
	}

	rawSeries, ok := raw["Time Series (Daily)"]
	if !ok {
		return nil, fmt.Errorf("missing 'Time Series (Daily)' in API response")
	}

	timeSeries, ok := rawSeries.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("'Time Series (Daily)' is not a valid object")
	}

	dates := make([]string, 0, len(timeSeries))
	for date := range timeSeries {
		dates = append(dates, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	result := []TimeSeries{}
	for _, date := range dates[:100] {
		entry := timeSeries[date].(map[string]interface{})
		closeVal := entry["4. close"].(string)
		var close float64
		fmt.Sscanf(closeVal, "%f", &close)
		result = append(result, TimeSeries{Date: date, Close: close})
	}

	return result, nil
}
