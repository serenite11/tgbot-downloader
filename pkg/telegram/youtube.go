package telegram

import (
	"github.com/kkdai/youtube/v2"
	"io"
	"net/http"
	"os"
)

func (b *Bot) Download(url string) string {
	client := youtube.Client{Debug: true}
	video, err := client.GetVideo(url)
	if err != nil {
		panic(err)
	}
	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	file, err := os.Create(video.Title + ".mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	return video.Title + ".mp3"
}
func (b *Bot) UrlIsValid(url string) string {
	path := "https://www.youtube.com/"
	if url[0:24] != path {
		return "Your url not from youtube"
	}
	_, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	return "Your url is valid"
}
