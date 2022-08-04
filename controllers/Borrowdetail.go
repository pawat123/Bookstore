package controllers

import (
	"Bookstore/config"
	"Bookstore/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BorrowBook(c *fiber.Ctx) error {
	db := config.DB

	input := model.BorrowInput{}
	c.BodyParser(&input)

	// Check member
	member := model.Member{}
	err := db.Model(&model.Member{}).Where("Member_id = ?", input.MemberId).First(&member).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Member not found",
		})
	}

	// Get book state

	book := model.Book{}
	err = db.Model(&model.Book{}).Where("id = ?", input.BookID).First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	state := book.State

	// Check state
	var message string

	switch state {
	case "damaged":
		message = "book has problem"
	case "not available":
		message = "Book not available"
	case "borrowed":
		message = "Book has been borrowed"
	case "available":
		db.Model(&model.Book{}).Where("id = ?", input.BookID).Update("state", "borrowed")

		borrow := model.Borrow{
			Date:     time.Now().Format("02-01-2006 15:04:05"),
			BookID:   input.BookID,
			MemberID: member.MemberID,
		}
		db.Model(&model.Borrow{}).Create(&borrow)

		message = "Borrow successfully"
	}

	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}

func GetBorrowList(c *fiber.Ctx) error {
	db := config.DB

	borrowedBooks := []model.BorrowResponse{}
	db.Model(&model.Borrow{}).
		Scan(&borrowedBooks)

	return c.Status(200).JSON(borrowedBooks)
}
