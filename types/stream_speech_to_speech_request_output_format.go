package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type StreamSpeechToSpeechRequestOutputFormat string

const (
	StreamSpeechToSpeechRequestOutputFormatMp32205032   StreamSpeechToSpeechRequestOutputFormat = "mp3_22050_32"
	StreamSpeechToSpeechRequestOutputFormatMp32400048   StreamSpeechToSpeechRequestOutputFormat = "mp3_24000_48"
	StreamSpeechToSpeechRequestOutputFormatMp34410032   StreamSpeechToSpeechRequestOutputFormat = "mp3_44100_32"
	StreamSpeechToSpeechRequestOutputFormatMp34410064   StreamSpeechToSpeechRequestOutputFormat = "mp3_44100_64"
	StreamSpeechToSpeechRequestOutputFormatMp34410096   StreamSpeechToSpeechRequestOutputFormat = "mp3_44100_96"
	StreamSpeechToSpeechRequestOutputFormatMp344100128  StreamSpeechToSpeechRequestOutputFormat = "mp3_44100_128"
	StreamSpeechToSpeechRequestOutputFormatMp344100192  StreamSpeechToSpeechRequestOutputFormat = "mp3_44100_192"
	StreamSpeechToSpeechRequestOutputFormatPcm8000      StreamSpeechToSpeechRequestOutputFormat = "pcm_8000"
	StreamSpeechToSpeechRequestOutputFormatPcm16000     StreamSpeechToSpeechRequestOutputFormat = "pcm_16000"
	StreamSpeechToSpeechRequestOutputFormatPcm22050     StreamSpeechToSpeechRequestOutputFormat = "pcm_22050"
	StreamSpeechToSpeechRequestOutputFormatPcm24000     StreamSpeechToSpeechRequestOutputFormat = "pcm_24000"
	StreamSpeechToSpeechRequestOutputFormatPcm32000     StreamSpeechToSpeechRequestOutputFormat = "pcm_32000"
	StreamSpeechToSpeechRequestOutputFormatPcm44100     StreamSpeechToSpeechRequestOutputFormat = "pcm_44100"
	StreamSpeechToSpeechRequestOutputFormatPcm48000     StreamSpeechToSpeechRequestOutputFormat = "pcm_48000"
	StreamSpeechToSpeechRequestOutputFormatUlaw8000     StreamSpeechToSpeechRequestOutputFormat = "ulaw_8000"
	StreamSpeechToSpeechRequestOutputFormatAlaw8000     StreamSpeechToSpeechRequestOutputFormat = "alaw_8000"
	StreamSpeechToSpeechRequestOutputFormatOpus4800032  StreamSpeechToSpeechRequestOutputFormat = "opus_48000_32"
	StreamSpeechToSpeechRequestOutputFormatOpus4800064  StreamSpeechToSpeechRequestOutputFormat = "opus_48000_64"
	StreamSpeechToSpeechRequestOutputFormatOpus4800096  StreamSpeechToSpeechRequestOutputFormat = "opus_48000_96"
	StreamSpeechToSpeechRequestOutputFormatOpus48000128 StreamSpeechToSpeechRequestOutputFormat = "opus_48000_128"
	StreamSpeechToSpeechRequestOutputFormatOpus48000192 StreamSpeechToSpeechRequestOutputFormat = "opus_48000_192"
)

func NewStreamSpeechToSpeechRequestOutputFormatFromString(s string) (StreamSpeechToSpeechRequestOutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return StreamSpeechToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return StreamSpeechToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return StreamSpeechToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return StreamSpeechToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return StreamSpeechToSpeechRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return StreamSpeechToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return StreamSpeechToSpeechRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return StreamSpeechToSpeechRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return StreamSpeechToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return StreamSpeechToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return StreamSpeechToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return StreamSpeechToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return StreamSpeechToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return StreamSpeechToSpeechRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return StreamSpeechToSpeechRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return StreamSpeechToSpeechRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return StreamSpeechToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return StreamSpeechToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return StreamSpeechToSpeechRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return StreamSpeechToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return StreamSpeechToSpeechRequestOutputFormatOpus48000192, nil
	}
	var t StreamSpeechToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StreamSpeechToSpeechRequestOutputFormat) Ptr() *StreamSpeechToSpeechRequestOutputFormat {
	return &e
}
