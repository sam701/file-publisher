package meta

import (
	"time"
)

type PublishingRequest struct {
	FileName       string
	FileHash       string
	ExpirationTime string
}

func (s *PublishingRequest) Expired() bool {
	ti, err := time.Parse(time.RFC3339, s.ExpirationTime)
	if err != nil {
		return true
	}
	return time.Now().After(ti)
}

type PublishingResponse struct {
	SharingURL string
}
