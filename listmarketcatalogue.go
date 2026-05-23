package betfair

import (
	"context"
	"time"
)

// PriceLadderType is the type of price ladder.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#PriceLadderType
type PriceLadderType string

const (
	PLT_CLASSIC    PriceLadderType = "CLASSIC"    // PLT_CLASSIC is a price ladder increment traditionally used for Odds Markets.
	PLT_FINEST     PriceLadderType = "FINEST"     // PLT_FINEST is a price ladder with the finest available increment, traditionally used for Asian Handicap markets.
	PLT_LINE_RANGE PriceLadderType = "LINE_RANGE" // PLT_LINE_RANGE is a price ladder used for LINE markets.
)

// MarketSort controls the order of the results. Will default to RANK if not passed.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#MarketSort
type MarketSort string

const (
	MS_MINIMUM_TRADED    MarketSort = "MINIMUM_TRADED"    // MS_MINIMUM_TRADED is minimum traded volume.
	MS_MAXIMUM_TRADED    MarketSort = "MAXIMUM_TRADED"    // MS_MAXIMUM_TRADED is maximum traded volume.
	MS_MINIMUM_AVAILABLE MarketSort = "MINIMUM_AVAILABLE" // MS_MINIMUM_AVAILABLE is minimum available to match.
	MS_MAXIMUM_AVAILABLE MarketSort = "MAXIMUM_AVAILABLE" // MS_MAXIMUM_AVAILABLE is maximum available to match.
	MS_FIRST_TO_START    MarketSort = "FIRST_TO_START"    // MS_FIRST_TO_START is the closest markets based on their expected start time.
	MS_LAST_TO_START     MarketSort = "LAST_TO_START"     // MS_LAST_TO_START is the most distant markets based on their expected start time.
	MS_RANK              MarketSort = "RANK"              // MS_RANK is the order of the results. Will default to RANK if not passed.
)

// MarketProjection controls the type and amount of data returned about the market.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#MarketProjection
type MarketProjection string

const (
	MP_COMPETITION        MarketProjection = "COMPETITION"        // MP_COMPETITION - If not selected then the competition will not be returned with marketCatalogue.
	MP_EVENT              MarketProjection = "EVENT"              // MP_EVENT - If not selected then the event will not be returned with marketCatalogue.
	MP_EVENT_TYPE         MarketProjection = "EVENT_TYPE"         // MP_EVENT_TYPE - If not selected then the eventType will not be returned with marketCatalogue.
	MP_MARKET_START_TIME  MarketProjection = "MARKET_START_TIME"  // MP_MARKET_START_TIME - If not selected then the start time will not be returned with marketCatalogue.
	MP_MARKET_DESCRIPTION MarketProjection = "MARKET_DESCRIPTION" // MP_MARKET_DESCRIPTION - If not selected then the description will not be returned with marketCatalogue.
	MP_RUNNER_DESCRIPTION MarketProjection = "RUNNER_DESCRIPTION" // MP_RUNNER_DESCRIPTION - If not selected then the runners will not be returned with marketCatalogue.
	MP_RUNNER_METADATA    MarketProjection = "RUNNER_METADATA"    // MP_RUNNER_METADATA - If not selected then the runner metadata will not be returned with marketCatalogue. If selected then RUNNER_DESCRIPTION will also be returned regardless of whether it is included as a market projection.
)

