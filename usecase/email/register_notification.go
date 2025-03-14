package email

import (
	"fmt"
	"net/smtp"

	"github.com/abiyyu03/siruta/helper"
	"github.com/gofiber/fiber/v2/log"
)

// func SendNotificationAfterRWProfileRegistration(ctx fiber.Ctx) error {}
type EmailRegistrationUsecase struct{}

func (e *EmailRegistrationUsecase) RtNotification(emailDestination string, token string) error {
	from := helper.GoDotEnv("EMAIL_FROM")
	password := helper.GoDotEnv("EMAIL_PASSWORD")
	smtpHost := helper.GoDotEnv("EMAIL_PROT_HOST")
	smtpPort := helper.GoDotEnv("EMAIL_PROT_PORT")

	url := "https://satuwarga.id/register/rt?token=" + token

	message := []byte(
		"Subject: SATUWARGA User Account (Ketua RT)\n\nBerikut adalah tauran registrasi untuk anda, jangan berikan url ini ke pihak manapun\n\nURL Pendaftaran : " + url,
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

func (e *EmailRegistrationUsecase) RwNotification(emailDestination string, token string) error {
	from := helper.GoDotEnv("EMAIL_FROM")
	password := helper.GoDotEnv("EMAIL_PASSWORD")
	smtpHost := helper.GoDotEnv("EMAIL_PROT_HOST")
	smtpPort := helper.GoDotEnv("EMAIL_PROT_PORT")

	url := "https://satuwarga.id/register/rw?token=" + token

	message := []byte(
		"Subject: SATUWARGA User Account (Ketua RW)\n\nBerikut adalah tauran registrasi untuk anda, jangan berikan url ini ke pihak manapun\n\nURL Pendaftaran : " + url,
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
