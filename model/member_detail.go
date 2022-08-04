package model

type Member struct {
	MemberID  uint   `json:"member_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
