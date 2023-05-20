# get-playlist-items-go

## Description
1

Get free token from [Spotify](https://developer.spotify.com/documentation/web-api/tutorials/getting-started#request-an-access-token) api. 

2

Select the spotify playlist you want to receive items for.

e.g. To get playlist url, right click on playlist and select `Share` and then `Copy Playlist Link`.

e.g. Playlist id is `5PzmUOcfEsyHMADFUxm555` of the url: https://open.spotify.com/playlist/5PzmUOcfEsyHMADFUxm555?si=e1b37580f2f54555

e.g. Or you can get playlist id from the url in the browser.

![Alt text](./image.png?raw=true "Example get playlist id")

3

Put your `PLAYLIST_ID` and `SPOTIFY_BEARER_TOKEN` in `config.yml` file.

## Usage
```go
go run main.go
```
