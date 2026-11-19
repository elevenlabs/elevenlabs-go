package types

import (
	"fmt"
)

type StreamMusicRequestOutputFormat string

const (
	StreamMusicRequestOutputFormatAuto         StreamMusicRequestOutputFormat = "auto"
	StreamMusicRequestOutputFormatMp348000128  StreamMusicRequestOutputFormat = "mp3_48000_128"
	StreamMusicRequestOutputFormatMp348000192  StreamMusicRequestOutputFormat = "mp3_48000_192"
	StreamMusicRequestOutputFormatMp348000240  StreamMusicRequestOutputFormat = "mp3_48000_240"
	StreamMusicRequestOutputFormatMp348000320  StreamMusicRequestOutputFormat = "mp3_48000_320"
	StreamMusicRequestOutputFormatMp32205032   StreamMusicRequestOutputFormat = "mp3_22050_32"
	StreamMusicRequestOutputFormatMp32400048   StreamMusicRequestOutputFormat = "mp3_24000_48"
	StreamMusicRequestOutputFormatMp34410032   StreamMusicRequestOutputFormat = "mp3_44100_32"
	StreamMusicRequestOutputFormatMp34410064   StreamMusicRequestOutputFormat = "mp3_44100_64"
	StreamMusicRequestOutputFormatMp34410096   StreamMusicRequestOutputFormat = "mp3_44100_96"
	StreamMusicRequestOutputFormatMp344100128  StreamMusicRequestOutputFormat = "mp3_44100_128"
	StreamMusicRequestOutputFormatMp344100192  StreamMusicRequestOutputFormat = "mp3_44100_192"
	StreamMusicRequestOutputFormatPcm8000      StreamMusicRequestOutputFormat = "pcm_8000"
	StreamMusicRequestOutputFormatPcm16000     StreamMusicRequestOutputFormat = "pcm_16000"
	StreamMusicRequestOutputFormatPcm22050     StreamMusicRequestOutputFormat = "pcm_22050"
	StreamMusicRequestOutputFormatPcm24000     StreamMusicRequestOutputFormat = "pcm_24000"
	StreamMusicRequestOutputFormatPcm32000     StreamMusicRequestOutputFormat = "pcm_32000"
	StreamMusicRequestOutputFormatPcm44100     StreamMusicRequestOutputFormat = "pcm_44100"
	StreamMusicRequestOutputFormatPcm48000     StreamMusicRequestOutputFormat = "pcm_48000"
	StreamMusicRequestOutputFormatUlaw8000     StreamMusicRequestOutputFormat = "ulaw_8000"
	StreamMusicRequestOutputFormatAlaw8000     StreamMusicRequestOutputFormat = "alaw_8000"
	StreamMusicRequestOutputFormatOpus4800032  StreamMusicRequestOutputFormat = "opus_48000_32"
	StreamMusicRequestOutputFormatOpus4800064  StreamMusicRequestOutputFormat = "opus_48000_64"
	StreamMusicRequestOutputFormatOpus4800096  StreamMusicRequestOutputFormat = "opus_48000_96"
	StreamMusicRequestOutputFormatOpus48000128 StreamMusicRequestOutputFormat = "opus_48000_128"
	StreamMusicRequestOutputFormatOpus48000192 StreamMusicRequestOutputFormat = "opus_48000_192"
)

func NewStreamMusicRequestOutputFormatFromString(s string) (StreamMusicRequestOutputFormat, error) {
	switch s {
	case "auto":
		return StreamMusicRequestOutputFormatAuto, nil
	case "mp3_48000_128":
		return StreamMusicRequestOutputFormatMp348000128, nil
	case "mp3_48000_192":
		return StreamMusicRequestOutputFormatMp348000192, nil
	case "mp3_48000_240":
		return StreamMusicRequestOutputFormatMp348000240, nil
	case "mp3_48000_320":
		return StreamMusicRequestOutputFormatMp348000320, nil
	case "mp3_22050_32":
		return StreamMusicRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return StreamMusicRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return StreamMusicRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return StreamMusicRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return StreamMusicRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return StreamMusicRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return StreamMusicRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return StreamMusicRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return StreamMusicRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return StreamMusicRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return StreamMusicRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return StreamMusicRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return StreamMusicRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return StreamMusicRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return StreamMusicRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return StreamMusicRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return StreamMusicRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return StreamMusicRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return StreamMusicRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return StreamMusicRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return StreamMusicRequestOutputFormatOpus48000192, nil
	}
	var t StreamMusicRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StreamMusicRequestOutputFormat) Ptr() *StreamMusicRequestOutputFormat {
	return &e
}
