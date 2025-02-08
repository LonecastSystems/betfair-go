package betfair

import (
	"math/rand/v2"
	"time"
)

type (
	MarketSubscriptionMessage struct {
		ID                  int                `json:"id"`
		Op                  string             `json:"op"`
		SegmentationEnabled bool               `json:"segmentationEnabled"`
		ConflateMs          int                `json:"conflateMs"`
		HeartbeatMs         int                `json:"heartbeatMs"`
		InitialClk          string             `json:"initialClk,omitempty"`
		Clk                 string             `json:"clk,omitempty"`
		MarketFilter        StreamMarketFilter `json:"marketFilter"`
		MarketDataFilter    MarketDataFilter   `json:"marketDataFilter"`
	}

	StreamMarketFilter struct {
		MarketIDs         []string `json:"marketIds,omitempty"`
		BspMarket         bool     `json:"bspMarket,omitempty"`
		BettingTypes      []string `json:"bettingTypes,omitempty"`
		EventTypeIDs      []string `json:"eventTypeIds,omitempty"`
		EventIDs          []string `json:"eventIds,omitempty"`
		TurnInPlayEnabled bool     `json:"turnInPlayEnabled,omitempty"`
		MarketTypes       []string `json:"marketTypes,omitempty"`
		Venues            []string `json:"venues,omitempty"`
		CountryCodes      []string `json:"countryCodes,omitempty"`
		RaceTypes         []string `json:"raceTypes,omitempty"`
	}

	MarketDataFilter struct {
		LadderLevels int      `json:"ladderLevels,omitempty"`
		Fields       []string `json:"fields,omitempty"`
	}

	MarketChangeMessage struct {
		ID            int            `json:"id"`
		Op            string         `json:"op"`
		ChangeType    string         `json:"ct"`
		SegmentType   string         `json:"segmentType"`
		ConflateMs    int            `json:"conflateMs"`
		Status        string         `json:"status"`
		HeartbeatMs   int            `json:"heartbeatMs"`
		PublishTime   int64          `json:"pt"`
		InitialClk    string         `json:"initialClk"`
		Clk           string         `json:"clk"`
		MarketChanges []MarketChange `json:"mc"`
	}

	MarketChange struct {
		Image            bool             `json:"img"`                        // Replace existing prices/data with the data supplied (null if delta)
		TotalVolume      float64          `json:"tv,omitempty"`               // Total amount matched across the market (null if unchanged)
		MarketDefinition MarketDefinition `json:"marketDefinition,omitempty"` // Full Market Definition if changed
		RunnerChanges    []RunnerChange   `json:"rc"`                         // List of Runner Changes
	}

	MarketDefinition struct {
		MarketID              string    `json:"id"`                    // Market Id - the id of the market
		Venue                 string    `json:"venue"`                 // The venue - applies to horse racing and greyhound markets only
		BspMarket             bool      `json:"bspMarket"`             // If 'true' the market supports Betfair SP betting
		TurnInPlayEnabled     bool      `json:"turnInPlayEnabled"`     // If 'true' the market is set to turn in-play
		PersistenceEnabled    bool      `json:"persistenceEnabled"`    // If 'true' the market supports 'Keep' bets if the market is to be turned in-play
		MarketBaseRate        float64   `json:"marketBaseRate"`        // The commission rate applicable to the market
		EventID               string    `json:"eventId"`               // The unique id for the event
		EventTypeID           string    `json:"eventTypeId"`           // The unique eventTypeId that the event belongs to
		NumberOfWinners       int       `json:"numberOfWinners"`       // The number of winners on a market
		BettingType           string    `json:"bettingType"`           // The market betting type i.e. ODDS, ASIAN_HANDICAP_DOUBLE_LINE, etc.
		MarketType            string    `json:"marketType"`            // Market base type
		MarketTime            string    `json:"marketTime"`            // The market start time
		SuspendTime           string    `json:"suspendTime"`           // The market suspend time
		BspReconciled         bool      `json:"bspReconciled"`         // True if the market starting price has been reconciled
		Complete              bool      `json:"complete"`              // If false, runners may be added to the market
		InPlay                bool      `json:"inPlay"`                // True if the market is currently in play
		CrossMatching         bool      `json:"crossMatching"`         // True if cross-matching is enabled for this market
		RunnersVoidable       bool      `json:"runnersVoidable"`       // True if runners in the market can be voided
		NumberOfActiveRunners int       `json:"numberOfActiveRunners"` // The number of runners that are currently active
		BetDelay              int       `json:"betDelay"`              // The number of seconds an order is held until it is submitted into the market
		Status                string    `json:"status"`                // The status of the market, for example, OPEN, SUSPENDED, CLOSED (settled), etc.
		Regulators            []string  `json:"regulators"`            // The market regulators
		DiscountAllowed       bool      `json:"discountAllowed"`       // Indicate whether or not the user's discount rate is taken into account in this market
		Timezone              string    `json:"timezone"`              // This is the timezone in which the event is taking place
		OpenDate              time.Time `json:"openDate"`              // The scheduled start date and time of the event (GMT by default)
		Version               int64     `json:"version"`               // A non-monotonically increasing number indicates market changes
	}

	RunnerChange struct {
		Conflated                  bool        `json:"con"`           // If true, more than one change is combined in this message
		TradedVolume               float64     `json:"tv,omitempty"`  // Traded Volume on this runner (only sent if changed)
		LastTradedPrice            float64     `json:"ltp,omitempty"` // Last Traded Price on this runner (only sent if changed)
		StartingPriceNear          interface{} `json:"spn,omitempty"` // Starting Price Near (only sent if changed)
		StartingPriceFar           interface{} `json:"spf,omitempty"` // Starting Price Far (only sent if changed)
		BestAvailableToBack        [][]float64 `json:"batb"`          // Best Available To Back
		BestAvailableToLay         [][]float64 `json:"batl"`          // Best Available To Lay
		BestDisplayAvailableToBack [][]float64 `json:"bdatb"`         // Best Display Available To Back
		BestDisplayAvailableToLay  [][]float64 `json:"bdatl"`         // Best Display Available To Back
		AvailableToBack            [][]float64 `json:"atb"`           // Available To Back (non-virtual prices)
		AvailableToLay             [][]float64 `json:"atl"`           // Available To Lay (non-virtual prices)
		StartingPriceBack          [][]float64 `json:"spb"`           // Starting Price (Available To) Back
		StartingPriceLay           [][]float64 `json:"spl"`           // Starting Price (Available To) Lay
		Traded                     [][]float64 `json:"trd"`           // Traded prices and sizes
	}
)

func (client *StreamingClient) SubscribeToMarkets(marketFilter StreamMarketFilter, marketDataFilter MarketDataFilter) (chan MarketChangeMessage, error) {
	ms := MarketSubscriptionMessage{
		ID:                  int(rand.Uint64()),
		Op:                  "marketSubscription",
		MarketFilter:        marketFilter,
		MarketDataFilter:    marketDataFilter,
		SegmentationEnabled: client.SegmentationEnabled,
		ConflateMs:          client.ConflateMs,
		HeartbeatMs:         client.HeartbeatMs,
		InitialClk:          client.InitialClk,
		Clk:                 client.Clk,
	}

	if err := client.Write(ms, true); err != nil {
		return nil, err
	}

	marketChanges := make(chan MarketChangeMessage)
	go ReadStream(client.Connection, marketChanges)

	return marketChanges, nil
}
