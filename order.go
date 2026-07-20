package betfair

import (
	"math/rand/v2"
	"time"
)

type (
	// OrderSubscriptionMessage subscribes to order changes on the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderSubscriptionMessage
	OrderSubscriptionMessage struct {
		ID                  int         `json:"id"`                    // Client generated unique id to link request with response (like json rpc).
		Op                  string      `json:"op"`                    // The operation type.
		SegmentationEnabled bool        `json:"segmentationEnabled"`   // Segmentation Enabled - allow the server to send large sets of data in segments, instead of a single block.
		ConflateMs          int         `json:"conflateMs"`            // Conflate Milliseconds - the conflation rate (looped back on initial image after validation: bounds are 0 to 120000).
		HeartbeatMs         int         `json:"heartbeatMs"`           // Heartbeat Milliseconds - the heartbeat rate (looped back on initial image after validation: bounds are 500 to 5000).
		InitialClk          string      `json:"initialClk,omitempty"`  // Token value (received in initial OrderChangeMessage) that should be passed to resume a subscription.
		Clk                 string      `json:"clk,omitempty"`         // Token value delta (received in OrderChangeMessage) that should be passed to resume a subscription.
		OrderFilter         OrderFilter `json:"orderFilter,omitempty"` // Filter for order subscription.
	}

	// OrderFilter restricts which orders are included in an order subscription.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderFilter
	OrderFilter struct {
		AccountIDs                    []int64  `json:"accountIds,omitempty"`                    // Internal use only & should not be set on your filter (your subscription is already locked to your account).
		IncludeOverallPosition        bool     `json:"includeOverallPosition,omitempty"`        // Returns overall / net position (See: OrderChange.mb / OrderChange.ml). Default=true.
		CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`          // Restricts to specified customerStrategyRefs; this will filter orders and StrategyMatchChanges accordingly (Note: overall position is not filtered).
		PartitionMatchedByStrategyRef bool     `json:"partitionMatchedByStrategyRef,omitempty"` // Returns strategy positions - these are sent in delta format as per overall position. Default=false.
	}

	// OrderChangeMessage contains order changes from the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderChangeMessage
	OrderChangeMessage struct {
		ID                 int                  `json:"id"`                    // Client generated unique id to link request with response (like json rpc).
		Op                 string               `json:"op"`                    // The operation type.
		ChangeType         ChangeType           `json:"ct,omitempty"`          // Change Type - set to indicate the type of change - if null this is a delta.
		SegmentType        SegmentType          `json:"segmentType,omitempty"` // Segment Type - if the change is split into multiple segments, this denotes the beginning and end of a change.
		ConflateMs         int                  `json:"conflateMs"`            // Conflate Milliseconds - the conflation rate (may differ from that requested if subscription is delayed).
		Status             string               `json:"status"`                // Stream status: set to null if the exchange stream data is up to date and 503 if the downstream services are experiencing latencies.
		HeartbeatMs        int                  `json:"heartbeatMs"`           // Heartbeat Milliseconds - the heartbeat rate (may differ from requested: bounds are 500 to 30000).
		PublishTime        int64                `json:"pt"`                    // Publish Time (in millis since epoch) that the changes were generated.
		InitialClk         string               `json:"initialClk"`            // Token value (non-null) should be stored and passed in an OrderSubscriptionMessage to resume subscription (in case of disconnect).
		Clk                string               `json:"clk"`                   // Token value (non-null) should be stored and passed in an OrderSubscriptionMessage to resume subscription (in case of disconnect).
		OrderAccountChange []OrderAccountChange `json:"oc"`                    // OrderMarketChanges - the modifications to account's orders (will be null on a heartbeat).
	}

	// OrderAccountChange contains order changes for a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderMarketChange
	OrderAccountChange struct {
		MarketID     string        `json:"id"`        // Market Id - the id of the market the order is on.
		Closed       bool          `json:"closed"`    // True if the market is closed.
		FullImage    bool          `json:"fullImage"` // True if this is a full image replacement.
		OrderChanges []OrderChange `json:"orc"`       // Order Changes - a list of changes to orders on a selection.
	}

	// OrderChange contains order changes for a runner.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderRunnerChange
	OrderChange struct {
		FullImage       bool             `json:"fullImage"`    // Replace existing data at runner level with the data supplied (null if delta).
		SelectionID     int64            `json:"id"`           // Selection Id - the id of the runner (selection).
		Handicap        float64          `json:"hc,omitempty"` // Handicap - the handicap of the runner (null if not applicable).
		UnmatchedOrders []UnmatchedOrder `json:"uo"`           // Unmatched Orders on this runner.
		MatchedBacks    [][]Float        `json:"mb"`           // Matched Backs - matched amounts by distinct matched price on the Back side for this runner.
		MatchedLays     [][]Float        `json:"ml"`           // Matched Lays - matched amounts by distinct matched price on the Lay side for this runner.
	}

	// UnmatchedOrder is an unmatched order on a runner from the Exchange Stream API.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#Order
	UnmatchedOrder struct {
		BetID               string    `json:"id"`             // Bet Id - the id of the order.
		Price               float64   `json:"p"`              // Price - the original placed price of the order.
		Size                float64   `json:"s"`              // Size - the original placed size of the order.
		Bsp                 float64   `json:"bsp,omitempty"`  // BSP Liability - the BSP liability of the order (null if the order is not a BSP order).
		Side                string    `json:"side"`           // Side - the side of the order.
		Status              string    `json:"status"`         // Status - the status of the order (E = EXECUTABLE, EC = EXECUTION_COMPLETE).
		PersistenceType     string    `json:"pt"`             // Persistence Type - whether the order will persist at in play or not.
		OrderType           string    `json:"ot"`             // Order Type - the type of the order (L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE).
		PlacedDate          time.Time `json:"pd"`             // Placed Date - the date the order was placed.
		MatchedDate         time.Time `json:"md,omitempty"`   // Matched Date - the date the order was matched (null if the order is not matched).
		CancelledDate       time.Time `json:"cd,omitempty"`   // Cancelled Date - the date the order was cancelled (null if the order is not cancelled).
		LapsedDate          time.Time `json:"ld,omitempty"`   // Lapsed Date - the date the order was lapsed (null if the order is not lapsed).
		LapseStatusReason   string    `json:"lsrc,omitempty"` // Lapse Status Reason Code - the reason for lapsed order (null if no portion of the order is lapsed).
		AveragePriceMatched float64   `json:"avp,omitempty"`  // Average Price Matched - the average price the order was matched at (null if the order is not matched).
		SizeMatched         float64   `json:"sm"`             // Size Matched - the amount of the order that has been matched.
		SizeRemaining       float64   `json:"sr"`             // Size Remaining - the amount of the order that is remaining unmatched.
		SizeLapsed          float64   `json:"sl"`             // Size Lapsed - the amount of the order that has been lapsed.
		SizeCancelled       float64   `json:"sc"`             // Size Cancelled - the amount of the order that has been cancelled.
		SizeVoided          float64   `json:"sv"`             // Size Voided - the amount of the order that has been voided.
		RegulatorAuthCode   string    `json:"rac,omitempty"`  // Regulator Auth Code - the auth code returned by the regulator.
		RegulatorCode       string    `json:"rc,omitempty"`   // Regulator Code - the regulator of the order.
		ReferenceOrder      string    `json:"rfo,omitempty"`  // Reference Order - the customer supplied order reference.
		ReferenceStrategy   string    `json:"rfs"`            // Reference Strategy - the customer-supplied strategy reference used to group orders together (default is "").
	}
)

// SubscribeToOrders sends an orderSubscription request to receive order changes.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#OrderSubscriptionMessage
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
