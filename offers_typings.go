package atomicmarket

// GetOffers response
type OffersProps struct {
	Success   bool              `json:"success"`
	Data      []OffersDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64             `json:"query_time"`
	Message   string            `json:"message,omitempty"` // error message will be here
}

// GetOfferID response
type OfferIDProps struct {
	Success   bool            `json:"success"`
	Data      OffersDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64           `json:"query_time"`
	Message   string          `json:"message,omitempty"` // error message will be here
}

// GetOfferIDLogs
type OfferIDLogsProps AssetIDLogsProps

type OffersDataProps struct {
	Contract            string        `json:"contract"`
	OfferID             string        `json:"offer_id"`
	SenderName          string        `json:"sender_name"`
	RecipientName       string        `json:"recipient_name"`
	Memo                string        `json:"memo"`
	State               int64         `json:"state"`
	SenderAssets        []SenderAsset `json:"sender_assets"`
	RecipientAssets     []interface{} `json:"recipient_assets"`
	IsSenderContract    bool          `json:"is_sender_contract"`
	IsRecipientContract bool          `json:"is_recipient_contract"`
	UpdatedAtBlock      string        `json:"updated_at_block"`
	UpdatedAtTime       string        `json:"updated_at_time"`
	CreatedAtBlock      string        `json:"created_at_block"`
	CreatedAtTime       string        `json:"created_at_time"`
}

type SenderAsset struct {
	Contract           string        `json:"contract"`
	AssetID            string        `json:"asset_id"`
	Owner              string        `json:"owner"`
	IsTransferable     bool          `json:"is_transferable"`
	IsBurnable         bool          `json:"is_burnable"`
	Collection         Collection    `json:"collection"`
	Schema             Schema        `json:"schema"`
	Template           Template      `json:"template"`
	MutableData        Data          `json:"mutable_data"`
	ImmutableData      Data          `json:"immutable_data"`
	TemplateMint       string        `json:"template_mint"`
	BackedTokens       []interface{} `json:"backed_tokens"`
	BurnedByAccount    interface{}   `json:"burned_by_account"`
	BurnedAtBlock      interface{}   `json:"burned_at_block"`
	BurnedAtTime       interface{}   `json:"burned_at_time"`
	UpdatedAtBlock     string        `json:"updated_at_block"`
	UpdatedAtTime      string        `json:"updated_at_time"`
	TransferredAtBlock string        `json:"transferred_at_block"`
	TransferredAtTime  string        `json:"transferred_at_time"`
	MintedAtBlock      string        `json:"minted_at_block"`
	MintedAtTime       string        `json:"minted_at_time"`
	Sales              []interface{} `json:"sales"`
	Auctions           []interface{} `json:"auctions"`
	Prices             []Price       `json:"prices"`
	Data               Data          `json:"data"`
	Name               string        `json:"name"`
}

type Data map[string]interface{}

type Price struct {
	MarketContract   string `json:"market_contract"`
	Token            Token  `json:"token"`
	Median           string `json:"median"`
	Average          string `json:"average"`
	SuggestedMedian  string `json:"suggested_median"`
	SuggestedAverage string `json:"suggested_average"`
	Min              string `json:"min"`
	Max              string `json:"max"`
	Sales            string `json:"sales"`
}

type Token struct {
	TokenSymbol    string `json:"token_symbol"`
	TokenPrecision int64  `json:"token_precision"`
	TokenContract  string `json:"token_contract"`
}
