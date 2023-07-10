package commands

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
	"gopkg.in/gomail.v2"
)

func init() {
	registerCommand(startTryService)
}

func startTryService(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "try",
		Short: "try service",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("try aja")

			key := "AKIAXJGZHWFT7DZLSQ4P"
			secreet := "hdciS7HQua9KYoMw2sNIIs8Hulhvz1/HO1t85ZVp"
			region := "us-east-1"

			const (
				sender    = "wisnuraga418@gmail.com"
				recipient = "wisnuraga170402@gmail.com"
				subject   = "Amazon SES Test (AWS SDK for Go)"

				charset  = "UTF-8"
				HtmlBody = `<h1> Amazon SES Test (AWS SDK for Go) </h1>
							<p> This email was sent with
							<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the
							<a href='https://aws.amazon.com/sdk-for-go/'>
							AWS SDK for Go</a>.</p>`
				TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
			)

			sess, err := session.NewSession(&aws.Config{
				Region:      aws.String(region),
				Credentials: credentials.NewStaticCredentials(key, secreet, ""),
			})
			svc := ses.New(sess)

			msg := gomail.NewMessage()
			msg.SetAddressHeader("From", sender, "Wisnu Raga Prawida")
			msg.SetHeader("From", sender)
			msg.SetHeader("To", recipient)
			msg.SetHeader("Subject", subject)
			msg.SetBody("text/html", HtmlBody)

			// Send the email to Bob, Cora and Dan.
			// msg.SetHeader("Bcc", "

			var emailRaw bytes.Buffer
			msg.WriteTo(&emailRaw)

			message := ses.RawMessage{Data: emailRaw.Bytes()}

			from := sender
			to := recipient
			input := &ses.SendRawEmailInput{Source: &from, Destinations: []*string{&to}, RawMessage: &message}

			result, err := svc.SendRawEmail(input)

			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case ses.ErrCodeMessageRejected:
						log.Error(ses.ErrCodeMessageRejected, aerr.Error())
					case ses.ErrCodeMailFromDomainNotVerifiedException:
						log.Error(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
					case ses.ErrCodeConfigurationSetDoesNotExistException:
						log.Error(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
					default:
						log.Error(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					log.Error(err.Error())
				}
				return
			}

			log.Info(result.MessageId)
			log.Info("Email Sent to address: " + recipient + "!")

		},
	}
}
