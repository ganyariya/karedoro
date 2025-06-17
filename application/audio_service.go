package application

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"time"

	"github.com/ebitengine/oto/v3"
	"karedoro/domain"
)

const (
	// Audio system constants
	SampleRate   = 44100
	ChannelCount = 2
	DefaultVolume = 0.7
	
	// Audio generation constants
	MaxAmplitude = 32767
	BaseAmplitude = 0.3
	Harmonic2Amplitude = 0.3
	Harmonic3Amplitude = 0.1
	
	// Sound frequencies and durations
	StartSoundFreq = 800
	StartSoundDuration = 200 * time.Millisecond
	
	EndSound1Freq = 600
	EndSound1Duration = 200 * time.Millisecond
	EndSound2Freq = 800
	EndSound2Duration = 200 * time.Millisecond
	EndSound3Freq = 1000
	EndSound3Duration = 300 * time.Millisecond
	EndSound4Freq = 1200
	EndSound4Duration = 400 * time.Millisecond
	EndSoundGap = 50 * time.Millisecond
	EndSoundLongGap = 100 * time.Millisecond
	
	WarningHighFreq = 1400
	WarningLowFreq = 800
	WarningDuration = 200 * time.Millisecond
	WarningGap = 100 * time.Millisecond
	WarningCycles = 5
	
	PauseBeepFreq = 400
	ResumeBeepFreq = 600
	BeepDuration = 100 * time.Millisecond
)

type AudioService struct {
	context      *oto.Context
	readyChannel chan struct{}
	isReady      bool
	volume       float64
	initError    error
}

func NewAudioService() *AudioService {
	service := &AudioService{
		readyChannel: make(chan struct{}),
		isReady:      false,
		volume:       DefaultVolume,
	}
	
	go service.initialize()
	
	return service
}

func (a *AudioService) initialize() {
	op := &oto.NewContextOptions{
		SampleRate:   SampleRate,
		ChannelCount: ChannelCount,
		Format:       oto.FormatSignedInt16LE,
	}
	
	context, readyChan, err := oto.NewContext(op)
	if err != nil {
		a.initError = domain.NewTimerError("initialize", fmt.Errorf("%w: %v", domain.ErrAudioInitialization, err))
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
	if a.initError != nil {
		return a.initError
	}
	if !a.isReady {
		return domain.ErrAudioNotReady
	}
	
	samples := int(float64(SampleRate) * duration.Seconds())
	
	buf := make([]byte, samples*4)
	
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(SampleRate)
		sample := int16(MaxAmplitude * a.volume * BaseAmplitude * 
			(math.Sin(2*math.Pi*frequency*t) + 
			 math.Sin(2*math.Pi*frequency*2*t)*Harmonic2Amplitude + 
			 math.Sin(2*math.Pi*frequency*3*t)*Harmonic3Amplitude))
		
		buf[i*4] = byte(sample)
		buf[i*4+1] = byte(sample >> 8)
		buf[i*4+2] = byte(sample)
		buf[i*4+3] = byte(sample >> 8)
	}
	
	player := a.context.NewPlayer(bytes.NewReader(buf))
	go func() {
		defer player.Close()
		player.Play()
		time.Sleep(duration + BeepDuration)
	}()
	
	return nil
}

func (a *AudioService) PlayStartSound() error {
	return a.PlayBeep(StartSoundFreq, StartSoundDuration)
}

func (a *AudioService) PlayEndSound() error {
	// セッション終了を強力に通知
	if err := a.PlayBeep(EndSound1Freq, EndSound1Duration); err != nil {
		return fmt.Errorf("%w: first beep failed: %v", domain.ErrAudioPlayback, err)
	}
	time.Sleep(EndSoundGap)
	if err := a.PlayBeep(EndSound2Freq, EndSound2Duration); err != nil {
		return fmt.Errorf("%w: second beep failed: %v", domain.ErrAudioPlayback, err)
	}
	time.Sleep(EndSoundGap)
	if err := a.PlayBeep(EndSound3Freq, EndSound3Duration); err != nil {
		return fmt.Errorf("%w: third beep failed: %v", domain.ErrAudioPlayback, err)
	}
	time.Sleep(EndSoundLongGap)
	// 追加の強調音
	if err := a.PlayBeep(EndSound4Freq, EndSound4Duration); err != nil {
		return fmt.Errorf("%w: fourth beep failed: %v", domain.ErrAudioPlayback, err)
	}
	return nil
}

func (a *AudioService) PlayWarningSound() error {
	// より強力で持続的な警告音を再生
	for i := 0; i < WarningCycles; i++ {
		if err := a.PlayBeep(WarningHighFreq, WarningDuration); err != nil {
			return fmt.Errorf("%w: warning high beep cycle %d failed: %v", domain.ErrAudioPlayback, i, err)
		}
		time.Sleep(WarningGap)
		if err := a.PlayBeep(WarningLowFreq, WarningDuration); err != nil {
			return fmt.Errorf("%w: warning low beep cycle %d failed: %v", domain.ErrAudioPlayback, i, err)
		}
		time.Sleep(WarningGap)
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

func (a *AudioService) IsReady() bool {
	return a.isReady && a.initError == nil
}

func (a *AudioService) Close() error {
	// oto.Context does not have a Close method in v3
	// The context will be garbage collected when no longer referenced
	return nil
}

