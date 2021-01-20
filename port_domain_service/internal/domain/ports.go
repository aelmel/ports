package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Port struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key         string             `json:"key" bson:"key,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	City        string             `json:"city" bson:"city,omitempty"`
	Country     string             `json:"country" bson:"country,omitempty"`
	Alias       []string           `json:"alias" bson:"alias"`
	Regions     []string           `json:"regions" bson:"regions"`
	Coordinates []float32          `json:"coordinates" bson:"coordinates"`
	Province    string             `json:"province" bson:"province"`
	Timezone    string             `json:"timezone" bson:"timezone"`
	Unlocs      []string           `json:"unlocs" bson:"unlocs"`
	Code        string             `json:"code" bson:"code"`
}
