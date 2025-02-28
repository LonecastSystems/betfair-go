package betfair

import "time"

type MarketBettingType string

const (
	MBT_ODDS                       MarketBettingType = "ODDS"
	MBT_LINE                       MarketBettingType = "LINE"
	MBT_RANGE                      MarketBettingType = "RANGE"
	MBT_ASIAN_HANDICAP_SINGLE_LINE MarketBettingType = "ASIAN_HANDICAP_SINGLE_LINE"
	MBT_ASIAN_HANDICAP_DOUBLE_LINE MarketBettingType = "ASIAN_HANDICAP_DOUBLE_LINE"
	MBT_FIXED_ODDS                 MarketBettingType = "FIXED_ODDS"
)

type (
	MarketFilter struct {
		TextQuery          string              `json:"textQuery,omitempty"`
		EventTypeIDs       []string            `json:"eventTypeIds,omitempty"`
		EventIDs           []string            `json:"eventIds,omitempty"`
		CompetitionIDs     []string            `json:"competitionIds,omitempty"`
		MarketIDs          []string            `json:"marketIds,omitempty"`
		Venues             []string            `json:"venues,omitempty"`
		BspOnly            bool                `json:"bspOnly,omitempty"`
		TurnInPlayEnabled  bool                `json:"turnInPlayEnabled,omitempty"`
		InPlayOnly         bool                `json:"inPlayOnly,omitempty"`
		MarketBettingTypes []MarketBettingType `json:"marketBettingTypes,omitempty"`
		MarketTypeCodes    []string            `json:"marketTypeCodes,omitempty"`
		MarketCountries    []string            `json:"marketCountries"`
		MarketStartTime    TimeRange           `json:"marketStartTime,omitempty"`
		WithOrders         []OrderStatus       `json:"withOrders,omitempty"`
		RaceTypes          []string            `json:"raceTypes,omitempty"`
	}

	TimeRange struct {
		From string `json:"from,omitempty"`
		To   string `json:"to,omitempty"`
	}
)

type Wallet string

const (
	W_UK = "UK"
)

type MatchProjection string

const (
	MP_NO_ROLLUP              MatchProjection = "NO_ROLLUP"
	MP_ROLLED_UP              MatchProjection = "ROLLED_UP_BY_PRICE"
	MP_ROLLED_UP_BY_AVG_PRICE MatchProjection = "ROLLED_UP_BY_AVG_PRICE"
)

type OrderProjection string

const (
	OP_ALL                OrderProjection = "ALL"
	OP_EXECUTABLE         OrderProjection = "EXECUTABLE"
	OP_EXECUTION_COMPLETE OrderProjection = "EXECUTION_COMPLETE"
)

type PriceData string

const (
	PD_SP_AVAILABLE   PriceData = "SP_AVAILABLE"
	PD_SP_TRADED      PriceData = "SP_TRADED"
	PD_EX_BEST_OFFERS PriceData = "EX_BEST_OFFERS"
	PD_EX_ALL_OFFERS  PriceData = "EX_ALL_OFFERS"
	PD_EX_TRADED      PriceData = "EX_TRADED"
)

type (
	PriceProjection struct {
		PriceData             []PriceData     `json:"priceData"`
		ExBestOffersOverrides OffersOverrides `json:"exBestOffersOverrides,omitempty"`
		Virtualise            bool            `json:"virtualise,omitempty"`
		RolloverStakes        bool            `json:"rolloverStakes,omitempty"`
	}

	OffersOverrides struct {
		BestPricesDepth       int         `json:"bestPricesDepth,omitempty"`
		RollupModel           RollupModel `json:"rollupModel,omitempty"`
		RollupLimit           int         `json:"rollupLimit,omitempty"`
		RollupLiability       float64     `json:"rollupLiability,omitempty"`
		RollupLiabilityFactor int         `json:"rollupLiabilityFactor,omitempty"`
	}
)

