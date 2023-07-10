package commands

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func init() {
	registerCommand(AWSSNS)
}

func AWSSNS(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "sns",
		Short: "try service",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("try aja")

			key := "AKIAXJGZHWFT7DZLSQ4P"
			secreet := "hdciS7HQua9KYoMw2sNIIs8Hulhvz1/HO1t85ZVp"
			region := "us-east-1"

			sess, err := session.NewSession(&aws.Config{
				Region:      aws.String(region),
				Credentials: credentials.NewStaticCredentials(key, secreet, ""),
			})
			// svc := ses.New(sess)

			if err != nil {
				log.Error(err, " Session ")

			}
			//create SES client
			svc := ses.New(sess)

			// create an SNS client
			snsClient := sns.New(sess)

			// create a new SNS topic
			createTopicInput := &sns.CreateTopicInput{
				Name: aws.String("my-sns-topic"),
			}
			createTopicOutput, err := snsClient.CreateTopic(createTopicInput)
			if err != nil {
				log.Error(err, " Create Topic Error")
			}
			topicArn := createTopicOutput.TopicArn
			log.Info("topicArn :", *topicArn)

			// create an SQS client
			sqsClient := sqs.New(sess)

			// create a new SQS queue
			createQueueInput := &sqs.CreateQueueInput{
				QueueName: aws.String("my-sqs-queue"),
			}
			createQueueOutput, err := sqsClient.CreateQueue(createQueueInput)
			if err != nil {
				log.Error(err, " Create Queue Error")
			}
			queueUrl := createQueueOutput.QueueUrl
			log.Info("queueUrl :", *queueUrl)
			// Get the queue ARN
			getQueueAttributesInput := &sqs.GetQueueAttributesInput{
				QueueUrl:       aws.String(*queueUrl),
				AttributeNames: []*string{aws.String("QueueArn")},
			}
			getQueueAttributesOutput, err := sqsClient.GetQueueAttributes(getQueueAttributesInput)
			if err != nil {
				log.Info("Error getting queue attributes", err)
				return
			}
			queueArn := *getQueueAttributesOutput.Attributes["QueueArn"]
			log.Info("Queue ARN:", queueArn)

			// subscribe the SQS queue to the SNS topic
			subscribeInput := &sns.SubscribeInput{
				TopicArn: topicArn,
				Protocol: aws.String("sqs"),
				Endpoint: aws.String(queueArn),
			}
			_, err = snsClient.Subscribe(subscribeInput)
			if err != nil {
				log.Error(err, " Subscribe Error")
			}

			//set permissions policy

			// set the permission policy
			// policy := `{
			// 	"Version":"2012-10-17",
			// 	"Statement":[
			// 		{
			// 			"Effect":"Allow",
			// 			"Principal":"*",
			// 			"Action":"sqs:SendMessage",
			// 			"Resource":"` + queueArn + `",
			// 			"Condition":{
			// 				"ArnEquals":{
			// 					"aws:SourceArn":"` + *topicArn + `"
			// 				}
			// 			}
			// 		}
			// 	]
			// }`

			// set the parameters for adding the permission
			// params := &sqs.AddPermissionInput{
			// 	QueueUrl: aws.String(*queueUrl),
			// 	Label:    aws.String("SNSPublishPermission"),
			// 	AWSAccountIds: []*string{
			// 		aws.String("*"),
			// 	},
			// 	Actions: []*string{
			// 		aws.String("sqs:SendMessage"),
			// 	},

			// 	}
			// }

			// add the permission
			_, err = sqsClient.AddPermission(&sqs.AddPermissionInput{
				Label:    aws.String("SNSPublishPermission"),
				QueueUrl: aws.String(*queueUrl),
				AWSAccountIds: []*string{
					aws.String("500819276135"),
				},
				Actions: []*string{
					aws.String("SendMessage"),
					aws.String("ReceiveMessage"),
					aws.String("DeleteMessage"),
					aws.String("GetQueueAttributes"),
					aws.String("GetQueueUrl"),
				},
			})

			if err != nil {
				// handle error
				log.Error(err, " Add Permission Error")
				return
			}

			configSetName := "my-config-set"

			// Create the configuration set.
			_, err = svc.CreateConfigurationSet(&ses.CreateConfigurationSetInput{
				ConfigurationSet: &ses.ConfigurationSet{
					Name: aws.String(configSetName),
				},
			})
			if err != nil {
				log.Error(err, " configset error")
			}

			// Specify the types of events to publish
			eventTypes := []*string{
				aws.String("send"),
				aws.String("reject"),
				aws.String("bounce"),
				aws.String("complaint"),
				aws.String("delivery"),
			}

			// configuration set input
			input := &ses.CreateConfigurationSetEventDestinationInput{
				ConfigurationSetName: aws.String(configSetName),
				EventDestination: &ses.EventDestination{
					Name:               aws.String("my-destination"),
					MatchingEventTypes: eventTypes,
					Enabled:            aws.Bool(true),
					SNSDestination: &ses.SNSDestination{
						TopicARN: topicArn,
					},
				},
			}

			// Create the configuration set and specify the destination for event publishing
			_, err = svc.CreateConfigurationSetEventDestination(input)
			if err != nil {
				log.Error(err, " config set Error")
			}

		},
	}
}
