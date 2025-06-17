package application

import (
	"bytes"
	"context"
	"io"
	"math"
	"time"

	"github.com/ebitengine/oto/v3"
)

type AudioService struct {
	context      *oto.Context
	readyChannel chan struct{}
	isReady      bool
	volume       float64
}

func NewAudioService() *AudioService {
	service := &AudioService{
		readyChannel: make(chan struct{}),
		isReady:      false,
		volume:       0.7,
	}
	
	go service.initialize()
	
	return service
}

func (a *AudioService) initialize() {
	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}
	
	context, readyChan, err := oto.NewContext(op)
	if err != nil {
		close(a.readyChannel)
		return
	}
	
	a.context = context
	
	<-readyChan
	a.isReady = true
	close(a.readyChannel)
}

func (a *AudioService) WaitForReady(timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	select {
	case <-a.readyChannel:
		return a.isReady
	case <-ctx.Done():
		return false
	}
}

func (a *AudioService) PlayBeep(frequency float64, duration time.Duration) error {
	if !a.isReady {
		return nil
	}
	
	sampleRate := 44100
	samples := int(float64(sampleRate) * duration.Seconds())
	
	buf := make([]byte, samples*4)
	
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		sample := int16(32767 * a.volume * 0.3 * 
			(math.Sin(2*math.Pi*frequency*t) + 
			 math.Sin(2*math.Pi*frequency*2*t)*0.3 + 
			 math.Sin(2*math.Pi*frequency*3*t)*0.1))
		
		buf[i*4] = byte(sample)
		buf[i*4+1] = byte(sample >> 8)
		buf[i*4+2] = byte(sample)
		buf[i*4+3] = byte(sample >> 8)
	}
	
	player := a.context.NewPlayer(bytes.NewReader(buf))
	go func() {
		defer player.Close()
		player.Play()
		time.Sleep(duration + 100*time.Millisecond)
	}()
	
	return nil
}

func (a *AudioService) PlayStartSound() error {
	return a.PlayBeep(800, 200*time.Millisecond)
}

func (a *AudioService) PlayEndSound() error {
	// セッション終了を強力に通知
	a.PlayBeep(600, 200*time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	a.PlayBeep(800, 200*time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	a.PlayBeep(1000, 300*time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	// 追加の強調音
	a.PlayBeep(1200, 400*time.Millisecond)
	return nil
}

func (a *AudioService) PlayWarningSound() error {
	// より強力で持続的な警告音を再生
	for i := 0; i < 5; i++ {
		a.PlayBeep(1400, 200*time.Millisecond) // 高い周波数で長時間
		time.Sleep(100 * time.Millisecond)
		a.PlayBeep(800, 200*time.Millisecond)  // 低い周波数で対比
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func (a *AudioService) SetVolume(volume float64) {
	if volume < 0 {
		volume = 0
	}
	if volume > 1 {
		volume = 1
	}
	a.volume = volume
}

func (a *AudioService) Close() error {
	return nil
}

func (a *AudioService) playFromReader(reader io.Reader) error {
	if !a.isReady {
		return nil
	}
	
	player := a.context.NewPlayer(reader)
	go func() {
		defer player.Close()
		player.Play()
		time.Sleep(2 * time.Second)
	}()
	
	return nil
}