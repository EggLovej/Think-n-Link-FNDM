# Think'n'Link Funderdome

Welcome to the Think'n'Link Funderdome — a Go-based microservice that fetches stock market data from Alpha Vantage.

This is the first microservice in a planned architecture for a stock viewing and analysis platform.

---

## Features

- Fetch latest 5-day closing prices for a given stock symbol
- Simple test endpoint returning a static number
- Written in idiomatic Go with chi router
- Environment-based API key loading

---

## Endpoints

### GET /stocks/{symbol}

Fetches the last 5 days of daily closing prices for the given stock symbol.

Example request:
curl http://localhost:8080/stocks/AAPL

Example response:
[
  {"date":"2024-05-10","close":182.23},
  {"date":"2024-05-09","close":179.76}
]

---

### GET /number

Returns a static JSON object with the number 62.

Example request:
curl http://localhost:8080/number

Example response:
{"number":62} 

---

## How to Run Locally

1. Clone the repository

git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
cd YOUR_REPO_NAME/data-gatherer

2. Set your environment variables

Create a file called `.env` in the `data-gatherer/` directory with the following contents:

ALPHAVANTAGE_API_KEY=your_api_key_here

Alternatively, export it in your terminal session:

export ALPHAVANTAGE_API_KEY=your_api_key_here

3. Install dependencies

go mod tidy

4. Run the service

go run ./cmd

The server will start at http://localhost:8080

---

## Test It

Using curl:

curl http://localhost:8080/stocks/MSFT
curl http://localhost:8080/number

Or just open http://localhost:8080/number in your browser — you should see:

{"number":62}

---

## Docker Support

Coming soon: Dockerfile and docker-compose for multi-service orchestration.

---

## License

MIT — Use it however you like. We take no responsibility for your financial decisions 🐒📉

---

## Author

@YOUR_USERNAME
