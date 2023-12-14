package resources

import (
	_ "embed"
)

var (
	//go:embed mail_templates/reset_password_otp.html
	ResetPasswordOtpTemplate string

	//go:embed mail_templates/password_changed_notification.html
	PasswordChangedNotificationTemplate string

	//go:embed mail_templates/email_verification_otp.html
	EmailVerificationOtpTemplate string
)
