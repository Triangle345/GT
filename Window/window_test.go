// window_test.go
package Window

import (
	"GT/Graphics"
	"GT/Window"
	"testing"
)

func Test_window(t *testing.T) {
	w := Window.NewWindowedWindow("test", 800, 600)

	s, _ := Graphics.NewBasicScene("test.png", 800, 600)

	for {
		s.Draw()
		w.Refresh()
	}

	// // e := w.Open()

	// // if e != nil {
	// // 	t.Error("Window open failure: " + e.Error())
	// // }

	// // if w.isOpen() == false {
	// // 	t.Error("Window should be open but it's not")
	// // }

	// // if w.Width != 800 {
	// // 	t.Error("Window width should be 800")
	// // }

	// // if w.Height != 600 {
	// // 	t.Error("Window height should be 600")
	// // }

	// running := true

	// for running == true {

	// 	w.Refresh()

	// }

	w.Close()

	// if w.isOpen() == true {
	// 	t.Error("Window should be closed but it's not")
	// }

}
