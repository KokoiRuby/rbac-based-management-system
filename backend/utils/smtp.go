package utils

import (
	"crypto/tls"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"gopkg.in/gomail.v2"
)

func SendEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Code)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: cfg.SkipVerify}
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}
