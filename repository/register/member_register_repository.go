package register

import (
	"log"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/repository/member"
	"github.com/abiyyu03/siruta/repository/user"
	"gorm.io/gorm"
)

type MemberRegisterRepository struct {
	userRepository   *user.UserRepository
	memberRepository *member.MemberRepository
}

func (m *MemberRegisterRepository) RegisterMember(memberData *model.Member, user *model.User) error {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		_, err := m.memberRepository.Store(tx, memberData)

		if err != nil {
			return err
		}

		_, err = m.userRepository.RegisterUser(tx, user, user.RoleID)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Transaction failed, rolled back due to error: %v", err)
		return err
	}

	log.Println("Transaction committed successfully.")

	return nil
}
