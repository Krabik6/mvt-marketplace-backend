package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mvt-marketplace-backend/models"
)

type SellsPostgres struct {
	db *sqlx.DB
}

func NewSellsPostgres(db *sqlx.DB) *SellsPostgres {
	return &SellsPostgres{db: db}
}

func (s *SellsPostgres) PutOnSell(input models.SellInput) error {
	db := s.db

	query := fmt.Sprintf(`INSERT INTO sales 
			("NftId", "NftCollection", "Seller", "Cost", "Token") 
			values (%d, '%s', '%s', %d, '%s');`, input.NftId, input.NftCollection, input.Seller, input.Cost, input.Token)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *SellsPostgres) GetForSaleNftInfo(input models.GetNftInfoInput) (models.GetNftInfoOutput, error) {
	db := s.db
	var output = models.GetNftInfoOutput{}

	query := fmt.Sprintf(`SELECT "Seller", "Cost", "Token" FROM sales WHERE "NftCollection"='%s' and "NftId"=%d;`,
		input.NftCollection,
		input.NftId)
	fmt.Println(query)

	err := db.Get(&output, query)
	fmt.Println(output)

	if err != nil {
		return output, err
	}
	fmt.Println(output)
	return output, err
}
