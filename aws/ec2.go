package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func EC2Infomation(region string) ([]Result, error) {
	fmt.Println(region)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstancesInput{}
	output, err := GetInstances(context.TODO(), client, input)
	if err != nil {
		return nil, err
	}

	result := make([]Result, 0)
	for _, r := range output.Reservations {
		for _, i := range r.Instances {
			ret := &Result{}
			ret.Name = nameInTag(i.Tags)
			ret.Status = string(i.State.Name)
			ret.LaunchDate = i.LaunchTime.String()
			result = append(result, *ret)
		}
	}
	return result, nil
}

func nameInTag(tags []types.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}
	return ""
}
