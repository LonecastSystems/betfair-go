package betfair

import (
	"context"
	"time"
)

// IncludeItem specifies which items to include, if not specified then defaults to ALL.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#IncludeItem
type IncludeItem string

const (
	IT_ALL                  IncludeItem = "ALL"                  // IT_ALL includes all items.
	IT_DEPOSITS_WITHDRAWALS IncludeItem = "DEPOSITS_WITHDRAWALS" // IT_DEPOSITS_WITHDRAWALS includes payments only.
	IT_EXCHANGE             IncludeItem = "EXCHANGE"             // IT_EXCHANGE includes exchange bets only.
	IT_POKER_ROOM           IncludeItem = "POKER_ROOM"           // IT_POKER_ROOM includes poker transactions only.
)

// ItemClass is the statement item class.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#ItemClass
type ItemClass string

const (
	IC_UNKNOWN                ItemClass = "UNKNOWN"                // IC_UNKNOWN indicates a statement item not mapped to a specific class. All values will be concatenated into a single key/value pair. The key will be 'unknownStatementItem' and the value will be a comma separated string. Please note: This is used to represent commission payment items.
	IC_DEPOSIT                ItemClass = "DEPOSIT"                // IC_DEPOSIT is a deposit item.
	IC_WITHDRAWAL             ItemClass = "WITHDRAWAL"             // IC_WITHDRAWAL is a withdrawal item.
	IC_COMMISSION             ItemClass = "COMMISSION"             // IC_COMMISSION is a commission item.
	IC_SETTLED_BET            ItemClass = "SETTLED_BET"            // IC_SETTLED_BET is a settled bet item.
	IC_VOIDED_BET             ItemClass = "VOIDED_BET"             // IC_VOIDED_BET is a voided bet item.
	IC_LAPSED_BET             ItemClass = "LAPSED_BET"             // IC_LAPSED_BET is a lapsed bet item.
	IC_CANCELLED_BET          ItemClass = "CANCELLED_BET"          // IC_CANCELLED_BET is a cancelled bet item.
	IC_TRANSFER               ItemClass = "TRANSFER"               // IC_TRANSFER is a transfer item.
	IC_TAX                    ItemClass = "TAX"                    // IC_TAX is a tax item.
	IC_DISCOUNT               ItemClass = "DISCOUNT"               // IC_DISCOUNT is a discount item.
	IC_BONUS                  ItemClass = "BONUS"                  // IC_BONUS is a bonus item.
	IC_ADJUSTMENT             ItemClass = "ADJUSTMENT"             // IC_ADJUSTMENT is an adjustment item.
	IC_ADMIN                  ItemClass = "ADMIN"                  // IC_ADMIN is an admin item.
	IC_DEPOSIT_FAIL           ItemClass = "DEPOSIT_FAIL"           // IC_DEPOSIT_FAIL is a failed deposit item.
	IC_WITHDRAWAL_FAIL        ItemClass = "WITHDRAWAL_FAIL"        // IC_WITHDRAWAL_FAIL is a failed withdrawal item.
	IC_TRANSFER_TO_EXCHANGE   ItemClass = "TRANSFER_TO_EXCHANGE"   // IC_TRANSFER_TO_EXCHANGE is a transfer to exchange item.
	IC_TRANSFER_FROM_EXCHANGE ItemClass = "TRANSFER_FROM_EXCHANGE" // IC_TRANSFER_FROM_EXCHANGE is a transfer from exchange item.
)

// WinLose is the settlement outcome label for a statement item.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#winLose
type WinLose string

const (
	WL_RESULT_ERR            WinLose = "RESULT_ERR"            // WL_RESULT_ERR indicates the record has been affected by a unsettlement. There is no impact on the balance for these records, this just a label to say that these are to be corrected.
	WL_RESULT_FIX            WinLose = "RESULT_FIX"            // WL_RESULT_FIX is a correction to the balance to reverse the impact of records shown as in error. If commission has been paid on the original settlement then there will be a second FIX record to reverse the commission.
	WL_RESULT_LOST           WinLose = "RESULT_LOST"           // WL_RESULT_LOST is a loss.
	WL_RESULT_NOT_APPLICABLE WinLose = "RESULT_NOT_APPLICABLE" // WL_RESULT_NOT_APPLICABLE indicates amounts relating to commission payments.
	WL_RESULT_WON            WinLose = "RESULT_WON"            // WL_RESULT_WON indicates a win.
	WL_COMMISSION_REVERSAL   WinLose = "COMMISSION_REVERSAL"   // WL_COMMISSION_REVERSAL indicates Betfair have restored the funds to your account that it previously received from you in commission.
)

// MarketType is the market type for legacy statement data.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687907/Accounts+Enums#MarketType
type MarketType string

