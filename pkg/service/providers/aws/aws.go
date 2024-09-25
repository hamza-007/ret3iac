package aws

import (
	resourceModel "github.com/hamza-007/ret3iac/pkg/model/resource"

	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	s3 "github.com/aws/aws-sdk-go/service/s3"
)

type awsService interface {
	CreateEC2Instance(*resourceModel.Resource) error
	CreateS3Bucket(*resourceModel.Resource) error
}

type awsWriter struct{}

func AwsWriter() awsService {
	return awsWriter{}
}

// Create EC2 instance
func (a awsWriter) CreateEC2Instance(node *resourceModel.Resource) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(node.Config["region"]),
	}))
	svc := ec2.New(sess)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(node.Config["ami"]),
		InstanceType: aws.String(node.Config["instance_type"]),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	_, err := svc.RunInstances(input)
	if err != nil {
		return err
	}
	return nil
}

// Create S3 bucket
func (a awsWriter) CreateS3Bucket(node *resourceModel.Resource) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(node.Config["region"]),
	}))
	svc := s3.New(sess)

	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(node.Config["bucket"]),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(node.Config["region"]),
		},
	})

	if err != nil {
		return err
	}
	return nil
}
