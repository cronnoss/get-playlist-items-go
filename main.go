package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	bearerToken := viper.GetString("SPOTIFY_BEARER_TOKEN")
	playlistID := viper.GetString("PLAYLIST_ID")

	url := "https://api.spotify.com/v1/playlists/" + playlistID + "/tracks"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the name and artist of each track in the playlist
	for _, item := range data["items"].([]interface{}) {
		track := item.(map[string]interface{})["track"].(map[string]interface{})
		fmt.Printf("%s - %s\n", track["name"], track["artists"].([]interface{})[0].(map[string]interface{})["name"])
	}

	// Save the output to a file
	f, err := os.Create("tracks.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	for _, item := range data["items"].([]interface{}) {
		track := item.(map[string]interface{})["track"].(map[string]interface{})
		fmt.Fprintf(f, "%s - %s\n", track["name"], track["artists"].([]interface{})[0].(map[string]interface{})["name"])
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
