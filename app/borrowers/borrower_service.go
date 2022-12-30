package borrowers

import (
	"errors"
	"strings"

	"github.com/HungOnBlog/libma/core/gen"
	"github.com/HungOnBlog/libma/core/models"
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

	if !IsCreateDtoValid(dto) {
		return &BorrowerGetDto{}, errors.New("InvalidBorrowerCreationInfo")
	}

	borrower := &models.Borrower{
		Id:       gen.GenBorrowerId(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	err := service.repo.Create(borrower)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return &BorrowerGetDto{}, errors.New("BorrowerAlreadyExists")
		}
		return &BorrowerGetDto{}, err
	}

	return &BorrowerGetDto{
		Id:    borrower.Id,
		Name:  borrower.Name,
		Email: borrower.Email,
	}, nil
}
