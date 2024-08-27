package domain

type PromoteDemoteRequest struct {
	Identifier string `json:"identifier" bson:"identifier"`
	Action     string `json:"action" bson:"action"`
}
