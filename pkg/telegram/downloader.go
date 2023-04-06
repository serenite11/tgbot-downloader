package telegram

import (
	"github.com/kkdai/youtube/v2"
	"io"
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
	defer os.Remove(video.Title + ".mp3")

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	return video.Title + ".mp3"
}
