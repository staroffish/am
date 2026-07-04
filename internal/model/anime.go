package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Anime struct {
	ObjectID     primitive.ObjectID `bson:"_id,omitempty"`
	ID           string             `bson:"-" json:"id"`
	AnimeNameCn  string    `bson:"animenamecn"   json:"animenamecn"`
	AnimeNameJp  string    `bson:"animenamejp"   json:"animenamejp"`
	Cast         string    `bson:"cast"          json:"cast"`
	SerialsDuri  string    `bson:"serialsduri"   json:"serialsduri"`
	Type         string    `bson:"type"          json:"type"`
	ImageBin  []byte    `bson:"image"      json:"-"`
	StorDir   string    `bson:"stordir"    json:"stordir"`
	Status    string    `bson:"status"     json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updatetime" json:"updated_at"`
}

type AnimeFile struct {
	Name     string `json:"name"`
	FullPath string `json:"full_path"`
}
