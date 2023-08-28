package audiostripper

import (
	"bytes"
	"context"
	"fmt"
)

type (
	// ExtractAudioInput defines the input for the ExtractAudio method.
	ExtractAudioInput struct {
		SampleRate string
		FilePath   string
	}

	// ExtractAudioOutput defines the output for the ExtractAudio method
	ExtractAudioOutput struct {
		FilePath string
	}

	ExtractCmdParams struct {
		InputFile, OutputFile, SampleRate string
		Stderr                            *bytes.Buffer
	}

	// ExtractCmd is a function that runs the extractor command.
	ExtractCmd func(params *ExtractCmdParams) error

	// Audiostripper provides methods for extracting audio from a video file.
	Audiostripper struct {
		cmd ExtractCmd
	}
)

// New creates a new audtiostripper instance.
func New(cmd ExtractCmd) *Audiostripper {
	return &Audiostripper{
		cmd: cmd,
	}
}

// ExtractAudio extracts audio from a video file.
func (a *Audiostripper) ExtractAudio(ctx context.Context, in *ExtractAudioInput) (*ExtractAudioOutput, error) {
	cmdParams := ExtractCmdParams{
		InputFile:  in.FilePath,
		OutputFile: outputFilePath(in.FilePath),
		SampleRate: in.SampleRate,
		Stderr:     &bytes.Buffer{},
	}

	if err := a.cmd(&cmdParams); err != nil {
		return nil, fmt.Errorf("could not run extractor command: %s", err)
	}

	return &ExtractAudioOutput{
		FilePath: cmdParams.OutputFile,
	}, nil
}

func outputFilePath(in string) string {
	return in[:len(in)-4] + ".wav"
}
