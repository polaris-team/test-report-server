package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/polaris-team/test-report-server/pkg/config"
)

func main() {

	config.LoadConfig("../../../../src/configs", "application")

	oc := config.GetOSSConfig()

	client, err := oss.New(oc.EndPoint, oc.AccessKeyId, oc.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	// Create the bucket with default parameters
	err = client.CreateBucket(oc.BucketName)
	if err != nil {
		fmt.Println(err)
	}

	// the policy string
	var policyInfo string
	policyInfo = `
	{
		"Version":"1",
		"Statement":[
			{
				"Action":[
					"oss:GetObject",
					"oss:PutObject"
				],
				"Effect":"Deny",
				"Principal":"[123456790]",
				"Resource":["acs:oss:*:1234567890:*/*"]
			}
		]
	}`

	// Set policy
	err = client.SetBucketPolicy(oc.BucketName, policyInfo)
	if err != nil {
		fmt.Println(err)
	}

	// Get Bucket policy
	ret, err := client.GetBucketPolicy(oc.BucketName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bucket policy:", ret)
}
