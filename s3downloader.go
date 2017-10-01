package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/url"
	//"strings"
)

// Downloads an item from an S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3_download_object.go BUCKET ITEM
func main() {
	if len(os.Args) != 2 {
		exitErrorf("file name or s3 path required\nUsage: %s",
			os.Args[0])
	}
	u,_ := url.Parse(os.Args[1])
print(os.Args[1])	
bucket := u.Host
	if u.Scheme != "s3" {
		return
	}
	item := u.Path[1:]
	// Inititalize a session that the SDK will use to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	awsconf := aws.Config{Region: aws.String("us-east-1")}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, Config:awsconf,
	}))

/*	paths := strings.Split(item, "/")
	print(paths)*/
	file, err := os.Create("run.go"/*paths[len(paths) - 1]*/)
print("fun.go")
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})

	if err != nil {
		exitErrorf("Unable to download item %q, %v", item)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
