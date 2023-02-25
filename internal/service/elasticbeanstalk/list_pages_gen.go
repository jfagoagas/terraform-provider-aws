// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeEnvironments -ContextOnly"; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func describeEnvironmentsPages(ctx context.Context, conn *elasticbeanstalk.ElasticBeanstalk, input *elasticbeanstalk.DescribeEnvironmentsInput, fn func(*elasticbeanstalk.EnvironmentDescriptionsMessage, bool) bool) error {
	for {
		output, err := conn.DescribeEnvironmentsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
