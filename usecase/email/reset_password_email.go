package email

import (
	"fmt"
	"net/smtp"

	"github.com/abiyyu03/siruta/helper"
	"github.com/gofiber/fiber/v2/log"
)

type ResetPasswordEmailUsecase struct{}

func (e *ResetPasswordEmailUsecase) ResetPasswordEmail(emailDestination string, token string) error {
	from := helper.GoDotEnv("EMAIL_FROM")
	password := helper.GoDotEnv("EMAIL_PASSWORD")
	smtpHost := helper.GoDotEnv("EMAIL_PROT_HOST")
	smtpPort := helper.GoDotEnv("EMAIL_PROT_PORT")

	url := "https://satuwarga.id/reset-password?token=" + token

	message := []byte(
		"Subject: SATUWARGA User Account (Ketua RW)\n\nBerikut adalah tautan reset password untuk anda, jangan berikan url ini ke pihak manapun\n\nURL Pendaftaran : " + url,
	)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{emailDestination}, message)
	if err != nil {
		log.Error("Error Email Notification:", err.Error())
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
