package mailer

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

type MailConfig struct {
	Mail string
	User string
	Pass string
}

func SendVerMail(mc MailConfig, to, token string) error {
	from := mc.Mail

	addr := "smtp.yandex.ru:587"
	host := "smtp.yandex.ru"

	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = "SingUp verification"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body := "Для подтверждения электонной почты пройдите по ссылке:<br><br><a href=\"http://localhost:8888/auth/activate?token=" + token + "\">Подтвердить адрес</a>"

	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	user := mc.User
	password := mc.Pass
	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(message))

	if err != nil {
		return err
	}

	return nil
}
