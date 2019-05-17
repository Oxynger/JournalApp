package model

type SearchResponse struct {
	ObjectName string `bson:"name" json:"name" binding:"required"`
}
