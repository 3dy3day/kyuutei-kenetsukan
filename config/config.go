package config

import (
	"errors"
	"os"
	"strings"
)

const DEFAULT_APP_PATH = "/etc/kyuutei-kenetsukan"
const CAMEL_APP_NAME = "KyuuteiKenetsukan"
const KEBAB_APP_NAME = "kyuutei-kenetsukan"
const SNAKE_APP_NAME = "KYUUTEI_KENETSUKAN"
const SHORT_APP_NAME = "kyuutei-kenetsukan"

func GetAppDir() string {
	dir := os.Getenv("HOME_DIR")
	if dir == "" {
		return DEFAULT_APP_PATH
	}
	if !strings.HasSuffix(dir, "/") && !strings.HasSuffix(dir, "\\") {
		dir += "/"
	}
	return dir
}

func GetBotToken() (string, error) {
	token := os.Getenv("DISCORD_TOKEN")
	if len(token) == 0 {
		return "", errors.New("DISCORD_TOKEN not set")
	}
	return token, nil
}

func GetAwsAccessKeyId() (string, error) {
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	if len(accessKeyId) == 0 {
		return "", errors.New("AWS_ACCESS_KEY_ID not set")
	}
	return accessKeyId, nil
}

func GetAwsSecretAccessKey() (string, error) {
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if len(secretAccessKey) == 0 {
		return "", errors.New("AWS_SECRET_ACCESS_KEY not set")
	}
	return secretAccessKey, nil
}

func GetAwsS3Region() (string, error) {
	region := os.Getenv("AWS_S3_REGION")
	if len(region) == 0 {
		return "", errors.New("AWS_S3_REGION not set")
	}
	return region, nil
}

func GetAwsS3Bucket() (string, error) {
	bucket := os.Getenv("AWS_S3_BUCKET")
	if len(bucket) == 0 {
		return "", errors.New("AWS_S3_BUCKET not set")
	}
	return bucket, nil
}
