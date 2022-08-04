package Routes

import (
	"Bookstore/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Book
	books := app.Group("/books")
	books.Get("/", controllers.GetBooks)
	books.Get("/:book_id", controllers.GetBookById)
	books.Post("/", controllers.CreateBook)
	books.Put("/:book_id", controllers.UpdateBook)
	books.Delete("/:book_id", controllers.DeleteBook)

	// Member
	members := app.Group("/members")
	members.Get("/", controllers.GetMembers)
	members.Get(("/:member_id"), controllers.GetMemberById)
	members.Post("/", controllers.CreateMember)
	members.Put("/:member_id", controllers.UpdateMember)
	members.Delete("/:member_id", controllers.DeleteMember)
	members.Get("/:member_id", controllers.CheckMember)

	// Borrow
	borrows := app.Group("/borrows")
	borrows.Get("/", controllers.GetBorrowList)
	borrows.Post("/", controllers.BorrowBook)

}
