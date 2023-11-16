package transcoder

import (
	"context"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go-streams/internal/transcoder/videoformat"
)

type FileConversionResult struct {
	ResultPath string
	Err        error
}

type FileConversionRequest struct {
	OriginalFilePath string
	ResultFilePath   string
	TargetFormat     videoformat.VideoFormat
}

func Mp4ToM3u8(ctx context.Context, request FileConversionRequest) chan FileConversionResult {
	result := make(chan FileConversionResult)
	go func() {
		defer close(result)
		err := ffmpeg.Input(
			request.OriginalFilePath,
		).Output(
			request.ResultFilePath, ffmpeg.KwArgs{"format": request.TargetFormat, "start_number": "0", "hls_time": "10", "hls_list_size": "0"},
		).OverWriteOutput().Run()

		if err != nil {
			result <- FileConversionResult{ResultPath: "", Err: err}
		} else {
			result <- FileConversionResult{ResultPath: request.ResultFilePath, Err: nil}
		}
	}()

	return result
}
