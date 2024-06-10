// tinyvectordb.go
package tinyvectordb

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/xyproto/ollamaclient/v2"
)

type Vector struct {
	ID        int
	Filename  string // filename of the image
	Type      string // "text" or "image"
	Content   string // original content
	Embedding []float64
}

var (
	vectors []Vector
	mutex   sync.Mutex
	nextID  int
	oc      *ollamaclient.Config
)

func init() {
	// Initialize the Ollama client configuration
	oc = ollamaclient.New()
	oc.ModelName = "llava-llama3" // Update to the llava-llama3 model
	oc.Verbose = true
	oc.HTTPTimeout = 10 * time.Second
}

// StoreEmbedding stores the embedding of a text or image content
func StoreEmbedding(embedType, filename string) (string, error) {
	if embedType != "text" && embedType != "image" {
		return "", fmt.Errorf("invalid type. Use 'text' or 'image'")
	}

	embedding, err := getEmbedding(embedType, filename)
	if err != nil {
		return "", fmt.Errorf("error getting embedding: %v", err)
	}

	mutex.Lock()
	vectors = append(vectors, Vector{ID: nextID, Filename: filename, Type: embedType, Content: filename, Embedding: embedding})
	log.Printf("Stored %s embedding with ID %d (filename: %s): %v", embedType, nextID, filename, embedding)
	nextID++
	mutex.Unlock()

	return fmt.Sprintf("Stored %s embedding with ID %d (filename: %s)", embedType, nextID-1, filename), nil
}

func getEmbedding(embedType, content string) ([]float64, error) {
	if embedType == "text" {
		return oc.Embeddings(content)
	} else {
		// Assuming content is the path to the image file
		base64image, err := ollamaclient.Base64EncodeFile(content)
		if err != nil {
			return nil, err
		}
		return oc.Embeddings(base64image)
	}
}

func GetVectors() []Vector {
	mutex.Lock()
	defer mutex.Unlock()
	return vectors
}
