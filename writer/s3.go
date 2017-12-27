package writer

// this is pretty much a clone of reader/s3.go and will be merged
// in to https://github.com/whosonfirst/go-whosonfirst-s3/
// see also: https://github.com/thisisaaronland/go-iiif/blob/master/aws/s3.go
// (20171217/thisisaaronland)

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"io/ioutil"
	_ "log"
	"os/user"
	"path/filepath"
	"strings"
)

type S3Writer struct {
	Writer
	prefix  string
	bucket  string
	service *s3.S3
}

type S3Config struct {
	Bucket      string
	Prefix      string
	Region      string
	Credentials string // see notes below
}

func NewS3Writer(s3cfg S3Config) (Writer, error) {

	// https://docs.aws.amazon.com/sdk-for-go/v1/developerguide/configuring-sdk.html
	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/

	cfg := aws.NewConfig()
	cfg.WithRegion(s3cfg.Region)

	if strings.HasPrefix(s3cfg.Credentials, "env:") {

		creds := credentials.NewEnvCredentials()
		cfg.WithCredentials(creds)

	} else if strings.HasPrefix(s3cfg.Credentials, "shared:") {

		details := strings.Split(s3cfg.Credentials, ":")

		if len(details) != 3 {
			return nil, errors.New("Shared credentials need to be defined as 'shared:CREDENTIALS_FILE:PROFILE_NAME'")
		}

		creds := credentials.NewSharedCredentials(details[1], details[2])
		cfg.WithCredentials(creds)

	} else if strings.HasPrefix(s3cfg.Credentials, "iam:") {

		// assume an IAM role suffient for doing whatever

	} else if s3cfg.Credentials != "" {

		// for backwards compatibility as of 05a6042dc5956c13513bdc5ab4969877013f795c
		// (20161203/thisisaaronland)

		whoami, err := user.Current()

		if err != nil {
			return nil, err
		}

		dotaws := filepath.Join(whoami.HomeDir, ".aws")
		creds_file := filepath.Join(dotaws, "credentials")

		profile := s3cfg.Credentials

		creds := credentials.NewSharedCredentials(creds_file, profile)
		cfg.WithCredentials(creds)

	} else {

		// for backwards compatibility as of 05a6042dc5956c13513bdc5ab4969877013f795c
		// (20161203/thisisaaronland)

		creds := credentials.NewEnvCredentials()
		cfg.WithCredentials(creds)
	}

	sess := session.New(cfg)

	if s3cfg.Credentials != "" {

		_, err := sess.Config.Credentials.Get()

		if err != nil {
			return nil, err
		}
	}

	service := s3.New(sess)

	w := S3Writer{
		service: service,
		prefix:  s3cfg.Prefix,
		bucket:  s3cfg.Bucket,
	}

	return &w, nil
}

func (w *S3Writer) Write(key string, fh io.ReadCloser) error {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return err
	}

	key = r.prepareKey(key)

	params := &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
		ACL:    aws.String("public-read"),
	}

	_, err := r.service.PutObject(params)

	if err != nil {
		return err
	}

	return nil
}

func (r *S3Writer) prepareKey(key string) string {

	if r.prefix == "" {
		return key
	}

	return filepath.Join(r.prefix, key)
}
