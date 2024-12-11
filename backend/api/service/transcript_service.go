package service

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

const MODEL = "large"

func Transcript(input, outputDir string) error {
	cmd := exec.Command(
		"whisper",
		input,
		"--task", "translate",
		"--language", "Arabic",
		"--model", MODEL,
		"--output_dir", outputDir,
	)

	fmt.Println("Running Whisper for transcription...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run whisper command: %w", err)
	}

	inputBase := filepath.Base(input)
	inputName := inputBase[:len(inputBase)-len(filepath.Ext(inputBase))]

	subtitleFile := filepath.Join(outputDir, inputName+".srt")

	outputVideo := filepath.Join(outputDir, inputName+"_with_subtitles.mp4")

	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", input,
		"-vf", fmt.Sprintf("subtitles=%s", subtitleFile),
		"-c:a", "copy",
		outputVideo,
	)

	fmt.Println("Running FFmpeg to combine subtitles with video...")
	if err := ffmpegCmd.Run(); err != nil {
		return fmt.Errorf("failed to run ffmpeg command: %w", err)
	}

	fmt.Printf("Subtitled video saved to: %s\n", outputVideo)
	return nil
}
