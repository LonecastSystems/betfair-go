package betfairgo

import (
	"time"
)

type RollupModel string

const (
	RM_STAKE  = "STAKE"
	RM_PAYOUT = "PAYOUT"
	RM_NONE   = "NONE"
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

type MarketStatus string

const (
	MS_INACTIVE  MarketStatus = "INACTIVE"
	MS_OPEN      MarketStatus = "OPEN"
	MS_SUSPENDED MarketStatus = "SUSPENDED"
	MS_CLOSED    MarketStatus = "CLOSED"
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
	MarketBookParams struct {
		MarketIDs                     []string        `json:"marketIds"`
		PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
		OrderProjection               OrderProjection `json:"orderProjection,omitempty"`
		MatchProjection               MatchProjection `json:"matchProjection,omitempty"`
		IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
		PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
		CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
		CurrencyCode                  string          `json:"currencyCode,omitempty"`
		Locale                        string          `json:"locale,omitempty"`
		MatchedSince                  time.Time       `json:"matchedSince,omitempty"`
		BetIDs                        []string        `json:"betIds,omitempty"`
	}

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

func (client *BetfairClient) ListMarketBook(params MarketBookParams) ([]MarketBook, error) {
	json := []MarketBook{}

	if err := client.GetSports("listMarketBook", params, &json); err != nil {
		return []MarketBook{}, err
	}

	return json, nil
}
