package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
)

type Config struct {
	CookiesPath        string `json:"cookies_path"`
	InstagramHost      string `json:"intagram_host"`
	InstagramPath      string `json:"instagram_path"`
	InstagramQueryHash string `json:"instagram_query_hash"`
}

func LoadConfig(path string) Config {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func UrlBuilder(url, path, queryHash, queryData string) (result string) {
	result = fmt.Sprintf("%s/%s/?query_hash=%s&variables=%s", url, path, queryHash, html.EscapeString(queryData))
	return
}
