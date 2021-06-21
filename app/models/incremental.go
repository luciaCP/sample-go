package models


type Incremental struct {
	Id int `json:"id"`
	Amount int `json:"amount"`
	Upgraded bool `json:"upgraded"`
}


func (*Incremental)Inflate() {
	return
}