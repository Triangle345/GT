// Scene
package Graphics

import (
//"errors"
//gl "github.com/chsc/gogl/gl21"
)

type Scene interface {
	Load()
	Clear()
	Draw()
}

type BaseScene struct {
	entities      []*Drawable
	width, height int
}

//func NewBasicScene(w, h int) (error, Scene) {
//	s := Scene{width: w, height: h}
//	err := s.init()

//	if err != nil {
//		return err, s
//	}

//	return nil, s

//}

//func (s BaseScene) load() error {
//	if err := gl.Init(); err != nil {
//		return errors.New("Cannot initialize OGL: " + err.Error())
//	}

//	gl.MatrixMode(gl.PROJECTION)
//	gl.Disable(gl.DEPTH_TEST)

//	gl.Ortho(0, gl.Double(s.width), gl.Double(s.height), 0, -1, 1)

//	return nil
//}

//func (s *BaseScene) Add(sprites ...*Drawable) {
//	s.entities = append(s.entities, sprites...)

//	//for _, v := range sprites {
//	//	s.entities = append(s.entities, v)
//	//}
//}

//func (s BaseScene) clear() {
//	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

//}

//func (s BaseScene) Draw() {

//	for _, v := range s.entities {
//		(*v).draw()
//	}

//}
