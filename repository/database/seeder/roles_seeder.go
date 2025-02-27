package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
)

type RoleSeederStruct struct{}

func (r *RoleSeederStruct) Execute() error {
	if err := r.RoleSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}
func (r *RoleSeederStruct) RoleSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "role-seeder",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			roles := []*model.Role{
				{ID: 1, Name: "Super Admin"},
				{ID: 2, Name: "Pengurus RW"},
				{ID: 3, Name: "Pengurus RT"},
				{ID: 4, Name: "Member"},
			}

			tx := config.DB.Begin()
			for _, role := range roles {
				if err := tx.Create(&role).Error; err != nil {
					tx.Rollback()
					fmt.Printf("Failed to create role %s: %v\n", role.Name, err)
				}
			}
			tx.Commit()
		},
	}
	// return nil
}
