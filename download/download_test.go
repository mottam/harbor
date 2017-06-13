package download

import (
	"reflect"
	"testing"

	"github.com/elo7/harbor/config"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

func Test_checkAWSRegions(t *testing.T) {
	tests := []struct {
		name    string
		args    config.S3HarborRepositoriesConfig
		wantErr bool
	}{
		{"success", config.S3HarborRepositoriesConfig{"default": config.S3HarborConfig{Bucket: "config.elo7.com.br", Region: "sa-east-1"}}, false},
		{"invalid region", config.S3HarborRepositoriesConfig{"default": config.S3HarborConfig{Bucket: "config.elo7.com.br", Region: "aaa-east-1"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkAWSRegions(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("checkAWSRegions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getBucket(t *testing.T) {
	awsAuth := aws.Auth{}

	type args struct {
		config  config.S3HarborConfig
		awsAuth aws.Auth
	}
	tests := []struct {
		name string
		args args
		want *s3.Bucket
	}{
		{"success", args{config.S3HarborConfig{Bucket: "config.elo7.com.br", Region: "sa-east-1"}, awsAuth}, s3.New(awsAuth, aws.SAEast).Bucket("config.elo7.com.br")},
		{"default region", args{config.S3HarborConfig{Bucket: "config.elo7.com.br", Region: ""}, awsAuth}, s3.New(awsAuth, aws.USEast).Bucket("config.elo7.com.br")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBucket(tt.args.config, tt.args.awsAuth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBucket() = %v (%v), want %v (%v)", got, got.Region.Name, tt.want, tt.want.Region.Name)
			}
		})
	}
}
