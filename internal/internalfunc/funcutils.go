package internalfunc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func InitGent() models.ConfigSpace {
	envPath := filepath.Dir(".env")
	_, err := os.Stat(envPath)
	if os.IsNotExist(err) {
		log.Fatal("The .env file was not found:" + envPath)
	}
	var config models.ConfigSpace

	config.Key = os.Getenv("SPACES_KEY")
	if config.Key == "" {
		panic("no key provided")
	}
	config.Secret = os.Getenv("SPACES_SECRET")
	if config.Secret == "" {
		panic("no secret provided")
	}
	config.Bucket = os.Getenv("BUCKET")
	if config.Bucket == "" {
		panic("no bucket provided")
	}
	config.Region = os.Getenv("REGION")
	if config.Region == "" {
		panic("no region provided")
	}
	config.Host = os.Getenv("URLOCEAN")
	if config.Host == "" {
		panic("no host provided")
	}

	return config
}

func DownloadURL(sess *session.Session, bucket string, filename string, duration time.Duration) ([]byte, error) {
	client := s3.New(sess)
	parsedURL, err := url.Parse(filename)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %v", err)
	}
	resp, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(parsedURL.Path),
	})

	if err != nil {
		return nil, fmt.Errorf("error downloading Spaces file: %v", err)
	}

	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading file content: %v", err)
	}

	return fileContent, nil
}

func DownloadFile(urlFile string) ([]byte, error) {
	configSpace := InitGent()
	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(configSpace.Key, configSpace.Secret, ""),
		Endpoint:    aws.String(configSpace.Host),
		Region:      aws.String(configSpace.Region),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	fileContent, err := DownloadURL(sess, configSpace.Bucket, urlFile, 0)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}
