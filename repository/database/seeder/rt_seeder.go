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

type RTProfileSeederStruct struct{}

func (s *RTProfileSeederStruct) Execute() error {
	if err := s.RTProfileSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func (s *RTProfileSeederStruct) RTProfileSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "rtprofile-seeder",
		Short: "Seeder untuk RT Profile dan RT Leader",
		Long:  `Seeder ini akan menambahkan data RT Profile beserta akun RT Leader default ke dalam database.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Buat RT Profile
			rtProfile := &model.RTProfile{
				ID:           "rt_leader_profile", // gunakan ID tetap untuk RT Profile
				RTNumber:     "01",
				RTLogo:       func() *string { s := "rt_logo.png"; return &s }(),
				Latitude:     -6.234567,
				Longitude:    106.987654,
				IsAuthorized: true,
				RTEmail:      "rt01@example.com",
				MobilePhone:  "081234567890",
				FullAddress:  "Jl. Contoh RT 01 No. 123",
				RWProfileId:  "rw_leader_profile", // sesuaikan!
			}

			// Buat akun User RT Leader
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Failed to hash password:", err)
				os.Exit(1)
			}

			user := &model.User{
				ID:       "rt_leader_user", // gunakan ID tetap untuk User RT Leader
				Email:    "rtleader@example.com",
				Password: string(hashedPassword),
				RoleID:   constant.ROLE_RT, // pastikan konstanta ada
			}

			// Buat RT Leader
			rtLeader := &model.RTLeader{
				ID:          "rt_leader_detail",
				Fullname:    "Pak RT Satu",
				NikNumber:   "3201234567890001",
				KKNumber:    "3201234567890002",
				RTProfileId: "rt_leader_profile",
				Photo:       func() *string { s := "pak_rt_photo.png"; return &s }(),
				UserId:      "rt_leader_user",
				FullAddress: "Jl. Contoh RT 01 No. 123",
			}

			// Simpan dengan transaction
			tx := config.DB.Begin()
			if err := tx.Create(&rtProfile).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RT Profile: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Create(&user).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RT Leader User: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Create(&rtLeader).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create RT Leader: %v\n", err)
				os.Exit(1)
			}
			tx.Commit()

			fmt.Println("RT Profile & RT Leader created successfully!")
		},
	}
}
