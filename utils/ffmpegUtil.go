package utils

import "os/exec"

func GetSnapshot(videoPath, snapshotPath string, frameNum int) {
	cmdString := []string{
		"-i",
		videoPath,
		snapshotPath,
		"-ss",
		"00:00:01",
		"-r",
		"1",
		"-vframes",
		string(rune(frameNum)),
		"-an",
		"-vcodec",
		"mjpeg",
	}
	cmd := exec.Command("ffmpeg", cmdString...)
	cmd.Run()

}
