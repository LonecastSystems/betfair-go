package betfair

import "time"

// MarketBettingType restricts markets by the betting type of the market.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#MarketBettingType
type MarketBettingType string

const (
	MBT_ODDS                       MarketBettingType = "ODDS"                       // Odds Market - Any market that doesn't fit any any of the below categories.
	MBT_LINE                       MarketBettingType = "LINE"                       // Line Market - LINE markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps. Customers either Buy a line (LAY bet, winning if outcome is greater than the taken line (price)) or Sell a line (BACK bet, winning if outcome is less than the taken line (price)). If settled outcome equals the taken line, stake is returned.
	MBT_RANGE                      MarketBettingType = "RANGE"                      // Range Market - Now Deprecated.
	MBT_ASIAN_HANDICAP_SINGLE_LINE MarketBettingType = "ASIAN_HANDICAP_SINGLE_LINE" // Asian Single Line Market - A market in which there can be 0 or multiple winners. e,.g marketType TOTAL_GOALS.
	MBT_ASIAN_HANDICAP_DOUBLE_LINE MarketBettingType = "ASIAN_HANDICAP_DOUBLE_LINE" // Asian Handicap Market - A traditional Asian handicap market. Can be identified by marketType ASIAN_HANDICAP.
	MBT_FIXED_ODDS                 MarketBettingType = "FIXED_ODDS"                 // Sportsbook Odds Market. This type is deprecated and will be removed in future releases, when Sportsbook markets will be represented as ODDS market but with a different product type.
)

// BetDelayModel indicates which bet delay models are applied to a market if applicable.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#BetDelayModel
type BetDelayModel string

const (
	BDM_PASSIVE BetDelayModel = "PASSIVE" // For in-play markets where betDelay > 0, orders that are guaranteed not to match immediately are accepted straight away, bypassing the bet delay wait. Order requirements (otherwise bets will be subject to the usual bet delay before being placed). Only plain LIMIT orders are supported. Allowed persistenceType: LAPSE. The following attributes are not supported and must be omitted: timeInForce, minFillSize, betTargetType.
	BDM_DYNAMIC BetDelayModel = "DYNAMIC" // Indicates market is subject to dynamic in-play bet delays. This mean that the in-play betDelay will vary while the market is turned in-play. Please note: Currently returned for Tennis markets only. Specifically, every game 3,5,7,9,11 or game which decides a set (potentially 6,8,10,12) the betDelay is reduced to 1 second.
)

type (
	// MarketFilter selects markets for list* catalogue operations.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketFilter
	MarketFilter struct {
		TextQuery          string              `json:"textQuery,omitempty"`          // TextQuery restricts markets by any text associated with the Event name. You can include a wildcard (*) character as long as it is not the first character. Please note - the textQuery field doesn't evaluate market or selection names.
		EventTypeIDs       []string            `json:"eventTypeIds,omitempty"`       // EventTypeIDs restricts markets by event type associated with the market. (i.e., Football, Hockey, etc)
		EventIDs           []string            `json:"eventIds,omitempty"`           // EventIDs restricts markets by the event id associated with the market.
		CompetitionIDs     []string            `json:"competitionIds,omitempty"`     // CompetitionIDs restricts markets by the competitions associated with the market.
		MarketIDs          []string            `json:"marketIds,omitempty"`          // MarketIDs restricts markets by the market id associated with the market.
		Venues             []string            `json:"venues,omitempty"`             // Venues restricts markets by the venue associated with the market. Currently, only Horse & Greyhound racing markets have venues.
		BspOnly            bool                `json:"bspOnly,omitempty"`            // BspOnly restricts to bsp markets only, if True or non-bsp markets if False. If not specified then returns both BSP and non-BSP markets
		TurnInPlayEnabled  bool                `json:"turnInPlayEnabled,omitempty"`  // TurnInPlayEnabled restricts to markets that will turn in play if True or will not turn in play if false. If not specified, returns both.
		InPlayOnly         bool                `json:"inPlayOnly,omitempty"`         // InPlayOnly restricts to markets that are currently in play if True or are not currently in play if false. If not specified, returns both.
		MarketBettingTypes []MarketBettingType `json:"marketBettingTypes,omitempty"` // MarketBettingTypes restricts to markets that match the betting type of the market (i.e. Odds, Asian Handicap Singles, Asian Handicap Doubles or Line)
		MarketTypeCodes    []string            `json:"marketTypeCodes,omitempty"`    // MarketTypeCodes restricts to markets that match the type of the market (i.e., MATCH_ODDS, HALF_TIME_SCORE). You should use this instead of relying on the market name as the market type codes are the same in all locales. Please note: All market types are available via the listMarketTypes operations.
		MarketCountries    []string            `json:"marketCountries,omitempty"`    // MarketCountries restricts to markets that are in the specified country or countries. Please note: the default value is 'GB' when the correct country code cannot be determined.
		MarketStartTime    TimeRange           `json:"marketStartTime,omitempty"`    // MarketStartTime restricts to markets with a market start time before or after the specified date
		WithOrders         []OrderStatus       `json:"withOrders,omitempty"`         // WithOrders restricts to markets where I have one or more orders in these states.
		RaceTypes          []RaceType          `json:"raceTypes,omitempty"`          // RaceTypes restricts to markets of a specific raceType. Valid values are - Harness, Flat, Hurdle, Chase, Bumper, NH Flat, Steeple (AUS/NZ races), and NO_VALUE (when no valid race type has been mapped).
		BetDelayModels     []BetDelayModel     `json:"betDelayModels,omitempty"`     // BetDelayModels restricts to markets specific betDelayModels. e.g. PASSIVE, DYNAMIC
	}

	// TimeRange restricts to markets with a market start time before or after the specified date.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#TimeRange
	TimeRange struct {
		From string `json:"from,omitempty"` // From is from.
		To   string `json:"to,omitempty"`   // To is to.
	}
)

