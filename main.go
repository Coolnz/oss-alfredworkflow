package main

import "C"
import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	aw "github.com/deanishe/awgo"
	"path/filepath"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(UploadFile)
}

func UploadFile() {

	var (
		accessKeyID     = wf.Config.GetString("access_key_id")
		accessKeySecret = wf.Config.GetString("access_key_secret")
		bucketName      = wf.Config.GetString("bucket_name")
		endpoint        = wf.Config.GetString("endpoint")
		prefix          = wf.Config.GetString("prefix")
	)

	args := wf.Args()
	if len(args) > 1 || len(args) == 0 {
		wf.Fatal("too many args")
	}
	filePath := args[0]
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		wf.FatalError(err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		wf.FatalError(err)
	}

	filename := filepath.Base(filePath)
	ossFile := prefix + filename
	err = bucket.UploadFile(ossFile, filePath, 100*1024, oss.Routines(3))
	if err != nil {
		wf.FatalError(err)
	}

	url := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, ossFile)
	ss := fmt.Sprintf("![%s](%s)", filename, url)
	// wf.NewItem(ss).Arg(ss).Valid(true).Autocomplete(ss).Icon(&aw.Icon{Value: "icon.png"})
	// wf.SendFeedback()
	fmt.Println(ss)
	return
}
