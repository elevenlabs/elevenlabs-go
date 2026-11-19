package types

import (
	"fmt"
)

type ComposeMusicRequestOutputFormat string

const (
	ComposeMusicRequestOutputFormatAuto         ComposeMusicRequestOutputFormat = "auto"
	ComposeMusicRequestOutputFormatMp348000128  ComposeMusicRequestOutputFormat = "mp3_48000_128"
	ComposeMusicRequestOutputFormatMp348000192  ComposeMusicRequestOutputFormat = "mp3_48000_192"
	ComposeMusicRequestOutputFormatMp348000240  ComposeMusicRequestOutputFormat = "mp3_48000_240"
	ComposeMusicRequestOutputFormatMp348000320  ComposeMusicRequestOutputFormat = "mp3_48000_320"
	ComposeMusicRequestOutputFormatMp32205032   ComposeMusicRequestOutputFormat = "mp3_22050_32"
	ComposeMusicRequestOutputFormatMp32400048   ComposeMusicRequestOutputFormat = "mp3_24000_48"
	ComposeMusicRequestOutputFormatMp34410032   ComposeMusicRequestOutputFormat = "mp3_44100_32"
	ComposeMusicRequestOutputFormatMp34410064   ComposeMusicRequestOutputFormat = "mp3_44100_64"
	ComposeMusicRequestOutputFormatMp34410096   ComposeMusicRequestOutputFormat = "mp3_44100_96"
	ComposeMusicRequestOutputFormatMp344100128  ComposeMusicRequestOutputFormat = "mp3_44100_128"
	ComposeMusicRequestOutputFormatMp344100192  ComposeMusicRequestOutputFormat = "mp3_44100_192"
	ComposeMusicRequestOutputFormatPcm8000      ComposeMusicRequestOutputFormat = "pcm_8000"
	ComposeMusicRequestOutputFormatPcm16000     ComposeMusicRequestOutputFormat = "pcm_16000"
	ComposeMusicRequestOutputFormatPcm22050     ComposeMusicRequestOutputFormat = "pcm_22050"
	ComposeMusicRequestOutputFormatPcm24000     ComposeMusicRequestOutputFormat = "pcm_24000"
	ComposeMusicRequestOutputFormatPcm32000     ComposeMusicRequestOutputFormat = "pcm_32000"
	ComposeMusicRequestOutputFormatPcm44100     ComposeMusicRequestOutputFormat = "pcm_44100"
	ComposeMusicRequestOutputFormatPcm48000     ComposeMusicRequestOutputFormat = "pcm_48000"
	ComposeMusicRequestOutputFormatUlaw8000     ComposeMusicRequestOutputFormat = "ulaw_8000"
	ComposeMusicRequestOutputFormatAlaw8000     ComposeMusicRequestOutputFormat = "alaw_8000"
	ComposeMusicRequestOutputFormatOpus4800032  ComposeMusicRequestOutputFormat = "opus_48000_32"
	ComposeMusicRequestOutputFormatOpus4800064  ComposeMusicRequestOutputFormat = "opus_48000_64"
	ComposeMusicRequestOutputFormatOpus4800096  ComposeMusicRequestOutputFormat = "opus_48000_96"
	ComposeMusicRequestOutputFormatOpus48000128 ComposeMusicRequestOutputFormat = "opus_48000_128"
	ComposeMusicRequestOutputFormatOpus48000192 ComposeMusicRequestOutputFormat = "opus_48000_192"
)

func NewComposeMusicRequestOutputFormatFromString(s string) (ComposeMusicRequestOutputFormat, error) {
	switch s {
	case "auto":
		return ComposeMusicRequestOutputFormatAuto, nil
	case "mp3_48000_128":
		return ComposeMusicRequestOutputFormatMp348000128, nil
	case "mp3_48000_192":
		return ComposeMusicRequestOutputFormatMp348000192, nil
	case "mp3_48000_240":
		return ComposeMusicRequestOutputFormatMp348000240, nil
	case "mp3_48000_320":
		return ComposeMusicRequestOutputFormatMp348000320, nil
	case "mp3_22050_32":
		return ComposeMusicRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ComposeMusicRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return ComposeMusicRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ComposeMusicRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ComposeMusicRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return ComposeMusicRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ComposeMusicRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return ComposeMusicRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return ComposeMusicRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ComposeMusicRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ComposeMusicRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ComposeMusicRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ComposeMusicRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ComposeMusicRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return ComposeMusicRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return ComposeMusicRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return ComposeMusicRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ComposeMusicRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ComposeMusicRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return ComposeMusicRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ComposeMusicRequestOutputFormatOpus48000192, nil
	}
	var t ComposeMusicRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ComposeMusicRequestOutputFormat) Ptr() *ComposeMusicRequestOutputFormat {
	return &e
}
