package borrowers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var borrowerService *BorrowerService

func init() {
	borrowerService = NewBorrowerService()
}

type BorrowerController struct{}

func (controller *BorrowerController) ApplyController(router fiber.Router) {
	borrowerGroup := router.Group("/borrowers")

	borrowerGroup.Get("/", controller.GetBorrowers)
	fmt.Println("GET    - /borrowers ")

	borrowerGroup.Get("/:id", controller.GetBorrowerById)
	fmt.Println("GET    - /borrowers/:id ")

	borrowerGroup.Post("/", controller.CreateBorrower)
	fmt.Println("POST   - /borrowers ")

	borrowerGroup.Put("/:id", controller.UpdateBorrower)
	fmt.Println("PUT    - /borrowers/:id ")

	borrowerGroup.Delete("/:id", controller.DeleteBorrower)
	fmt.Println("DELETE - /borrowers/:id ")

	borrowerGroup.Post("/login", controller.Login)
	fmt.Println("POST   - /borrowers/login ")
}

func (controller *BorrowerController) GetBorrowers(c *fiber.Ctx) error {
	borrowers, err := borrowerService.GetAllBorrowers()
	if err != nil {
		return err
	}

	return c.JSON(borrowers)
}

func (controller *BorrowerController) GetBorrowerById(c *fiber.Ctx) error {
	id := c.Params("id")

	borrower, err := borrowerService.GetBorrowerById(id)
	if err != nil {
		return err
	}

	return c.JSON(borrower)
}

func (controller *BorrowerController) CreateBorrower(c *fiber.Ctx) error {
	dto := new(BorrowerDto)
	if err := c.BodyParser(dto); err != nil {
		return err
	}

	borrower, err := borrowerService.CreateBorrower(dto)
	if err != nil {
		return err
	}

	return c.JSON(borrower)
}

func (controller *BorrowerController) UpdateBorrower(c *fiber.Ctx) error {
	id := c.Params("id")

	dto := new(BorrowerUpdateDto)
	if err := c.BodyParser(dto); err != nil {
		return err
	}

	borrower, err := borrowerService.UpdateBorrower(id, dto)
	if err != nil {
		return err
	}

	return c.JSON(borrower)
}

func (controller *BorrowerController) DeleteBorrower(c *fiber.Ctx) error {
	id := c.Params("id")

	err := borrowerService.DeleteBorrower(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Borrower deleted"})
}

func (controller *BorrowerController) Login(c *fiber.Ctx) error {
	dto := new(BorrowerLoginDto)
	if err := c.BodyParser(dto); err != nil {
		return err
	}

	err := borrowerService.Login(dto)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
}
