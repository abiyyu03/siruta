package email

import (
	"fmt"
	"net/smtp"

	"github.com/abiyyu03/siruta/helper"
	"github.com/abiyyu03/siruta/repository/register"
	"github.com/gofiber/fiber/v2/log"
)

// func SendNotificationAfterRWProfileRegistration(ctx fiber.Ctx) error {}
type EmailRegistrationRepository struct {
	regToken *register.RegTokenRepository
}

func (e *EmailRegistrationRepository) RtNotification(emailDestination string) error {
	from := helper.GoDotEnv("EMAIL_FROM")
	password := helper.GoDotEnv("EMAIL_PASSWORD")
	smtpHost := helper.GoDotEnv("EMAIL_PROT_HOST")
	smtpPort := helper.GoDotEnv("EMAIL_PROT_PORT")

	token, _ := e.regToken.CreateToken()

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

func (e *EmailRegistrationRepository) RwNotification(emailDestination string) error {
	from := helper.GoDotEnv("EMAIL_FROM")
	password := helper.GoDotEnv("EMAIL_PASSWORD")
	smtpHost := helper.GoDotEnv("EMAIL_PROT_HOST")
	smtpPort := helper.GoDotEnv("EMAIL_PROT_PORT")

	token, _ := e.regToken.CreateToken()

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

// func (e *EmailRegistrationRepository) RejectNotification(emailDestination string) {}
