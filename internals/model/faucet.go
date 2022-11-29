package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faucet struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Chain          string             `json:"chain,omitempty" validate:"required"`
	Public_address string             `json:"public_address" validate:"required"`
	Txn_count      string             `json:"txn_count,omitempty" validate:"required"`
	Last_txn_Time  string             `json:"last_txn_time validate:"required"`
}
