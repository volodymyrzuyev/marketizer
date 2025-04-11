package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	custSql "github.com/volodymyrzuyev/marketizer/internal/database/sql"
	"github.com/volodymyrzuyev/marketizer/internal/steam/listings"
	"github.com/volodymyrzuyev/marketizer/internal/steam/listings/assets"
)

// Service represents a service that interacts with a database.
type Service interface {
	GetUser(email string) (custSql.User, error)
	AddUser(email, password, name string) error
	AddItems(listingInfo []byte, assetInfo []byte)
	GetItems(orderBy, sortBy, searchString string) ([]custSql.Item, error)
	GetItemsToNotify(email string) (map[int64]bool, error)
	AddToFollows(email, marketHashName string, assetId int64) error
	GetFollows(email string) (map[string]bool, error)
	GetFollowItems(orderBy, sortBy, searchString, email string) ([]custSql.Item, error)
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	q  *custSql.Queries
	db *sql.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	dbInstance = &service{
		q:  custSql.New(db),
		db: db,
	}
	dbInstance.q.Create_table1(context.TODO())
	dbInstance.q.Create_table2(context.TODO())
	dbInstance.q.Create_table3(context.TODO())
	dbInstance.q.Create_table4(context.TODO())
	return dbInstance
}

func (s *service) AddToFollows(email, marketHashName string, assetID int64) error {
	arg := custSql.AddToFollowsParams{
		Email:          email,
		MarketHashName: marketHashName,
	}

	arg2 := custSql.SetItemAsNotifiedParams{
		AssetID: assetID,
		Email:   email,
	}
	err := s.q.SetItemAsNotified(context.TODO(), arg2)
	if err != nil {
		return err
	}

	return s.q.AddToFollows(context.TODO(), arg)
}

func (s *service) GetFollows(email string) (map[string]bool, error) {
	m, err := s.q.GetFollows(context.TODO(), email)
	if err != nil || errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	retMap := make(map[string]bool)

	for _, n := range m {
		retMap[n] = true
	}

	return retMap, nil
}

func (s *service) GetFollowItems(orderBy, sortBy, searchString, email string) ([]custSql.Item, error) {
	args := custSql.GetFollowItemsNameASCParams{
		Email:          email,
		MarketHashName: "%" + searchString + "%",
	}
	var err error
	var items []custSql.Item
	switch orderBy {
	case "asc":
		switch sortBy {
		case "time":
			items, err = s.q.GetFollowItemsTimeDSC(context.TODO(), custSql.GetFollowItemsTimeDSCParams(args))
		case "price":
			items, err = s.q.GetFollowItemsPriceASC(context.TODO(), custSql.GetFollowItemsPriceASCParams(args))
		case "name":
			items, err = s.q.GetFollowItemsNameASC(context.TODO(), args)
		default:
			items, err = []custSql.Item{}, fmt.Errorf("Invalid ordering")
		}
	case "dsc":
		switch sortBy {
		case "time":
			items, err = s.q.GetFollowItemsTimeASC(context.TODO(), custSql.GetFollowItemsTimeASCParams(args))
		case "price":
			items, err = s.q.GetFollowItemsPriceDSC(context.TODO(), custSql.GetFollowItemsPriceDSCParams(args))
		case "name":
			items, err = s.q.GetFollowItemsNameDSC(context.TODO(), custSql.GetFollowItemsNameDSCParams(args))
		default:
			items, err = []custSql.Item{}, fmt.Errorf("Invalid ordering")
		}
	}

	if err != nil || errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return items, nil
}

func (s *service) GetItemsToNotify(email string) (map[int64]bool, error) {
	assetsIDs, err := s.q.GetItemsToNotify(context.TODO(), email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	retMap := make(map[int64]bool)

	tx, err := s.db.BeginTx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	p := s.q.WithTx(tx)

	for _, i := range assetsIDs {
		retMap[i] = true
		args := custSql.SetItemAsNotifiedParams{
			AssetID: i,
			Email:   email,
		}
		p.SetItemAsNotified(context.TODO(), args)
	}

	return retMap, tx.Commit()
}

func (s *service) GetUser(email string) (custSql.User, error) {
	return s.q.Get_User(context.TODO(), email)
}

func (s *service) AddUser(email, password, name string) error {
	args := custSql.Add_UserParams{
		Name:     sql.NullString{name, name != ""},
		Email:    email,
		Password: sql.NullString{password, password != ""},
	}
	return s.q.Add_User(context.TODO(), args)
}

func (s *service) GetItems(orderBy, sortBy, searchString string) ([]custSql.Item, error) {
	switch orderBy {
	case "asc":
		switch sortBy {
		case "time":
			return s.q.GetItemsTimeDESC(context.TODO(), "%"+searchString+"%")
		case "price":
			return s.q.GetItemsPriceASC(context.TODO(), "%"+searchString+"%")
		case "name":
			return s.q.GetItemsNameASC(context.TODO(), "%"+searchString+"%")
		default:
			return []custSql.Item{}, fmt.Errorf("Invalid ordering")
		}
	case "dsc":
		switch sortBy {
		case "time":
			return s.q.GetItemsTimeASC(context.TODO(), "%"+searchString+"%")
		case "price":
			return s.q.GetItemsPriceDESC(context.TODO(), "%"+searchString+"%")
		case "name":
			return s.q.GetItemsNameDESC(context.TODO(), "%"+searchString+"%")
		default:
			return []custSql.Item{}, fmt.Errorf("Invalid ordering")
		}
	}

	return []custSql.Item{}, fmt.Errorf("Invalid ordering")
}

func (s *service) AddItems(listingInfo []byte, assetInfo []byte) {
	asset := assets.BaseAsset{}
	listing := listings.BaseListing{}

	err := json.Unmarshal(listingInfo, &listing)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}

	err = json.Unmarshal(assetInfo, &asset)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}

	dbPrams := custSql.AddAssetParams{
		AssetID:        int64(asset.ID),
		MarketHashName: asset.MarketName,
		Price:          int64(listing.Price + listing.Fee),
		Appid:          int64(listing.Asset.AppID),
		Time:           time.Now().Unix(),
		Image:          "https://community.fastly.steamstatic.com/economy/image/" + asset.IconUrl,
	}

	if dbPrams.Price == 0 {
		return
	}

	err = s.q.AddAsset(context.TODO(), dbPrams)
	if err != nil {
		fmt.Println("DB add error", err)
	}

}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}
