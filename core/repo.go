package core

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type IRepository interface {
	SetDb(db *gorm.DB)
	AutoMigrate() error
}

func AutoMigrate(repos ...IRepository) {
	for _, repo := range repos {
		err := repo.AutoMigrate()
		if err != nil {
			panic(err)
		}

		fmt.Println("AutoMigrate: ", reflect.TypeOf(repo).Elem().Name())
	}
}
