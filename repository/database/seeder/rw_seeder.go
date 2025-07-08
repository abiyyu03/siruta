package seeder

import (
	"fmt"
	"os"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

type RWProfileSeederStruct struct{}

func (s *RWProfileSeederStruct) Execute() error {
	if err := s.RWProfileSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func (s *RWProfileSeederStruct) RWProfileSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "rwprofile-seeder",
		Short: "Seeder untuk RW Profile dan RW Leader",
		Long:  `Seeder ini akan menambahkan data RW Profile beserta akun RW Leader default ke dalam database.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Buat RW Profile
			rwProfile := &model.RWProfile{
				ID:           "rw_leader_profile",
				RWNumber:     "01",
				VillageID:    1,
				RWLogo:       func() *string { s := "rw_logo.png"; return &s }(),
				Latitude:     -6.234567,
				Longitude:    106.987654,
				IsAuthorized: true,
				RwEmail:      "rw01@example.com",
				MobilePhone:  "081234567891",
				RegencyLogo:  func() *string { s := "regency_logo.png"; return &s }(),
				FullAddress:  "Jl. Contoh RW 01 No. 456",
			}

			// Buat User RW Leader
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Failed to hash password:", err)
				os.Exit(1)
			}

			user := &model.User{
				ID:       "rw_leader_user",
				Email:    "rwleader@example.com",
				Password: string(hashedPassword),
				RoleID:   constant.ROLE_RW, // pastikan konstanta ada
			}

			// Buat RW Leader
			rwLeader := &model.RWLeader{
				ID:          "rw_leader_detail",
				Fullname:    "Pak RW Satu",
				NikNumber:   "3201234567899999",
				KKNumber:    "3201234567898888",
				Photo:       func() *string { s := "pak_rw_photo.png"; return &s }(),
				RWProfileId: "rw_leader_profile",
				UserId:      "rw_leader_user",
				FullAddress: "Jl. Contoh RW 01 No. 456",
			}

			// Simpan transaction
			tx := config.DB.Begin()
			if err := tx.Create(&rwProfile).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RW Profile: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Create(&user).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RW Leader User: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Create(&rwLeader).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RW Leader: %v\n", err)
				os.Exit(1)
			}
			tx.Commit()

			fmt.Println("RW Profile & RW Leader created successfully!")
		},
	}
}
