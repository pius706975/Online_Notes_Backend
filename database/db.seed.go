package database

import (
	"fmt"
	"log"

	"github.com/pius706975/backend/database/models"
	"github.com/pius706975/backend/database/seeders"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type seedData struct {
	name  string
	model interface{}
	size  int
}

var SeedCMD = &cobra.Command{
	Use: "seed",
	Short: "For running db seeder",
	RunE: Seed,
}

var seedUP bool
var seedDOWN bool

func init()  {
	SeedCMD.Flags().BoolVarP(&seedUP, "seedUP", "u", true, "run seed up")

	SeedCMD.Flags().BoolVarP(&seedDOWN, "seedDOWN", "d", false, "run seed down")
}

func Seed(cmd *cobra.Command, args []string) error {
	
	var err error

	db, err := NewDB()
	if err != nil {
		return err
	}

	if seedDOWN {
		err = seedDown(db)
		return err
	}

	if seedUP {
		err = seedUp(db)
	}

	return err
}

func seedUp(db *gorm.DB) error {
	
	var err error

	var seedModel = []seedData{
		{
			name: "user",
			model: seeders.UserSeeds,
			size: cap(seeders.UserSeeds),
		},
	}

	for _, data := range seedModel {

		log.Println("create seeding data for ", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err
}

func seedDown(db *gorm.DB) error {
	
	var err error

	var seedModel = []seedData{
		{
			name: models.User{}.TableName(),
			model: models.User{},
		},
	}

	for _, data := range seedModel {
		log.Println("Delete seeding data for ", data.name)
		sql := fmt.Sprintf("DELETE FROM %v ", data.name)
		err = db.Exec(sql).Error
	}

	return err
}