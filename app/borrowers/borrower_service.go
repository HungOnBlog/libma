package borrowers

import (
	"github.com/HungOnBlog/libma/core/gen"
	"github.com/HungOnBlog/libma/core/models"
	"github.com/HungOnBlog/libma/infra/db"
)

type BorrowerService struct {
	repo *BorrowerRepo
}

func NewBorrowerService() *BorrowerService {
	return &BorrowerService{
		repo: GetBorrowerRepo(),
	}
}

func (service *BorrowerService) CreateBorrower(dto *BorrowerDto) (*BorrowerGetDto, error) {
	dbClient := db.GetDb()
	if dbClient == nil {
		return &BorrowerGetDto{}, nil
	}

	borrower := &models.Borrower{
		Id:       gen.GenBorrowerId(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	err := service.repo.Create(borrower)
	if err != nil {
		return &BorrowerGetDto{}, err
	}

	return &BorrowerGetDto{
		Id:    borrower.Id,
		Name:  borrower.Name,
		Email: borrower.Email,
	}, nil
}
