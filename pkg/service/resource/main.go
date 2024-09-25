package service

import (
	"fmt"

	resourceModel "github.com/hamza-007/ret3iac/pkg/model/resource"
	awsService "github.com/hamza-007/ret3iac/pkg/service/providers/aws"
)

type resourceWriter struct{ }

func ResourceWriter() resourceWriter {
	return resourceWriter{}
}

func (r resourceWriter) Interpret(nodes []*resourceModel.Resource) error {
	for _, node := range nodes {
		switch node.Type {
		case "aws_instance": {
			if err := awsService.AwsWriter().CreateEC2Instance(node); err != nil {
				return err
			}
			break;
		}
		case "aws_s3_bucket":
			{
				if err := awsService.AwsWriter().CreateS3Bucket(node); err != nil {
					return err
				}
				break;
			}
		default:
			return fmt.Errorf("unknown resource type: %s", node.Type)
		}
	}
	return nil
}
