package betfair

import (
	"context"
	"time"
)

type IncludeItem string

const (
	IT_ALL                  = "ALL"
	IT_DEPOSITS_WITHDRAWALS = "DEPOSITS_WITHDRAWALS"
	IT_EXCHANGE             = "EXCHANGE"
	IT_POKER_ROOM           = "POKER_ROOM"
)

type ItemClass string

const (
	IC_UNKNOWN                ItemClass = "UNKNOWN"
	IC_DEPOSIT                ItemClass = "DEPOSIT"
	IC_WITHDRAWAL             ItemClass = "WITHDRAWAL"
	IC_COMMISSION             ItemClass = "COMMISSION"
	IC_SETTLED_BET            ItemClass = "SETTLED_BET"
	IC_VOIDED_BET             ItemClass = "VOIDED_BET"
	IC_LAPSED_BET             ItemClass = "LAPSED_BET"
	IC_CANCELLED_BET          ItemClass = "CANCELLED_BET"
	IC_TRANSFER               ItemClass = "TRANSFER"
	IC_TAX                    ItemClass = "TAX"
	IC_DISCOUNT               ItemClass = "DISCOUNT"
	IC_BONUS                  ItemClass = "BONUS"
	IC_ADJUSTMENT             ItemClass = "ADJUSTMENT"
	IC_ADMIN                  ItemClass = "ADMIN"
	IC_DEPOSIT_FAIL           ItemClass = "DEPOSIT_FAIL"
	IC_WITHDRAWAL_FAIL        ItemClass = "WITHDRAWAL_FAIL"
	IC_TRANSFER_TO_EXCHANGE   ItemClass = "TRANSFER_TO_EXCHANGE"
	IC_TRANSFER_FROM_EXCHANGE ItemClass = "TRANSFER_FROM_EXCHANGE"
)

type WinLose string

const (
	WL_RESULT_ERR            WinLose = "RESULT_ERR"
	WL_RESULT_FIX            WinLose = "RESULT_FIX"
	WL_RESULT_LOST           WinLose = "RESULT_LOST"
	WL_RESULT_NOT_APPLICABLE WinLose = "RESULT_NOT_APPLICABLE"
	WL_RESULT_WON            WinLose = "RESULT_WON"
	WL_COMMISSION_REVERSAL   WinLose = "COMMISSION_REVERSAL"
)

type MarketType string

const (
	MT_ASIAN_HANDICAP MarketType = "A"
	MT_LINE_MARKET    MarketType = "L"
	MT_ODDS_MARKET    MarketType = "O"
	MT_RANGE_MARKET   MarketType = "R"
	MT_NOT_APPLICABLE MarketType = "NOT_APPLICABLE"
)

type (
	GetAccountStatementParams struct {
		Locale        string      `json:"locale,omitempty"`
		FromRecord    int         `json:"fromRecord,omitempty"`
		RecordCount   int         `json:"recordCount,omitempty"`
		ItemDateRange TimeRange   `json:"itemDateRange,omitempty"`
		IncludeItem   IncludeItem `json:"includeItem,omitempty"`
		Wallet        Wallet      `json:"wallet,omitempty"`
	}

	GetAccountStatementReport struct {
		AccountStatements []GetAccountStatementResult `json:"accountStatement"`
		MoreAvailable     bool                        `json:"moreAvailable"`
	}

	GetAccountStatementResult struct {
		RefID         string              `json:"refId"`
		ItemDate      time.Time           `json:"itemDate"`
		Amount        float64             `json:"amount"`
		Balance       float64             `json:"balance"`
		ItemClass     ItemClass           `json:"itemClass"`
		ItemClassData map[string]string   `json:"itemClassData"`
		LegacyData    StatementLegacyData `json:"legacyData"`
	}

	StatementLegacyData struct {
		AvgPrice             float64    `json:"avgPrice"`
		BetSize              float64    `json:"betSize"`
		BetType              string     `json:"betType"`
		BetCategoryType      string     `json:"betCategoryType"`
		CommissionRate       string     `json:"commissionRate"`
		EventID              float64    `json:"eventId"`
		EventTypeID          float64    `json:"eventTypeId"`
		FullMarketName       string     `json:"fullMarketName"`
		GrossBetAmount       float64    `json:"grossBetAmount"`
		MarketName           string     `json:"marketName"`
		MarketType           MarketType `json:"marketType"`
		PlacedDate           time.Time  `json:"placedDate"`
		SelectionID          int64      `json:"selectionId"`
		SelectionName        string     `json:"selectionName"`
		StartDate            time.Time  `json:"startDate"`
		TransactionType      string     `json:"transactionType"`
		TransactionID        float64    `json:"transactionId"`
		WinLose              WinLose    `json:"winLose"`
		DeadHeatPriceDivisor float64    `json:"deadHeatPriceDivisor"`
		AvgPriceRaw          float64    `json:"avgPriceRaw"`
	}
)

func (client *Client) GetAccountStatement(ctx context.Context, params GetAccountStatementParams) (GetAccountStatementReport, error) {
	json := GetAccountStatementReport{}

	return json, client.GetAccounts(ctx, "getAccountStatement", params, &json)
}
