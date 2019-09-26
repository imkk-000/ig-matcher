package main_test

import (
	. "ig-matcher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	expectedConfig := Config{
		CookiesPath:        "cookies.txt",
		InstagramHost:      "https://www.instagram.com",
		InstagramPath:      "/graphql/query/",
		InstagramQueryHash: "qwertyuiop1234567890asdfghjkl123",
	}

	actualConfig := LoadConfig("./test/resources/settings_success.json")

	assert.Equal(t, expectedConfig, actualConfig)
}

func TestUrlBuilder(t *testing.T) {
	expectedUrl := "https://www.instagram.com/graphql/query//?query_hash=c76146de99bb02f6415203be841dd25a&variables={&#34;id&#34;:&#34;123455667&#34;,&#34;include_reel&#34;:true,&#34;fetch_mutual&#34;:false,&#34;first&#34;:12,&#34;after&#34;:&#34;that&#34;}"

	actualUrl := UrlBuilder("https://www.instagram.com", "graphql/query/", "c76146de99bb02f6415203be841dd25a",
		`{"id":"123455667","include_reel":true,"fetch_mutual":false,"first":12,"after":"that"}`)

	assert.Equal(t, expectedUrl, actualUrl)
}
