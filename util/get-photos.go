package util

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Asset struct {
	Key      string
	Name     string
	Location string
}

func GetPhotos() ([]Asset, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("z-dev-profile"))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	response, err := client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String("z-photos-bucket"),
	})

	if err != nil {
		log.Fatal(err)
	}

	contents := response.Contents

	var data []Asset

	for i := 0; i < len(contents); i++ {
		metadata, err := client.HeadObject(context.Background(), &s3.HeadObjectInput{
			Bucket: aws.String("z-photos-bucket"),
			Key:    contents[i].Key,
		})
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Asset{Key: *contents[i].Key, Name: metadata.Metadata["name"], Location: metadata.Metadata["location"]})
	}

	return data, err
}
