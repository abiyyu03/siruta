package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
)

type VillageSeederStruct struct{}

func (s *VillageSeederStruct) Execute() error {
	if err := s.VillageSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func (s *VillageSeederStruct) VillageSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "village-seeder",
		Short: "Seeder untuk Village (Kelurahan/Desa)",
		Long:  `Seeder ini akan menambahkan data desa/kelurahan contoh di wilayah Bogor.`,
		Run: func(cmd *cobra.Command, args []string) {
			villages := []model.Village{
				{
					Name:       "Desa Bojonggede",
					AltName:    "Bojonggede",
					Latitude:   -6.485456,
					Longitude:  106.821228,
					CodePostal: "16920",
				},
				{
					Name:       "Desa Citayam",
					AltName:    "Citayam",
					Latitude:   -6.441249,
					Longitude:  106.787430,
					CodePostal: "16924",
				},
				{
					Name:       "Desa Cibinong",
					AltName:    "Cibinong",
					Latitude:   -6.485880,
					Longitude:  106.836379,
					CodePostal: "16914",
				},
				{
					Name:       "Desa Sukahati",
					AltName:    "Sukahati",
					Latitude:   -6.479192,
					Longitude:  106.836231,
					CodePostal: "16913",
				},
				{
					Name:       "Desa Nanggewer",
					AltName:    "Nanggewer",
					Latitude:   -6.471892,
					Longitude:  106.840958,
					CodePostal: "16911",
				},
			}

			tx := config.DB.Begin()
			for _, village := range villages {
				if err := tx.Create(&village).Error; err != nil {
					tx.Rollback()
					fmt.Printf("Failed to insert village %s: %v\n", village.Name, err)
					os.Exit(1)
				}
			}
			tx.Commit()

			fmt.Println("Villages inserted successfully!")
		},
	}
}
