package download

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/elo7/harbor/config"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

//FromS3 get files from s3
func FromS3(harborConfig config.HarborConfig) error {
	if len(harborConfig.Files) < 0 {
		return nil
	}

	if err := checkAWSRegions(harborConfig.S3Repositories); err != nil {
		return err
	}

	// region := aws.USEast
	awsAuth, err := aws.GetAuth("", "", "", time.Time{})
	if err != nil {
		return err
	}

	downloadPath := fmt.Sprintf("%s/%s", harborConfig.ProjectPath, harborConfig.DownloadPath)
	fmt.Printf("--- %d Files to be downloaded on path: %s\r\n", len(harborConfig.Files), downloadPath)

	for _, file := range harborConfig.Files {
		if err := downloadFile(awsAuth, harborConfig.S3Repositories.Get(file.Repository), file, downloadPath); err != nil {
			return err
		}
	}
	return nil
}

func downloadFile(awsAuth aws.Auth, s3config config.S3HarborConfig, file config.HarborFile, downloadPath string) error {
	s3FilePath := filepath.Join(s3config.BasePath, file.S3Path)
	outputFilePath := filepath.Join(downloadPath, file.FileName)
	bucket := getBucket(s3config, awsAuth)

	fmt.Printf("--- Downloading file %s from bucket %s (%s) to %s...\r\n", s3FilePath, bucket.Name, bucket.Region.Name, outputFilePath)
	outputDirectory := filepath.Dir(outputFilePath)
	os.MkdirAll(outputDirectory, 0755)

	// FIXME: Use GetReader to stream file contents instead of loading all the file to memory before writing
	contents, err := bucket.Get(s3FilePath)
	if err != nil {
		return err
	}

	// Sets default permission if not configured in YAML
	if file.Permission == 0 {
		file.Permission = 0644
	}

	filemode := os.FileMode(file.Permission & 0777)
	err = ioutil.WriteFile(outputFilePath, contents, filemode)
	if err != nil {
		return err
	}

	err = os.Chmod(outputFilePath, filemode)
	if err != nil {
		return err
	}

	return nil
}

func checkAWSRegions(configs config.S3HarborRepositoriesConfig) error {
	for name, config := range configs {
		if config.Region != "" {
			if _, exists := aws.Regions[config.Region]; !exists {
				return fmt.Errorf("The region: %s is not valid on %s config", config.Region, name)
			}
		}
	}
	return nil
}

func getBucket(config config.S3HarborConfig, awsAuth aws.Auth) *s3.Bucket {
	region := aws.USEast
	if config.Region != "" {
		region = aws.Regions[config.Region]
	}
	return s3.New(awsAuth, region).Bucket(config.Bucket)
}
