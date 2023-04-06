package mailKit

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/richelieu42/chimera/v2/consts"
	"github.com/richelieu42/chimera/v2/core/errorKit"
	"github.com/richelieu42/chimera/v2/core/strKit"
	"net/smtp"
	"sync"
	"time"
)

type (
	SmtpConfig struct {
		// e.g. "smtp.163.com:25"
		Address string
		// e.g. "smtp.163.com"
		Host string

		// 邮箱账号
		Account string
		// 邮箱的授权码
		Password string
		// 发件人的昵称（默认: consts.OwnName）
		NickName string
	}
)

var smtpLock = &sync.RWMutex{}
var smtpPool *email.Pool
var defaultFrom string

// InitializeSmtp
/*
e.g.
初始化传参count为2，同时发10条邮件，并不会丢失邮件，10条邮件都会被正常发出去.
*/
func InitializeSmtp(config *SmtpConfig, count int) error {
	smtpLock.Lock()
	defer smtpLock.Unlock()

	if config == nil {
		return errorKit.Simple("config == nil")
	}
	config.NickName = strKit.EmptyToDefault(config.NickName, consts.OwnName, true)
	defaultFrom = fmt.Sprintf("%s <%s>", config.NickName, config.Account)

	if count <= 0 {
		return errorKit.Simple("count(%d) is invalid", count)
	}

	auth := smtp.PlainAuth("", config.Account, config.Password, config.Host)
	if p, err := email.NewPool(config.Address, count, auth); err != nil {
		smtpPool = nil
		return err
	} else {
		smtpPool = p
		return nil
	}
}

// Dispose 释放资源
func Dispose() {
	smtpLock.Lock()
	defer smtpLock.Unlock()

	if smtpPool != nil {
		smtpPool.Close()
		smtpPool = nil
	}
}

func NewMail(from string, to []string, subject string, text, html []byte, cc, bcc []string) *email.Email {
	mail := email.NewEmail()
	mail.From = from
	mail.To = to
	mail.Subject = subject
	mail.Text = text
	mail.HTML = html
	mail.Cc = cc
	mail.Bcc = bcc
	return mail
}

func SendMail(mail *email.Email) error {
	smtpLock.RLock()
	defer smtpLock.RUnlock()

	if smtpPool == nil {
		return errorKit.Simple("smtp pool hasn't been initialized")
	}
	if mail == nil {
		return errorKit.Simple("mail == nil")
	}
	mail.From = strKit.EmptyToDefault(mail.From, defaultFrom, true)
	return smtpPool.Send(mail, time.Second*6)
}
