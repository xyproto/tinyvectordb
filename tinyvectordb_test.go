// tinyvectordb_test.go
package tinyvectordb

import (
	"fmt"
	"testing"
)

func TestEmbeddings(t *testing.T) {
	// Assume we have images cat1.png, cat2.png, and dog1.png in the img directory
	imageFiles := []string{"img/cat1.png", "img/cat2.png", "img/dog1.png"}

	for _, imageFile := range imageFiles {
		msg, err := StoreEmbedding(ImageContent, imageFile)
		if err != nil {
			t.Fatalf("Failed to store embedding for %s: %v", imageFile, err)
		}
		fmt.Println(msg)
	}

	// Check if embeddings are valid and distinct
	vectors := GetVectors()
	if !CheckEmbeddingQuality(vectors) {
		t.Fatalf("Embeddings are not valid or distinct")
	}

	// Check distances
	cat1Cat2ID, cat1Cat2Dist := FindMostSimilar(vectors[0].ID, vectors[0].Embedding, vectors)
	fmt.Printf("Closest to cat1.png: ID %d, Distance %f\n", cat1Cat2ID, cat1Cat2Dist)

	cat1Dog1ID, cat1Dog1Dist := FindMostSimilar(vectors[2].ID, vectors[2].Embedding, vectors)
	fmt.Printf("Closest to dog1.png: ID %d, Distance %f\n", cat1Dog1ID, cat1Dog1Dist)

	// Ensure that the distance between cat images is smaller than the distance between cat and dog images
	if cat1Cat2Dist >= cat1Dog1Dist {
		t.Fatalf("Distance between cat1.png and cat2.png (%f) is not smaller than distance between cat1.png and dog1.png (%f)", cat1Cat2Dist, cat1Dog1Dist)
	}

	fmt.Println("Test passed: cat1.png and cat2.png are closer to each other than to dog1.png")
}
