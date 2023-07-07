package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SeriesReference string             `bson:"Series_reference" json:"seriesReference"`
	Period          string             `bson:"Period" json:"period"`
	DataValue       float64            `bson:"Data_value" json:"dataValue"`
	Suppressed      string             `bson:"Suppressed" json:"suppressed"`
	Status          string             `bson:"STATUS" json:"status"`
	Units           string             `bson:"UNITS" json:"units"`
	Magnitude       int                `bson:"Magnitude" json:"magnitude"`
	Subject         string             `bson:"Subject" json:"subject"`
	Group           string             `bson:"Group" json:"group"`
	SeriesTitle1    string             `bson:"Series_title_1" json:"seriesTitle1"`
	SeriesTitle2    string             `bson:"Series_title_2" json:"seriesTitle2"`
	SeriesTitle3    string             `bson:"Series_title_3" json:"seriesTitle3"`
	SeriesTitle4    string             `bson:"Series_title_4" json:"seriesTitle4"`
	SeriesTitle5    string             `bson:"Series_title_5" json:"seriesTitle5"`
}
