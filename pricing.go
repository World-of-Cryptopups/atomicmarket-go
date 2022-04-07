package atomicmarket

import (
	"fmt"

	"github.com/TheBoringDude/go-urljoin"
	"github.com/google/go-querystring/query"
)

const PRICES_API = "prices"

type GetPriceSalesQuery struct {
	CollectionName string `url:"collection_name,omitempty"`
	SchemaName     string `url:"schema_name,omitempty"`
	TemplateID     int64  `url:"template_id,omitempty"`
	Burned         bool   `url:"burned,omitempty"`
	Symbol         string `url:"symbol,omitempty"`
}

// Gets price history for a template or schema.
func (a *AtomicMarket) GetPriceSales(options *GetPriceSalesQuery) (*PriceSalesProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "sales", fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "sales")
	}

	var result = &PriceSalesProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Gets price history for a template or schema.
func (a *AtomicMarket) GetPriceSalesDays(options *GetPriceSalesQuery) (*PriceSalesDaysProps, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "sales", "days", fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "sales", "days")
	}

	var result = &PriceSalesDaysProps{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetPriceTemplatesQuery struct {
	CollectionName string `url:"collection_name,omitempty"`
	SchemaName     string `url:"schema_name,omitempty"`
	TemplateID     int64  `url:"template_id,omitempty"`
	Burned         bool   `url:"burned,omitempty"`
	Symbol         string `url:"symbol,omitempty"`
	Page           int64  `url:"page,omitempty"`
	Limit          int64  `url:"limit,omitempty"`
	Order          string `url:"order,omitempty"`
}

// Get template price stats.
func (a *AtomicMarket) GetPriceTemplates(options *GetPriceTemplatesQuery) (*PriceTemplates, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "templates", fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "templates")
	}

	var result = &PriceTemplates{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetPriceAssetsQuery struct {
	CollectionName      string   `url:"collection_name,omitempty"`
	SchemaName          string   `url:"schema_name,omitempty"`
	TemplateID          int64    `url:"template_id,omitempty"`
	Burned              bool     `url:"burned,omitempty"`
	Owner               string   `url:"owner,omitempty"`
	Match               string   `url:"match,omitempty"`
	Search              string   `url:"search,omitempty"`
	MatchImmutableName  string   `url:"match_immutable_name,omitempty"`
	MatchMutableName    string   `url:"match_mutable_name,omitempty"`
	IsTransferable      bool     `url:"is_transferable,omitempty"`
	IsBurnable          bool     `url:"is_burnable,omitempty"`
	CollectionBlacklist []string `url:"collection_blacklist,omitempty" del:","`
	CollectionWhitelist []string `url:"collection_whitelist,omitempty" del:","`
	AuthorizedAccount   bool     `url:"authorized_account,omitempty"`
	HideOffers          bool     `url:"hide_offers,omitempty"`
	IDs                 []string `url:"ids,omitempty" del:","`
	LowerBound          []string `url:"lower_bound,omitempty" del:","`
	UpperBound          []string `url:"upper_bound,omitempty" del:","`
}

// Get price history for a template or schema.
func (a *AtomicMarket) GetPriceAssets(options *GetPriceAssetsQuery) (*PriceAssets, error) {
	var u string

	if options != nil {
		v, _ := query.Values(options)
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "assets", fmt.Sprintf("?%s", v.Encode()))
	} else {
		u = urljoin.UrlJoin(a.apiUrl, PRICES_API, "assets")
	}

	var result = &PriceAssets{}
	if err := a._requester(u, result); err != nil {
		return nil, err
	}

	return result, nil
}
