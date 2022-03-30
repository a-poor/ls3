package lib

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	ListObjects(context.Context, *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	GetObject(context.Context, *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	PutObject(context.Context, *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(context.Context, *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

type s3ClientMock struct {
	listRes *s3.ListObjectsV2Output
	listErr error

	getRes *s3.GetObjectOutput
	getErr error

	putRes *s3.PutObjectOutput
	putErr error

	delRes *s3.DeleteObjectOutput
	delErr error
}

func NewS3ClientMock() S3Client {
	return &s3ClientMock{}
}

func (c *s3ClientMock) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return c.listRes, c.listErr
}

func (c *s3ClientMock) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return c.getRes, c.getErr
}

func (c *s3ClientMock) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return c.putRes, c.putErr
}

func (c *s3ClientMock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return c.delRes, c.delErr
}

type s3Client struct {
}

func NewS3Client() S3Client {
	return &s3Client{}
}

func (c *s3Client) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return nil, nil
}

func (c *s3Client) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return nil, nil
}

func (c *s3Client) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return nil, nil
}

func (c *s3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}

type S3FS struct {
	WorkPath string
}
