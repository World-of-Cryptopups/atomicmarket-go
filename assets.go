package atomicmarket

import (
	"fmt"

	"github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const ASSETS_API = "assets"

type GetAssetsQuery struct {
	CollectionName          string   `url:"collection_name,omitempty"`
	SchemaName              string   `url:"schema_name,omitempty"`
	TemplateID              int      `url:"template_id,omitempty"`
	IsTransferable          bool     `url:"is_transferable,omitempty"`
	IsBurnable              bool     `url:"is_burnable,omitempty"`
	Burned                  bool     `url:"burned,omitempty"`
	Owner                   string   `url:"owner,omitempty"`
	Match                   string   `url:"match,omitempty"`
	CollectionBlacklist     []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist     []string `url:"collection_whitelist,omitempty" del:","`
	OnlyDuplicateTemplates  bool     `url:"only_duplicate_templates,omitempty"`
	HasBackedTokens         bool     `url:"has_backed_tokens,omitempty"`
	AuthorizedAccount       bool     `url:"authorized_account,omitempty"`
	TemplateWhitelist       []string `url:"template_whitelist,omitempty" del:","`
	TemplateBlacklist       []string `url:"template_blacklist,omitempty" del:","`
	HideTemplatesByAccounts string   `url:"hide_templates_by_accounts,omitempty"`
	HideOffers              bool     `url:"hide_offers,omitempty"`
	IDs                     []string `url:"ids,omitempty" del:","`
	LowerBound              []string `url:"lower_bound,omitempty" del:","`
	UpperBound              []string `url:"upper_bound,omitempty" del:","`
	Before                  int64    `url:"before,omitempty"`
	After                   int64    `url:"after,omitempty"`
	Page                    int      `url:"page,omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Order                   string   `url:"order,omitempty"`
	Sort                    string   `url:"sort,omitempty"`
}

// Fetch assets.
func (a *AtomicMarket) GetAssets(options *GetAssetsQuery) (*AssetsProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, ASSETS_API, fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, ASSETS_API)
	}

	var result = &AssetsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Fetch asset by id.
func (a *AtomicMarket) GetAssetID(id string) (*AssetIDProps, error) {
	u := urljoin.UrlJoin(a.apiUrl, ASSETS_API, id)

	var result = &AssetIDProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Fetch asset stats.
func (a *AtomicMarket) GetAssetIDStats(id string) (*AssetIDStatsProps, error) {
	u := urljoin.UrlJoin(a.apiUrl, ASSETS_API, id, "stats")

	var result = &AssetIDStatsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetAssetLogsQuery struct {
	Page            int    `url:"page,omitempty"`
	Limit           int    `url:"limit,omitempty"`
	Order           string `url:"order,omitempty"`
	ActionWhitelist string `url:"action_whitelist,omitempty"`
	ActionBlacklist string `url:"action_blacklist,omitempty"`
}

// Fetch asset logs.
func (a *AtomicMarket) GetAssetIDLogs(id string, options *GetAssetLogsQuery) (*AssetIDLogsProps, error) {
	var opts = &GetAssetLogsQuery{
		Limit: 100,
		Page:  1,
		Order: "desc",
	}

	if options != nil {
		opts = &GetAssetLogsQuery{
			ActionWhitelist: options.ActionWhitelist,
			ActionBlacklist: options.ActionBlacklist,
		}

		if options.Limit < 1 {
			opts.Limit = 100
		}
		if options.Page < 1 {
			opts.Page = 1
		}
		if options.Order == "desc" || options.Order == "asc" {
			opts.Order = options.Order
		} else {
			// default to `desc`
			opts.Order = "desc"
		}
	}

	v, _ := query.Values(opts)
	u := urljoin.UrlJoin(a.apiUrl, ASSETS_API, id, "logs", fmt.Sprintf("?%s", v.Encode()))

	var result = &AssetIDLogsProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetAssetSalesQuery struct {
	Buyer  string `url:"buyer,omitempty" `
	Seller string `url:"seller,omitempty"`
	Symbol string `url:"symbol,omitempty"`
	Order  string `url:"order,omitempty"`
}

// Gets price history for a specific asset.
func (a *AtomicMarket) GetAssetIDSales(id string, options *GetAssetSalesQuery) (*AssetIDSalesProps, error) {
	var opts = &GetAssetSalesQuery{
		Order: "asc",
	}

	if options != nil {
		if options.Order == "desc" || options.Order == "asc" {
			opts.Order = options.Order
		} else {
			// default to `asc`
			opts.Order = "asc"
		}
	}

	v, _ := query.Values(opts)
	u := urljoin.UrlJoin(a.apiUrl, ASSETS_API, id, "sales", fmt.Sprintf("?%s", v.Encode()))

	var result = &AssetIDSalesProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
