package betfairgo

import (
	"time"
)

type PriceLadderType string

const (
	PLT_CLASSIC    PriceLadderType = "CLASSIC"
	PLT_FINEST     PriceLadderType = "FINEST"
	PLT_LINE_RANGE PriceLadderType = "LINE_RANGE"
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

type (
	MarketCatalogueParams struct {
		Filter           MarketFilter     `json:"filter"`
		MarketProjection MarketProjection `json:"marketProjection,omitempty"`
		Sort             MarketSort       `json:"sort,omitempty"`
		MaxResults       int              `json:"maxResults"`
		Locale           string           `json:"locale,omitempty"`
	}

	MarketCatalogueResult struct {
		MarketID          string            `json:"marketId"`
		MarketName        string            `json:"marketName"`
		MarketStartTime   time.Time         `json:"marketStartTime,omitempty"`
		MarketDescription MarketDescription `json:"description,omitempty"`
		TotalMatched      float64           `json:"totalMatched,omitempty"`
		Runners           []RunnerCatalog   `json:"runners,omitempty"`
		EventType         EventType         `json:"eventType,omitempty"`
		Competition       Competition       `json:"competition,omitempty"`
		Event             Event             `json:"event,omitempty"`
	}

	MarketDescription struct {
		PersistenceEnabled     bool                   `json:"persistenceEnabled"`
		BspMarket              bool                   `json:"bspMarket"`
		MarketTime             time.Time              `json:"marketTime"`
		SuspendTime            time.Time              `json:"suspendTime"`
		SettleTime             time.Time              `json:"settleTime,omitempty"`
		BettingType            MarketBettingType      `json:"bettingType"`
		TurnInPlayEnabled      bool                   `json:"turnInPlayEnabled"`
		MarketType             string                 `json:"marketType"`
		Regulator              string                 `json:"regulator"`
		MarketBaseRate         float64                `json:"marketBaseRate"`
		DiscountAllowed        bool                   `json:"discountAllowed"`
		Wallet                 string                 `json:"wallet,omitempty"`
		Rules                  string                 `json:"rules,omitempty"`
		RulesHasDate           bool                   `json:"rulesHasDate,omitempty"`
		EachWayDivisor         float64                `json:"eachWayDivisor,omitempty"`
		Clarifications         string                 `json:"clarifications,omitempty"`
		LineRangeInfo          MarketLineRangeInfo    `json:"lineRangeInfo,omitempty"`
		RaceType               string                 `json:"raceType,omitempty"`
		PriceLadderDescription PriceLadderDescription `json:"priceLadderDescription,omitempty"`
	}

	RunnerCatalog struct {
		SelectionID  int64             `json:"selectionId"`
		RunnerName   string            `json:"runnerName"`
		Handicap     float64           `json:"handicap"`
		SortPriority int               `json:"sortPriority"`
		Metadata     map[string]string `json:"metadata,omitempty"`
	}

	MarketLineRangeInfo struct {
		MaxUnitValue float64 `json:"maxUnitValue"`
		MinUnitValue float64 `json:"minUnitValue"`
		Interval     float64 `json:"interval"`
		MarketUnit   string  `json:"marketUnit"`
	}

	PriceLadderDescription struct {
		Type PriceLadderType `json:"priceLadderType,omitempty"`
	}
)

func (client *BetfairClient) ListMarketCatalogue(params MarketCatalogueParams) ([]MarketCatalogueResult, error) {
	json := []MarketCatalogueResult{}

	if err := client.GetSports("listMarketCatalogue", params, &json); err != nil {
		return []MarketCatalogueResult{}, err
	}

	return json, nil
}
