package main

type address struct {
	State   string `json:"state" bson:"ad_state"`
	City    string `json:"city" bson:"ad_city"`
	Pincode int    `json:"pincode" bson:"ad_pincode"`
}

// Iniatialising all the structure of data
type User struct {
	ID        int     `json:"id" bson:"user_id"`
	Name      string  `json:"name" bson:"name"`
	Age       int     `json:"age" bson:"user_age"`
	Address   address `json:"address" bson:"user_address"`
	Desc      string  `json:"desc" bson:"user_desc"`
	CreatedAt string  `json:"created" bson:"user_created"`
}
