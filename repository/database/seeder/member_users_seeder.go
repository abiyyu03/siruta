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

type MemberSeederStruct struct{}

func (a *MemberSeederStruct) Execute() error {
	if err := a.MemberSeeder().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func (a *MemberSeederStruct) MemberSeeder() *cobra.Command {
	return &cobra.Command{
		Use:   "admin-seeder",
		Short: "Seeder untuk akun member",
		Long:  `Seeder ini akan menambahkan akun member default ke dalam database.`,
		Run: func(cmd *cobra.Command, args []string) {
			adminID, _ := uuid.NewV7()
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Failed to hash password:", err)
				os.Exit(1)
			}

			user := &model.User{
				ID:       adminID.String(),
				Password: string(hashedPassword),
				RoleID:   constant.ROLE_MEMBER,
				Email:    "member1@satuwarga.com",
			}

			member := &model.Member{
				ID:             "member1",
				Fullname:       "Abiyyu Cakra",
				NikNumber:      func() *string { s := "166567127010003"; return &s }(),
				KKNumber:       func() *string { s := "1665671270100012"; return &s }(),
				BornPlace:      "Bogor",
				BirthDate:      "20-12-2000",
				Gender:         "L",
				HomeAddress:    func() *string { s := "Perum sana sini jaya"; return &s }(),
				MaritalStatus:  func() *string { s := "Abiyyu Cakra"; return &s }(),
				ReligionId:     1,
				MemberStatusId: 1,
				UserId:         adminID.String(),
				Photo:          func() *string { s := "Abiyyu Cakra.png"; return &s }(),
				Occupation:     func() *string { s := "asdasd"; return &s }(),
				Status:         "resident",
				RTProfileId:    "rt_leader_profile",
			}

			tx := config.DB.Begin()
			if err := tx.Create(&user).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create member account: %v\n", err)
				os.Exit(1)
			}
			if err := tx.Create(&member).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Failed to create member account: %v\n", err)
				os.Exit(1)
			}
			tx.Commit()

			fmt.Println("Admin account created successfully!")
		},
	}
}
