package util

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type Asset struct {
	Key      string
	Name     string
	Location string
}

func GetPhotos() ([]Asset, error) {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		godotenv.Load()
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")),
		config.WithRegion("us-east-1"),
	)

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
