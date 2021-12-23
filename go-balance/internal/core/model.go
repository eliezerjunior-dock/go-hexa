package core

import (
	"time"

)

type Balance struct {
	Id					string		`json:"id"`
    Account 			string 		`json:"account"`
	Amount				int32 		`json:"amount"`
    DateBalance  		time.Time 	`json:"date_balance"`
	Description			string 		`json:"description"`
}

type Customer struct {
	Id      			int     	`json:"id,omitempty"`
	Name    			string  	`json:"name,omitempty"`
    Email   			string  	`json:"email,omitempty"`
    Address 			string  	`json:"address,omitempty"`
	Status  			bool    	`json:"status,omitempty"`
    CreditCardNumber 	string  	`json:"credit_card,omitempty"` 
    Phones 				[]PhoneArray `json:"phones,omitmepty"` 
	Obs     			string  	`json:"obs,omitempty"` 
}

type PhoneArray struct {
    Number 				string 		`json:"number,omitempty"`
}

type Pod struct {
	Name		string `json:"name"`
	PID			string `json:"os_pid"`
	Host		string `json:"host"`
    Ip 			string `json:"ip"`
    Port		string `json:"port"`
	Setup		Setup  `json:"setup"`
}

type Setup struct {
    ResponseTime 		int `json:"response_time"`
    ResponseStatusCode  int `json:"response_status_code"`
	IsRandomTime		bool `json:"is_random_time"`
	Count				int `json:"count"`
}
