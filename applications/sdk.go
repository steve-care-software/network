package applications

import (
	"steve.care/network/domain/frames"
	"steve.care/network/domain/programs"
)

// Application represents an application
type Application interface {
	Execute(programm programs.Program, frame frames.Frames) (frames.Frames, error)
}