const (
	MT_ASIAN_HANDICAP MarketType = "A"              // MT_ASIAN_HANDICAP is an Asian Handicap market.
	MT_LINE_MARKET    MarketType = "L"              // MT_LINE_MARKET is a line market.
	MT_ODDS_MARKET    MarketType = "O"              // MT_ODDS_MARKET is an odds market.
	MT_RANGE_MARKET   MarketType = "R"              // MT_RANGE_MARKET is a range market.
	MT_NOT_APPLICABLE MarketType = "NOT_APPLICABLE" // MT_NOT_APPLICABLE indicates the market does not have an applicable marketType.
)

type (
	// GetAccountStatementParams are the parameters for getAccountStatement.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687876/getAccountStatement
	GetAccountStatementParams struct {
		Locale        string      `json:"locale,omitempty"`        // The language to be used where applicable. If not specified, the customer account default is returned.
		FromRecord    int         `json:"fromRecord,omitempty"`    // Specifies the first record that will be returned. Records start at index zero. If not specified then it will default to 0.
		RecordCount   int         `json:"recordCount,omitempty"`   // Specifies the maximum number of records to be returned. Note that there is a page size limit of 100.
		ItemDateRange TimeRange   `json:"itemDateRange,omitempty"` // Return items with an itemDate within this date range. Both from and to date times are inclusive.
		IncludeItem   IncludeItem `json:"includeItem,omitempty"`   // Which items to include, if not specified then defaults to ALL.
		Wallet        Wallet      `json:"wallet,omitempty"`        // Which wallet to return statementItems for. If unspecified then the UK wallet will be selected.
	}

	// GetAccountStatementReport is a list of statement items chronologically ordered plus moreAvailable boolean to facilitate paging.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#AccountStatementReport
	GetAccountStatementReport struct {
		AccountStatements []GetAccountStatementResult `json:"accountStatement"` // List of statement items.
		MoreAvailable     bool                        `json:"moreAvailable"`    // Indicates whether more records are available.
	}

	// GetAccountStatementResult is a single account statement item.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#StatementItem
	GetAccountStatementResult struct {
		RefID         string              `json:"refId"`         // The reference identifier for the statement item.
		ItemDate      time.Time           `json:"itemDate"`      // The date of the statement item.
		Amount        float64             `json:"amount"`        // The amount of the statement item.
		Balance       float64             `json:"balance"`       // The account balance after the statement item.
		ItemClass     ItemClass           `json:"itemClass"`     // The class of the statement item.
		ItemClassData map[string]string   `json:"itemClassData"` // Additional key/value data for the statement item.
		LegacyData    StatementLegacyData `json:"legacyData"`    // Legacy data for exchange statement items.
	}

	// StatementLegacyData contains legacy exchange bet data for a statement item.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#StatementLegacyData
	StatementLegacyData struct {
		AvgPrice             float64    `json:"avgPrice"`             // Average price matched.
		BetSize              float64    `json:"betSize"`              // Bet size.
		BetType              string     `json:"betType"`              // Bet type.
		BetCategoryType      string     `json:"betCategoryType"`      // Bet category type.
		CommissionRate       string     `json:"commissionRate"`       // Commission rate.
		EventID              float64    `json:"eventId"`              // Event id.
		EventTypeID          float64    `json:"eventTypeId"`          // Event type id.
		FullMarketName       string     `json:"fullMarketName"`       // Full market name.
		GrossBetAmount       float64    `json:"grossBetAmount"`       // Gross bet amount.
		MarketName           string     `json:"marketName"`           // Market name.
		MarketType           MarketType `json:"marketType"`           // Market type.
		PlacedDate           time.Time  `json:"placedDate"`           // Date the bet was placed.
		SelectionID          int64      `json:"selectionId"`          // Selection id.
		SelectionName        string     `json:"selectionName"`        // Selection name.
		StartDate            time.Time  `json:"startDate"`            // Event start date.
		TransactionType      string     `json:"transactionType"`      // Transaction type.
		TransactionID        float64    `json:"transactionId"`        // Transaction id.
		WinLose              WinLose    `json:"winLose"`              // Settlement outcome label.
		DeadHeatPriceDivisor float64    `json:"deadHeatPriceDivisor"` // Dead heat price divisor.
		AvgPriceRaw          float64    `json:"avgPriceRaw"`          // Raw average price matched.
	}
)

// GetAccountStatement returns account statement items chronologically ordered plus moreAvailable boolean to facilitate paging. Please note: You can only retrieve account statement items for the last 90 days.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687876/getAccountStatement
func (client *Client) GetAccountStatement(ctx context.Context, params GetAccountStatementParams) (GetAccountStatementReport, error) {
	json := GetAccountStatementReport{}

	return json, client.GetAccounts(ctx, "getAccountStatement", params, &json)
}
