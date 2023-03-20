package models

type SellInput struct {
	NftId         int    `json:"NftId,omitempty"`
	NftCollection string `json:"nftCollection,omitempty"`
	Seller        string `json:"seller,omitempty"`
	Cost          int    `json:"cost,omitempty"`
	Token         string `json:"token,omitempty"`
	//Token - address
}

type GetNftInfoInput struct {
	NftId         int    `json:"nftId,omitempty" db:"NftId"`
	NftCollection string `json:"nftCollection,omitempty" db:"NftCollection"`
}

type GetNftInfoOutput struct {
	Seller string `json:"Seller,omitempty" db:"Seller"`
	Cost   int    `json:"Cost,omitempty" db:"Cost"`
	Token  string `json:"Token,omitempty" db:"Token"`
}
