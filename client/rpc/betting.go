package rpc

import "time"

type (
	MarketParams struct {
		Filter     MarketFilter `json:"filter"`
		MaxResults string       `json:"maxResults"`
	}

	MarketFilter struct {
		TextQuery          string    `json:"textQuery,omitempty"`
		EventTypeIds       []string  `json:"eventTypeIds,omitempty"`
		EventIds           []string  `json:"eventIds,omitempty"`
		CompetitionIds     []string  `json:"competitionIds,omitempty"`
		MarketIds          []string  `json:"marketIds,omitempty"`
		Venues             []string  `json:"venues,omitempty"`
		BspOnly            bool      `json:"bspOnly,omitempty"`
		TurnInPlayEnabled  bool      `json:"turnInPlayEnabled,omitempty"`
		InPlayOnly         bool      `json:"inPlayOnly,omitempty"`
		MarketBettingTypes []string  `json:"marketBettingTypes,omitempty"`
		MarketTypeCodes    []string  `json:"marketTypeCodes,omitempty"`
		MarketCountries    []string  `json:"marketCountries"`
		MarketStartTime    TimeRange `json:"marketStartTime,omitempty"`
		WithOrders         []string  `json:"withOrders,omitempty"`
		RaceTypes          []string  `json:"raceTypes,omitempty"`
	}

	TimeRange struct {
		From string `json:"from,omitempty"`
		To   string `json:"to,omitempty"`
	}
)

type EventTypeResult struct {
	EventType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"eventType"`
	MarketCount int `json:"marketCount"`
}

func (client *RpcClient) ListEventTypes(params MarketParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := GetSports(client, 1, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}

type EventResult struct {
	Event struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		CountryCode string    `json:"countryCode"`
		Timezone    string    `json:"timezone"`
		OpenDate    time.Time `json:"openDate"`
	} `json:"event"`
	MarketCount int `json:"marketCount"`
}

func (client *RpcClient) ListEvents(params MarketParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := GetSports(client, 2, "listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}

type CompetitionResult struct {
	Competition struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"competition"`
	MarketCount       int    `json:"marketCount"`
	CompetitionRegion string `json:"competitionRegion"`
}

func (client *RpcClient) ListCompetitions(params MarketParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := GetSports(client, 3, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}

type MarketTypeResult struct {
	MarketType  string `json:"marketType"`
	MarketCount int    `json:"marketCount"`
}

func (client *RpcClient) ListMarketTypes(params MarketParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := GetSports(client, 4, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}

type MarketCatalogueResult struct {
	MarketID     string  `json:"marketId"`
	MarketName   string  `json:"marketName"`
	TotalMatched float64 `json:"totalMatched"`
}

func (client *RpcClient) ListMarketCatalogue(params MarketParams) ([]MarketCatalogueResult, error) {
	json := []MarketCatalogueResult{}

	if err := GetSports(client, 5, "listMarketCatalogue", params, &json); err != nil {
		return []MarketCatalogueResult{}, err
	}

	return json, nil
}

type (
	MarketBookParams struct {
		MarketIds       []string        `json:"marketIds,omitempty"`
		PriceProjection PriceProjection `json:"priceProjection,omitempty"`
	}

	PriceProjection struct {
		PriceData []string `json:"priceData"`
	}

	MarketBookResult struct {
		MarketID              string    `json:"marketId"`
		IsMarketDataDelayed   bool      `json:"isMarketDataDelayed"`
		Status                string    `json:"status"`
		BetDelay              int       `json:"betDelay"`
		BspReconciled         bool      `json:"bspReconciled"`
		Complete              bool      `json:"complete"`
		Inplay                bool      `json:"inplay"`
		NumberOfWinners       int       `json:"numberOfWinners"`
		NumberOfRunners       int       `json:"numberOfRunners"`
		NumberOfActiveRunners int       `json:"numberOfActiveRunners"`
		LastMatchTime         time.Time `json:"lastMatchTime"`
		TotalMatched          float64   `json:"totalMatched"`
		TotalAvailable        float64   `json:"totalAvailable"`
		CrossMatching         bool      `json:"crossMatching"`
		RunnersVoidable       bool      `json:"runnersVoidable"`
		Version               int64     `json:"version"`
		Runners               []Runner  `json:"runners"`
	}

	Runner struct {
		SelectionID     int     `json:"selectionId"`
		Handicap        float64 `json:"handicap"`
		Status          string  `json:"status"`
		LastPriceTraded float64 `json:"lastPriceTraded"`
		TotalMatched    float64 `json:"totalMatched"`
		Ex              Ex      `json:"ex"`
	}

	Ex struct {
		AvailableToBack []PriceSize   `json:"availableToBack"`
		AvailableToLay  []PriceSize   `json:"availableToLay"`
		TradedVolume    []interface{} `json:"tradedVolume"`
	}

	PriceSize struct {
		Price float64 `json:"price"`
		Size  float64 `json:"size"`
	}
)

func (client *RpcClient) ListMarketBook(params MarketBookParams) ([]MarketBookResult, error) {
	json := []MarketBookResult{}

	if err := GetSports(client, 6, "listMarketBook", params, &json); err != nil {
		return []MarketBookResult{}, err
	}

	return json, nil
}

type (
	CurrentOrdersParams struct {
		FromRecord  int `json:"fromRecord"`
		RecordCount int `json:"recordCount"`
	}

	CurrentOrderSummaryReport struct {
		Orders        []CurrentOrderSummary `json:"currentOrders"`
		MoreAvailable bool                  `json:"moreAvailable"`
	}

	CurrentOrderSummary struct {
		BetID               string    `json:"betId"`
		MarketID            string    `json:"marketId"`
		SelectionID         int       `json:"selectionId"`
		Handicap            float64   `json:"handicap"`
		PriceSize           PriceSize `json:"priceSize"`
		BspLiability        float64   `json:"bspLiability"`
		Side                string    `json:"side"`
		Status              string    `json:"status"`
		PersistenceType     string    `json:"persistenceType"`
		OrderType           string    `json:"orderType"`
		PlacedDate          time.Time `json:"placedDate"`
		MatchedDate         time.Time `json:"matchedDate"`
		AveragePriceMatched float64   `json:"averagePriceMatched"`
		SizeMatched         float64   `json:"sizeMatched"`
		SizeRemaining       float64   `json:"sizeRemaining"`
		SizeLapsed          float64   `json:"sizeLapsed"`
		SizeCancelled       float64   `json:"sizeCancelled"`
		SizeVoided          float64   `json:"sizeVoided"`
		RegulatorCode       string    `json:"regulatorCode"`
	}
)

func (client *RpcClient) ListCurrentOrders(params CurrentOrdersParams) (CurrentOrderSummaryReport, error) {
	json := CurrentOrderSummaryReport{}

	if err := GetSports(client, 7, "listCurrentOrders", params, &json); err != nil {
		return CurrentOrderSummaryReport{}, err
	}

	return json, nil
}
