package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
)

type MemberStatusSeederStruct struct{}

func (r *MemberStatusSeederStruct) Execute() error {
	if err := r.MemberStatusSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}
func (r *MemberStatusSeederStruct) MemberStatusSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "member-status-seeder",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			memberStatuses := []*model.MemberStatus{
				{ID: 1, Status: "Kepala Keluarga"},
				{ID: 2, Status: "Suami"},
				{ID: 3, Status: "Istri"},
				{ID: 4, Status: "Anak"},
				{ID: 5, Status: "Menantu"},
				{ID: 6, Status: "Cucu"},
				{ID: 7, Status: "Orangtua"},
				{ID: 8, Status: "Mertua"},
				{ID: 9, Status: "Famili Lain"},
				{ID: 10, Status: "Pembantu"},
				{ID: 11, Status: "Lainnya"},
			}

			tx := config.DB.Begin()
			for _, memberStatus := range memberStatuses {
				if err := tx.Create(&memberStatus).Error; err != nil {
					tx.Rollback()
					fmt.Printf("Failed to create role %s: %v\n", memberStatus.Status, err)
				}
			}
			tx.Commit()
		},
	}
	// return nil
}