type (
	MarketBook struct {
		MarketID              string             `json:"marketId"`
		IsMarketDataDelayed   bool               `json:"isMarketDataDelayed"`
		Status                MarketStatus       `json:"status,omitempty"`
		BetDelay              int                `json:"betDelay,omitempty"`
		BspReconciled         bool               `json:"bspReconciled,omitempty"`
		Complete              bool               `json:"complete,omitempty"`
		Inplay                bool               `json:"inplay,omitempty"`
		NumberOfWinners       int                `json:"numberOfWinners,omitempty"`
		NumberOfRunners       int                `json:"numberOfRunners,omitempty"`
		NumberOfActiveRunners int                `json:"numberOfActiveRunners,omitempty"`
		LastMatchTime         time.Time          `json:"lastMatchTime,omitempty"`
		TotalMatched          float64            `json:"totalMatched,omitempty"`
		TotalAvailable        float64            `json:"totalAvailable,omitempty"`
		CrossMatching         bool               `json:"crossMatching,omitempty"`
		RunnersVoidable       bool               `json:"runnersVoidable,omitempty"`
		Version               int64              `json:"version,omitempty"`
		Runners               []Runner           `json:"runners,omitempty"`
		KeyLineDescription    KeyLineDescription `json:"keyLineDescription,omitempty"`
	}

	Runner struct {
		SelectionID       int64              `json:"selectionId"`
		Handicap          float64            `json:"handicap"`
		Status            RunnerStatus       `json:"status"`
		AdjustmentFactor  float64            `json:"adjustmentFactor,omitempty"`
		LastPriceTraded   float64            `json:"lastPriceTraded,omitempty"`
		TotalMatched      float64            `json:"totalMatched,omitempty"`
		RemovalDate       time.Time          `json:"removalDate,omitempty"`
		StartingPrices    StartingPrices     `json:"sp,omitempty"`
		ExchangePrices    ExchangePrices     `json:"ex,omitempty"`
		Orders            []Order            `json:"orders,omitempty"`
		Matches           []Match            `json:"matches,omitempty"`
		MatchesByStrategy map[string][]Match `json:"matchesByStrategy,omitempty"`
	}

	StartingPrices struct {
		NearPrice         float64     `json:"nearPrice,omitempty"`
		FarPrice          float64     `json:"farPrice,omitempty"`
		BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
		LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
		ActualSP          float64     `json:"actualSP,omitempty"`
	}

	ExchangePrices struct {
		AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
		AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
		TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
	}

	PriceSize struct {
		Price float64 `json:"price"`
		Size  float64 `json:"size"`
	}

	Order struct {
		BetID               string          `json:"betId"`
		OrderType           OrderType       `json:"orderType"`
		Status              OrderStatus     `json:"status"`
		PersistenceType     PersistenceType `json:"persistenceType"`
		Side                Side            `json:"side"`
		Price               float64         `json:"price"`
		Size                float64         `json:"size"`
		BspLiability        float64         `json:"bspLiability"`
		PlacedDate          time.Time       `json:"placedDate"`
		AvgPriceMatched     float64         `json:"avgPriceMatched,omitempty"`
		SizeMatched         float64         `json:"sizeMatched,omitempty"`
		SizeRemaining       float64         `json:"sizeRemaining,omitempty"`
		SizeLapsed          float64         `json:"sizeLapsed,omitempty"`
		SizeCancelled       float64         `json:"sizeCancelled,omitempty"`
		SizeVoided          float64         `json:"sizeVoided,omitempty"`
		CustomerOrderRef    string          `json:"customerOrderRef,omitempty"`
		CustomerStrategyRef string          `json:"customerStrategyRef,omitempty"`
	}

	Match struct {
		BetID     string    `json:"betId,omitempty"`
		MatchID   string    `json:"matchId,omitempty"`
		Side      Side      `json:"side"`
		Price     float64   `json:"price"`
		Size      float64   `json:"size"`
		MatchDate time.Time `json:"matchDate,omitempty"`
	}

	KeyLineDescription struct {
		KeyLine []KeyLineSelection `json:"keyLine,omitempty"`
	}

	KeyLineSelection struct {
		SelectionID int64   `json:"selectionId,omitempty"`
		Handicap    float64 `json:"handicap,omitempty"`
	}
)

type Competition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MarketVersion struct {
	Version int64 `json:"version,omitempty"`
}

type (
	Event struct {
		ID          string    `json:"id,omitempty"`
		Name        string    `json:"name,omitempty"`
		CountryCode string    `json:"countryCode,omitempty"`
		Timezone    string    `json:"timezone,omitempty"`
		Venue       string    `json:"venue,omitempty"`
		OpenDate    time.Time `json:"openDate,omitempty"`
	}

	EventType struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
)

