package lib_test

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/a-poor/ls3/lib"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// s3ClientSuccessMock is a mock implementation of S3Client that always returns
// the specified success results (with a nil error).
type s3ClientFuncMock struct {
	list func(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	get  func(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	put  func(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	del  func(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

func (c *s3ClientFuncMock) asS3Client() lib.S3Client {
	return c
}

func (c *s3ClientFuncMock) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return c.list(ctx, params)
}

func (c *s3ClientFuncMock) GetObject(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return c.get(ctx, params)
}

func (c *s3ClientFuncMock) PutObject(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return c.put(ctx, params)
}

func (c *s3ClientFuncMock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return c.del(ctx, params)
}

func TestS3FSListContents(t *testing.T) {
	t.Run("successful-response", func(t *testing.T) {
		expects := []lib.FileObject{
			{"_dir1", true, 0},
			{"dir2", true, 0},
			{".foo.txt", false, 10},
			{"bar.json", false, 20},
		}

		s3Mock := s3ClientFuncMock{list: func(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
			return &s3.ListObjectsV2Output{
				CommonPrefixes: []types.CommonPrefix{
					{Prefix: aws.String("_dir1")},
					{Prefix: aws.String("dir2")},
				},
				Contents: []types.Object{
					{Key: aws.String(".foo.txt"), Size: 10},
					{Key: aws.String("bar.json"), Size: 20},
				},
			}, nil
		}}
		client := lib.NewS3FSWithClient(s3Mock.asS3Client(), "", "/")

		res, err := client.ListContents()
		if err != nil {
			t.Errorf("Error listing contents: %s", err)
		}

		if len(res) != len(expects) {
			t.Errorf("Expected %d results, got %d", len(expects), len(res))
		}

		for _, expect := range expects {
			var found bool
			for _, got := range res {
				if got.Name == expect.Name && got.IsDir == expect.IsDir && got.Size == expect.Size {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Didn't find %q in res", expect.Name)
			}
		}
	})

	t.Run("error-response", func(t *testing.T) {
		expect := errors.New("something went wrong")
		s3Mock := s3ClientFuncMock{list: func(ctx context.Context, params *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
			return nil, expect
		}}
		client := lib.NewS3FSWithClient(s3Mock.asS3Client(), "", "/")
		res, err := client.ListContents()
		if err == nil {
			t.Errorf("Expecting an error but didn't get one")
		}
		if res != nil {
			t.Errorf("Expecting nil results, got %v", res)
		}
		if !errors.Is(err, expect) {
			t.Errorf("Expecting error %q, got %q", expect, err)
		}
	})
}

func TestS3FSChangeDir(t *testing.T) {
	t.Run("successful-response", func(t *testing.T) {
		// TODO - Add tests for success responses...
	})
	t.Run("error-response", func(t *testing.T) {
		// TODO - Add tests for success responses...
	})
}

func TestS3FSGetFile(t *testing.T) {
	t.Run("successful-response", func(t *testing.T) {
		data := "Hello, World!"
		s3Mock := s3ClientFuncMock{get: func(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			// Create an io.ReadCloser from the data
			r := strings.NewReader(data)
			rc := io.NopCloser(r)
			return &s3.GetObjectOutput{
				Body: rc,
			}, nil
		}}
		client := lib.NewS3FSWithClient(s3Mock.asS3Client(), "", "/")
		b, err := client.GetFile("foo.txt")
		if err != nil {
			t.Errorf("Error getting file: %s", err)
		}
		if string(b) != data {
			t.Errorf("Expected %q, got %q", data, string(b))
		}
	})
	t.Run("error-response", func(t *testing.T) {
		expect := errors.New("something went wrong")
		s3Mock := s3ClientFuncMock{get: func(ctx context.Context, params *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
			return nil, expect
		}}
		client := lib.NewS3FSWithClient(s3Mock.asS3Client(), "", "/")
		res, err := client.GetFile("foo.txt")
		if err == nil {
			t.Errorf("Expecting an error but didn't get one")
		}
		if res != nil {
			t.Errorf("Expecting nil results, got %v", res)
		}
		if !errors.Is(err, expect) {
			t.Errorf("Expecting error %q, got %q", expect, err)
		}
	})
}

func TestS3FSWriteFile(t *testing.T) {
	t.Run("successful-response", func(t *testing.T) {
		// TODO - Add tests for success responses...
	})
	t.Run("error-response", func(t *testing.T) {
		// TODO - Add tests for success responses...
	})
}
