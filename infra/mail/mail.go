package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/db"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	gomail "gopkg.in/mail.v2"
)

type MailInterface interface {
	SendChangePasswordOtp(data objects.SendChangePasswordOtpMailData) *constants.ErrorResponse
	SendFormVersionApproval(data objects.SendFormVersionApprovalMailData) *constants.ErrorResponse
	SendCustom(data objects.SendCustomMailData) *constants.ErrorResponse
}

type mailService struct {
	*db.DB
	*repository.RepoCtx
	*config.Config
}

func NewMailService(dbApp *db.DB, repo *repository.RepoCtx, config *config.Config) MailInterface {
	return &mailService{
		dbApp,
		repo,
		config,
	}
}

func (a mailService) sendMail(recipients []string, subject, body string) *constants.ErrorResponse {
	message := gomail.NewMessage()

	message.SetHeader("From", a.Mailer.Sender)
	message.SetHeader("To", recipients...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(a.Mailer.Host, a.Mailer.Port, a.Mailer.Username, a.Mailer.Password)
	if err := dialer.DialAndSend(message); err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a mailService) createMailTemplate(files []string, data any) (string, *constants.ErrorResponse) {
	var result string

	if len(files) == 0 {
		return result, utils.ErrorInternalServer("no template files provided")
	}

	ctx := context.Background()
	defer ctx.Done()
	tx, err := a.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	templateFuncMap := template.FuncMap{
		"year": func() string { return time.Now().Format("2006") },
		"logo": func() template.HTML {
			return template.HTML(a.Server.LogoUrl)
		},
	}

	files = append([]string{constants.BaseMailTemplate}, files...)
	t, err := template.New(filepath.Base(constants.BaseMailTemplate)).Funcs(templateFuncMap).ParseFiles(files...)
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	result = buf.String()

	if err := tx.Commit(); err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a mailService) SendChangePasswordOtp(data objects.SendChangePasswordOtpMailData) *constants.ErrorResponse {
	templateFiles := []string{
		constants.ForgotPasswordMailTemplate,
	}

	body, errs := a.createMailTemplate(templateFiles, data)
	if errs != nil {
		return errs
	}

	if errs := a.sendMail([]string{data.To}, constants.MailChangePasswordSubject, body); errs != nil {
		return errs
	}

	return nil
}

func (a mailService) SendFormVersionApproval(data objects.SendFormVersionApprovalMailData) *constants.ErrorResponse {
	templateFiles := []string{
		constants.FormVersionVerificationRequestMailTemplate,
	}

	body, errs := a.createMailTemplate(templateFiles, data)
	if errs != nil {
		return errs
	}

	if errs := a.sendMail([]string{data.To}, fmt.Sprintf(constants.MailFormVersionVerificationRequestSubject, data.FormName), body); errs != nil {
		return errs
	}

	return nil
}

func (a mailService) SendCustom(data objects.SendCustomMailData) *constants.ErrorResponse {
	templateFiles := []string{
		constants.CustomMailTemplate,
	}

	body, errs := a.createMailTemplate(templateFiles, data)
	if errs != nil {
		return errs
	}

	recipients := utils.RemoveDuplicate(data.To)
	for i := 0; i < len(recipients); i += constants.MaximumSendBatch {
		lastIndex := i + constants.MaximumSendBatch
		if lastIndex > len(recipients) {
			lastIndex = len(recipients)
		}

		go func() {
			if errs := a.sendMail(recipients[i:lastIndex], fmt.Sprintf(constants.MailCustomSubject, data.Title), body); errs != nil {
				utils.PrintError(*errs)
			}
		}()
	}

	return nil
}
