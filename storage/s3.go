package storage

import (
	"context"
	"errors"
	"io"
	"net/url"
	"strings"

	"blog-server/config"
	"blog-server/errs"
	"blog-server/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Storage struct {
	client *s3.Client
	log    logger.Logger
}

func (s *S3Storage) Copy(ctx context.Context, bucket string, srcKey string, dstKey string) error {
	escapedKey := url.PathEscape(strings.TrimPrefix(srcKey, "/"))
	copySource := bucket + "/" + escapedKey

	_, err := s.client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     &bucket,
		Key:        &dstKey,
		CopySource: &copySource,
	})
	if err != nil {
		return errs.New(errs.CodeStorageError, "storage error", err)
	}
	return nil
}

func (s *S3Storage) Exists(ctx context.Context, bucket string, key string) (bool, error) {
	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err == nil {
		return true, nil
	}

	var nf *types.NotFound
	if errors.As(err, &nf) {
		return false, nil
	}

	return false, errs.New(errs.CodeStorageError, "storage error", err)
}

func (s *S3Storage) Download(ctx context.Context, bucket string, key string) (io.ReadCloser, string, error) {
	out, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, "", err
	}
	ct := "application/octet-stream"
	if out.ContentType != nil {
		ct = *out.ContentType
	}
	return out.Body, ct, nil
}

func (s *S3Storage) Delete(ctx context.Context, bucket, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	return err
}

func (s *S3Storage) Upload(ctx context.Context, bucket, key string, body io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        body,
		ContentType: &contentType,
	})
	return err
}

func NewS3Storage(cfg *config.Config, log logger.Logger) Storage {
	ctx := context.Background()
	s3Cfg, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(s3Cfg, func(o *s3.Options) {
		o.Region = cfg.Rustfs.Region
		o.BaseEndpoint = aws.String(cfg.Rustfs.Endpoint)
		o.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(cfg.Rustfs.AccessKeyID, cfg.Rustfs.SecretAccessKey, ""))
		o.UsePathStyle = true
	})
	return &S3Storage{client: client, log: log}
}
