package constants

import "time"

const (
	DefaultMailResendTimePeriod = 2 * time.Minute
)

const (
	MailChangePasswordSubject                 = "Primaya Hospital - Ubah Kata Sandi"
	MailFormVersionVerificationRequestSubject = "Pengajuan Verifikasi - %s"
	MailCustomSubject                         = "Primaya Hospital - %s"
)

const (
	DefaultChangePasswordOtpLength     = 6
	DefaultChangePasswordOtpLifeMinute = 30
)

const (
	ChangePasswordPurpose                 = "change-password"
	FormVersionVerificationRequestPurpose = "form-version-verification-request"
)
