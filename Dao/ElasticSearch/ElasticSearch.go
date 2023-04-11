package ElasticSearch

import (
	"FastWiki/Setting"
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
	"time"
)

var esClient *elasticsearch.Client

func Init(config *Setting.ElaSearchConfig) (err error) {
	cfg := elasticsearch.Config{Addresses: []string{
		config.Host,
	},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   50,
			MaxConnsPerHost:       200,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	esClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	return nil
}
