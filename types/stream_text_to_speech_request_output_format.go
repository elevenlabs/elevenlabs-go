package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type StreamTextToSpeechRequestOutputFormat string

const (
	StreamTextToSpeechRequestOutputFormatMp32205032   StreamTextToSpeechRequestOutputFormat = "mp3_22050_32"
	StreamTextToSpeechRequestOutputFormatMp32400048   StreamTextToSpeechRequestOutputFormat = "mp3_24000_48"
	StreamTextToSpeechRequestOutputFormatMp34410032   StreamTextToSpeechRequestOutputFormat = "mp3_44100_32"
	StreamTextToSpeechRequestOutputFormatMp34410064   StreamTextToSpeechRequestOutputFormat = "mp3_44100_64"
	StreamTextToSpeechRequestOutputFormatMp34410096   StreamTextToSpeechRequestOutputFormat = "mp3_44100_96"
	StreamTextToSpeechRequestOutputFormatMp344100128  StreamTextToSpeechRequestOutputFormat = "mp3_44100_128"
	StreamTextToSpeechRequestOutputFormatMp344100192  StreamTextToSpeechRequestOutputFormat = "mp3_44100_192"
	StreamTextToSpeechRequestOutputFormatPcm8000      StreamTextToSpeechRequestOutputFormat = "pcm_8000"
	StreamTextToSpeechRequestOutputFormatPcm16000     StreamTextToSpeechRequestOutputFormat = "pcm_16000"
	StreamTextToSpeechRequestOutputFormatPcm22050     StreamTextToSpeechRequestOutputFormat = "pcm_22050"
	StreamTextToSpeechRequestOutputFormatPcm24000     StreamTextToSpeechRequestOutputFormat = "pcm_24000"
	StreamTextToSpeechRequestOutputFormatPcm32000     StreamTextToSpeechRequestOutputFormat = "pcm_32000"
	StreamTextToSpeechRequestOutputFormatPcm44100     StreamTextToSpeechRequestOutputFormat = "pcm_44100"
	StreamTextToSpeechRequestOutputFormatPcm48000     StreamTextToSpeechRequestOutputFormat = "pcm_48000"
	StreamTextToSpeechRequestOutputFormatUlaw8000     StreamTextToSpeechRequestOutputFormat = "ulaw_8000"
	StreamTextToSpeechRequestOutputFormatAlaw8000     StreamTextToSpeechRequestOutputFormat = "alaw_8000"
	StreamTextToSpeechRequestOutputFormatOpus4800032  StreamTextToSpeechRequestOutputFormat = "opus_48000_32"
	StreamTextToSpeechRequestOutputFormatOpus4800064  StreamTextToSpeechRequestOutputFormat = "opus_48000_64"
	StreamTextToSpeechRequestOutputFormatOpus4800096  StreamTextToSpeechRequestOutputFormat = "opus_48000_96"
	StreamTextToSpeechRequestOutputFormatOpus48000128 StreamTextToSpeechRequestOutputFormat = "opus_48000_128"
	StreamTextToSpeechRequestOutputFormatOpus48000192 StreamTextToSpeechRequestOutputFormat = "opus_48000_192"
)

func NewStreamTextToSpeechRequestOutputFormatFromString(s string) (StreamTextToSpeechRequestOutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return StreamTextToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return StreamTextToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return StreamTextToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return StreamTextToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return StreamTextToSpeechRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return StreamTextToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return StreamTextToSpeechRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return StreamTextToSpeechRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return StreamTextToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return StreamTextToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return StreamTextToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return StreamTextToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return StreamTextToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return StreamTextToSpeechRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return StreamTextToSpeechRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return StreamTextToSpeechRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return StreamTextToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return StreamTextToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return StreamTextToSpeechRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return StreamTextToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return StreamTextToSpeechRequestOutputFormatOpus48000192, nil
	}
	var t StreamTextToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StreamTextToSpeechRequestOutputFormat) Ptr() *StreamTextToSpeechRequestOutputFormat {
	return &e
}