// Wallet is the wallet from which funds are taken or to which statement items belong.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#Wallet
type Wallet string

const (
	W_UK        Wallet = "UK"        // The Global Exchange wallet.
	W_AUSTRALIA Wallet = "AUSTRALIA" // Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#Wallet
	W_SPAIN     Wallet = "SPAIN"     // Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#Wallet
	W_ITALY     Wallet = "ITALY"     // Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#Wallet
)

// MatchProjection specifies the representation of matches if you ask for orders.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#MatchProjection
type MatchProjection string

const (
	MP_NO_ROLLUP              MatchProjection = "NO_ROLLUP"              // No rollup, return raw fragments.
	MP_ROLLED_UP              MatchProjection = "ROLLED_UP_BY_PRICE"     // Rollup matched amounts by distinct matched prices per side.
	MP_ROLLED_UP_BY_AVG_PRICE MatchProjection = "ROLLED_UP_BY_AVG_PRICE" // Rollup matched amounts by average matched price per side.
)

// OrderProjection is the orders you want to receive in the response.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#OrderProjection
type OrderProjection string

const (
	OP_ALL                OrderProjection = "ALL"                // EXECUTABLE and EXECUTION_COMPLETE orders.
	OP_EXECUTABLE         OrderProjection = "EXECUTABLE"         // An order that has a remaining unmatched portion. This is either a fully unmatched or partially matched bet (order).
	OP_EXECUTION_COMPLETE OrderProjection = "EXECUTION_COMPLETE" // An order that does not have any remaining unmatched portion. This is a fully matched bet (order).
)

// PriceData is the basic price data you want to receive in the response.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#PriceData
type PriceData string

const (
	PD_SP_AVAILABLE   PriceData = "SP_AVAILABLE"   // Amount available for the BSP auction.
	PD_SP_TRADED      PriceData = "SP_TRADED"      // Amount traded in the BSP auction.
	PD_EX_BEST_OFFERS PriceData = "EX_BEST_OFFERS" // Only the best prices available for each runner, to requested price depth.
	PD_EX_ALL_OFFERS  PriceData = "EX_ALL_OFFERS"  // EX_ALL_OFFERS trumps EX_BEST_OFFERS if both settings are present.
	PD_EX_TRADED      PriceData = "EX_TRADED"      // Amount traded on the exchange.
)

type (
	// PriceProjection is the selection criteria of the returning price data.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PriceProjection
	PriceProjection struct {
		PriceData             []PriceData     `json:"priceData"`                       // PriceData is the basic price data you want to receive in the response.
		ExBestOffersOverrides OffersOverrides `json:"exBestOffersOverrides,omitempty"` // ExBestOffersOverrides are options to alter the default representation of best offer prices Applicable to EX_BEST_OFFERS priceData selection
		Virtualise            bool            `json:"virtualise,omitempty"`            // Virtualise indicates if the returned prices should include virtual prices. Applicable to EX_BEST_OFFERS and EX_ALL_OFFERS priceData selections, default value is false. Please note: This must be set to 'true' replicate the display of prices on the Betfair Exchange website.
		RolloverStakes        bool            `json:"rolloverStakes,omitempty"`        // RolloverStakes indicates if the volume returned at each price point should be the absolute value or a cumulative sum of volumes available at the price and all better prices. If unspecified defaults to false. Applicable to EX_BEST_OFFERS and EX_ALL_OFFERS price projections. Not supported as yet.
	}

	// OffersOverrides are options to alter the default representation of best offer prices.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#OffersOverrides
	OffersOverrides struct {
		BestPricesDepth       int         `json:"bestPricesDepth,omitempty"`       // BestPricesDepth is the maximum number of prices to return on each side for each runner. If unspecified defaults to 3. The maximum returned price depth returned is 10.
		RollupModel           RollupModel `json:"rollupModel,omitempty"`           // RollupModel is the model to use when rolling up available sizes. If unspecified defaults to STAKE rollup model with rollupLimit of minimum stake in the specified currency.
		RollupLimit           int         `json:"rollupLimit,omitempty"`           // RollupLimit is the volume limit to use when rolling up returned sizes. The exact definition of the limit depends on the rollupModel. If no limit is provided it will use minimum stake as default the value. Ignored if no rollup model is specified.
		RollupLiability       float64     `json:"rollupLiability,omitempty"`       // RollupLiability is only applicable when rollupModel is MANAGED_LIABILITY. The rollup model switches from being stake-based to liability-based at the smallest lay price which is >= rollupLiabilityThreshold.service level default (TBD). Not supported as yet.
		RollupLiabilityFactor int         `json:"rollupLiabilityFactor,omitempty"` // RollupLiabilityFactor is only applicable when rollupModel is MANAGED_LIABILITY. (rollupLiabilityFactor * rollupLimit) is the minimum liability the user is deemed to be comfortable with. After the rollupLiabilityThreshold price subsequent volumes will be rolled up to minimum value such that the liability >= the minimum liability.service level default (5). Not supported as yet.
	}
)

