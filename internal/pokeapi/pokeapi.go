package pokeapi

import (
	"net/http"
	"time"

	"github.com/pfalzergbr/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Second * 15,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
