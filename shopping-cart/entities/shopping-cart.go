package entities

import "gopkg.in/mgo.v2/bson"

// Prod information
type Prod struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Prize       float64       `bson:"prize" json:"prize"`
	Discount    float64       `bson:"discount" json:"discount"`
	CoverImage  string        `bson:"coverImage" json:"coverImage"`
	Description string        `bson:"description" json:"description"`
}

// Add Prod in cart information
type AddProd struct {
	ProdID        string `json:"prodId" example:"Prod Id"`
	DiscountPass  string `json:"discountPass" example:"discount pass"`
}