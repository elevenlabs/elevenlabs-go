package types

import (
	"fmt"
)

type ComposeDetailedMusicRequestOutputFormat string

const (
	ComposeDetailedMusicRequestOutputFormatAuto         ComposeDetailedMusicRequestOutputFormat = "auto"
	ComposeDetailedMusicRequestOutputFormatMp348000128  ComposeDetailedMusicRequestOutputFormat = "mp3_48000_128"
	ComposeDetailedMusicRequestOutputFormatMp348000192  ComposeDetailedMusicRequestOutputFormat = "mp3_48000_192"
	ComposeDetailedMusicRequestOutputFormatMp348000240  ComposeDetailedMusicRequestOutputFormat = "mp3_48000_240"
	ComposeDetailedMusicRequestOutputFormatMp348000320  ComposeDetailedMusicRequestOutputFormat = "mp3_48000_320"
	ComposeDetailedMusicRequestOutputFormatMp32205032   ComposeDetailedMusicRequestOutputFormat = "mp3_22050_32"
	ComposeDetailedMusicRequestOutputFormatMp32400048   ComposeDetailedMusicRequestOutputFormat = "mp3_24000_48"
	ComposeDetailedMusicRequestOutputFormatMp34410032   ComposeDetailedMusicRequestOutputFormat = "mp3_44100_32"
	ComposeDetailedMusicRequestOutputFormatMp34410064   ComposeDetailedMusicRequestOutputFormat = "mp3_44100_64"
	ComposeDetailedMusicRequestOutputFormatMp34410096   ComposeDetailedMusicRequestOutputFormat = "mp3_44100_96"
	ComposeDetailedMusicRequestOutputFormatMp344100128  ComposeDetailedMusicRequestOutputFormat = "mp3_44100_128"
	ComposeDetailedMusicRequestOutputFormatMp344100192  ComposeDetailedMusicRequestOutputFormat = "mp3_44100_192"
	ComposeDetailedMusicRequestOutputFormatPcm8000      ComposeDetailedMusicRequestOutputFormat = "pcm_8000"
	ComposeDetailedMusicRequestOutputFormatPcm16000     ComposeDetailedMusicRequestOutputFormat = "pcm_16000"
	ComposeDetailedMusicRequestOutputFormatPcm22050     ComposeDetailedMusicRequestOutputFormat = "pcm_22050"
	ComposeDetailedMusicRequestOutputFormatPcm24000     ComposeDetailedMusicRequestOutputFormat = "pcm_24000"
	ComposeDetailedMusicRequestOutputFormatPcm32000     ComposeDetailedMusicRequestOutputFormat = "pcm_32000"
	ComposeDetailedMusicRequestOutputFormatPcm44100     ComposeDetailedMusicRequestOutputFormat = "pcm_44100"
	ComposeDetailedMusicRequestOutputFormatPcm48000     ComposeDetailedMusicRequestOutputFormat = "pcm_48000"
	ComposeDetailedMusicRequestOutputFormatUlaw8000     ComposeDetailedMusicRequestOutputFormat = "ulaw_8000"
	ComposeDetailedMusicRequestOutputFormatAlaw8000     ComposeDetailedMusicRequestOutputFormat = "alaw_8000"
	ComposeDetailedMusicRequestOutputFormatOpus4800032  ComposeDetailedMusicRequestOutputFormat = "opus_48000_32"
	ComposeDetailedMusicRequestOutputFormatOpus4800064  ComposeDetailedMusicRequestOutputFormat = "opus_48000_64"
	ComposeDetailedMusicRequestOutputFormatOpus4800096  ComposeDetailedMusicRequestOutputFormat = "opus_48000_96"
	ComposeDetailedMusicRequestOutputFormatOpus48000128 ComposeDetailedMusicRequestOutputFormat = "opus_48000_128"
	ComposeDetailedMusicRequestOutputFormatOpus48000192 ComposeDetailedMusicRequestOutputFormat = "opus_48000_192"
)

func NewComposeDetailedMusicRequestOutputFormatFromString(s string) (ComposeDetailedMusicRequestOutputFormat, error) {
	switch s {
	case "auto":
		return ComposeDetailedMusicRequestOutputFormatAuto, nil
	case "mp3_48000_128":
		return ComposeDetailedMusicRequestOutputFormatMp348000128, nil
	case "mp3_48000_192":
		return ComposeDetailedMusicRequestOutputFormatMp348000192, nil
	case "mp3_48000_240":
		return ComposeDetailedMusicRequestOutputFormatMp348000240, nil
	case "mp3_48000_320":
		return ComposeDetailedMusicRequestOutputFormatMp348000320, nil
	case "mp3_22050_32":
		return ComposeDetailedMusicRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ComposeDetailedMusicRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return ComposeDetailedMusicRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ComposeDetailedMusicRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ComposeDetailedMusicRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return ComposeDetailedMusicRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ComposeDetailedMusicRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return ComposeDetailedMusicRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return ComposeDetailedMusicRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ComposeDetailedMusicRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ComposeDetailedMusicRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ComposeDetailedMusicRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ComposeDetailedMusicRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ComposeDetailedMusicRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return ComposeDetailedMusicRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return ComposeDetailedMusicRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return ComposeDetailedMusicRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ComposeDetailedMusicRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ComposeDetailedMusicRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return ComposeDetailedMusicRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ComposeDetailedMusicRequestOutputFormatOpus48000192, nil
	}
	var t ComposeDetailedMusicRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ComposeDetailedMusicRequestOutputFormat) Ptr() *ComposeDetailedMusicRequestOutputFormat {
	return &e
}
