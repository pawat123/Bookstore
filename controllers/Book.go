package controllers

import (
	"Bookstore/config"
	"Bookstore/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := config.DB

	books := []model.BookResponse{}
	db.Model(&model.Book{}).
		Scan(&books)
	//statusok
	return c.Status(200).JSON(books)
}

func GetBookById(c *fiber.Ctx) error {
	db := config.DB

	bookId := c.Params("book_id")

	book := model.BookResponse{}
	err := db.Model(&model.Book{}).
		Where("books.id = ?", bookId).
		First(&book).Error
	//statusnotfound
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No this book",
		})
	}

	return c.Status(200).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	db := config.DB

	var input model.NewBook
	c.BodyParser(&input)

	// Find book detail
	book := model.Book{}
	err := db.Model(&model.Book{}).Where("name = ?", input.Name).First(&book).Error
	if err != nil {
		book.Name = input.Name
		book.Amount = 0
		db.Model(&model.Book{}).Create(&book)
	}

	// Add book
	newBook := model.Book{
		BookID: book.BookID,
		Date:   time.Now().Format("02-01-2006 15:04:05"),
	}
	db.Model(&model.Book{}).Create(&newBook)

	// Update book amount
	book.Amount = book.Amount + 1
	db.Updates(&book)

	return c.Status(200).JSON(fiber.Map{
		"message": "Create book successfully",
		"book":    newBook,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	db := config.DB

	updateBook := model.UpdateBook{}
	c.BodyParser(&updateBook)

	// Find book
	bookId := c.Params("book_id")
	book := model.Book{}
	err := db.Model(&model.Book{}).Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// Update book

	db.Model(&model.Book{}).Where("id = ?", bookId).Updates(&book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update book successfully",
		"book":    book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	db := config.DB

	updateBook := model.UpdateBook{}
	c.BodyParser(&updateBook)

	// Find book
	bookId := c.Params("book_id")
	book := model.Book{}
	err := db.Model(&model.Book{}).Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// Delete book
	db.Model(&model.Book{}).Delete(&book)

	// Decrease book amount
	book = model.Book{}
	db.Model(&model.Book{}).Where("id = ?", book.BookID).First(&book)

	if book.Amount == 1 {
		book.Amount = 0
		db.Model(&model.Book{}).Where("id = ?", book.BookID).Select("amount").Update("amount", 0)
	} else {
		book.Amount = book.Amount - 1
		db.Model(&model.Book{}).Where("id = ?", book.BookID).Updates(&book)
	}

	// Delete book from db
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete book successfully",
		"book":    book,
	})
}
