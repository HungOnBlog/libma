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