type (
	// MarketBook is the dynamic data in a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketBook
	MarketBook struct {
		MarketID              string             `json:"marketId"`                        // MarketID is the unique identifier for the market. MarketId's are prefixed with '1.'.
		IsMarketDataDelayed   bool               `json:"isMarketDataDelayed"`             // IsMarketDataDelayed is true if the data returned by listMarketBook will be delayed. The data may be delayed because you are not logged in with a funded account or you are using an Application Key that does not allow up to date data.
		Status                MarketStatus       `json:"status,omitempty"`                // Status is the status of the market, for example OPEN, SUSPENDED, CLOSED (settled), etc.
		SuspendReason         string             `json:"suspendReason,omitempty"`         // SuspendReason is currently returned only for Soccer markets, when status = SUSPENDED. Possible values are Goal, Third Party Unavailable, Penalty, Red Card, Non In Play Market,
		BetDelay              int                `json:"betDelay,omitempty"`              // BetDelay is the number of seconds an order is held until it is submitted into the market. Orders are usually delayed when the market is in-play
		BspReconciled         bool               `json:"bspReconciled,omitempty"`         // BspReconciled is true if the market starting price has been reconciled
		Complete              bool               `json:"complete,omitempty"`              // Complete is false if runners may be added to the market
		Inplay                bool               `json:"inplay,omitempty"`                // Inplay is true if the market is currently in play
		NumberOfWinners       int                `json:"numberOfWinners,omitempty"`       // NumberOfWinners is the number of selections that could be settled as winners
		NumberOfRunners       int                `json:"numberOfRunners,omitempty"`       // NumberOfRunners is the number of runners in the market
		NumberOfActiveRunners int                `json:"numberOfActiveRunners,omitempty"` // NumberOfActiveRunners is the number of runners that are currently active. An active runner is a selection available for betting
		LastMatchTime         time.Time          `json:"lastMatchTime,omitempty"`         // LastMatchTime is the most recent time an order was executed
		TotalMatched          float64            `json:"totalMatched,omitempty"`          // TotalMatched is the total amount matched
		TotalAvailable        float64            `json:"totalAvailable,omitempty"`        // TotalAvailable is the total amount of orders that remain unmatched
		CrossMatching         bool               `json:"crossMatching,omitempty"`         // CrossMatching is true if cross matching is enabled for this market.
		RunnersVoidable       bool               `json:"runnersVoidable,omitempty"`       // RunnersVoidable is true if runners in the market can be voided. Please note - this doesn't include horse racing markets under which bets are voided on non-runners with any applicable reduction factor applied/
		Version               int64              `json:"version,omitempty"`               // Version is the version of the market. The version increments whenever the market status changes, for example, turning in-play, or suspended when a goal is scored.
		Runners               []Runner           `json:"runners,omitempty"`               // Runners is information about the runners (selections) in the market.
		KeyLineDescription    KeyLineDescription `json:"keyLineDescription,omitempty"`    // KeyLineDescription is description of a markets key line for valid market types
		BetDelayModels        []BetDelayModel    `json:"betDelayModels,omitempty"`        // BetDelayModels indicates which bet delay models are applied to a market if applicable. i.e. PASSIVE, DYNAMIC or both.
	}

	// Runner is the dynamic data about runners in a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#Runner
	Runner struct {
		SelectionID       int64              `json:"selectionId"`                 // SelectionID is the unique id of the runner (selection). Please note - the same selectionId and runnerName pairs are used accross all Betfair markets which contain them.
		Handicap          float64            `json:"handicap"`                    // Handicap is the handicap. Enter the specific handicap value (returned by RUNNER in listMaketBook) if the market is an Asian handicap market.
		Status            RunnerStatus       `json:"status"`                      // Status is the status of the selection (i.e., ACTIVE, REMOVED, WINNER, PLACED, LOSER, HIDDEN) Runner status information is available for 90 days following market settlement.
		AdjustmentFactor  float64            `json:"adjustmentFactor,omitempty"`  // AdjustmentFactor is the adjustment factor applied if the selection is removed
		LastPriceTraded   float64            `json:"lastPriceTraded,omitempty"`   // LastPriceTraded is the price of the most recent bet matched on this selection
		TotalMatched      float64            `json:"totalMatched,omitempty"`      // TotalMatched is the total amount matched on this runner
		RemovalDate       time.Time          `json:"removalDate,omitempty"`       // RemovalDate is if date and time the runner was removed
		StartingPrices    StartingPrices     `json:"sp,omitempty"`                // StartingPrices are the BSP related prices for this runner
		ExchangePrices    ExchangePrices     `json:"ex,omitempty"`                // ExchangePrices are the Exchange prices available for this runner
		Orders            []Order            `json:"orders,omitempty"`            // Orders is list of orders in the market
		Matches           []Match            `json:"matches,omitempty"`           // Matches is list of matches (i.e, orders that have been fully or partially executed)
		MatchesByStrategy map[string][]Match `json:"matchesByStrategy,omitempty"` // MatchesByStrategy is list of matches for each strategy, ordered by matched data
	}

	// StartingPrices is information about the Betfair Starting Price. Only available in BSP markets.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#StartingPrices
	StartingPrices struct {
		NearPrice         Float       `json:"nearPrice,omitempty"`         // NearPrice is what the starting price would be if the market was reconciled now taking into account the SP bets as well as unmatched exchange bets on the same selection in the exchange. This data is cached and update every 60 seconds. Please note: Type Double may contain numbers, INF, -INF, and NaN.
		FarPrice          Float       `json:"farPrice,omitempty"`          // FarPrice is what the starting price would be if the market was reconciled now taking into account only the currently place SP bets. The Far Price is not as complicated but not as accurate and only accounts for money on the exchange at SP. This data is cached and updated every 60 seconds. Please note: Type Double may contain numbers, INF, -INF, and NaN.
		BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`    // BackStakeTaken is the total amount of back bets matched at the actual Betfair Starting Price. Pre-reconciliation, this field is zero for all prices except 1.01 (for Market on Close bets) and at the limit price for any Limit on Close bets.
		LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"` // LayLiabilityTaken is the lay amount matched at the actual Betfair Starting Price. Pre-reconciliation, this field is zero for all prices except 1000 (for Market on Close bets) and at the limit price for any Limit on Close bets.
		ActualSP          Float       `json:"actualSP,omitempty"`          // ActualSP is the final BSP price for this runner. Only available for a BSP market that has been reconciled. Please note: for REMOVED runners the actualSP will be returned as 'NaN. Value may be returned as 'Infinity' if no BSP can be calculated.
	}

	// ExchangePrices are exchange prices available for a runner.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ExchangePrices
	ExchangePrices struct {
		AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
		AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
		TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
	}

	// PriceSize is a price and stake available at that price.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PriceSize
	PriceSize struct {
		Price Float `json:"price"` // Price is the price available
		Size  Float `json:"size"`  // Size is the stake available. Virtualised stream sizes may be Infinity.
	}

	// Order is an order in the market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#Order
	Order struct {
		BetID               string          `json:"betId"`                         // BetID is the bet id.
		OrderType           OrderType       `json:"orderType"`                     // OrderType is BSP Order type.
		Status              OrderStatus     `json:"status"`                        // Status is either EXECUTABLE (an unmatched amount remains) or EXECUTION_COMPLETE (no unmatched amount remains).
		PersistenceType     PersistenceType `json:"persistenceType"`               // PersistenceType is what to do with the order at turn-in-play
		Side                Side            `json:"side"`                          // Side indicates if the bet is a Back or a LAY. For LINE markets customers either Buy a line (LAY bet, winning if outcome is greater than the taken line (price)) or Sell a line (BACK bet, winning if outcome is less than the taken line (price))
		Price               float64         `json:"price"`                         // Price is the price of the bet. Please note: LINE markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps
		Size                float64         `json:"size"`                          // Size is the size of the bet.
		BspLiability        float64         `json:"bspLiability"`                  // BspLiability is not to be confused with size. This is the liability of a given BSP bet.
		PlacedDate          time.Time       `json:"placedDate"`                    // PlacedDate is the date, to the second, the bet was placed.
		AvgPriceMatched     float64         `json:"avgPriceMatched,omitempty"`     // AvgPriceMatched is the average price matched at. Voided match fragments are removed from this average calculation. For MARKET_ON_CLOSE BSP bets this reports the matched SP price following the SP reconciliation process. This value is not meaningful for activity on LINE markets and is not guaranteed to be returned or maintained for these markets.
		SizeMatched         float64         `json:"sizeMatched,omitempty"`         // SizeMatched is the current amount of this bet that was matched.
		SizeRemaining       float64         `json:"sizeRemaining,omitempty"`       // SizeRemaining is the current amount of this bet that is unmatched.
		SizeLapsed          float64         `json:"sizeLapsed,omitempty"`          // SizeLapsed is the current amount of this bet that was lapsed.
		SizeCancelled       float64         `json:"sizeCancelled,omitempty"`       // SizeCancelled is the current amount of this bet that was cancelled.
		SizeVoided          float64         `json:"sizeVoided,omitempty"`          // SizeVoided is the current amount of this bet that was voided.
		CustomerOrderRef    string          `json:"customerOrderRef,omitempty"`    // CustomerOrderRef is the customer order reference sent for this bet
		CustomerStrategyRef string          `json:"customerStrategyRef,omitempty"` // CustomerStrategyRef is the customer strategy reference sent for this bet
	}

	// Match is an individual bet Match, or rollup by price or avg price. Rollup depends on the requested MatchProjection.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#Match
	Match struct {
		BetID     string    `json:"betId,omitempty"`     // BetID is only present if no rollup
		MatchID   string    `json:"matchId,omitempty"`   // MatchID is only present if no rollup
		Side      Side      `json:"side"`                // Side indicates if the bet is a Back or a LAY
		Price     float64   `json:"price"`               // Price is either actual match price or avg match price depending on rollup. This value is not meaningful for activity on LINE markets and is not guaranteed to be returned or maintained for these markets.
		Size      float64   `json:"size"`                // Size is size matched at in this fragment, or at this price or avg price depending on rollup
		MatchDate time.Time `json:"matchDate,omitempty"` // MatchDate is only present if no rollup
	}

	// KeyLineDescription is a list of KeyLineSelection objects describing the key line for the market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#KeyLineDescription
	KeyLineDescription struct {
		KeyLine []KeyLineSelection `json:"keyLine,omitempty"` // KeyLine is a list of KeyLineSelection objects
	}

	// KeyLineSelection describes a markets key line selection, comprising the selectionId and handicap of the team it is applied to.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#KeyLineSelection
	KeyLineSelection struct {
		SelectionID int64   `json:"selectionId,omitempty"` // SelectionID is selection ID of the runner in the key line handicap.
		Handicap    float64 `json:"handicap,omitempty"`    // Handicap is handicap value of the key line.
	}
)

