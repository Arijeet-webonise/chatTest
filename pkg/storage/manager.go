package storage

import (
	"bytes"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// StorageManager encapsulates simple storage related methods
type StorageManager interface {
	NewSession() (*session.Session, error)
	Upload(filePath string, fileName string, fileType string) (*string, error)
	Download(fileName string, fileType string) (string, error)
	Delete(fileName string, fileType string) error
	ObjectRequest(fileName string, fileType string) (string, error)
	DeleteImages(imagePath string) error
	Credentials() *credentials.Credentials
}

// StorageManagerServiceImpl provides simple storage essential configurations
type StorageManagerServiceImpl struct {
	Bucket         string
	Endpoint       string
	Region         string
	AccessKey      string
	SecretKey      string
	UploadFilePath string
}

// Credentials returns credentials for simple storage connectivity
func (serviceImpl *StorageManagerServiceImpl) Credentials() *credentials.Credentials {
	return credentials.NewStaticCredentials(serviceImpl.AccessKey, serviceImpl.SecretKey, "")
}

// NewSession returns new session for provided configurations
func (serviceImpl *StorageManagerServiceImpl) NewSession() (*session.Session, error) {
	var sess *session.Session
	var err error

	// Use following configuration for locally configured minio
	if serviceImpl.AccessKey != `None` && serviceImpl.SecretKey != `None` {
		sess, err = session.NewSession(&aws.Config{
			Credentials:      serviceImpl.Credentials(),
			Endpoint:         aws.String(serviceImpl.Endpoint),
			Region:           aws.String(serviceImpl.Region),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		})
	} else {

		// Use following configuration for EC2 instance configured with IAM Role
		sess, err = session.NewSession(&aws.Config{
			Endpoint:         aws.String(serviceImpl.Endpoint),
			Region:           aws.String(serviceImpl.Region),
			DisableSSL:       aws.Bool(false),
			S3ForcePathStyle: aws.Bool(true),
		})
	}
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// Upload uploads file to simple storage
func (serviceImpl *StorageManagerServiceImpl) Upload(filePath string, fileName string, fileType string) (*string, error) {
	session, err := serviceImpl.NewSession()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	storageFilePath := storageFilePath(fileType) + fileName
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(serviceImpl.Bucket),
		Key:           &storageFilePath,
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
	})
	if err != nil {
		return nil, err
	}

	// Remove asset from local once it is pushed to simple storage
	err = os.Remove(filePath)
	if err != nil {
		return nil, err
	}
	return &storageFilePath, nil
}

// Download downloads file from simple storage
func (serviceImpl *StorageManagerServiceImpl) Download(fileName string, fileType string) (string, error) {
	session, err := serviceImpl.NewSession()
	if err != nil {
		return "", err
	}
	file, err := os.Create(serviceImpl.UploadFilePath + fileName)
	if err != nil {
		return "", err
	}
	storageFilePath := storageFilePath(fileType) + fileName
	_, err = s3manager.NewDownloader(session).Download(file, &s3.GetObjectInput{
		Bucket: aws.String(serviceImpl.Bucket),
		Key:    &storageFilePath,
	})
	if err != nil {
		return "", err
	}
	return serviceImpl.UploadFilePath + fileName, nil
}

// Delete deletes object from simple storage
func (serviceImpl *StorageManagerServiceImpl) Delete(fileName string, fileType string) error {
	session, err := serviceImpl.NewSession()
	if err != nil {
		return err
	}
	svc := s3.New(session)
	storageFilePath := storageFilePath(fileType) + fileName
	content := &s3.DeleteObjectInput{
		Bucket: aws.String(serviceImpl.Bucket),
		Key:    &storageFilePath,
	}
	_, err = svc.DeleteObject(content)
	if err != nil {
		return err
	}
	return nil
}

// ObjectRequest returns presigned URL for object present in simple storage
func (serviceImpl *StorageManagerServiceImpl) ObjectRequest(fileName string, fileType string) (string, error) {
	session, err := serviceImpl.NewSession()
	if err != nil {
		return "", err
	}
	svc := s3.New(session)
	storageFilePath := storageFilePath(fileType) + fileName
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(serviceImpl.Bucket),
		Key:    &storageFilePath,
	})
	presignURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}
	return presignURL, nil
}

// storageFilePath determines file path on simple storage
func storageFilePath(fileType string) string {
	switch fileType {
	case `excel`:
		return `/excel-files/`
	case `image`:
		return `/images/`
	case `doc`:
		return `/documents/`
	}
	return ``
}

// DeleteImages deletes original and resized images from simple storage service
func (serviceImpl *StorageManagerServiceImpl) DeleteImages(imagePath string) error {
	if err := serviceImpl.Delete(`resized_`+imagePath, `image`); err != nil {
		return err
	}
	if err := serviceImpl.Delete(imagePath, `image`); err != nil {
		return err
	}
	return nil
}
