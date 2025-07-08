package tgbotapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSendPhoto(t *testing.T) {
	bot, _ := getBot(t)
	msg := &PhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/image.jpg"),
		},
	}
	_, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

func TestSendAudio(t *testing.T) {
	bot, _ := getBot(t)
	msg := &AudioConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/audio.mp3"),
			Thumb:    FilePath("./tests/image.jpg"),
		},
		Caption: "Music file",
	}
	_, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

func printResponse(resp Message) {
	fmt.Println("\n------------------------------------------")
	res, er := json.Marshal(resp)
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(string(res))
}

func TestSendDocument(t *testing.T) {
	bot, _ := getBot(t)
	msg := &DocumentConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/chatgpt.pdf"),
			Thumb:    FilePath("./tests/image.jpg"),
		},
		Caption: "ChatGPT PDF",
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendVideo(t *testing.T) {
	bot, _ := getBot(t)
	msg := &VideoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/video.mp4"),
			Thumb:    FilePath("./tests/image.jpg"),
		},
		Cover:   FilePath("./tests/cover.jpg"),
		Caption: "Video file",
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendAnimation(t *testing.T) {
	bot, _ := getBot(t)
	msg := &AnimationConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/video.mp4"),
			Thumb:    FilePath("./tests/image.jpg"),
		},
		Caption: "Animation file",
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendVoice(t *testing.T) {
	bot, _ := getBot(t)
	msg := &VoiceConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/voice.ogg"),
		},
		Caption: "Voice file",
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendVideoNote(t *testing.T) {
	bot, _ := getBot(t)
	msg := &VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/video.mp4"),
			Thumb:    FilePath("./tests/image.jpg"),
		},
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendSticker(t *testing.T) {
	bot, _ := getBot(t)
	msg := &StickerConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: PremiumChatId},
			File:     FilePath("./tests/sticker.webp"),
		},
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}

func TestSendAction(t *testing.T) {
	bot, _ := getBot(t)
	msg := &ChatActionConfig{
		BaseChat: BaseChat{ChatID: PremiumChatId},
		Action:   "typing",
	}
	resp, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
	printResponse(resp)
}
