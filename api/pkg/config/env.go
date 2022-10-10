package config

import "os"

func InfluxAPIToken() string {
	return os.Getenv("INFLUX_API_TOKEN")
}

func InfluxAPIURL() string {
	return os.Getenv("INFLUX_API_URL")
}

func InfluxAuthBucketName() string {
	return os.Getenv("INFLUX_AUTH_BUCKET_NAME")
}

func InfluxBucketID() string {
	return os.Getenv("INFLUX_BUCKET_ID")
}

func InfluxBucketName() string {
	return os.Getenv("INFLUX_BUCKET_NAME")
}

func InfluxOrganizationID() string {
	return os.Getenv("INFLUX_ORGANIZATION_ID")
}
