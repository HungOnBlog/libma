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

func (service *BorrowerService) GetBorrowerById(id string) (*BorrowerGetDto, error) {
	borrower, err := service.repo.FindById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &BorrowerGetDto{}, errors.New("BorrowerNotFound")
		}
		return &BorrowerGetDto{}, err
	}

	return &BorrowerGetDto{
		Id:    borrower.Id,
		Name:  borrower.Name,
		Email: borrower.Email,
	}, nil
}

func (service *BorrowerService) GetBorrowerByEmail(email string) (*BorrowerGetDto, error) {
	borrower, err := service.repo.FindByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &BorrowerGetDto{}, errors.New("BorrowerNotFound")
		}
		return &BorrowerGetDto{}, err
	}

	return &BorrowerGetDto{
		Id:    borrower.Id,
		Name:  borrower.Name,
		Email: borrower.Email,
	}, nil
}

func (service *BorrowerService) GetAllBorrowers() ([]*BorrowerGetDto, error) {
	borrowers, err := service.repo.FindAll()
	if err != nil {
		return []*BorrowerGetDto{}, err
	}

	var borrowerDtos []*BorrowerGetDto
	for _, borrower := range borrowers {
		borrowerDtos = append(borrowerDtos, &BorrowerGetDto{
			Id:    borrower.Id,
			Name:  borrower.Name,
			Email: borrower.Email,
		})
	}

	return borrowerDtos, nil
}

func (service *BorrowerService) UpdateBorrower(id string, dto *BorrowerUpdateDto) (*BorrowerGetDto, error) {
	if !IsUpdateDtoValid(dto) {
		return &BorrowerGetDto{}, errors.New("InvalidBorrowerUpdateInfo")
	}

	borrower, err := service.repo.FindById(id)
	if err != nil {
		return &BorrowerGetDto{}, err
	}

	borrower.Name = dto.Name
	borrower.Email = dto.Email
	borrower.Password = dto.Password

	err = service.repo.Update(borrower)
	if err != nil {
		return &BorrowerGetDto{}, err
	}

	return &BorrowerGetDto{
		Id:    borrower.Id,
		Name:  borrower.Name,
		Email: borrower.Email,
	}, nil
}

func (service *BorrowerService) DeleteBorrower(id string) error {
	err := service.repo.DeleteById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("BorrowerNotFound")
		}
		return err
	}

	return nil
}

func (service *BorrowerService) Login(dto *BorrowerLoginDto) error {
	if !IsLoginDtoValid(dto) {
		return errors.New("InvalidBorrowerLoginInfo")
	}

	borrower, err := service.repo.FindByEmail(dto.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("BorrowerLoginFailed")
		}
		return err
	}

	if borrower.Password != dto.Password {
		return errors.New("BorrowerLoginFailed")
	}

	return nil
}
