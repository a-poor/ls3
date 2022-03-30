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

type s3ClientSuccessMock struct {
	listRes *s3.ListObjectsV2Output
	getRes  *s3.GetObjectOutput
	putRes  *s3.PutObjectOutput
	delRes  *s3.DeleteObjectOutput
}

func NewS3ClientSuccessMock(l *s3.ListObjectsV2Output, g *s3.GetObjectOutput, p *s3.PutObjectOutput, d *s3.DeleteObjectOutput) S3Client {
	return &s3ClientSuccessMock{l, g, p, d}
}

func (c *s3ClientSuccessMock) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return c.listRes, nil
}

func (c *s3ClientSuccessMock) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return c.getRes, nil
}

func (c *s3ClientSuccessMock) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return c.putRes, nil
}

func (c *s3ClientSuccessMock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return c.delRes, nil
}

type s3ClientErrMock struct {
	err error
}

func NewS3ClientErrMock(err error) S3Client {
	return &s3ClientErrMock{}
}

func (c *s3ClientErrMock) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return nil, c.err
}

func (c *s3ClientErrMock) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return nil, c.err
}

func (c *s3ClientErrMock) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return nil, c.err
}

func (c *s3ClientErrMock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return nil, c.err
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
