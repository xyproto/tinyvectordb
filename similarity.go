// similarity.go
package tinyvectordb

import (
	"log"
	"math"
)

// EuclideanDistance calculates the Euclidean distance between two vectors
func euclideanDistance(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Printf("Embedding length mismatch: %d vs %d", len(a), len(b))
		return math.MaxFloat64
	}
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// FindMostSimilar returns the ID and distance of the most similar vector to the given vector
func FindMostSimilar(query []float64, vectors []Vector) (int, float64) {
	closestID := -1
	closestDist := math.MaxFloat64

	for _, v := range vectors {
		dist := euclideanDistance(query, v.Embedding)
		log.Printf("Distance from query to ID %d (filename: %s): %f", v.ID, v.Filename, dist)
		if dist < closestDist {
			closestDist = dist
			closestID = v.ID
		}
	}

	return closestID, closestDist
}

// CheckEmbeddingQuality checks if the embeddings are valid and distinct
func CheckEmbeddingQuality(vectors []Vector) bool {
	for i, v1 := range vectors {
		for j, v2 := range vectors {
			if i != j {
				dist := euclideanDistance(v1.Embedding, v2.Embedding)
				if dist == 0 {
					log.Printf("Zero distance found between ID %d (filename: %s) and ID %d (filename: %s)", v1.ID, v1.Filename, v2.ID, v2.Filename)
					return false
				}
			}
		}
	}
	return true
}