// Competition is a competition the market is contained within.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#Competition
type Competition struct {
	ID   string `json:"id"`   // ID is id.
	Name string `json:"name"` // Name is name.
}

// MarketVersion is the market version.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketVersion
type MarketVersion struct {
	Version int64 `json:"version,omitempty"` // Version is a non-monotonically increasing number indicating market changes
}

type (
	// Event is an event the market is contained within.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#Event
	Event struct {
		ID          string    `json:"id,omitempty"`          // ID is the unique id for the event
		Name        string    `json:"name,omitempty"`        // Name is the name of the event
		CountryCode string    `json:"countryCode,omitempty"` // CountryCode is the ISO-2 code for the event. A list of ISO-2 codes is available via http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2 Please note: default value is 'GB' when the correct country code cannot be determined.
		Timezone    string    `json:"timezone,omitempty"`    // Timezone is this is timezone in which the event is taking place.
		Venue       string    `json:"venue,omitempty"`       // Venue is venue.
		OpenDate    time.Time `json:"openDate,omitempty"`    // OpenDate is the scheduled start date and time of the event. This is Europe/London (GMT) by default
	}

	// EventType is the Event Type the market is contained within.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#EventType
	EventType struct {
		ID   string `json:"id,omitempty"`   // ID is id.
		Name string `json:"name,omitempty"` // Name is name.
	}
)

