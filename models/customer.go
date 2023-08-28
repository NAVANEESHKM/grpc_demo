package models



type Customer struct{

	Customer_ID int  `json:"customer_id" bson:"customer_id"`
	Customer_Name string `json:"customer_name" bson:"customer_name"`
	Account_ID int `json:"account_id" bson:"account_id"`
	Balance int `json:"balance" bson:"balance"`
	Bank_ID int `json:"bank_id" bson:"bank_id"`
   
}