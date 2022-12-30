package librarians

import (
	"errors"
	"strings"

	"github.com/HungOnBlog/libma/core/gen"
	"github.com/HungOnBlog/libma/core/models"
)

type LibrarianService struct {
	repo *LibrarianRepo
}

func NewLibrarianService() *LibrarianService {
	return &LibrarianService{
		repo: GetLibrarianRepo(),
	}
}

func (service *LibrarianService) CreateLibrarian(dto *LibrarianCreateDto) (*LibrarianDto, error) {

	if !IsLibrarianCreateDtoValid(dto) {
		return &LibrarianDto{}, errors.New("InvalidLibrarianCreationInfo")
	}

	librarian := &models.Librarian{
		Id:       gen.GenLibrarianId(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	err := service.repo.Create(librarian)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return &LibrarianDto{}, errors.New("LibrarianAlreadyExists")
		}
		return &LibrarianDto{}, err
	}

	return &LibrarianDto{
		Id:    librarian.Id,
		Name:  librarian.Name,
		Email: librarian.Email,
	}, nil
}

func (service *LibrarianService) GetLibrarianById(id string) (*LibrarianDto, error) {
	librarian, err := service.repo.FindById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &LibrarianDto{}, errors.New("LibrarianNotFound")
		}
		return &LibrarianDto{}, err
	}

	return &LibrarianDto{
		Id:    librarian.Id,
		Name:  librarian.Name,
		Email: librarian.Email,
	}, nil
}

func (service *LibrarianService) UpdateLibrarian(id string, dto *LibrarianUpdateDto) (*LibrarianDto, error) {
	if !IsLibrarianUpdateDtoValid(dto) {
		return &LibrarianDto{}, errors.New("InvalidLibrarianUpdateInfo")
	}

	librarian, err := service.repo.FindById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &LibrarianDto{}, errors.New("LibrarianNotFound")
		}
		return &LibrarianDto{}, err
	}

	librarian.Name = dto.Name
	librarian.Email = dto.Email

	err = service.repo.Update(id, librarian)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return &LibrarianDto{}, errors.New("LibrarianAlreadyExists")
		}
		return &LibrarianDto{}, err
	}

	return &LibrarianDto{
		Id:    librarian.Id,
		Name:  librarian.Name,
		Email: librarian.Email,
	}, nil
}

func (service *LibrarianService) DeleteLibrarian(id string) error {
	err := service.repo.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("LibrarianNotFound")
		}
		return err
	}

	return nil
}

func (service *LibrarianService) GetLibrarians() ([]*LibrarianDto, error) {
	librarians, err := service.repo.FindAll()
	if err != nil {
		return []*LibrarianDto{}, err
	}

	librarianDtos := []*LibrarianDto{}
	for _, librarian := range librarians {
		librarianDtos = append(librarianDtos, &LibrarianDto{
			Id:    librarian.Id,
			Name:  librarian.Name,
			Email: librarian.Email,
		})
	}

	return librarianDtos, nil
}

func (service *LibrarianService) GetLibrarianByEmail(email string) (*LibrarianDto, error) {
	librarian, err := service.repo.FindByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &LibrarianDto{}, errors.New("LibrarianNotFound")
		}
		return &LibrarianDto{}, err
	}

	return &LibrarianDto{
		Id:    librarian.Id,
		Name:  librarian.Name,
		Email: librarian.Email,
	}, nil
}

func (service *LibrarianService) Login(dto *LibrarianLoginDto) (*LibrarianDto, error) {
	if !IsLibrarianLoginDtoValid(dto) {
		return &LibrarianDto{}, errors.New("InvalidLibrarianLoginInfo")
	}

	librarian, err := service.repo.FindByEmail(dto.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &LibrarianDto{}, errors.New("LibrarianLoginFailed")
		}
		return &LibrarianDto{}, err
	}

	if librarian.Password != dto.Password {
		return &LibrarianDto{}, errors.New("LibrarianLoginFailed")
	}

	return &LibrarianDto{
		Id:    librarian.Id,
		Name:  librarian.Name,
		Email: librarian.Email,
	}, nil
}
