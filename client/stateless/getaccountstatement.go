package stateless

import "time"

type (
	IncludeItem string
	ItemClass   string
	MarketType  string
	WinLose     string

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

func (client *StatelessClient) GetAccountStatement(params AccountStatementParams) (AccountStatementReport, error) {
	json := AccountStatementReport{}

	if err := GetAccounts(client, "getAccountStatement", params, &json); err != nil {
		return AccountStatementReport{}, err
	}

	return json, nil
}
