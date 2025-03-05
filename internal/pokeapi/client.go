package pokeapi

import (
	"net/http"
	"time"

	"github.com/agustin-carnevale/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	baseUrl    string
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		baseUrl: "https://pokeapi.co/api/v2",
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
