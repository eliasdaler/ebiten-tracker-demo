package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
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

var sampleLength int64
var trueSample io.ReadSeeker

type Game struct {
	audioContext *audio.Context
	musicPlayer  *audio.Player

	trackPlayer playback.Playback
	m           mixing.Mixer
	panMixer    mixing.PanMixer

	rb *RingBuffer

	time int
}

type RingBuffer struct {
	head     int
	tail     int
	capacity int

	buf []byte
}

func NewRingBuffer(capacity int) *RingBuffer {
	if capacity <= 0 {
		panic("invalid buffer capacity")
	}
	return &RingBuffer{
		head:     0,
		tail:     0,
		capacity: capacity + 1,
		buf:      make([]byte, capacity+1),
	}
}

func (rb *RingBuffer) Append(b byte) {
	if (rb.tail+1)%rb.capacity == rb.head {
		panic("buffer overflow")
	}

	rb.buf[rb.tail] = b
	rb.tail = (rb.tail + 1) % rb.capacity
}

func (rb *RingBuffer) Empty() bool {
	return rb.Size() == 0
}

func (rb *RingBuffer) Size() int {
	if rb.tail == rb.head {
		return 0
	}

	if rb.tail > rb.head {
		return rb.tail - rb.head
	}
	return rb.capacity - rb.head + rb.tail
}

func (rb *RingBuffer) Pop() byte {
	if rb.Empty() {
		return 0
	}

	b := rb.buf[rb.head]
	rb.head = (rb.head + 1) % rb.capacity
	return b
}

func (rb *RingBuffer) Read(b []byte) (n int, err error) {
	fmt.Println("want", len(b), " got", rb.Size())

	if rb.Empty() {
		for i := 0; i < len(b); i++ {
			b[i] = 0
		}
		return len(b), nil
	}

	for i := 0; i < len(b); i++ {
		b[i] = rb.Pop()
		if rb.Empty() {
			panic("WHAT")
		}
	}

	return len(b), nil
}

func (g *Game) Update() error {
	g.time++

	if !g.musicPlayer.IsPlaying() && g.time > 100 {
		g.musicPlayer.Play()
	}

	for i := 0; i < 5; i++ {
		if g.rb.Size() > sampleRate*2 {
			break
		}
		g.GenerateSamples()
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
	g.rb = NewRingBuffer(sampleRate * 4)

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

	for i := 0; i < 40; i++ {
		g.GenerateSamples()
	}

	g.audioContext = audio.NewContext(sampleRate)
	g.musicPlayer, err = g.audioContext.NewPlayer(g.rb)
	g.musicPlayer.SetBufferSize(time.Second / 20)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
