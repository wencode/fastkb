package audio

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

const (
	sampleRate = 44100

	introLengthInSecond = 5
	loopLengthInSecond  = 4
)

var (
	context    *audio.Context = audio.NewContext(sampleRate)
	hitPlayer  *audio.Player
	dangPlayer *audio.Player
	bgPlayer   *audio.Player
)

func init() {
	d, err := wav.Decode(context, bytes.NewReader(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}

	hitPlayer, err = audio.NewPlayer(context, d)
	if err != nil {
		log.Fatal(err)
	}
	hitPlayer.SetVolume(0.75)

}

func init() {
	file, err := os.Open("./dang.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	d, err := wav.Decode(context, bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	dangPlayer, err = audio.NewPlayer(context, d)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	// Decode an Ogg file.
	// oggS is a decoded io.ReadCloser and io.Seeker.
	oggS, err := vorbis.Decode(context, bytes.NewReader(raudio.Ragtime_ogg))
	if err != nil {
		log.Fatal(err)
	}

	s := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*4*sampleRate, loopLengthInSecond*4*sampleRate)
	bgPlayer, err = audio.NewPlayer(context, s)
	if err != nil {
		log.Fatal(err)
	}
}

func PlayBG() {
	bgPlayer.Play()
}

func Hit() {
	hitPlayer.Rewind()
	hitPlayer.Play()
}

func Miss() {
	dangPlayer.Rewind()
	dangPlayer.Play()
}
