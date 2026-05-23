package betfair

import (
	"math/rand/v2"
	"time"
)

// MarketDataFilterField selects Exchange Stream API market data fields in a MarketDataFilter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketDataFilter
type MarketDataFilterField string

const (
	MDFF_EX_BEST_OFFERS_DISP MarketDataFilterField = "EX_BEST_OFFERS_DISP" // Best prices including Virtual Bets - depth is controlled by ladderLevels (1 to 10).
	MDFF_EX_BEST_OFFERS      MarketDataFilterField = "EX_BEST_OFFERS"      // Best prices not including Virtual Bets - depth is controlled by ladderLevels (1 to 10).
	MDFF_EX_ALL_OFFERS       MarketDataFilterField = "EX_ALL_OFFERS"       // Full available to BACK/LAY ladder.
	MDFF_EX_TRADED           MarketDataFilterField = "EX_TRADED"           // Full traded ladder. This is the amount traded at any price on any selection in the market.
	MDFF_EX_TRADED_VOL       MarketDataFilterField = "EX_TRADED_VOL"       // Market and runner level traded volume.
	MDFF_EX_LTP              MarketDataFilterField = "EX_LTP"              // The "Last Price Matched" on a selection.
	MDFF_EX_MARKET_DEF       MarketDataFilterField = "EX_MARKET_DEF"       // Send market definitions. To receive updates to any of the MarketDefinition fields.
	MDFF_SP_TRADED           MarketDataFilterField = "SP_TRADED"           // Starting price ladder.
	MDFF_SP_PROJECTED        MarketDataFilterField = "SP_PROJECTED"        // Starting price projection prices. To receive any update to the Betfair SP Near and Far price.
)

// RaceType restricts markets by race type in MarketFilter.raceTypes.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketFilter
type RaceType string

const (
	RC_Harness  RaceType = "Harness"  // Harness racing (Aus/NZ harness racing events).
	RC_Flat     RaceType = "Flat"     // Flat racing (standard ANZ thoroughbred races).
	RC_Hurdle   RaceType = "Hurdle"   // Hurdle racing (Aus/NZ hurdle races).
	RC_Chase    RaceType = "Chase"    // Chase racing.
	RC_Bumper   RaceType = "Bumper"   // Bumper racing.
	RC_NHFlat   RaceType = "NH Flat"  // National Hunt flat racing.
	RC_Steeple  RaceType = "Steeple"  // Steeple chase racing (Aus/NZ steeple chase races).
	RC_NO_VALUE RaceType = "NO_VALUE" // Used when no valid race type has been mapped.
)

// BettingType restricts Exchange Stream API markets by betting type.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketFilter
type BettingType string

const (
	BT_ODDS                       BettingType = "ODDS"                       // Odds Market - Any market that doesn't fit any any of the below categories.
	BT_LINE                       BettingType = "LINE"                       // Line Market - LINE markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps. Customers either Buy a line (LAY bet, winning if outcome is greater than the taken line (price)) or Sell a line (BACK bet, winning if outcome is less than the taken line (price)). If settled outcome equals the taken line, stake is returned.
	BT_RANGE                      BettingType = "RANGE"                      // Range Market - Now Deprecated.
	BT_ASIAN_HANDICAP_DOUBLE_LINE BettingType = "ASIAN_HANDICAP_DOUBLE_LINE" // Asian Handicap Market - A traditional Asian handicap market. Can be identified by marketType ASIAN_HANDICAP.
	BT_ASIAN_HANDICAP_SINGLE_LINE BettingType = "ASIAN_HANDICAP_SINGLE_LINE" // Asian Single Line Market - A market in which there can be 0 or multiple winners. e,.g marketType TOTAL_GOALS.
	BT_FIXED_ODDS                 BettingType = "FIXED_ODDS"                 // Sportsbook Odds Market. This type is deprecated and will be removed in future releases, when Sportsbook markets will be represented as ODDS market but with a different product type.
)

