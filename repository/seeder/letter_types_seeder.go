package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
)

type LetterTypeSeeder struct{}

func (r *LetterTypeSeeder) Execute() error {
	if err := r.LetterTypeSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}
func (r *LetterTypeSeeder) LetterTypeSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "letter-type-seeder",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			types := []*model.LetterType{
				{ID: 1, TypeName: "Surat Pengantar", Code: "SP"},
				{ID: 2, TypeName: "Surat Keterangan Domisili", Code: "SKD"},
				{ID: 3, TypeName: "Surat Keterangan Tidak Mampu", Code: "SKTM"},
				{ID: 4, TypeName: "Surat Keterangan Usaha", Code: "SKU"},
			}

			tx := config.DB.Begin()
			for _, letterType := range types {
				if err := tx.Create(&letterType).Error; err != nil {
					tx.Rollback()
					fmt.Printf("Failed to create role %s: %v\n", letterType.TypeName, err)
				}
			}
			tx.Commit()
		},
	}
	// return nil
}
