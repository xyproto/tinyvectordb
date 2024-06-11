package tinyvectordb

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/xyproto/ollamaclient/v2"
)

const (
	TextContent = iota
	ImageContent
)

type ContentType int

type Vector struct {
	ID        int
	Filename  string      // filename of the image
	Type      ContentType // TextContent or ImageContent
	Content   string      // original content
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
	oc.ModelName = "llava-llama3"
	oc.Verbose = true
	oc.HTTPTimeout = 60 * time.Second
}

// StoreEmbedding stores the embedding of a text or image content
func StoreEmbedding(embedType ContentType, filename string) (string, error) {
	if embedType != TextContent && embedType != ImageContent {
		return "", fmt.Errorf("invalid type. Use TextContent or ImageContent")
	}

	embedding, err := getEmbedding(embedType, filename)
	if err != nil {
		return "", fmt.Errorf("error getting embedding: %v", err)
	}

	mutex.Lock()
	vectors = append(vectors, Vector{ID: nextID, Filename: filename, Type: embedType, Content: filename, Embedding: embedding})
	if embedType == ImageContent {
		log.Printf("Stored image embedding with ID %d (filename: %s): %v", nextID, filename, embedding)
	} else {
		log.Printf("Stored text embedding with ID %d (filename: %s): %v", nextID, filename, embedding)
	}
	nextID++
	mutex.Unlock()

	if embedType == ImageContent {
		return fmt.Sprintf("Stored image embedding with ID %d (filename: %s)", nextID-1, filename), nil
	}
	return fmt.Sprintf("Stored text embedding with ID %d (filename: %s)", nextID-1, filename), nil
}

func getEmbedding(embedType ContentType, content string) ([]float64, error) {
	switch embedType {
	case ImageContent:
		// Assuming content is the path to the image file
		base64image, err := ollamaclient.Base64EncodeFile(content)
		if err != nil {
			return nil, err
		}
		return oc.Embeddings(base64image)
	case TextContent:
		fallthrough
	default:
		return oc.Embeddings(content)
	}
}

func GetVectors() []Vector {
	mutex.Lock()
	defer mutex.Unlock()
	return vectors
}
