// window_test.go
package Window

import (
	"testing"
)

func Test_window(t *testing.T) {
	w := NewWindowedWindow("test", 800, 600)
	e := w.Open()

	if e != nil {
		t.Error("Window open failure: " + e.Error())
	}

	if w.isOpen() == false {
		t.Error("Window should be open but it's not")
	}

	if w.width != 800 {
		t.Error("Window width should be 800")
	}

	if w.height != 600 {
		t.Error("Window height should be 600")
	}

	w.Close()

	if w.isOpen() == true {
		t.Error("Window should be closed but it's not")
	}

}
