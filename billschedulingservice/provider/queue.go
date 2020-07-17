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
