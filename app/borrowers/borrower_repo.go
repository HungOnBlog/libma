package borrowers

import (
	"github.com/HungOnBlog/libma/core/models"
	"github.com/HungOnBlog/libma/infra/db"
	"gorm.io/gorm"
)

type BorrowerRepo struct {
	Db *gorm.DB
}

var borrowerRepo *BorrowerRepo

func init() {
	borrowerRepo = &BorrowerRepo{
		Db: db.GetDb(),
	}
}

func GetBorrowerRepo() *BorrowerRepo {
	if borrowerRepo == nil {
		borrowerRepo = &BorrowerRepo{
			Db: db.GetDb(),
		}
	}
	return borrowerRepo
}

func (repo *BorrowerRepo) SetDb(db *gorm.DB) {
	repo.Db = db
}

func (repo *BorrowerRepo) AutoMigrate() error {
	return repo.Db.AutoMigrate(&models.Borrower{})
}

func (repo *BorrowerRepo) Create(borrower *models.Borrower) error {
	return repo.Db.Create(borrower).Error
}

func (repo *BorrowerRepo) FindById(id string) (*models.Borrower, error) {
	borrower := &models.Borrower{}
	err := repo.Db.Where("id = ?", id).First(borrower).Error
	return borrower, err
}

func (repo *BorrowerRepo) FindByEmail(email string) (*models.Borrower, error) {
	borrower := &models.Borrower{}
	err := repo.Db.Where("email = ?", email).First(borrower).Error
	return borrower, err
}

func (repo *BorrowerRepo) FindAll() ([]*models.Borrower, error) {
	borrowers := []*models.Borrower{}
	err := repo.Db.Find(&borrowers).Error
	return borrowers, err
}

func (repo *BorrowerRepo) Update(borrower *models.Borrower) error {
	return repo.Db.Model(&models.Borrower{}).Where("id = ?", borrower.Id).Updates(borrower).Error
}

func (repo *BorrowerRepo) DeleteById(id string) error {
	return repo.Db.Where("id = ?", id).Delete(&models.Borrower{}).Error
}
