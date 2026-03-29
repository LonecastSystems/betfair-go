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
	IC_UNKNOWN = "UNKNOWN"
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
	AccountStatementParams struct {
		Locale        string      `json:"locale,omitempty"`
		FromRecord    int         `json:"fromRecord,omitempty"`
		RecordCount   int         `json:"recordCount,omitempty"`
		ItemDateRange TimeRange   `json:"itemDateRange,omitempty"`
		IncludeItem   IncludeItem `json:"includeItem,omitempty"`
		Wallet        Wallet      `json:"wallet,omitempty"`
	}

	AccountStatementReport struct {
		AccountStatements []AccountStatement `json:"accountStatement"`
		MoreAvailable     bool               `json:"moreAvailable"`
	}

	AccountStatement struct {
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

func (client *Client) GetAccountStatement(ctx context.Context, params AccountStatementParams) (AccountStatementReport, error) {
	json := AccountStatementReport{}

	return json, client.GetAccounts(ctx, "getAccountStatement", params, &json)
}