type (
	// CancelOrdersInstructionReport is the response to a CancelInstruction.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CancelInstructionReport
	CancelOrdersInstructionReport struct {
		Status        InstructionReportStatus    `json:"status"`                  // Status is whether the command succeeded or failed
		ErrorCode     InstructionReportErrorCode `json:"errorCode,omitempty"`     // ErrorCode is cause of failure, or null if command succeeds
		Instruction   CancelInstruction          `json:"instruction,omitempty"`   // Instruction is the instruction that was requested
		SizeCancelled float64                    `json:"sizeCancelled"`           // SizeCancelled is sizeCancelled.
		CancelledDate time.Time                  `json:"cancelledDate,omitempty"` // CancelledDate is cancelledDate.
	}

	// PlaceOrdersInstructionReport is the response to a PlaceInstruction.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PlaceInstructionReport
	PlaceOrdersInstructionReport struct {
		Status              InstructionReportStatus    `json:"status"`                        // Status is whether the command succeeded or failed
		ErrorCode           InstructionReportErrorCode `json:"errorCode,omitempty"`           // ErrorCode is cause of failure, or null if the command succeeds
		OrderStatus         OrderStatus                `json:"orderStatus,omitempty"`         // OrderStatus is the status of the order, if the instruction succeeded. If the instruction is unsuccessful, no value is provided. Please note: by default, this field is not returned for MARKET_ON_CLOSE and LIMIT_ON_CLOSE orders.
		Instruction         PlaceInstruction           `json:"instruction"`                   // Instruction is the instruction that was requested
		BetID               string                     `json:"betId,omitempty"`               // BetID is the bet ID of the new bet. Will be null on failure or if the order was placed asynchronously.
		PlacedDate          time.Time                  `json:"placedDate,omitempty"`          // PlacedDate will be null if the order was placed asynchronously
		AveragePriceMatched float64                    `json:"averagePriceMatched,omitempty"` // AveragePriceMatched will be null if the order was placed asynchronously. This value is not meaningful for activity on LINE markets and is not guaranteed to be returned or maintained for these markets.
		SizeMatched         float64                    `json:"sizeMatched,omitempty"`         // SizeMatched will be null if the order was placed asynchronously
	}
)

// Side indicates if the bet is a Back or a LAY.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#Side
type Side string

const (
	SD_BACK Side = "BACK" // To back a team, horse or outcome is to bet on the selection to win. For LINE markets a Back bet refers to a SELL line. A SELL line will win if the outcome is LESS THAN the taken line (price).
	SD_LAY  Side = "LAY"  // To lay a team, horse, or outcome is to bet on the selection to lose. For LINE markets a Lay bet refers to a BUY line. A BUY line will win if the outcome is MORE THAN the taken line (price).
)

// OrderStatus is the status of the order.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#OrderStatus
type OrderStatus string

