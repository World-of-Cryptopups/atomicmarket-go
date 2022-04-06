package atomicmarket

// GetTransfers response
type TransfersProps struct {
	Success   bool                 `json:"success"`
	Data      []TransfersDataProps `json:"data,omitempty"` // if error from api, this is omitted
	QueryTime int64                `json:"query_time"`
	Message   string               `json:"message,omitempty"` // error message will be here
}

type TransfersDataProps struct {
	TransferID     string  `json:"transfer_id"`
	Contract       string  `json:"contract"`
	SenderName     string  `json:"sender_name"`
	RecipientName  string  `json:"recipient_name"`
	Memo           string  `json:"memo"`
	Txid           string  `json:"txid"`
	Assets         []Asset `json:"assets"`
	CreatedAtBlock string  `json:"created_at_block"`
	CreatedAtTime  string  `json:"created_at_time"`
}

type Asset struct {
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
