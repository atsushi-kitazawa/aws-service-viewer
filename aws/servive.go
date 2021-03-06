package aws

import (
	"errors"
	"fmt"
)

type target struct {
	service string
	region  string
}

type Result struct {
	Name       string
	Status     string
	LaunchDate string
}

func NewTarget() *target {
	return &target{}
}

func (t *target) SetService(name string) {
	t.service = name
}

func (t *target) SetRegion(name string) {
	t.region = name
}

func (t *target) DescribeTarget() ([][]string, error) {
	fmt.Println(t.service)
	switch t.service {
	case "instance":
		fmt.Println("ec2 instance region=" + t.region)
		result, err := EC2Infomation(t.region)
		if err != nil {
			fmt.Println(err)
			return [][]string{}, nil
		}
		return result2string(result), nil
	default:
		return nil, errors.New("please specify service in tree")
	}
}

func result2string(result []Result) [][]string {
	s := make([][]string, 0)
	s = append(s, []string{"name", "status", "launchdate"})
	for _, r := range result {
		ss := []string{r.Name, r.Status, r.LaunchDate}
		s = append(s, ss)
	}
	return s
}

func dummy() [][]string {
	return [][]string{{"name", "status", "launchdate"},
		{"-0000000000a", "running", "2022/05/01"},
		{"-1111111111a", "running", "2022/05/05"},
		{"-2222222222a", "stop", "2022/06/01"},
		{"-3333333333a", "stop", "2022/06/01"}}
}