type (
	// ListMarketCatalogueParams are the request parameters for listMarketCatalogue.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687517/listMarketCatalogue
	ListMarketCatalogueParams struct {
		Filter            MarketFilter       `json:"filter"`                     // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		MarketProjections []MarketProjection `json:"marketProjection,omitempty"` // MarketProjections controls the type and amount of data returned about the market.
		Sort              MarketSort         `json:"sort,omitempty"`             // Sort controls the order of the results. Will default to RANK if not passed.
		MaxResults        int                `json:"maxResults"`                 // MaxResults is the limit on the total number of results returned, must be greater than 0 and less than or equal to 1000.
		Locale            string             `json:"locale,omitempty"`           // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListMarketCatalogueResult holds information about a market (MarketCatalogue).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketCatalogue
	ListMarketCatalogueResult struct {
		MarketID          string            `json:"marketId"`                  // MarketID is the unique identifier for the market. MarketId's are prefixed with '1.'.
		MarketName        string            `json:"marketName"`                // MarketName is the name of the market.
		MarketStartTime   time.Time         `json:"marketStartTime,omitempty"` // MarketStartTime is the time this market starts at, only returned when the MARKET_START_TIME enum is passed in the marketProjections.
		MarketDescription MarketDescription `json:"description,omitempty"`     // MarketDescription contains details about the market.
		TotalMatched      float64           `json:"totalMatched,omitempty"`    // TotalMatched is the total amount of money matched on the market. Please note: The returned value is cached. For the live total matched value please use listMarketBook.
		Runners           []RunnerCatalog   `json:"runners,omitempty"`         // Runners are the runners (selections) contained in the market.
		EventType         EventType         `json:"eventType,omitempty"`       // EventType is the Event Type the market is contained within.
		Competition       Competition       `json:"competition,omitempty"`     // Competition is the competition the market is contained within. Usually only applies to Football competitions.
		Event             Event             `json:"event,omitempty"`           // Event is the event the market is contained within.
	}

	// MarketDescription contains details about the market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketDescription
	MarketDescription struct {
		PersistenceEnabled     bool                   `json:"persistenceEnabled"`               // PersistenceEnabled indicates whether the market supports persistence.
		BspMarket              bool                   `json:"bspMarket"`                        // BspMarket indicates whether the market is an BSP betting market.
		MarketTime             time.Time              `json:"marketTime"`                       // MarketTime is the time this market starts at.
		SuspendTime            time.Time              `json:"suspendTime"`                      // SuspendTime is the time this market is suspended at (or off time if turnInPlayEnabled is true).
		SettleTime             time.Time              `json:"settleTime,omitempty"`             // SettleTime is the settled time (may not be present).
		BettingType            MarketBettingType      `json:"bettingType"`                      // BettingType is the betting type of the market.
		TurnInPlayEnabled      bool                   `json:"turnInPlayEnabled"`                // TurnInPlayEnabled indicates whether the market will turn in play.
		MarketType             string                 `json:"marketType"`                       // MarketType is the market type e.g. MATCH_ODDS.
		Regulator              string                 `json:"regulator"`                        // Regulator is the regulator for the market.
		MarketBaseRate         float64                `json:"marketBaseRate"`                   // MarketBaseRate is the commission rate applicable to the market.
		DiscountAllowed        bool                   `json:"discountAllowed"`                  // DiscountAllowed indicates whether the user's discount rate is applied.
		Wallet                 string                 `json:"wallet,omitempty"`                 // Wallet is the wallet from which the prices should be taken.
		Rules                  string                 `json:"rules,omitempty"`                  // Rules are the rules for the market.
		RulesHasDate           bool                   `json:"rulesHasDate,omitempty"`           // RulesHasDate indicates whether the rules include a date.
		EachWayDivisor         float64                `json:"eachWayDivisor,omitempty"`         // EachWayDivisor is the divisor for each way markets.
		Clarifications         string                 `json:"clarifications,omitempty"`         // Clarifications are any clarifications for the market.
		LineRangeInfo          MarketLineRangeInfo    `json:"lineRangeInfo,omitempty"`          // LineRangeInfo is the range of prices for a line market.
		RaceType               string                 `json:"raceType,omitempty"`               // RaceType is the race type for horse racing markets.
		PriceLadderDescription PriceLadderDescription `json:"priceLadderDescription,omitempty"` // PriceLadderDescription is the description of the price ladder for this market.
	}

	// RunnerCatalog holds information about the Runners (selections) in a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#RunnerCatalog
	RunnerCatalog struct {
		SelectionID  int64             `json:"selectionId"`        // SelectionID is the unique id for the selection. The same selectionId and runnerName pairs are used accross all Betfair markets which contain them.
		RunnerName   string            `json:"runnerName"`         // RunnerName is the name of the runner.
		Handicap     float64           `json:"handicap"`           // Handicap applies to markets with MarketBettingType ASIAN_HANDICAP_SINGLE_LINE & ASIAN_HANDICAP_DOUBLE_LINE only otherwise '0'.
		SortPriority int               `json:"sortPriority"`       // SortPriority indicates the order in which the runners are displayed on the Betfair Exchange website.
		Metadata     map[string]string `json:"metadata,omitempty"` // Metadata is metadata associated with the runner. For a description of this data for Horse Racing, please see Runner Metadata Description.
	}

	// MarketLineRangeInfo describes the range of prices for a line market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketLineRangeInfo
	MarketLineRangeInfo struct {
		MaxUnitValue float64 `json:"maxUnitValue"` // MaxUnitValue is the maximum value for the line market.
		MinUnitValue float64 `json:"minUnitValue"` // MinUnitValue is the minimum value for the line market.
		Interval     float64 `json:"interval"`     // Interval is the increments in which the market shall be displayed.
		MarketUnit   string  `json:"marketUnit"`   // MarketUnit is the unit for the market (e.g. runs, goals).
	}

	// PriceLadderDescription describes the price ladder for a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PriceLadderDescription
	PriceLadderDescription struct {
		Type PriceLadderType `json:"type,omitempty"` // Type is the type of price ladder.
	}
)

// ListMarketCatalogue returns a list of information about published (ACTIVE/SUSPENDED) markets that does not change (or changes very rarely). You use listMarketCatalogue to retrieve the name of the market, the names of selections and other information about markets. Market Data Request Limits apply to requests made to listMarketCatalogue. Please note: listMarketCatalogue does not return markets that are CLOSED.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687517/listMarketCatalogue
func (client *Client) ListMarketCatalogue(ctx context.Context, params ListMarketCatalogueParams) ([]ListMarketCatalogueResult, error) {
	json := []ListMarketCatalogueResult{}

	return json, client.GetSports(ctx, "listMarketCatalogue", params, &json)
}
