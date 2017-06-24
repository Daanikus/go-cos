package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"net/http"

	"bitbucket.org/mozillazg/go-cos"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cn-north.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &cos.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test_multipart.txt"
	v, _, err := c.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", v.UploadID)

	resp, err := c.Object.AbortMultipartUpload(context.Background(), name, v.UploadID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resp.Status)
}
