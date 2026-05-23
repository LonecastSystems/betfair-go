package betfair

import (
	"math/rand/v2"
	"time"
)

type (
	OrderSubscriptionMessage struct {
		ID                  int         `json:"id"`
		Op                  string      `json:"op"`
		SegmentationEnabled bool        `json:"segmentationEnabled"`
		ConflateMs          int         `json:"conflateMs"`
		HeartbeatMs         int         `json:"heartbeatMs"`
		InitialClk          string      `json:"initialClk,omitempty"`
		Clk                 string      `json:"clk,omitempty"`
		OrderFilter         OrderFilter `json:"orderFilter,omitempty"`
	}

	OrderFilter struct {
		AccountIDs                    []int64  `json:"accountIds,omitempty"` // Internal-only, do not set this
		IncludeOverallPosition        bool     `json:"includeOverallPosition,omitempty"`
		CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`
		PartitionMatchedByStrategyRef bool     `json:"partitionMatchedByStrategyRef,omitempty"`
	}

	OrderChangeMessage struct {
		ID                 int                  `json:"id"`
		Op                 string               `json:"op"`
		ChangeType         ChangeType           `json:"ct,omitempty"`
		SegmentType        SegmentType          `json:"segmentType,omitempty"`
		ConflateMs         int                  `json:"conflateMs"`
		Status             string               `json:"status"`
		HeartbeatMs        int                  `json:"heartbeatMs"`
		PublishTime        int64                `json:"pt"`
		InitialClk         string               `json:"initialClk"`
		Clk                string               `json:"clk"`
		OrderAccountChange []OrderAccountChange `json:"oc"`
	}

	OrderAccountChange struct {
		MarketID     string        `json:"id"`
		Closed       bool          `json:"closed"`
		FullImage    bool          `json:"fullImage"`
		OrderChanges []OrderChange `json:"orc"`
	}

	OrderChange struct {
		FullImage       bool             `json:"fullImage"`    // Replace existing data at runner level with the data supplied (null if delta)
		SelectionID     int64            `json:"id"`           // Selection Id - the id of the runner (selection)
		Handicap        float64          `json:"hc,omitempty"` // Handicap - the handicap of the runner (null if not applicable)
		UnmatchedOrders []UnmatchedOrder `json:"uo"`           // Unmatched Orders on this runner
		MatchedBacks    [][]float64      `json:"mb"`           // Matched Backs - matched amounts by distinct matched price on the Back side for this runner
		MatchedLays     [][]float64      `json:"ml"`           // Matched Lays - matched amounts by distinct matched price on the Lay side for this runner
	}

	UnmatchedOrder struct {
		BetID               string    `json:"id"`             // Bet Id - the id of the order
		Price               float64   `json:"p"`              // Price - the original placed price of the order
		Size                float64   `json:"s"`              // Size - the original placed size of the order
		Bsp                 float64   `json:"bsp,omitempty"`  // BSP Liability - the BSP liability of the order (null if the order is not a BSP order)
		Side                string    `json:"side"`           // Side - the side of the order
		Status              string    `json:"status"`         // Status - the status of the order (E = EXECUTABLE, EC = EXECUTION_COMPLETE)
		PersistenceType     string    `json:"pt"`             // Persistence Type - whether the order will persist at in play or not
		OrderType           string    `json:"ot"`             // Order Type - the type of the order (L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE)
		PlacedDate          time.Time `json:"pd"`             // Placed Date - the date the order was placed
		MatchedDate         time.Time `json:"md,omitempty"`   // Matched Date - the date the order was matched (null if the order is not matched)
		CancelledDate       time.Time `json:"cd,omitempty"`   // Cancelled Date - the date the order was cancelled (null if the order is not cancelled)
		LapsedDate          time.Time `json:"ld,omitempty"`   // Lapsed Date - the date the order was lapsed (null if the order is not lapsed)
		LapseStatusReason   string    `json:"lsrc,omitempty"` // Lapse Status Reason Code - the reason for lapsed order (null if no portion of the order is lapsed)
		AveragePriceMatched float64   `json:"avp,omitempty"`  // Average Price Matched - the average price the order was matched at (null if the order is not matched)
		SizeMatched         float64   `json:"sm"`             // Size Matched - the amount of the order that has been matched
		SizeRemaining       float64   `json:"sr"`             // Size Remaining - the amount of the order that is remaining unmatched
		SizeLapsed          float64   `json:"sl"`             // Size Lapsed - the amount of the order that has been lapsed
		SizeCancelled       float64   `json:"sc"`             // Size Cancelled - the amount of the order that has been cancelled
		SizeVoided          float64   `json:"sv"`             // Size Voided - the amount of the order that has been voided
		RegulatorAuthCode   string    `json:"rac,omitempty"`  // Regulator Auth Code - the auth code returned by the regulator
		RegulatorCode       string    `json:"rc,omitempty"`   // Regulator Code - the regulator of the order
		ReferenceOrder      string    `json:"rfo,omitempty"`  // Reference Order - the customer supplied order reference
		ReferenceStrategy   string    `json:"rfs"`            // Reference Strategy - the customer-supplied strategy reference used to group orders together (default is "")
	}
)

func (client *StreamingClient) SubscribeToOrders(orderFilter OrderFilter) (chan OrderChangeMessage, error) {
	config := client.Config

	ms := OrderSubscriptionMessage{
		ID:                  int(rand.UintN(16)),
		Op:                  "orderSubscription",
		OrderFilter:         orderFilter,
		SegmentationEnabled: config.SegmentationEnabled,
		ConflateMs:          config.ConflateMs,
		HeartbeatMs:         config.HeartbeatMs,
		InitialClk:          config.InitialClk,
		Clk:                 config.Clk,
	}

	if err := client.Write(ms, true); err != nil {
		return nil, err
	}

	return ReadStream[OrderChangeMessage](client.Connection), nil
}
