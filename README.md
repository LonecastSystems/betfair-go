# betfair-go

Go wrapper library for interacting with the Betfair exchange, including both the standard and the streaming API.

Currently only supports non-interactive (certificate) login.

Requires Go 1.22 or later.

## Documentation

- [pkg.go.dev](https://pkg.go.dev/github.com/LonecastSystems/betfair-go) — API reference (types include inline links to Betfair docs)
- [Betfair Exchange API](https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/) — upstream API documentation

## Prerequisites

Before using this library you need:

1. **A Betfair exchange account** with API access enabled.
2. **An application key** from the [Betfair Developer Program](https://developer.betfair.com/).
   - **Delayed app key** — free; market data is delayed. Fine for development.
   - **Live app key** — requires approval; needed for real-time prices and production trading.
3. **An SSL client certificate** for non-interactive login. Generate one in your Betfair account settings and download the `.crt` / `.key` pair (or use in-memory bytes via `GetTLSConfigFromBytes`).

See [Non-Interactive (bot) login](https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login) for certificate setup details.

## Installation

Pin a release:

```sh
go get github.com/LonecastSystems/betfair-go@v0.1.0
```

Or track the latest on `main`:

```sh
go get github.com/LonecastSystems/betfair-go@latest
```

## Configuration

`ClientConfig` fields worth knowing:

| Field | Description |
|-------|-------------|
| `Tls` | TLS config from `GetTLSConfig` or `GetTLSConfigFromBytes` (required for login) |
| `ApplicationKey` | Your Betfair app key (sent as `X-Application`) |
| `Jurisdiction` | API domain suffix. Defaults to `JD_GLOBAL` (`.com`). Also: `JD_AUSTRALIA`, `JD_ITALY`, `JD_SPAIN`, `JD_ROMANIA`, `JD_SWEDEN` |
| `Rest` | `false` (default) uses JSON-RPC; `true` uses the REST API instead |

## Usage

### Standard client

#### Login and session

```go
package main

import (
	"context"
	"log"

	"github.com/LonecastSystems/betfair-go"
)

func main() {
	tlsConfig, err := betfair.GetTLSConfig("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &betfair.ClientConfig{
		Tls:            tlsConfig,
		ApplicationKey: "YourAppKey",
		Jurisdiction:   betfair.JD_GLOBAL, // optional; default is .com
		Rest:           false,             // JSON-RPC (default)
	}
	client := betfair.NewClient(config)

	ctx := context.Background()
	resp, err := client.Login(ctx, "username", "password")
	if err != nil {
		log.Fatal(err)
	}
	if resp.LoginStatus != betfair.LS_SUCCESS {
		log.Fatalf("login failed: %s", resp.LoginStatus)
	}

	// Extend the session periodically while your app runs.
	if _, err := client.KeepAlive(ctx); err != nil {
		log.Fatal(err)
	}
}
```

#### Certificate from bytes

```go
tlsConfig, err := betfair.GetTLSConfigFromBytes(certPEM, keyPEM)
```

#### Market data

Typical flow: discover markets with `ListMarketCatalogue`, then fetch prices with `ListMarketBook`.

```go
catalogue, err := client.ListMarketCatalogue(ctx, betfair.ListMarketCatalogueParams{
	Filter: betfair.MarketFilter{
		EventTypeIDs: []string{"7"}, // horse racing
	},
	MaxResults: 10,
})
if err != nil {
	log.Fatal(err)
}

marketIDs := make([]string, len(catalogue))
for i, m := range catalogue {
	marketIDs[i] = m.MarketID
}

books, err := client.ListMarketBook(ctx, betfair.ListMarketBookParams{
	MarketIDs: marketIDs,
	PriceProjection: betfair.PriceProjection{
		PriceData: []betfair.PriceData{betfair.PD_EX_BEST_OFFERS},
	},
})
if err != nil {
	log.Fatal(err)
}
log.Println(len(books), "market books")
```

### Streaming client

Requires an authenticated standard client (certificate login and session token), unless using direct authentication below.

#### Order subscription

```go
streamClient, err := client.GetStream(&betfair.StreamingClientConfig{
	HeartbeatMs: 5000,
})
if err != nil {
	log.Fatal(err)
}
defer streamClient.Close()

orderChanges, err := streamClient.SubscribeToOrders(betfair.OrderFilter{
	IncludeOverallPosition: true,
})
if err != nil {
	log.Fatal(err)
}

for msg := range orderChanges {
	// Process order deltas; heartbeats arrive when there is no data.
	log.Printf("publish time: %d, markets: %d", msg.PublishTime, len(msg.OrderAccountChange))
}
```

#### Market subscription

```go
marketChanges, err := streamClient.SubscribeToMarkets(
	betfair.StreamMarketFilter{MarketIDs: []string{"1.234567890"}},
	betfair.MarketDataFilter{
		Fields:       []betfair.MarketDataFilterField{betfair.MDFF_EX_BEST_OFFERS},
		LadderLevels: 3,
	},
)
if err != nil {
	log.Fatal(err)
}

for msg := range marketChanges {
	log.Printf("market changes: %d", len(msg.MarketChanges))
}
```

#### Direct authentication

Alternatively, connect with `StreamingClient.Authenticate` using TLS config, app key, and session token directly — without going through `Client.GetStream`.

#### Integration endpoint

Set `StreamingClientConfig.PreProduction` to `true` to connect to the integration stream endpoint instead of production.

## API overview

All RPC methods take a `context.Context` as the first argument.

### Session

| Method | Description |
|--------|-------------|
| `Login` | Certificate-based non-interactive login |
| `Logout` | End the session |
| `KeepAlive` | Extend session lifetime |
| `Resume` | Resume an existing session token |
| `GetStream` | Open an authenticated Exchange Stream connection |

### Account

| Method | Description |
|--------|-------------|
| `GetAccountDetails` | Account profile |
| `GetAccountFunds` | Available balance, exposure, commission |
| `GetAccountStatement` | Account statement / transaction history |
| `GetDeveloperAppKeys` | Application keys for the developer account |
| `ListCurrencyRates` | Currency conversion rates |

### Betting — markets

| Method | Description |
|--------|-------------|
| `ListMarketCatalogue` | Static market metadata |
| `ListMarketBook` | Dynamic prices and orders |
| `ListRunnerBook` | Runner-level book for a market |
| `ListEvents` | Events matching a filter |
| `ListEventTypes` | Event types (e.g. Football, Horse Racing) |
| `ListCompetitions` | Competitions matching a filter |
| `ListCountries` | Countries with markets |
| `ListVenues` | Venues (racing) |
| `ListMarketTypes` | Market type codes |
| `ListTimeRanges` | Granularity buckets for market volume |
| `ListMarketProfitAndLoss` | P&L by market / runner |

### Betting — orders

| Method | Description |
|--------|-------------|
| `PlaceOrders` | Place new bets |
| `CancelOrders` | Cancel existing bets |
| `ReplaceOrders` | Cancel and replace in one instruction |
| `UpdateOrders` | Update persistence or price |
| `ListCurrentOrders` | Open / recently matched orders |
| `ListClearedOrders` | Settled order history |

### Other

| Method | Description |
|--------|-------------|
| `ListRaceDetails` | Race status (Scores API) |
| `Heartbeat` | Concurrency / timing probe |

### Exchange Stream

| Method | Description |
|--------|-------------|
| `SubscribeToMarkets` | Live market price stream |
| `SubscribeToOrders` | Live order stream |
| `Authenticate` | Connect and authenticate directly |

## License

Apache 2.0 — see [LICENSE](LICENSE).
