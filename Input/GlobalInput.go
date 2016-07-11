// Package GlobalInput project GlobalInput.go
package GlobalInput

import (
	//"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

// map our key / buttons with true / false state
var inputMap map[string]bool

// TODO: map our positioning inputs (mouse / touch) with coordinates OR map by string

func init() {
	if inputMap == nil {
		inputMap = make(map[string]bool)
	}
}

// CheckForUpdates polls sdl events and updates the map accordingly
func CheckForUpdates() bool {
	changeDetected := false

	// poll any current events
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			inputMap["Quit"] = true
		case *sdl.KeyDownEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				inputMap["Esc"] = true
			} else {
				// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				// t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		default:
			// fmt.Printf("Some event\n")
		}

		changeDetected = true
	}

	return changeDetected
}

// GetInputStatus returns true if the desired key is pressed or false otherwise
// consider changing to string for coords etc.?
func GetInputStatus(keyToLookFor string) bool {
	result := false
	result, exists := inputMap[keyToLookFor]
	if !exists {
		//fmt.Printf(keyToLookFor + " not found, please try a valid key to look for\n")
	}
	return result
}
