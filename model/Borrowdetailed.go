package model

type Borrow struct {
	ID       uint   `json:"id"`
	MemberID uint   `json:"member_id"`
	BookID   uint   `json:"book_id"`
	Date     string `json:"date"`
}

type BorrowInput struct {
	BookID   uint   `json:"book_id"`
	MemberId uint   `json:"member_id"`
	BookName string `json:"book_name"`
}

type BorrowResponse struct {
	ID       uint   `json:"id"`
	BookID   uint   `json:"book_id"`
	BookName string `json:"book_name"`
	MemberID uint   `json:"member_id"`
	Date     string `json:"date"`
}
