package stateless

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

type MarketProjection string

const (
	MP_COMPETITION        MarketProjection = "COMPETITION"
	MP_EVENT              MarketProjection = "EVENT"
	MP_EVENT_TYPE         MarketProjection = "EVENT_TYPE"
	MP_MARKET_START_TIME  MarketProjection = "MARKET_START_TIME"
	MP_MARKET_DESCRIPTION MarketProjection = "MARKET_DESCRIPTION"
	MP_RUNNER_DESCRIPTION MarketProjection = "RUNNER_DESCRIPTION"
	MP_RUNNER_METADATA    MarketProjection = "RUNNER_METADATA"
)

type PriceData string

const (
	PD_SP_AVAILABLE   PriceData = "SP_AVAILABLE"
	PD_SP_TRADED      PriceData = "SP_TRADED"
	PD_EX_BEST_OFFERS PriceData = "EX_BEST_OFFERS"
	PD_EX_ALL_OFFERS  PriceData = "EX_ALL_OFFERS"
	PD_EX_TRADED      PriceData = "EX_TRADED"
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

type MarketStatus string

const (
	MS_INACTIVE  MarketStatus = "INACTIVE"
	MS_OPEN      MarketStatus = "OPEN"
	MS_SUSPENDED MarketStatus = "SUSPENDED"
	MS_CLOSED    MarketStatus = "CLOSED"
)

type RunnerStatus string

const (
	RS_ACTIVE         RunnerStatus = "ACTIVE"
	RS_WINNER         RunnerStatus = "WINNER"
	RS_LOSER          RunnerStatus = "LOSER"
	RS_PLACED         RunnerStatus = "PLACED"
	RS_REMOVED_VACANT RunnerStatus = "REMOVED_VACANT"
	RS_REMOVED        RunnerStatus = "REMOVED"
	RS_HIDDEN         RunnerStatus = "HIDDEN"
)

type TimeGranularity string

const (
	TG_DAYS    TimeGranularity = "DAYS"
	TG_HOURS   TimeGranularity = "HOURS"
	TG_MINUTES TimeGranularity = "MINUTES"
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

type OrderBy string

const (
	OB_BY_BET          OrderBy = "BY_BET"
	OB_BY_MARKET       OrderBy = "BY_MARKET"
	OB_BY_MATCH_TIME   OrderBy = "BY_MATCH_TIME"
	OB_BY_PLACE_TIME   OrderBy = "BY_PLACE_TIME"
	OB_BY_SETTLED_TIME OrderBy = "BY_SETTLED_TIME"
	OB_BY_VOID_TIME    OrderBy = "BY_VOID_TIME"
)

type SortDir string

const (
	SD_EARLIEST_TO_LATEST SortDir = "EARLIEST_TO_LATEST"
	SD_LATEST_TO_EARLIEST SortDir = "LATEST_TO_EARLIEST"
)

type OrderType string

const (
	OT_LIMIT           OrderType = "LIMIT"
	OT_LIMIT_ON_CLOSE  OrderType = "LIMIT_ON_CLOSE"
	OT_MARKET_ON_CLOSE OrderType = "MARKET_ON_CLOSE"
)

type MarketSort string

const (
	MS_MINIMUM_TRADED    MarketSort = "MINIMUM_TRADED"
	MS_MAXIMUM_TRADED    MarketSort = "MAXIMUM_TRADED"
	MS_MINIMUM_AVAILABLE MarketSort = "MINIMUM_AVAILABLE"
	MS_MAXIMUM_AVAILABLE MarketSort = "MAXIMUM_AVAILABLE"
	MS_FIRST_TO_START    MarketSort = "FIRST_TO_START"
	MS_LAST_TO_START     MarketSort = "LAST_TO_START"
)

type MarketBettingType string

const (
	MBT_ODDS       MarketBettingType = "ODDS"
	MBT_LINE       MarketBettingType = "LINE"
	MBT_RANGE      MarketBettingType = "RANGE"
	MBT_ASIAN      MarketBettingType = "ASIAN_HANDICAP_DOUBLE_LINE"
	MBT_FIXED      MarketBettingType = "FIXED_ODDS"
	MBT_FIXED_ODDS MarketBettingType = "FIXED_ODDS"
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
	IRS_ERR_BET_LAPSED                      InstructionReportErrorCode = "BET_LAPSED"
)

type GroupBy string

const (
	GB_EVENT_TYPE GroupBy = "EVENT_TYPE"
	GB_EVENT      GroupBy = "EVENT"
	GB_MARKET     GroupBy = "MARKET"
	GB_SIDE       GroupBy = "SIDE"
	GB_BET        GroupBy = "BET"
)

type BetStatus string

const (
	BS_SETTLED   BetStatus = "SETTLED"
	BS_VOIDED    BetStatus = "VOIDED"
	BS_LAPSED    BetStatus = "LAPSED"
	BS_CANCELLED BetStatus = "CANCELLED"
)

type TimeInForce string

const (
	TIF_FILL_OR_KILL TimeInForce = "FILL_OR_KILL"
)

type BetTargetType string

const (
	BTT_BACKER_PROFIT BetTargetType = "BACKER_PROFIT"
	BTT_PAYOUT        BetTargetType = "PAYOUT"
)

type PriceLadderType string

const (
	PLT_CLASSIC    PriceLadderType = "CLASSIC"
	PLT_FINEST     PriceLadderType = "FINEST"
	PLT_LINE_RANGE PriceLadderType = "LINE_RANGE"
)