const (
	OS_PENDING            OrderStatus = "PENDING"            // An asynchronous order is yet to be processed. Once the bet has been processed by the exchange (including waiting for any in-play delay), the result will be reported and available on the Exchange Stream API and API NG. Not a valid search criteria on MarketFilter.
	OS_EXECUTION_COMPLETE OrderStatus = "EXECUTION_COMPLETE" // An order that does not have any remaining unmatched portion.
	OS_EXECUTABLE         OrderStatus = "EXECUTABLE"         // An order that has a remaining unmatched portion.
	OS_EXPIRED            OrderStatus = "EXPIRED"            // The order is no longer available for execution due to its time in force constraint. In the case of FILL_OR_KILL orders, this means the order has been killed because it could not be filled to your specifications. Not a valid search criteria on MarketFilter.
)

// OrderType is BSP Order type.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#OrderType
type OrderType string

const (
	OT_LIMIT           OrderType = "LIMIT"           // A normal exchange limit order for immediate execution.
	OT_LIMIT_ON_CLOSE  OrderType = "LIMIT_ON_CLOSE"  // Limit order for the auction (SP).
	OT_MARKET_ON_CLOSE OrderType = "MARKET_ON_CLOSE" // Market order for the auction (SP).
)

// ExecutionReportStatus is the status of the execution report.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#ExecutionReportStatus
type ExecutionReportStatus string

const (
	ERS_SUCCESS               ExecutionReportStatus = "SUCCESS"               // Order processed successfully.
	ERS_FAILURE               ExecutionReportStatus = "FAILURE"               // Order failed.
	ERS_PROCESSED_WITH_ERRORS ExecutionReportStatus = "PROCESSED_WITH_ERRORS" // The order itself has been accepted, but at least one (possibly all) actions have generated errors. This error only occurs for replaceOrders, cancelOrders and updateOrders operations. In normal circumstances the placeOrders operation will not return PROCESSED_WITH_ERRORS status as it is an atomic operation. PLEASE NOTE: if the 'Best Execution' features is switched off, placeOrders can return 'PROCESSED_WITH_ERRORS' meaning that some bets can be rejected and other placed when submitted in the same PlaceInstruction.
	ERS_TIMEOUT               ExecutionReportStatus = "TIMEOUT"               // The order timed out & the status of the bet is unknown. If a TIMEOUT error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please Note: Timeouts will occur after 5 seconds of attempting to process the bet but please allow up to 15 seconds for a timed out order to appear. After this time any unprocessed bets will automatically be Lapsed and no longer be available on the Exchange.
)

// ExecutionReportErrorCode is the error code if the operation failed.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#ExecutionReportErrorCode
type ExecutionReportErrorCode string

const (
	ERS_ERR_ERROR_IN_MATCHER            ExecutionReportErrorCode = "ERROR_IN_MATCHER"            // The matcher is not healthy. Please note: The error will also be returned is you attempt concurrent 'cancel all' bets requests using cancelOrders which isn't permitted.
	ERS_ERR_PROCESSED_WITH_ERRORS       ExecutionReportErrorCode = "PROCESSED_WITH_ERRORS"       // The order itself has been accepted, but at least one (possibly all) actions have generated errors.
	ERS_ERR_BET_ACTION_ERROR            ExecutionReportErrorCode = "BET_ACTION_ERROR"            // There is an error with an action that has caused the entire order to be rejected. Check the instructionReports errorCode for the reason for the rejection of the order.
	ERS_ERR_INVALID_ACCOUNT_STATE       ExecutionReportErrorCode = "INVALID_ACCOUNT_STATE"       // Order rejected due to the account's status (suspended, inactive, dup cards).
	ERS_ERR_INVALID_WALLET_STATUS       ExecutionReportErrorCode = "INVALID_WALLET_STATUS"       // Order rejected due to the account's wallet's status.
	ERS_ERR_INSUFFICIENT_FUNDS          ExecutionReportErrorCode = "INSUFFICIENT_FUNDS"          // Account has exceeded its exposure limit or available to bet limit.
	ERS_ERR_LOSS_LIMIT_EXCEEDED         ExecutionReportErrorCode = "LOSS_LIMIT_EXCEEDED"         // The account has exceed the self imposed loss limit.
	ERS_ERR_MARKET_SUSPENDED            ExecutionReportErrorCode = "MARKET_SUSPENDED"            // Market is suspended.
	ERS_ERR_MARKET_NOT_OPEN_FOR_BETTING ExecutionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING" // Market is not open for betting. It is either not yet active, suspended or closed awaiting settlement.
	ERS_ERR_DUPLICATE_TRANSACTION       ExecutionReportErrorCode = "DUPLICATE_TRANSACTION"       // Duplicate customer reference data submitted - Please note: There is a time window associated with the de-duplication of duplicate submissions which is 60 second.
	ERS_ERR_INVALID_ORDER               ExecutionReportErrorCode = "INVALID_ORDER"               // Order cannot be accepted by the matcher due to the combination of actions. For example, bets being edited are not on the same market, or order includes both edits and placement.
	ERS_ERR_INVALID_MARKET_ID           ExecutionReportErrorCode = "INVALID_MARKET_ID"           // Market doesn't exist.
	ERS_ERR_PERMISSION_DENIED           ExecutionReportErrorCode = "PERMISSION_DENIED"           // Business rules do not allow order to be placed. You are either attempting to place the order using a Delayed Application Key or from a restricted jurisdiction (i.e. USA).
	ERS_ERR_DUPLICATE_BETIDS            ExecutionReportErrorCode = "DUPLICATE_BETIDS"            // Duplicate bet ids found. For example, you've included the same betId more than once in a single cancelOrders request.
	ERS_ERR_NO_ACTION_REQUIRED          ExecutionReportErrorCode = "NO_ACTION_REQUIRED"          // Order hasn't been passed to matcher as system detected there will be no state change.
	ERS_ERR_SERVICE_UNAVAILABLE         ExecutionReportErrorCode = "SERVICE_UNAVAILABLE"         // The requested service is unavailable.
	ERS_ERR_REJECTED_BY_REGULATOR       ExecutionReportErrorCode = "REJECTED_BY_REGULATOR"       // The regulator rejected the order. On the Italian Exchange this error will occur if more than 50 bets are sent in a single placeOrders request.
	ERS_ERR_NO_CHASING                  ExecutionReportErrorCode = "NO_CHASING"                  // A specific error code that relates to Spanish Exchange markets only which indicates that the bet placed contravenes the Spanish regulatory rules relating to loss chasing.
	ERS_ERR_REGULATOR_IS_NOT_AVAILABLE  ExecutionReportErrorCode = "REGULATOR_IS_NOT_AVAILABLE"  // The underlying regulator service is not available.
	ERS_ERR_TOO_MANY_INSTRUCTIONS       ExecutionReportErrorCode = "TOO_MANY_INSTRUCTIONS"       // The amount of orders exceeded the maximum amount allowed to be executed.
	ERS_ERR_INVALID_MARKET_VERSION      ExecutionReportErrorCode = "INVALID_MARKET_VERSION"      // The supplied market version is invalid. Max length allowed for market version is 12.
	ERS_ERR_INVALID_PROFIT_RATIO        ExecutionReportErrorCode = "INVALID_PROFIT_RATIO"        // The order falls outside the permitted price and size combination.
)

