package main

import (
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	"os"
)

/*
sox --i Submarine.aiff

Input File     : 'Submarine.aiff'
Channels       : 2
Sample Rate    : 48000
Precision      : 24-bit
Duration       : 00:00:01.49 = 71638 samples ~ 111.934 CDDA sectors
File Size      : 430k
Bit Rate       : 2.30M
Sample Encoding: 24-bit Signed Integer PCM

oto doesn't support 24-bit PCM

to convert 24-bit PCM to 16-bit PCM

sox /System/Library/Sounds/Submarine.aiff --rate 44100 --bits 16 submarine.wav

for file in *.aiff; do sox "$file" --rate 44100 --bits 16 "${file%.aiff}.wav"; done

for file in *.wav; do sox --i "$file"; done

play submarine.wav

https://www.zapsplat.com/sound-effect-categories/


*/

var (
	soundPath = "sounds"

	sounds = []string{
		"Basso.wav",
		"Frog.wav",
		"Pop.wav",
		"Submarine.wav",
		"Blow.wav",
		"Funk.wav",
		"Morse.wav",
		"Purr.wav",
		"Tink.wav",
		"Bottle.wav",
		"Glass.wav",
		"Ping.wav",
		"Sosumi.wav",
	}
)

func PlaySound(p string, sampleRate int) {
	// Configure the audio player
	//sampleRate := 44100
	//sampleRate := 48000
	channelNum := 2
	bitDepthInBytes := 2
	bufferSize := 65536

	log.Println("Playing sound:", p)
	file, err := os.Open(p)
	if err != nil {
		log.Printf("Failed to open audio file: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read audio file: %v", err)
		return
	}
	log.Printf("data len %d", len(data))

	context, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, bufferSize)
	if err != nil {
		log.Printf("Failed to create audio context: %v", err)
		return
	}
	player := context.NewPlayer()
	// Play the audio
	if _, err := player.Write(data); err != nil {
		log.Printf("Failed to play audio: %v", err)
	}
	//time.Sleep(time.Second * 2)
	player.Close()
	context.Close()
}
func main() {
	//p := "/Users/ivostoyanov/tradebot/data/sounds/prompt1.wav"
	//p := "/Users/ivostoyanov/tradebot/data/sounds/bell1.wav"
	p := "/Users/ivostoyanov/tradebot/data/sounds/beep1.wav"
	//p := "/Users/ivostoyanov/tradebot/data/sounds/Frog.wav"
	//p := "/Users/ivostoyanov/tradebot/data/sounds/Submarine.wav"
	PlaySound(p, 48000)
	//p = "sounds/Frog.wav"
	//PlaySound(p)
}

func OLD() {
	// Configure the audio player
	sampleRate := 44100
	sampleRate = 48000
	channelNum := 2
	bitDepthInBytes := 2
	bufferSize := 65536

	for _, sound := range sounds {
		log.Println("Playing sound:", sound)
		path := soundPath + "/" + sound
		file, err := os.Open(path)
		if err != nil {
			log.Printf("Failed to open audio file: %v", err)
			continue
		}
		data, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Failed to read audio file: %v", err)
			continue
		}
		context, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, bufferSize)
		if err != nil {
			log.Printf("Failed to create audio context: %v", err)
			file.Close()
			continue
		}
		player := context.NewPlayer()
		// Play the audio
		if _, err := player.Write(data); err != nil {
			log.Printf("Failed to play audio: %v", err)
		}
		player.Close()
		context.Close()
		file.Close()
	}

}
