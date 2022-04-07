package atomicmarket

// GetPriceSales response
type PriceSalesProps struct {
	Success   bool                  `json:"success"`
	Data      []PriceSalesDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                 `json:"query_time"`
	Message   string                `json:"message,omitempty"` // error message will be here
}

type PriceSalesDataProps struct {
	SaleID         string      `json:"sale_id"`
	AuctionID      interface{} `json:"auction_id"`
	BuyofferID     interface{} `json:"buyoffer_id"`
	Price          string      `json:"price"`
	TemplateMint   int64       `json:"template_mint"`
	TokenSymbol    string      `json:"token_symbol"`
	TokenPrecision int64       `json:"token_precision"`
	TokenContract  string      `json:"token_contract"`
	BlockTime      string      `json:"block_time"`
}

// GetPriceSalesDars response
type PriceSalesDaysProps struct {
	Success   bool                      `json:"success"`
	Data      []PriceSalesDaysDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                     `json:"query_time"`
	Message   string                    `json:"message,omitempty"` // error message will be here
}

type PriceSalesDaysDataProps struct {
	Median         string `json:"median"`
	Average        string `json:"average"`
	Sales          string `json:"sales"`
	TokenSymbol    string `json:"token_symbol"`
	TokenPrecision int64  `json:"token_precision"`
	TokenContract  string `json:"token_contract"`
	Time           int64  `json:"time"`
}

// GetPriceTemplates response
type PriceTemplates struct {
	Success   bool                      `json:"success"`
	Data      []PriceTemplatesDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                     `json:"query_time"`
	Message   string                    `json:"message,omitempty"` // error message will be here
}

type PriceTemplatesDataProps struct {
	MarketContract   string `json:"market_contract"`
	AssetsContract   string `json:"assets_contract"`
	CollectionName   string `json:"collection_name"`
	TemplateID       string `json:"template_id"`
	TokenSymbol      string `json:"token_symbol"`
	TokenContract    string `json:"token_contract"`
	TokenPrecision   int64  `json:"token_precision"`
	Median           string `json:"median"`
	Average          string `json:"average"`
	Min              string `json:"min"`
	Max              string `json:"max"`
	Sales            string `json:"sales"`
	SuggestedMedian  string `json:"suggested_median"`
	SuggestedAverage string `json:"suggested_average"`
}

// GetPriceAssets response
type PriceAssets struct {
	Success   bool                   `json:"success"`
	Data      []PriceAssetsDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                  `json:"query_time"`
	Message   string                 `json:"message,omitempty"` // error message will be here
}

type PriceAssetsDataProps struct {
	TokenSymbol      string `json:"token_symbol"`
	TokenPrecision   int64  `json:"token_precision"`
	TokenContract    string `json:"token_contract"`
	Median           string `json:"median"`
	Average          string `json:"average"`
	Min              string `json:"min"`
	Max              string `json:"max"`
	SuggestedMedian  string `json:"suggested_median"`
	SuggestedAverage string `json:"suggested_average"`
}
