package assets

import (
	"encoding/json"
	"strconv"
)

type BaseAsset struct {
	ID           int
	ClassID      int
	InstanceID   int
	Status       int
	IconUrl      string
	IconUrlLarge string
	Descriptions []Descriptions
	NameColor    string
	Type         string
	MarketName   string
}

type Descriptions struct {
	Type  string
	Value string
	Color string
	Name  string
}

type baseAssetRaw struct {
	ID             string
	ClassID        string
	InstanceId     string
	Status         int
	Icon_url       string
	Icon_url_large string
	Descriptions   []struct {
		Type  string
		Value string
		Color string
		Name  string
	}
	Name_color       string
	Type             string
	Market_hash_name string
}

func (i *BaseAsset) UnmarshalJSON(data []byte) error {
	raw := baseAssetRaw{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	dummyAsset := BaseAsset{}

	dummyAsset.ID, err = strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}

	dummyAsset.ClassID, err = strconv.Atoi(raw.ClassID)
	if err != nil {
		return err
	}

	dummyAsset.InstanceID, err = strconv.Atoi(raw.InstanceId)
	if err != nil {
		return err
	}

	dummyAsset.Status = raw.Status

	dummyAsset.IconUrl = raw.Icon_url
	dummyAsset.IconUrlLarge = raw.Icon_url_large

	for _, v := range raw.Descriptions {
		dummyAsset.Descriptions = append(dummyAsset.Descriptions, Descriptions{v.Type, v.Value, v.Color, v.Name})
	}

	dummyAsset.NameColor = raw.Name_color
	dummyAsset.Type = raw.Type
	dummyAsset.MarketName = raw.Market_hash_name

	*i = dummyAsset

	return nil
}
