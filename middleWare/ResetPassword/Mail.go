package ResetPassword

import (
	"log"
	"net/smtp"
	"os"
)

func SendPinCode(UserMail string, pin string) error {
	// Настройка SMTP-сервера
	from := os.Getenv("Email")
	password := os.Getenv("Mail_Password")
	to := []string{UserMail}
	smtpAddr := os.Getenv("SMTP")

	// Создание сообщения
	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: Восстановление пароля\r\n\r\n" +
			"Ваш пинкод для восстановления пароля: " + pin)

	// Определение аутентификации (SSL/TLS)
	auth := smtp.PlainAuth("", from, password, os.Getenv("Host"))
	// Отправка письма
	err := smtp.SendMail(smtpAddr, auth, from, to, msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения: ", err)
		return err
	}
	log.Println("Email sent successfully")
	return nil
}
