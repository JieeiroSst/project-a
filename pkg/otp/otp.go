package otp

import (
	"github.com/twilio/twilio-go"
	"github.com/JieeiroSst/itjob/config"
)

type Otp struct {
	twilio *twilio.RestClient
}

func NewOtp(config config.Config) *Otp{
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.Twilio.TwilioAccountSid,
		Password: config.Twilio.TwilioAuthToken,
	})

	return &Otp{twilio:client}
}