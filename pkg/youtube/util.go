package youtube

import (
	"net/url"

	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"
)

const baseURL = "https://www.youtube.com/watch"

// BuildURLsFromItems return accessible urls
func BuildURLsFromItems(items []*youtube.SearchResult) ([]string, error) {
	urls := []string{}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse baseURL")
	}

	for _, item := range items {
		_u := u
		videoID := item.Id.VideoId

		q := _u.Query()
		q.Set("v", videoID)
		_u.RawQuery = q.Encode()

		urls = append(urls, _u.String())
	}

	return urls, nil
}
