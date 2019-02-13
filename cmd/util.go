package cmd

import (
	"math/rand"
	"time"

	"google.golang.org/api/youtube/v3"
)

// Shuffle returns shuffled array
func Shuffle(rs []*youtube.SearchResult) {
	rand.Seed(time.Now().UnixNano())
	for i := range rs {
		j := rand.Intn(i + 1)
		rs[i], rs[j] = rs[j], rs[i]
	}
}
