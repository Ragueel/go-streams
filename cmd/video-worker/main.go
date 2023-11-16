package main

import (
	"context"
	"go-streams/internal/transcoder"
	"go-streams/internal/transcoder/videoformat"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	resultChan := transcoder.Mp4ToM3u8(ctx, transcoder.FileConversionRequest{
		OriginalFilePath: "/home/ali/Documents/PersonalProjects/go-streams/tests/data/test_video.mp4",
		ResultFilePath:   "/home/ali/Documents/PersonalProjects/go-streams/tests/data/result/test_video.m3u8",
		TargetFormat:     videoformat.HLS,
	})
	select {
	case fullResult := <-resultChan:
		if fullResult.Err != nil {
			println("Failed %s", fullResult.Err.Error())
		} else {
			println("Finished successfully")
		}
	case <-ctx.Done():
		println("Done here")
	}
}
