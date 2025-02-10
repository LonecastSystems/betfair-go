package betfair

type BetTargetType string

const (
	BTT_BACKER_PROFIT BetTargetType = "BACKER_PROFIT"
	BTT_PAYOUT        BetTargetType = "PAYOUT"
)

type TimeInForce string

const (
	TIF_FILL_OR_KILL TimeInForce = "FILL_OR_KILL"
)

type (
	PlaceOrdersParams struct {
		MarketID            string             `json:"marketId"`
		Instructions        []PlaceInstruction `json:"instructions"`
		CustomerRef         string             `json:"customerRef,omitempty"`
		MarketVersion       MarketVersion      `json:"marketVersion,omitempty"`
		CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
		Async               bool               `json:"async,omitempty"`
	}

	PlaceInstruction struct {
		OrderType          OrderType          `json:"orderType"`
		SelectionID        int64              `json:"selectionId"`
		Handicap           float64            `json:"handicap,omitempty"`
		Side               Side               `json:"side"`
		LimitOrder         LimitOrder         `json:"limitOrder,omitempty"`
		LimitOnCloseOrder  LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
		MarketOnCloseOrder MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
		CustomerOrderRef   string             `json:"customerOrderRef,omitempty"`
	}

	LimitOrder struct {
		Size            float64         `json:"size"`
		Price           float64         `json:"price"`
		PersistenceType PersistenceType `json:"persistenceType"`
		TimeInForce     TimeInForce     `json:"timeInForce,omitempty"`
		MinFillSize     float64         `json:"minFillSize,omitempty"`
		BetTargetType   BetTargetType   `json:"betTargetType,omitempty"`
		BetTargetSize   float64         `json:"betTargetSize,omitempty"`
	}

	LimitOnCloseOrder struct {
		Liability float64 `json:"liability"`
		Price     float64 `json:"price"`
	}

	MarketOnCloseOrder struct {
		Liability float64 `json:"liability"`
	}

	PlaceExecutionReport struct {
		CustomerRef        string                   `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus    `json:"status"`
		ErrorCode          ExecutionReportErrorCode `json:"errorCode,omitempty"`
		MarketID           string                   `json:"marketId,omitempty"`
		InstructionReports []PlaceInstructionReport `json:"instructionReports,omitempty"`
	}
)

func (client *Client) PlaceOrders(params PlaceOrdersParams) (PlaceExecutionReport, error) {
	json := PlaceExecutionReport{}

	if err := client.GetSports("placeOrders", params, &json); err != nil {
		return PlaceExecutionReport{}, err
	}

	return json, nil
}
