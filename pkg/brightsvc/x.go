package brightsvc

import (
	"math/rand"
	"time"
)

func randomHash(hashes []int) int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hashes), func(i, j int) {
		hashes[i], hashes[j] = hashes[j], hashes[i]
	})
	return hashes[0]
}
