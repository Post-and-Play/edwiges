package models

import "gopkg.in/validator.v2"

type Sender struct {
	SenderMail string
	SenderPass string
}

type Receiver struct {
	ReceiverMail string
}

type MailRequest struct {
	Mail string `json:"mail" validate:"nonzero"`
}

func MailValidator(mail *MailRequest) error {
	if err := validator.Validate(mail); err != nil {
		return err
	}

	return nil
}