// PersistenceType is what to do with the order at turn-in-play.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#PersistenceType
type PersistenceType string

const (
	PT_LAPSE           PersistenceType = "LAPSE"           // Lapse (cancel) the order automatically when the market is turned in play if the bet is unmatched.
	PT_PERSIST         PersistenceType = "PERSIST"         // Persist the unmatched order to in-play. The bet will be placed automatically into the in-play market at the start of the event. Once in play, the bet won't be cancelled by Betfair if a material event takes place and will be available until matched or cancelled by the user.
	PT_MARKET_ON_CLOSE PersistenceType = "MARKET_ON_CLOSE" // Put the order into the auction (SP) at turn-in-play.
)

// InstructionReportStatus is whether the command succeeded or failed.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#InstructionReportStatus
type InstructionReportStatus string

const (
	IRS_SUCCESS InstructionReportStatus = "SUCCESS" // The instruction was successful.
	IRS_FAILURE InstructionReportStatus = "FAILURE" // The instruction failed.
	IRS_TIMEOUT InstructionReportStatus = "TIMEOUT" // The order timed out & the status of the bet is unknown. If a TIMEOUT error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please Note: Timeouts will occur after 5 seconds of attempting to process the bet but please allow up to 15 seconds for a timed out order to appear. After this time any unprocessed bets will automatically be Lapsed and no longer be available on the Exchange.
)

// InstructionReportErrorCode is cause of failure, or null if the command succeeds.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#InstructionReportErrorCode
type InstructionReportErrorCode string

