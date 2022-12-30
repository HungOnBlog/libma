package librarians

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var librarianService *LibrarianService

func init() {
	librarianService = NewLibrarianService()
}

type LibrarianController struct {
}

func (controller *LibrarianController) ApplyController(router fiber.Router) {
	librarianGroup := router.Group("/librarians")

	librarianGroup.Get("/", controller.GetLibrarians)
	fmt.Println("GET    - /librarians")

	librarianGroup.Get("/:id", controller.GetLibrarianById)
	fmt.Println("GET    - /librarians/:id")

	librarianGroup.Post("/", controller.CreateLibrarian)
	fmt.Println("POST   - /librarians")

	librarianGroup.Put("/:id", controller.UpdateLibrarian)
	fmt.Println("PUT    - /librarians/:id")

	librarianGroup.Delete("/:id", controller.DeleteLibrarian)
	fmt.Println("DELETE - /librarians/:id")

	librarianGroup.Post("/login", controller.Login)
	fmt.Println("POST   - /librarians/login")
}

// GetLibrarians ...
//
//	@Summary		Get all librarians
//	@Description	Return all librarians (active) in the system
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians [get]
//	@Success		200	{object}	[]LibrarianDto
func (controller *LibrarianController) GetLibrarians(c *fiber.Ctx) error {
	librarians, err := librarianService.GetLibrarians()
	if err != nil {
		return err
	}

	return c.JSON(librarians)
}

// GetLibrarianById ...
//
//	@Summary		Get librarian by id
//	@Description	Return librarian by id
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians/{id} [get]
//	@Param			id	path	string	true	"Librarian id"
func (controller *LibrarianController) GetLibrarianById(c *fiber.Ctx) error {
	id := c.Params("id")

	librarian, err := librarianService.GetLibrarianById(id)
	if err != nil {
		return err
	}

	return c.JSON(librarian)
}

// CreateLibrarian ...
//
//	@Summary		Create librarian
//	@Description	Create new librarian
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians [post]
//	@Param			librarian	body		LibrarianDto	true	"Librarian"
//	@Success		200			{object}	LibrarianDto
func (controller *LibrarianController) CreateLibrarian(c *fiber.Ctx) error {
	librarian := LibrarianCreateDto{}
	err := c.BodyParser(&librarian)
	if err != nil {
		return err
	}

	result, err := librarianService.CreateLibrarian(&librarian)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// UpdateLibrarian ...
//
//	@Summary		Update librarian
//	@Description	Update librarian
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians/{id} [put]
//	@Param			id			path		string			true	"Librarian id"
//	@Param			librarian	body		LibrarianDto	true	"Librarian"
//	@Success		200			{object}	LibrarianDto
func (controller *LibrarianController) UpdateLibrarian(c *fiber.Ctx) error {
	id := c.Params("id")

	librarian := LibrarianUpdateDto{}
	err := c.BodyParser(&librarian)
	if err != nil {
		return err
	}

	result, err := librarianService.UpdateLibrarian(id, &librarian)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// DeleteLibrarian ...
//
//	@Summary		Delete librarian
//	@Description	Delete librarian
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians/{id} [delete]
//	@Param			id	path	string	true	"Librarian id"
func (controller *LibrarianController) DeleteLibrarian(c *fiber.Ctx) error {
	id := c.Params("id")

	err := librarianService.DeleteLibrarian(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Librarian deleted"})
}

// Login ...
//
//	@Summary		Login
//	@Description	Login
//	@Tags			librarians
//	@Accept			json
//	@Produce		json
//	@Router			/librarians/login [post]
//	@Param			librarian	body		LibrarianLoginDto	true	"Librarian"
//	@Success		200			{object}	LibrarianDto
func (controller *LibrarianController) Login(c *fiber.Ctx) error {
	librarian := LibrarianLoginDto{}
	err := c.BodyParser(&librarian)
	if err != nil {
		return err
	}

	result, err := librarianService.Login(&librarian)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
