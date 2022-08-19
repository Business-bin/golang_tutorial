package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"path/filepath"
	"strings"
)

func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// 파일 정보
		orgFile, fileHeader, err1 := c.Request().FormFile("file")
		if err1 != nil {
			fmt.Println("업로드 111 fileHeader error")
		}
		/*
			fileName := c.FormValue("filename")
			realFile, err2 := c.FormFile("file")
			if err2 != nil {
				fmt.Println("업로드 222 realFile error")
			}
			file, err3 := realFile.Open()
			if err3 != nil {
				fmt.Errorf("os.Open - filename: %s, err: %v", fileName, err3)
			}
			defer file.Close()*/

		// Create an uploader with the session and default options
		// s3 업로더가 사용 할 세션
		uploader := s3manager.NewUploader(newSession())
		// uuid 생성
		fileNameUUID := GenerateUUID()
		// s3에 저장될 파일명 생성
		fileNamaSave := fileNameUUID + filepath.Ext(fileHeader.Filename)
		// 파일 저장소 경로
		bucketPathFullName := "cf.krms.file/krms3/app/"
		// Upload the file to S3.
		result, err4 := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketPathFullName),
			Key:    aws.String(fileNamaSave),
			Body:   orgFile,
		})
		if err4 != nil {
			fmt.Println("업로드 333 file")
			return fmt.Errorf("failed to upload file, %v", err4)
		}
		fmt.Println("업로드 success")
		fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
		return nil
	}
}

const (
	region          = "us-east-1"
	accessKeyID     = "AKIARJETX2AKL4NNDJ2H"
	secretAccessKey = "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"
)

// aws s3 리전과 session
func newSession() *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
	}))
	return sess
}

// UUID 생성
func GenerateUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