const (
	IRS_ERR_INVALID_BET_SIZE                       InstructionReportErrorCode = "INVALID_BET_SIZE"                       // bet size is invalid for your currency or your regulator.
	IRS_ERR_INVALID_RUNNER                         InstructionReportErrorCode = "INVALID_RUNNER"                         // Runner does not exist, includes vacant traps in greyhound racing.
	IRS_ERR_BET_TAKEN_OR_LAPSED                    InstructionReportErrorCode = "BET_TAKEN_OR_LAPSED"                    // Bet cannot be cancelled or modified as it has already been taken or has been cancelled/lapsed Includes attempts to cancel/modify market on close BSP bets and cancelling limit on close BSP bets. The error may be returned on placeOrders request if for example a bet is placed at the point when a market admin event takes place (i.e. market is turned in-play). The error will also be returned if a market version is submitted and a material change has taken place since the bet was submitted causing the bet to be rejected.
	IRS_ERR_BET_IN_PROGRESS                        InstructionReportErrorCode = "BET_IN_PROGRESS"                        // No result was received from the matcher in a timeout configured for the system.
	IRS_ERR_RUNNER_REMOVED                         InstructionReportErrorCode = "RUNNER_REMOVED"                         // Runner has been removed from the event.
	IRS_ERR_MARKET_NOT_OPEN_FOR_BETTING            InstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"            // Attempt to edit a bet on a market that has closed.
	IRS_ERR_LOSS_LIMIT_EXCEEDED                    InstructionReportErrorCode = "LOSS_LIMIT_EXCEEDED"                    // The action has caused the account to exceed the self imposed loss limit.
	IRS_ERR_MARKET_NOT_OPEN_FOR_BSP_BETTING        InstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BSP_BETTING"        // Market now closed to bsp betting. Turned in-play or has been reconciled.
	IRS_ERR_INVALID_PRICE_EDIT                     InstructionReportErrorCode = "INVALID_PRICE_EDIT"                     // Attempt to edit down the price of a bsp limit on close lay bet, or edit up the price of a limit on close back bet.
	IRS_ERR_INVALID_ODDS                           InstructionReportErrorCode = "INVALID_ODDS"                           // Odds not on price ladder - either edit or placement.
	IRS_ERR_INSUFFICIENT_FUNDS                     InstructionReportErrorCode = "INSUFFICIENT_FUNDS"                     // Insufficient funds available to cover the bet action. Either the exposure limit or available to bet limit would be exceeded.
	IRS_ERR_INVALID_PERSISTENCE_TYPE               InstructionReportErrorCode = "INVALID_PERSISTENCE_TYPE"               // Invalid persistence type for this market, e.g. KEEP for a non in-play market or KEEP for markets with PASSIVE betDelayModels.
	IRS_ERR_ERROR_IN_MATCHER                       InstructionReportErrorCode = "ERROR_IN_MATCHER"                       // A problem with the matcher prevented this action completing successfully.
	IRS_ERR_INVALID_BACK_LAY_COMBINATION           InstructionReportErrorCode = "INVALID_BACK_LAY_COMBINATION"           // The order contains a back and a lay for the same runner at overlapping prices. This would guarantee a self match. This also applies to BSP limit on close bets.
	IRS_ERR_ERROR_IN_ORDER                         InstructionReportErrorCode = "ERROR_IN_ORDER"                         // The action failed because the parent order failed.
	IRS_ERR_INVALID_BID_TYPE                       InstructionReportErrorCode = "INVALID_BID_TYPE"                       // Bid type is mandatory.
	IRS_ERR_INVALID_BET_ID                         InstructionReportErrorCode = "INVALID_BET_ID"                         // Bet for id supplied has not been found.
	IRS_ERR_CANCELLED_NOT_PLACED                   InstructionReportErrorCode = "CANCELLED_NOT_PLACED"                   // Bet cancelled but replacement bet was not placed.
	IRS_ERR_RELATED_ACTION_FAILED                  InstructionReportErrorCode = "RELATED_ACTION_FAILED"                  // Action failed due to the failure of a action on which this action is dependent.
	IRS_ERR_NO_ACTION_REQUIRED                     InstructionReportErrorCode = "NO_ACTION_REQUIRED"                     // The action does not result in any state change. eg changing a persistence to it's current value.
	IRS_ERR_TIME_IN_FORCE_CONFLICT                 InstructionReportErrorCode = "TIME_IN_FORCE_CONFLICT"                 // You may only specify a time in force on either the place request OR on individual limit order instructions (not both), since the implied behaviors are incompatible.
	IRS_ERR_UNEXPECTED_PERSISTENCE_TYPE            InstructionReportErrorCode = "UNEXPECTED_PERSISTENCE_TYPE"            // You have specified a persistence type for a FILL_OR_KILL order, which is nonsensical because no unmatched portion can remain after the order has been placed.
	IRS_ERR_INVALID_ORDER_TYPE                     InstructionReportErrorCode = "INVALID_ORDER_TYPE"                     // You have specified a time in force of FILL_OR_KILL, but have included a non-LIMIT order type.
	IRS_ERR_UNEXPECTED_MIN_FILL_SIZE               InstructionReportErrorCode = "UNEXPECTED_MIN_FILL_SIZE"               // You have specified a minFillSize on a limit order, where the limit order's time in force is not FILL_OR_KILL. Using minFillSize is not supported where the time in force of the request (as opposed to an order) is FILL_OR_KILL.
	IRS_ERR_INVALID_CUSTOMER_ORDER_REF             InstructionReportErrorCode = "INVALID_CUSTOMER_ORDER_REF"             // The supplied customer order reference is too long.
	IRS_ERR_INVALID_MIN_FILL_SIZE                  InstructionReportErrorCode = "INVALID_MIN_FILL_SIZE"                  // The minFillSize must be greater than zero and less than or equal to the order's size. The minFillSize cannot be less than the minimum bet size for your currency.
	IRS_ERR_BET_LAPSED_PRICE_IMPROVEMENT_TOO_LARGE InstructionReportErrorCode = "BET_LAPSED_PRICE_IMPROVEMENT_TOO_LARGE" // Your bet is lapsed. There is better odds than requested available in the market, but your preferences don't allow the system to match your bet against better odds. Change your betting preferences to accept better odds if you don't want to receive this error. Please see https://support.betfair.com/app/answers/detail/a_id/404/ for more details regarding Best Execution and how to update your settings.
)

// ChangeType identifies the type of change in an Exchange Stream API ChangeMessage.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#ChangeType
type ChangeType string

const (
	CT_SUB_IMAGE   ChangeType = "SUB_IMAGE"   // The initial image returned from a subscription. May also happen while the subscription is ongoing and should replace local cache entirely.
	CT_RESUB_DELTA ChangeType = "RESUB_DELTA" // A patch returned from a resubscribe.
	CT_HEARTBEAT   ChangeType = "HEARTBEAT"   // An empty message published if no data has been sent within heartbeatMs.
)

// SegmentType identifies multi-part segmented Exchange Stream API messages.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#SegmentType
type SegmentType string

const (
	ST_SEG_START SegmentType = "SEG_START" // Start of a segmented message.
	ST_SEG       SegmentType = "SEG"       // Middle part of a segmented message.
	ST_SEG_END   SegmentType = "SEG_END"   // Last part of a segmented message.
)