type (
	CancelInstructionReport struct {
		Status        InstructionReportStatus    `json:"status"`
		ErrorCode     InstructionReportErrorCode `json:"errorCode,omitempty"`
		Instruction   CancelInstruction          `json:"instruction,omitempty"`
		SizeCancelled float64                    `json:"sizeCancelled"`
		CanceledDate  time.Time                  `json:"canceledDate,omitempty"`
	}

	PlaceInstructionReport struct {
		Status              InstructionReportStatus    `json:"status"`
		ErrorCode           InstructionReportErrorCode `json:"errorCode,omitempty"`
		OrderStatus         OrderStatus                `json:"orderStatus,omitempty"`
		Instruction         PlaceInstruction           `json:"instruction"`
		BetID               string                     `json:"betId,omitempty"`
		PlacedDate          time.Time                  `json:"placedDate,omitempty"`
		AveragePriceMatched float64                    `json:"avgPrice,omitempty"`
		SizeMatched         float64                    `json:"sizeMatched,omitempty"`
	}
)

type Side string

const (
	SD_BACK Side = "BACK"
	SD_LAY  Side = "LAY"
)

type OrderStatus string

const (
	OS_PENDING            OrderStatus = "PENDING"
	OS_EXECUTION_COMPLETE OrderStatus = "EXECUTION_COMPLETE"
	OS_EXECUTABLE         OrderStatus = "EXECUTABLE"
	OS_EXPIRED            OrderStatus = "EXPIRED"
)

type OrderType string

const (
	OT_LIMIT           OrderType = "LIMIT"
	OT_LIMIT_ON_CLOSE  OrderType = "LIMIT_ON_CLOSE"
	OT_MARKET_ON_CLOSE OrderType = "MARKET_ON_CLOSE"
)

type ExecutionReportStatus string

const (
	ERS_SUCCESS               ExecutionReportStatus = "SUCCESS"
	ERS_FAILURE               ExecutionReportStatus = "FAILURE"
	ERS_PROCESSED_WITH_ERRORS ExecutionReportStatus = "PROCESSED_WITH_ERRORS"
	ERS_TIMEOUT               ExecutionReportStatus = "TIMEOUT"
)

type ExecutionReportErrorCode string

const (
	ERS_ERR_ERROR_IN_MATCHER                    ExecutionReportErrorCode = "ERROR_IN_MATCHER"
	ERS_ERR_PROCESSED_WITH_ERRORS               ExecutionReportErrorCode = "PROCESSED_WITH_ERRORS"
	ERS_ERR_BET_ACTION_ERROR                    ExecutionReportErrorCode = "BET_ACTION_ERROR"
	ERS_ERR_INVALID_ACCOUNT_STATE               ExecutionReportErrorCode = "INVALID_ACCOUNT_STATE"
	ERS_ERR_INVALID_WALLET_STATUS               ExecutionReportErrorCode = "INVALID_WALLET_STATUS"
	ERS_ERR_INSUFFICIENT_FUNDS                  ExecutionReportErrorCode = "INSUFFICIENT_FUNDS"
	ERS_ERR_LOSS_LIMIT_EXCEEDED                 ExecutionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	ERS_ERR_MARKET_SUSPENDED                    ExecutionReportErrorCode = "MARKET_SUSPENDED"
	ERS_ERR_EXE_ERR_MARKET_NOT_OPEN_FOR_BETTING ExecutionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	ERS_ERR_DUPLICATE_TRANSACTION               ExecutionReportErrorCode = "DUPLICATE_TRANSACTION"
	ERS_ERR_INVALID_ORDER                       ExecutionReportErrorCode = "INVALID_ORDER"
	ERS_ERR_INVALID_MARKET_ID                   ExecutionReportErrorCode = "INVALID_MARKET_ID"
	ERS_ERR_PERMISSION_DENIED                   ExecutionReportErrorCode = "PERMISSION_DENIED"
	ERS_ERR_DUPLICATE_BETIDS                    ExecutionReportErrorCode = "DUPLICATE_BETIDS"
	ERS_ERR_NO_ACTION_REQUIRED                  ExecutionReportErrorCode = "NO_ACTION_REQUIRED"
	ERS_ERR_SERVICE_UNAVAILABLE                 ExecutionReportErrorCode = "SERVICE_UNAVAILABLE"
	ERS_ERR_REJECTED_BY_REGULATOR               ExecutionReportErrorCode = "REJECTED_BY_REGULATOR"
	ERS_ERR_NO_CHASING                          ExecutionReportErrorCode = "NO_CHASING"
	ERS_ERR_REGULATOR_IS_NOT_AVAILABLE          ExecutionReportErrorCode = "REGULATOR_IS_NOT_AVAILABLE"
	ERS_ERR_TOO_MANY_INSTRUCTIONS               ExecutionReportErrorCode = "TOO_MANY_INSTRUCTIONS"
	ERS_ERR_INVALID_MARKET_VERSION              ExecutionReportErrorCode = "INVALID_MARKET_VERSION"
	ERS_ERR_INVALID_PROFIT_RATIO                ExecutionReportErrorCode = "INVALID_PROFIT_RATIO"
)

