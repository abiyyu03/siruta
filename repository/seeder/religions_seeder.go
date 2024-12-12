package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
)

type ReligionSeederStruct struct{}

func (r *ReligionSeederStruct) Execute() error {
	if err := r.ReligionSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}
func (r *ReligionSeederStruct) ReligionSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "religion-seeder",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			religions := []*model.Religion{
				{ID: 1, ReligionName: "Islam"},
				{ID: 2, ReligionName: "Kristen"},
				{ID: 3, ReligionName: "Katholik"},
				{ID: 4, ReligionName: "Hindu"},
				{ID: 5, ReligionName: "Budha"},
				{ID: 6, ReligionName: "Konghucu"},
				{ID: 7, ReligionName: "Lainnya"},
			}

			tx := config.DB.Begin()
			for _, religion := range religions {
				if err := tx.Create(&religion).Error; err != nil {
					tx.Rollback()
					fmt.Printf("Failed to create role %s: %v\n", religion.ReligionName, err)
				}
			}
			tx.Commit()
		},
	}
	// return nil
}
