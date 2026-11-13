package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type StreamWithTimestampsTextToSpeechRequestOutputFormat string

const (
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp32205032   StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_22050_32"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp32400048   StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_24000_48"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410032   StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_32"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410064   StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_64"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410096   StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_96"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100128  StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_128"
	StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100192  StreamWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_192"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm8000      StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_8000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm16000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_16000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm22050     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_22050"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm24000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_24000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm32000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_32000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm44100     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_44100"
	StreamWithTimestampsTextToSpeechRequestOutputFormatPcm48000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "pcm_48000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatUlaw8000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "ulaw_8000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatAlaw8000     StreamWithTimestampsTextToSpeechRequestOutputFormat = "alaw_8000"
	StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800032  StreamWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_32"
	StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800064  StreamWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_64"
	StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800096  StreamWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_96"
	StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000128 StreamWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_128"
	StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000192 StreamWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_192"
)

func NewStreamWithTimestampsTextToSpeechRequestOutputFormatFromString(s string) (StreamWithTimestampsTextToSpeechRequestOutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000192, nil
	}
	var t StreamWithTimestampsTextToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StreamWithTimestampsTextToSpeechRequestOutputFormat) Ptr() *StreamWithTimestampsTextToSpeechRequestOutputFormat {
	return &e
}