type PersistenceType string

const (
	PT_LAPSE           PersistenceType = "LAPSE"
	PT_PERSIST         PersistenceType = "PERSIST"
	PT_MARKET_ON_CLOSE PersistenceType = "MARKET_ON_CLOSE"
)

type InstructionReportStatus string

const (
	IRS_SUCCESS InstructionReportStatus = "SUCCESS"
	IRS_FAILURE InstructionReportStatus = "FAILURE"
	IRS_TIMEOUT InstructionReportStatus = "TIMEOUT"
)

type InstructionReportErrorCode string

const (
	IRS_ERR_INVALID_BET_SIZE                InstructionReportErrorCode = "INVALID_BET_SIZE"
	IRS_ERR_INVALID_RUNNER                  InstructionReportErrorCode = "INVALID_RUNNER"
	IRS_ERR_BET_TAKEN_OR_LAPSED             InstructionReportErrorCode = "BET_TAKEN_OR_LAPSED"
	IRS_ERR_BET_IN_PROGRESS                 InstructionReportErrorCode = "BET_IN_PROGRESS"
	IRS_ERR_RUNNER_REMOVED                  InstructionReportErrorCode = "RUNNER_REMOVED"
	IRS_ERR_MARKET_NOT_OPEN_FOR_BETTING     InstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	IRS_ERR_LOSS_LIMIT_EXCEEDED             InstructionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	IRS_ERR_MARKET_NOT_OPEN_FOR_BSP_BETTING InstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BSP_BETTING"
	IRS_ERR_INVALID_PRICE_EDIT              InstructionReportErrorCode = "INVALID_PRICE_EDIT"
	IRS_ERR_INVALID_ODDS                    InstructionReportErrorCode = "INVALID_ODDS"
	IRS_ERR_INSUFFICIENT_FUNDS              InstructionReportErrorCode = "INSUFFICIENT_FUNDS"
	IRS_ERR_INVALID_PERSISTENCE_TYPE        InstructionReportErrorCode = "INVALID_PERSISTENCE_TYPE"
	IRS_ERR_ERROR_IN_MATCHER                InstructionReportErrorCode = "ERROR_IN_MATCHER"
	IRS_ERR_INVALID_BACK_LAY_COMBINATION    InstructionReportErrorCode = "INVALID_BACK_LAY_COMBINATION"
	IRS_ERR_ERROR_IN_ORDER                  InstructionReportErrorCode = "ERROR_IN_ORDER"
	IRS_ERR_INVALID_BID_TYPE                InstructionReportErrorCode = "INVALID_BID_TYPE"
	IRS_ERR_INVALID_BET_ID                  InstructionReportErrorCode = "INVALID_BET_ID"
	IRS_ERR_CANCELLED_NOT_PLACED            InstructionReportErrorCode = "CANCELLED_NOT_PLACED"
	IRS_ERR_RELATED_ACTION_FAILED           InstructionReportErrorCode = "RELATED_ACTION_FAILED"
	IRS_ERR_NO_ACTION_REQUIRED              InstructionReportErrorCode = "NO_ACTION_REQUIRED"
	IRS_ERR_TIME_IN_FORCE_CONFLICT          InstructionReportErrorCode = "TIME_IN_FORCE_CONFLICT"
	IRS_ERR_UNEXPECTED_PERSISTENCE_TYPE     InstructionReportErrorCode = "UNEXPECTED_PERSISTENCE_TYPE"
	IRS_ERR_INVALID_ORDER_TYPE              InstructionReportErrorCode = "INVALID_ORDER_TYPE"
	IRS_ERR_UNEXPECTED_MIN_FILL_SIZE        InstructionReportErrorCode = "UNEXPECTED_MIN_FILL_SIZE"
	IRS_ERR_INVALID_CUSTOMER_ORDER_REF      InstructionReportErrorCode = "INVALID_CUSTOMER_ORDER_REF"
	IRS_ERR_INVALID_MIN_FILL_SIZE           InstructionReportErrorCode = "BET_LAPSED_PRICE_IMPROVEMENT_TOO_LARGE"
)
