package provider

import (
	"sqs/config"
	"sqs/utils"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetService() *sqs.SQS {
	svc, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewSharedCredentials("", config.CredProfile),
	})
	utils.CheckError(err)

	_, err = svc.Config.Credentials.Get()
	utils.CheckError(err)

	return sqs.New(svc)
}

func SendMessage(svc *sqs.SQS, message string) {

	qURL := config.QueueURL

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:               &qURL,
		MessageBody:            aws.String(message),
		MessageGroupId:         aws.String("bills"),
		MessageDeduplicationId: aws.String(now),
	})
	utils.CheckError(err)

}

func GetMessage(svc *sqs.SQS) *sqs.Message {

	qURL := config.QueueURL

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames:        []*string{aws.String(sqs.MessageSystemAttributeNameSentTimestamp)},
		MessageAttributeNames: []*string{aws.String(sqs.QueueAttributeNameAll)},
		QueueUrl:              &qURL,
		MaxNumberOfMessages:   aws.Int64(1),
		VisibilityTimeout:     aws.Int64(10), // seconds
		WaitTimeSeconds:       aws.Int64(0),
	})
	utils.CheckError(err)

	if len(result.Messages) == 0 {
		return nil
	}

	return result.Messages[0]

}

func DeleteMessage(svc *sqs.SQS, message *sqs.Message) {

	qURL := config.QueueURL

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: message.ReceiptHandle,
	})
	utils.CheckError(err)

}
