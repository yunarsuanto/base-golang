package constants

import "time"

const (
	DefaultMailResendTimePeriod = 2 * time.Minute
)

const (
	MailChangePasswordSubject                 = "Serenus - Ubah Kata Sandi"
	MailFormVersionVerificationRequestSubject = "Pengajuan Verifikasi - %s"
	MailCustomSubject                         = "Serenus - %s"
)

const (
	DefaultChangePasswordOtpLength     = 6
	DefaultChangePasswordOtpLifeMinute = 30
)

const (
	ChangePasswordPurpose                 = "change-password"
	FormVersionVerificationRequestPurpose = "form-version-verification-request"
)
