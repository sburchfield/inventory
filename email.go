
package main

import (
	mailer "gnardex/mailer"
)

func sendPasswordResetEmail(emailTo string, payload interface{}) error {

	mailer.From = "saburchfield@gmail.com"
	mailer.ReplyTo = "saburchfield@gmail.com"
	subject := "Inventory Password Reset Link"

	tag := "password_reset"

	htmlpayload, err := mailer.GetTemplateHtml(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	textpayload, err := mailer.GetTemplateText(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	msg := mailer.PrepareMailgunMessage(emailTo, subject, htmlpayload, textpayload)
	if err := mailer.SendMailgunMessage(msg); err != nil {

		handleCriticalError(err)
		return err

	}

	return nil

}

func sendSignupEmail(emailTo string) error {

	mailer.From = "saburchfield@gmail.com"
	mailer.ReplyTo = "saburchfield@gmail.com"
	subject := "Inventory Password Reset Link"

	payload := struct {
		AppVersion string
	}{
		AppVersion: "inventory-backend",
	}

	tag := "signup"

	htmlpayload, err := mailer.GetTemplateHtml(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	textpayload, err := mailer.GetTemplateText(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	msg := mailer.PrepareMailgunMessage(emailTo, subject, htmlpayload, textpayload)
	if err := mailer.SendMailgunMessage(msg); err != nil {

		handleCriticalError(err)
		return err

	}

	return nil

}


func sendOrdersEmail(emailTo string, payload interface{}) error {

	mailer.From = "saburchfield@gmail.com"
	mailer.ReplyTo = "saburchfield@gmail.com"
	subject := "Inventory Nightly Order"

	tag := "order"

	htmlpayload, err := mailer.GetTemplateHtml(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	textpayload, err := mailer.GetTemplateText(tag, payload)
	if err != nil {

		handleCriticalError(err)
		return err

	}

	msg := mailer.PrepareMailgunMessage(emailTo, subject, htmlpayload, textpayload)
	if err := mailer.SendMailgunMessage(msg); err != nil {

		handleCriticalError(err)
		return err

	}

	return nil

}
