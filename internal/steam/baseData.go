package steam

import (
	"encoding/json"
	"strconv"
)

func NewParser(handler func(listingInfo []byte, assetInfo []byte)) BaseParser {
	parser := BaseParser{}
	parser.handler = handler

	return parser
}

type BaseParser struct {
	handler func(listingInfo []byte, assetInfo []byte)
}

type responseStruct struct {
	Success      bool
	More         bool
	ResultsHtml  bool
	ListingInfo  map[string]json.RawMessage
	Purchaseinfo []string
	Assets       map[string]map[string]map[string]json.RawMessage
	Currency     []string
	Hovers       bool
	AppData      map[string]json.RawMessage
	LastTime     int64
	LastListing  string
}

type baseInfoStruct struct {
	ListingID string `json:"listingid"`
	Asset     struct {
		AppID     int
		ContextID string
		Amount    string
		AssetID   string `json:"id"`
	}
}

func (h *BaseParser) RunParsers(data []byte) error {
	response := responseStruct{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	for _, listingInfo := range response.ListingInfo {
		baseInfo := baseInfoStruct{}
		// no biggie is fails, just move on to the next item
		err = json.Unmarshal(listingInfo, &baseInfo)
		if err != nil {
			// TODO: add a logger message
			continue
		}

		asset := baseInfo.Asset

		appID := strconv.Itoa(asset.AppID)

		assetInfo, ok := response.Assets[appID][asset.ContextID][asset.AssetID]
		if !ok {
			// TODO: need something good if this fails
			continue
		}

		h.handler(listingInfo, assetInfo)
	}
	return nil
}
