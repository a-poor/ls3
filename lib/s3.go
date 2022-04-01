package lib

import (
	"context"
	"errors"
	"path"

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
func NewS3FS(cfg aws.Config, bucket, workDir string) *S3FS {
	return &S3FS{
		Client:  NewS3Client(cfg),
		Bucket:  bucket,
		WorkDir: workDir,
	}
}

func NewS3FSWithClient(client S3Client, bucket, workDir string) *S3FS {
	return &S3FS{
		Client:  client,
		Bucket:  bucket,
		WorkDir: workDir,
	}
}

func (fs *S3FS) fmtPath(p string) string {
	// Special case: Moving up from root directory
	if fs.WorkDir == "" && p == ".." {
		return RootDirS3
	}

	p2 := path.Join(fs.WorkDir, p)
	p2 = path.Clean(p2)

	// Special case: Moved up to root directory
	if p2 == "/" {
		return RootDirS3
	}

	return p2
}

func (fs *S3FS) ListContents() ([]FileObject, error) {
	return nil, errors.New("not implemented") // TODO – Not implemented...
}

func (fs *S3FS) ChangeDir(string) error {
	return errors.New("not implemented") // TODO – Not implemented...
}

func (fs *S3FS) GetFile(string) ([]byte, error) {
	return nil, errors.New("not implemented") // TODO – Not implemented...
}

func (fs *S3FS) WriteFile(string, []byte) error {
	return errors.New("not implemented") // TODO – Not implemented...
}

func (fs *S3FS) PathExists(name string) bool {
	return false // TODO - Not implemented...
}

func (fs *S3FS) IsFile(name string) bool {
	return false // TODO - Not implemented...
}

func (fs *S3FS) IsDir(name string) bool {
	return false // TODO - Not implemented...
}

func (fs *S3FS) IsAtRoot() bool {
	return fs.WorkDir == RootDirS3
}
