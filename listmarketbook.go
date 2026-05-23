package betfair

import (
	"context"
	"time"
)

// RollupModel specifies how rollup should be performed.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#OffersOverrides
type RollupModel string

const (
	RM_STAKE  = "STAKE"
	RM_PAYOUT = "PAYOUT"
	RM_NONE   = "NONE"
)

// RunnerStatus is the status of the selection (i.e., ACTIVE, REMOVED, WINNER, PLACED, LOSER, HIDDEN).
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#RunnerStatus
type RunnerStatus string

const (
	RS_ACTIVE         RunnerStatus = "ACTIVE"         // RS_ACTIVE is ACTIVE.
	RS_WINNER         RunnerStatus = "WINNER"         // RS_WINNER is WINNER.
	RS_LOSER          RunnerStatus = "LOSER"          // RS_LOSER is LOSER.
	RS_PLACED         RunnerStatus = "PLACED"         // RS_PLACED indicates the runner was placed, applies to EACH_WAY marketTypes only.
	RS_REMOVED_VACANT RunnerStatus = "REMOVED_VACANT" // RS_REMOVED_VACANT applies to Greyhounds. Greyhound markets always return a fixed number of runners (traps). If a dog has been removed, the trap is shown as vacant.
	RS_REMOVED        RunnerStatus = "REMOVED"        // RS_REMOVED is REMOVED.
	RS_HIDDEN         RunnerStatus = "HIDDEN"         // RS_HIDDEN indicates the selection is hidden from the market. This occurs in Horse Racing markets were runners is hidden when it is doesn't hold an official entry following an entry stage. This could be because the horse was never entered or because they have been scratched from a race at a declaration stage. All matched customer bet prices are set to 1.0 even if there are later supplementary stages. Should it appear likely that a specific runner may actually be supplemented into the race this runner will be reinstated with all matched customer bets set back to the original prices.
)

// MarketStatus is the status of the market, for example OPEN, SUSPENDED, CLOSED (settled), etc.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#MarketStatus
type MarketStatus string

const (
	MS_INACTIVE  MarketStatus = "INACTIVE"  // MS_INACTIVE indicates the market has been created but isn't yet available.
	MS_OPEN      MarketStatus = "OPEN"      // MS_OPEN indicates the market is open for betting.
	MS_SUSPENDED MarketStatus = "SUSPENDED" // MS_SUSPENDED indicates the market is suspended and not available for betting.
	MS_CLOSED    MarketStatus = "CLOSED"    // MS_CLOSED indicates the market has been settled and is no longer available for betting.
)

type (
	// ListMarketBookParams are the request parameters for listMarketBook.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687510/listMarketBook
	ListMarketBookParams struct {
		MarketIDs                     []string        `json:"marketIds"`                               // MarketIDs are one or more market ids. The number of markets returned depends on the amount of data you request via the price projection.
		PriceProjection               PriceProjection `json:"priceProjection,omitempty"`               // PriceProjection is the projection of price data you want to receive in the response.
		OrderProjection               OrderProjection `json:"orderProjection,omitempty"`               // OrderProjection is the orders you want to receive in the response.
		MatchProjection               MatchProjection `json:"matchProjection,omitempty"`               // MatchProjection specifies the representation of matches if you ask for orders.
		IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`        // IncludeOverallPosition returns matches for each selection if you ask for orders. Defaults to true if unspecified.
		PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"` // PartitionMatchedByStrategyRef returns the breakdown of matches by strategy for each selection if you ask for orders. Defaults to false if unspecified.
		CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`          // CustomerStrategyRefs restricts the results to orders matching any of the specified set of customer defined strategies.
		CurrencyCode                  string          `json:"currencyCode,omitempty"`                  // CurrencyCode is a Betfair standard currency code. If not specified, the default currency code is used.
		Locale                        string          `json:"locale,omitempty"`                        // Locale is the language used for the response. If not specified, the default is returned.
		MatchedSince                  time.Time       `json:"matchedSince,omitempty"`                  // MatchedSince restricts the results to orders that have at least one fragment matched since the specified date.
		BetIDs                        []string        `json:"betIds,omitempty"`                        // BetIDs restricts the results to orders with the specified bet IDs. A maximum of 250 betId's can be provided at a time.
	}
)

// ListMarketBookResult is the dynamic data in a market (MarketBook).
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketBook
type ListMarketBookResult = MarketBook

// ListMarketBook returns a list of dynamic data about markets. Dynamic data includes prices, the status of the market, the status of selections, the traded volume, and the status of any orders you have placed in the market. Please note: Separate requests should be made for OPEN & CLOSED markets. Requests that include both OPEN & CLOSED markets will only return those markets that are OPEN.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687510/listMarketBook
func (client *Client) ListMarketBook(ctx context.Context, params ListMarketBookParams) ([]ListMarketBookResult, error) {
	json := []ListMarketBookResult{}

	return json, client.GetSports(ctx, "listMarketBook", params, &json)
}