type (
	// MarketSubscriptionMessage subscribes to market price changes on the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketSubscriptionMessage
	MarketSubscriptionMessage struct {
		ID                  int                `json:"id"`                   // Client generated unique id to link request with response (like json rpc).
		Op                  string             `json:"op"`                   // The operation type.
		SegmentationEnabled bool               `json:"segmentationEnabled"`  // Segmentation Enabled - allow the server to send large sets of data in segments, instead of a single block.
		ConflateMs          int                `json:"conflateMs"`           // Conflate Milliseconds - the conflation rate (looped back on initial image after validation: bounds are 0 to 120000).
		HeartbeatMs         int                `json:"heartbeatMs"`          // Heartbeat Milliseconds - the heartbeat rate (looped back on initial image after validation: bounds are 500 to 5000).
		InitialClk          string             `json:"initialClk,omitempty"` // Token value (received in initial MarketChangeMessage) that should be passed to resume a subscription.
		Clk                 string             `json:"clk,omitempty"`        // Token value delta (received in MarketChangeMessage) that should be passed to resume a subscription.
		MarketFilter        StreamMarketFilter `json:"marketFilter"`         // Horizontal filter for market subscription.
		MarketDataFilter    MarketDataFilter   `json:"marketDataFilter"`     // Vertical filter for market data fields.
	}

	// StreamMarketFilter restricts which markets are included in a market subscription.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketFilter
	StreamMarketFilter struct {
		MarketIDs         []string      `json:"marketIds,omitempty"`         // Restrict to specified market ids.
		BspMarket         bool          `json:"bspMarket,omitempty"`         // Restrict to BSP markets.
		BettingTypes      []BettingType `json:"bettingTypes,omitempty"`      // Restrict to specified betting types.
		EventTypeIDs      []string      `json:"eventTypeIds,omitempty"`      // Restrict to specified event type ids.
		EventIDs          []string      `json:"eventIds,omitempty"`          // Restrict to specified event ids.
		TurnInPlayEnabled bool          `json:"turnInPlayEnabled,omitempty"` // Restrict to markets that will turn in-play.
		MarketTypes       []string      `json:"marketTypes,omitempty"`       // Restrict to specified market types.
		Venues            []string      `json:"venues,omitempty"`            // Restrict to specified venues.
		CountryCodes      []string      `json:"countryCodes,omitempty"`      // Restrict to specified country codes.
		RaceTypes         []RaceType    `json:"raceTypes,omitempty"`         // Restrict to specified race types.
	}

	// MarketDataFilter selects which market data fields are returned in a market subscription.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketDataFilter
	MarketDataFilter struct {
		LadderLevels int                     `json:"ladderLevels,omitempty"` // Number of price ladder levels to return (1 to 10).
		Fields       []MarketDataFilterField `json:"fields,omitempty"`       // Market data fields to include in the subscription.
	}

	// MarketChangeMessage contains market price changes from the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketChangeMessage
	MarketChangeMessage struct {
		ID            int            `json:"id"`                    // Client generated unique id to link request with response (like json rpc).
		Op            string         `json:"op"`                    // The operation type.
		ChangeType    ChangeType     `json:"ct,omitempty"`          // Change Type - set to indicate the type of change - if null this is a delta.
		SegmentType   SegmentType    `json:"segmentType,omitempty"` // Segment Type - if the change is split into multiple segments, this denotes the beginning and end of a change.
		ConflateMs    int            `json:"conflateMs"`            // Conflate Milliseconds - the conflation rate (may differ from that requested if subscription is delayed).
		Status        string         `json:"status"`                // Stream status: set to null if the exchange stream data is up to date and 503 if the downstream services are experiencing latencies.
		HeartbeatMs   int            `json:"heartbeatMs"`           // Heartbeat Milliseconds - the heartbeat rate (may differ from requested: bounds are 500 to 30000).
		PublishTime   int64          `json:"pt"`                    // Publish Time (in millis since epoch) that the changes were generated.
		InitialClk    string         `json:"initialClk"`            // Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect).
		Clk           string         `json:"clk"`                   // Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect).
		MarketChanges []MarketChange `json:"mc"`                    // MarketChanges - the modifications to markets (will be null on a heartbeat).
	}

	// MarketChange contains changes to a single market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketChange
	MarketChange struct {
		MarketID         string           `json:"id,omitempty"`               // Market Id - the id of the market.
		Image            bool             `json:"img,omitempty"`              // Image - replace existing prices/data with the data supplied: it is not a delta (or null if delta).
		TotalVolume      float64          `json:"tv,omitempty"`               // The total amount matched across the market. This value is truncated at 2dp (or null if un-changed).
		MarketDefinition MarketDefinition `json:"marketDefinition,omitempty"` // Full Market Definition if changed.
		RunnerChanges    []RunnerChange   `json:"rc"`                         // Runner Changes - a list of changes to runners (or null if un-changed).
	}

	// MarketDefinition is the full market definition from the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketDefinition
	MarketDefinition struct {
		MarketID              string                `json:"id"`                          // Market Id - the id of the market.
		Venue                 string                `json:"venue"`                       // The venue - applies to horse racing and greyhound markets only.
		RaceType              RaceType              `json:"raceType"`                    // Harness, Flat, Hurdle, Chase, Bumper, NH Flat, Steeple (AUS/NZ races), and NO_VALUE (when no valid race type has been mapped).
		SettledTime           time.Time             `json:"settledTime"`                 // Market settled time.
		TimeZone              string                `json:"timeZone"`                    // This is the timezone in which the event is taking place.
		EachWayDivisor        float64               `json:"eachWayDivisor"`              // The divisor is returned for the marketType EACH_WAY only and refers to the fraction of the win odds at which the place portion of an each way bet is settled.
		BspMarket             bool                  `json:"bspMarket"`                   // If 'true' the market supports Betfair SP betting.
		TurnInPlayEnabled     bool                  `json:"turnInPlayEnabled"`           // If 'true' the market is set to turn in-play.
		PriceLadderDefinition PriceLadderDefinition `json:"priceLadderDefinition"`       // Definition of the price ladder type "CLASSIC", "FINEST", "LINE_RANGE".
		KeyLineDefinition     KeyLineDescription    `json:"keyLineDefinition,omitempty"` // Definition of a markets key line selection (for valid markets).
		PersistenceEnabled    bool                  `json:"persistenceEnabled"`          // If 'true' the market supports 'Keep' bets if the market is to be turned in-play.
		MarketBaseRate        float64               `json:"marketBaseRate"`              // The commission rate applicable to the market.
		EventID               string                `json:"eventId"`                     // The unique id for the event.
		EventTypeID           string                `json:"eventTypeId"`                 // The unique eventTypeId that the event belongs to.
		NumberOfWinners       int                   `json:"numberOfWinners"`             // The number of winners on a market.
		CountryCode           string                `json:"countryCode"`                 // The events ISO 3166-2 country code.
		LineMaxUnit           float64               `json:"lineMaxUnit"`                 // For Handicap and Line markets, the maximum value for the outcome, in market units for this market (eg 100 runs).
		BettingType           string                `json:"bettingType"`                 // The market betting type i.e. ODDS, ASIAN_HANDICAP_DOUBLE_LINE, etc.
		MarketType            string                `json:"marketType"`                  // Market base type.
		MarketTime            string                `json:"marketTime"`                  // The market start time.
		SuspendTime           string                `json:"suspendTime"`                 // The market suspend time.
		BspReconciled         bool                  `json:"bspReconciled"`               // True if the market starting price has been reconciled.
		Complete              bool                  `json:"complete"`                    // If false, runners may be added to the market.
		InPlay                bool                  `json:"inPlay"`                      // True if the market is currently in play.
		CrossMatching         bool                  `json:"crossMatching"`               // True if cross-matching is enabled for this market.
		RunnersVoidable       bool                  `json:"runnersVoidable"`             // True if runners in the market can be voided.
		NumberOfActiveRunners int                   `json:"numberOfActiveRunners"`       // The number of runners that are currently active.
		LineMinUnit           float64               `json:"lineMinUnit"`                 // For Handicap and Line markets, the minimum value for the outcome, in market units for this market (eg 0 runs).
		BetDelay              int                   `json:"betDelay"`                    // The number of seconds an order is held until it is submitted into the market.
		Status                string                `json:"status"`                      // The status of the market, for example, OPEN, SUSPENDED, CLOSED (settled), etc.
		Regulators            []string              `json:"regulators"`                  // The market regulators.
		DiscountAllowed       bool                  `json:"discountAllowed"`             // Indicate whether or not the user's discount rate is taken into account in this market.
		OpenDate              time.Time             `json:"openDate"`                    // The scheduled start date and time of the event (GMT by default).
		Version               int64                 `json:"version"`                     // A non-monotonically increasing number indicates market changes.
	}

	// PriceLadderDefinition defines the price ladder type for a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#PriceLadderDefinition
	PriceLadderDefinition struct {
		Type string `json:"type"` // Price ladder type (CLASSIC, FINEST, LINE_RANGE).
	}

	// RunnerChange contains price and volume changes for a runner.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#RunnerChange
	RunnerChange struct {
		Conflated                  bool        `json:"con"`           // Conflated - have more than a single change been combined (or null if not conflated).
		TradedVolume               float64     `json:"tv,omitempty"`  // The total amount matched. This value is truncated at 2dp.
		LastTradedPrice            float64     `json:"ltp,omitempty"` // Last Traded Price - The last traded price (or null if un-changed).
		StartingPriceNear          any         `json:"spn,omitempty"` // Starting Price Near - The near starting price (or null if un-changed).
		StartingPriceFar           any         `json:"spf,omitempty"` // Starting Price Far - The far starting price (or null if un-changed).
		BestAvailableToBack        [][]float64 `json:"batb"`          // Best Available To Back - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove).
		BestAvailableToLay         [][]float64 `json:"batl"`          // Best Available To Lay - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove).
		BestDisplayAvailableToBack [][]float64 `json:"bdatb"`         // Best Display Available To Back (includes virtual prices) - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove).
		BestDisplayAvailableToLay  [][]float64 `json:"bdatl"`         // Best Display Available To Lay (includes virtual prices) - LevelPriceVol triple delta of price changes, keyed by level (0 vol is remove).
		AvailableToBack            [][]float64 `json:"atb"`           // Available To Back - PriceVol tuple delta of price changes (0 vol is remove).
		AvailableToLay             [][]float64 `json:"atl"`           // Available To Lay - PriceVol tuple delta of price changes (0 vol is remove).
		StartingPriceBack          [][]float64 `json:"spb"`           // Starting Price Back - PriceVol tuple delta of price changes (0 vol is remove).
		StartingPriceLay           [][]float64 `json:"spl"`           // Starting Price Lay - PriceVol tuple delta of price changes (0 vol is remove).
		Traded                     [][]float64 `json:"trd"`           // Traded - PriceVol tuple delta of price changes (0 vol is remove).
	}
)

// SubscribeToMarkets sends a marketSubscription request to receive price changes for one or more markets.
// Subscription criteria are determined by marketFilter (horizontal filter) and marketDataFilter (vertical field filter).
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketSubscriptionMessage
func (client *StreamingClient) SubscribeToMarkets(marketFilter StreamMarketFilter, marketDataFilter MarketDataFilter) (chan MarketChangeMessage, error) {
	config := client.Config

	ms := MarketSubscriptionMessage{
		ID:                  int(rand.UintN(16)),
		Op:                  "marketSubscription",
		MarketFilter:        marketFilter,
		MarketDataFilter:    marketDataFilter,
		SegmentationEnabled: config.SegmentationEnabled,
		ConflateMs:          config.ConflateMs,
		HeartbeatMs:         config.HeartbeatMs,
		InitialClk:          config.InitialClk,
		Clk:                 config.Clk,
	}

	if err := client.Write(ms, true); err != nil {
		return nil, err
	}

	return ReadStream[MarketChangeMessage](client.Connection), nil
}
