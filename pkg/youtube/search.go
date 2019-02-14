package youtube

import (
	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

var (
	maxResult    int64 = 50
	safeSearch         = "none"
	resourceType       = "video"
)

// SearchWithQuery returns SearchListResponse
func (svc *Service) SearchWithQuery(query string) (*youtube.SearchListResponse, error) {
	call := svc.Search.List("id,snippet").
		Q(query).
		MaxResults(maxResult).
		SafeSearch(safeSearch).
		Type(resourceType)

	res, err := call.Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to call search method")
	}
	return res, nil
}
