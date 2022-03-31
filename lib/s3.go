package lib

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client is an interface for working with AWS S3
type S3Client interface {
	// ListObjects returns a list of objects in the the specified S3 bucket
	ListObjects(context.Context, *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)

	// GetObject returns the contents of the specified S3 object
	GetObject(context.Context, *s3.GetObjectInput) (*s3.GetObjectOutput, error)

	// PutObject uploads the specified file to S3
	PutObject(context.Context, *s3.PutObjectInput) (*s3.PutObjectOutput, error)

	// DeleteObject deletes the specified S3 object
	DeleteObject(context.Context, *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

// s3ClientSuccessMock is a mock implementation of S3Client that always returns
// the specified success results (with a nil error).
type s3ClientSuccessMock struct {
	list *s3.ListObjectsV2Output // Always return this response for ListObjects
	get  *s3.GetObjectOutput     // Always return this response for GetObject
	put  *s3.PutObjectOutput     // Always return this response for PutObject
	del  *s3.DeleteObjectOutput  // Always return this response for DeleteObject
}

// NewS3ClientSuccessMock returns a mock S3Client that always returns the specified
// success results (with a nil error) – used for testing.
//
//   - "list" is returned by "client.ListObjects"
//   - "get"  is returned by "client.GetObject"
//   - "put"  is returned by "client.PutObject"
//   - "del"  is returned by "client.DeleteObject"
//
func NewS3ClientSuccessMock(list *s3.ListObjectsV2Output, get *s3.GetObjectOutput, put *s3.PutObjectOutput, del *s3.DeleteObjectOutput) S3Client {
	return &s3ClientSuccessMock{list, get, put, del}
}

func (c *s3ClientSuccessMock) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return c.list, nil
}

func (c *s3ClientSuccessMock) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return c.get, nil
}

func (c *s3ClientSuccessMock) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return c.put, nil
}

func (c *s3ClientSuccessMock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return c.del, nil
}

// s3ClientErrMock is a mock implementation of S3Client that always returns
// the specified error, regardless of the input parameters.
type s3ClientErrMock struct {
	err error // Always return this error for all methods
}

// NewS3ClientErrMock returns a new mock S3Client that always returns the specified error.
func NewS3ClientErrMock(err error) S3Client {
	return &s3ClientErrMock{err}
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

// s3Client is a real implementation of S3Client that uses the AWS SDK v2
// to interact with the S3 service.
type s3Client struct {
	client *s3.Client // AWS SDK v2 client
}

// Create a new client for accessing S3
func NewS3Client(cfg aws.Config) S3Client {
	return &s3Client{
		client: s3.NewFromConfig(cfg),
	}
}

func (c *s3Client) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return c.client.ListObjectsV2(ctx, params)
}

func (c *s3Client) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return c.client.GetObject(ctx, params)
}

func (c *s3Client) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return c.client.PutObject(ctx, params)
}

func (c *s3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return c.client.DeleteObject(ctx, params)
}

// S3FS is an implementation of FileSystem that uses AWS S3 as the backend.
type S3FS struct {
	Client  S3Client // Client for accessing S3
	Bucket  string   // S3 Bucket name
	WorkDir string   // Current working "directory" in S3
}

// NewS3FS creates a new S3FS instance using the specified S3 client, bucket,
// and working directory.
func NewS3FS(client S3Client, bucket, workDir string) *S3FS {
	return &S3FS{
		Client:  client,
		Bucket:  bucket,
		WorkDir: workDir,
	}
}

func (fs *S3FS) ListContents() ([]FileObject, error) {
	return nil, nil
}

func (fs *S3FS) ChangeDir(string) error {
	return nil
}

func (fs *S3FS) GetFile(string) ([]byte, error) {
	return nil, nil
}

func (fs *S3FS) WriteFile(string, []byte) error {
	return nil
}
