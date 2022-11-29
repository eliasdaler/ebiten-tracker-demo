package main

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/gotracker/gomixing/mixing"
	"github.com/gotracker/gomixing/sampling"
	"github.com/gotracker/playback"
	"github.com/gotracker/playback/format"
	"github.com/gotracker/playback/player/feature"
)

//go:embed theme.xm
var fileBytes []byte

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 44100
)

type Game struct {
	audioContext *audio.Context
	musicPlayer  *audio.Player

	trackPlayer playback.Playback
	m           mixing.Mixer
	panMixer    mixing.PanMixer

	rb *RingBuffer
}

var start bool = true

func (g *Game) Update() error {

	for i := 0; i < 5; i++ {
		if g.rb.Size() > sampleRate*2 {
			break
		}
		g.GenerateSamples()
	}

	if !g.musicPlayer.IsPlaying() {
		g.musicPlayer.Play()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Now playing... Cave Story - Toroko Theme")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) GenerateSamples() {
	premix, err := g.trackPlayer.Generate(0)
	if err != nil {
		// log.Fatal(err)
	}

	const sampleFormat = sampling.Format16BitLESigned
	data := g.m.Flatten(g.panMixer, premix.SamplesLen, premix.Data, premix.MixerVolume, sampleFormat)
	for j := 0; j < len(data); j++ {
		g.rb.Append(data[j])
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tracker (Demo)")
	// ebiten.SetRunnableOnUnfocused(false)

	g := &Game{}
	g.rb = NewRingBuffer(sampleRate * 40)

	var features []feature.Feature
	features = append(features, feature.UseNativeSampleFormat(true))
	features = append(features, feature.IgnoreUnknownEffect{Enabled: true})
	features = append(features, feature.SongLoop{Count: 0})

	player, _, err := format.LoadFromReader("xm", bytes.NewReader(fileBytes), features)
	if err != nil {
		panic(err)
	}

	const channels = 2
	if err := player.SetupSampler(sampleRate, channels); err != nil {
		panic(err)
	}

	if err := player.Configure(features); err != nil {
		panic(err)
	}
	g.trackPlayer = player

	g.m = mixing.Mixer{
		Channels: channels,
	}
	g.panMixer = mixing.GetPanMixer(channels)

	g.audioContext = audio.NewContext(sampleRate)
	g.musicPlayer, err = g.audioContext.NewPlayer(g.rb)
	g.musicPlayer.SetBufferSize(time.Second / 20)

	if start {
		for i := 0; i < 40*10; i++ {
			g.GenerateSamples()
		}
		start = false
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
