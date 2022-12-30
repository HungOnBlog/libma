package librarians

import (
	"github.com/HungOnBlog/libma/core/models"
	"github.com/HungOnBlog/libma/infra/db"
	"gorm.io/gorm"
)

type LibrarianRepo struct {
	Db *gorm.DB
}

var librarianRepo *LibrarianRepo

func init() {
	librarianRepo = &LibrarianRepo{
		Db: db.GetDb(),
	}
}

func GetLibrarianRepo() *LibrarianRepo {
	if librarianRepo == nil {
		librarianRepo = &LibrarianRepo{
			Db: db.GetDb(),
		}
	}
	return librarianRepo
}

func (repo *LibrarianRepo) SetDb(db *gorm.DB) {
	repo.Db = db
}

func (repo *LibrarianRepo) AutoMigrate() error {
	return repo.Db.AutoMigrate(&models.Librarian{})
}

func (repo *LibrarianRepo) Create(librarian *models.Librarian) error {
	return repo.Db.Create(librarian).Error
}

func (repo *LibrarianRepo) FindById(id string) (*models.Librarian, error) {
	librarian := &models.Librarian{}
	err := repo.Db.Where("id = ?", id).First(librarian).Error
	return librarian, err
}

func (repo *LibrarianRepo) FindByEmail(email string) (*models.Librarian, error) {
	librarian := &models.Librarian{}
	err := repo.Db.Where("email = ?", email).First(librarian).Error
	return librarian, err
}

func (repo *LibrarianRepo) FindAll() ([]*models.Librarian, error) {
	librarians := []*models.Librarian{}
	err := repo.Db.Find(&librarians).Error
	return librarians, err
}

func (repo *LibrarianRepo) Update(id string, librarian *models.Librarian) error {
	return repo.Db.Model(&models.Librarian{}).Where("id = ?", id).Updates(librarian).Error
}

func (repo *LibrarianRepo) Delete(id string) error {
	return repo.Db.Where("id = ?", id).Delete(&models.Librarian{}).Error
}
