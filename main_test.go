package main

import (
	"bytes"
	"testing"

	"github.com/gotracker/gomixing/mixing"
	"github.com/gotracker/gomixing/sampling"
	"github.com/gotracker/playback/format"
	"github.com/gotracker/playback/output"
	"github.com/gotracker/playback/player/feature"
)

func BenchmarkGenerate(b *testing.B) {
	b.ReportAllocs()

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

	const premixChannelSize = 8
	premixDataChannel := make(chan *output.PremixData, premixChannelSize)
	defer close(premixDataChannel)

	const sampleFormat = sampling.Format16BitLESigned

	m := mixing.Mixer{
		Channels: channels,
	}

	panMixer := mixing.GetPanMixer(channels)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		premix, err := player.Generate(0)
		if err != nil {
			// log.Fatal(err)
		}

		data := m.Flatten(panMixer, premix.SamplesLen, premix.Data, premix.MixerVolume, sampleFormat)
		_ = data
	}
}
