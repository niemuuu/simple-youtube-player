package youtube

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const (
	apiKey = "AIzaSyApJchZ-k6egHcMsnzvTAA144omq4d62fI"

	maxResult    = 3
	safeSearch   = "none"
	resourceType = "video"
)

// Service is wrapping youtube.Service
type Service struct {
	*youtube.Service
}

// NewService returns new youtube-service
func NewService() (*Service, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	svc, err := youtube.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create YoutubeService")
	}
	return &Service{svc}, nil
}

// SearchItems returns items
func (svc *Service) SearchItems(query string) error {
	call := svc.Search.List("id,snippet").
		Q(query).
		MaxResults(maxResult).
		SafeSearch(safeSearch).
		Type(resourceType)

	res, err := call.Do()
	if err != nil {
		return errors.Wrap(err, "failed to call search method")
	}

	pp.Print(res)
	return nil
}
