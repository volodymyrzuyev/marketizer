package listings

import (
	"encoding/json"
	"strconv"
)

type BaseListing struct {
	ListingID           int
	Price               int
	Fee                 int
	PublisherFeePercent float64
	Asset               struct {
		Currency  int
		AppID     int
		ContextID int
		AssetID   int
		Amount    int
	}
}

type baseListingRaw struct {
	ListingID                    string `json:"listingid"`
	Price                        int    `json:"price"`
	Fee                          int    `json:"fee"`
	PublisherFeeApp              int    `json:"publisher_fee_app"`
	PublisherFeePercent          string `json:"publisher_fee_percent"`
	CurrencyID                   int    `json:"currencyid"`
	SteamFee                     int    `json:"steam_fee"`
	PublisherFee                 int    `json:"publisher_fee"`
	ConvertedPrice               int    `json:"converted_price"`
	ConvertedFee                 int    `json:"converted_fee"`
	ConvertedCurrencyID          int    `json:"converted_currencyid"`
	ConvertedSteamFee            int    `json:"converted_steam_fee"`
	ConvertedPublisherFee        int    `json:"converted_publisher_fee"`
	ConvertedPricePerUnit        int    `json:"converted_price_per_unit"`
	ConvertedFeePerUnit          int    `json:"converted_fee_per_unit"`
	ConvertedSteamFeePerUnit     int    `json:"converted_steam_fee_per_unit"`
	ConvertedPublisherFeePerUnit int    `json:"converted_publisher_fee_per_unit"`
	Asset                        struct {
		Currency  int    `json:"currency"`
		AppID     int    `json:"appid"`
		ContextID string `json:"contextid"`
		AssetID   string `json:"id"`
		Amount    string `json:"amount"`
	}
}

func (i *BaseListing) UnmarshalJSON(data []byte) error {
	rawItem := baseListingRaw{}

	err := json.Unmarshal(data, &rawItem)
	if err != nil {
		return err
	}

	dummy := BaseListing{}

	dummy.ListingID, err = strconv.Atoi(rawItem.ListingID)
	if err != nil {
		return err
	}
	dummy.Price = rawItem.Price
	if rawItem.ConvertedPrice != 0 {
		dummy.Price = rawItem.ConvertedPrice
	}
	dummy.Fee = rawItem.Fee
	if rawItem.ConvertedFee != 0 {
		dummy.Fee = rawItem.ConvertedFee
	}
	dummy.PublisherFeePercent, err = strconv.ParseFloat(rawItem.PublisherFeePercent, 64)
	if err != nil {
		return err
	}
	dummy.Asset.Currency = rawItem.Asset.Currency
	dummy.Asset.AppID = rawItem.Asset.AppID
	dummy.Asset.ContextID, err = strconv.Atoi(rawItem.Asset.ContextID)
	if err != nil {
		return err
	}
	dummy.Asset.AssetID, err = strconv.Atoi(rawItem.Asset.AssetID)
	if err != nil {
		return err
	}
	dummy.Asset.Amount, err = strconv.Atoi(rawItem.Asset.Amount)
	if err != nil {
		return err
	}

	*i = dummy

	return nil
}
