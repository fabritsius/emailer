package emailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

// Mail represents a mail content
type Mail struct {
	Subject  string
	Template string
}

// Config defines parameters for SMTP connection
type Config struct {
	Name     string `env:"MAIL_NAME" default:""`
	Mail     string `env:"MAIL_ADDR"`
	Password string `env:"MAIL_PASS"`
	Server   string `env:"MAIL_SERV"`
	Port     string `env:"MAIL_PORT"`
}

// New allocates a new Mail object with template and subject
func New(template string, subject string) *Mail {
	return &Mail{
		Subject:  subject,
		Template: template,
	}
}

// SendToMany sends Mail to every recipient
func (m *Mail) SendToMany(recipients []map[string]string, cfg *Config) error {
	for _, recipient := range recipients {
		if err := m.SendTo(recipient, cfg); err != nil {
			return err
		}
	}

	return nil
}

// SendTo sends Mail to a recipient
func (m *Mail) SendTo(recipient map[string]string, cfg *Config) error {
	from := mail.Address{Name: cfg.Name, Address: cfg.Mail}
	to := mail.Address{Name: recipient["NAME"], Address: recipient["MAIL"]}

	headers := buildHeaders(from, to, m.Subject)
	messageBody, err := parseMessage(m.Template, &recipient)
	if err != nil {
		return nil
	}

	message := headers + "\r\n" + messageBody

	auth := smtp.PlainAuth("", cfg.Mail, cfg.Password, cfg.Server)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cfg.Server,
	}

	addr := fmt.Sprintf("%s:%s", cfg.Server, cfg.Port)
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, cfg.Server)
	defer client.Quit()
	if err != nil {
		return err
	}

	client.Auth(auth)

	client.Mail(from.Address)
	client.Rcpt(to.Address)

	w, err := client.Data()
	if err != nil {
		return nil
	}

	w.Write([]byte(message))

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

// parseMessage creates a message from a template
func parseMessage(temp string, data *map[string]string) (string, error) {
	t, err := template.New("email").Parse(temp)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	t.Execute(&tpl, &data)
	return tpl.String(), nil
}

// buildHeaders creates a header string
func buildHeaders(from mail.Address, to mail.Address, subject string) string {
	headers := map[string]string{
		"From":         from.String(),
		"To":           to.String(),
		"Subject":      subject,
		"Content-Type": "text/html",
		"charset":      "UTF-8",
	}

	result := ""
	for k, v := range headers {
		result += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	return result
}
