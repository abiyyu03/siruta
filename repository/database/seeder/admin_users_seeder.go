package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

type AdminSeederStruct struct{}

func (a *AdminSeederStruct) Execute() error {
	if err := a.AdminSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func (a *AdminSeederStruct) AdminSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "admin-seeder",
		Short: "Seeder untuk akun admin",
		Long:  `Seeder ini akan menambahkan akun admin default ke dalam database.`,
		Run: func(cmd *cobra.Command, args []string) {
			adminID, _ := uuid.NewV7()
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Failed to hash password:", err)
				os.Exit(1)
			}

			admin := &model.User{
				ID:       adminID.String(),
				Password: string(hashedPassword),
				RoleID:   constant.ROLE_SADMIN,
				Email:    "anotheriyyu29@gmail.com",
			}

			tx := config.DB.Begin()
			if err := tx.Create(&admin).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create admin account: %v\n", err)
				os.Exit(1)
			}
			tx.Commit()

			fmt.Println("Admin account created successfully!")
		},
	}
}
