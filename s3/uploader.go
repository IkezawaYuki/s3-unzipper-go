package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/sync/errgroup"
)

type Uploader struct {
	manager s3manager.Uploader
	src     string
	dest    string
}

func NewUploader(s *session.Session, src string, dest string) *Uploader {
	return &Uploader{
		manager: *s3manager.NewUploader(s),
		src:     src,
		dest:    dest,
	}
}

func (u Uploader) Upload() error {
	eg := errgroup.Group{}

	// todo

	return nil
}
