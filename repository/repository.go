package repository

import (
	"github.com/jmoiron/sqlx"
	"mvt-marketplace-backend/models"
)

type Sells interface {
	PutOnSell(input models.SellInput) error
	GetForSaleNftInfo(input models.GetNftInfoInput) (models.GetNftInfoOutput, error)
}

type Repository struct {
	Sells
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Sells: NewSellsPostgres(db),
	}
}
