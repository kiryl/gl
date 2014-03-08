// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// AttribLocation

type AttribLocation int

func (indx AttribLocation) Attrib1f(x float32) {
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}

func (indx AttribLocation) Attrib1fv(values *[1]float32) {
	C.glVertexAttrib1fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}

func (indx AttribLocation) Attrib2f(x float32, y float32) {
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}

func (indx AttribLocation) Attrib2fv(values *[2]float32) {
	C.glVertexAttrib2fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}

func (indx AttribLocation) Attrib3f(x float32, y float32, z float32) {
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (indx AttribLocation) Attrib3fv(values *[3]float32) {
	C.glVertexAttrib3fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}

func (indx AttribLocation) Attrib4f(x float32, y float32, z float32, w float32) {
	C.glVertexAttrib4f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (indx AttribLocation) Attrib4fv(values *[4]float32) {
	C.glVertexAttrib4fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}

func (indx AttribLocation) Attrib1s(x int16) {
	C.glVertexAttrib1s(C.GLuint(indx), C.GLshort(x))
}

func (indx AttribLocation) Attrib1sv(values *[1]int16) {
	C.glVertexAttrib1sv(C.GLuint(indx), (*C.GLshort)(&values[0]))
}

func (indx AttribLocation) Attrib2s(x int16, y int16) {
	C.glVertexAttrib2s(C.GLuint(indx), C.GLshort(x), C.GLshort(y))
}

func (indx AttribLocation) Attrib2sv(values *[2]int16) {
	C.glVertexAttrib2sv(C.GLuint(indx), (*C.GLshort)(&values[0]))
}

func (indx AttribLocation) Attrib3s(x int16, y int16, z int16) {
	C.glVertexAttrib3s(C.GLuint(indx), C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

func (indx AttribLocation) Attrib3sv(values *[3]int16) {
	C.glVertexAttrib3sv(C.GLuint(indx), (*C.GLshort)(&values[0]))
}

func (indx AttribLocation) Attrib4s(x int16, y int16, z int16, w int16) {
	C.glVertexAttrib4s(C.GLuint(indx), C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

func (indx AttribLocation) Attrib4sv(values *[4]int16) {
	C.glVertexAttrib4sv(C.GLuint(indx), (*C.GLshort)(&values[0]))
}

func (indx AttribLocation) Attrib1d(x float64) {
	C.glVertexAttrib1d(C.GLuint(indx), C.GLdouble(x))
}

func (indx AttribLocation) Attrib1dv(values *[1]float64) {
	C.glVertexAttrib1dv(C.GLuint(indx), (*C.GLdouble)(&values[0]))
}

func (indx AttribLocation) Attrib2d(x float64, y float64) {
	C.glVertexAttrib2d(C.GLuint(indx), C.GLdouble(x), C.GLdouble(y))
}

func (indx AttribLocation) Attrib2dv(values *[2]float64) {
	C.glVertexAttrib2dv(C.GLuint(indx), (*C.GLdouble)(&values[0]))
}

func (indx AttribLocation) Attrib3d(x float64, y float64, z float64) {
	C.glVertexAttrib3d(C.GLuint(indx), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func (indx AttribLocation) Attrib3dv(values *[3]float64) {
	C.glVertexAttrib3dv(C.GLuint(indx), (*C.GLdouble)(&values[0]))
}

func (indx AttribLocation) Attrib4d(x float64, y float64, z float64, w float64) {
	C.glVertexAttrib4d(C.GLuint(indx), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

func (indx AttribLocation) Attrib4dv(values *[4]float64) {
	C.glVertexAttrib4dv(C.GLuint(indx), (*C.GLdouble)(&values[0]))
}

func (indx AttribLocation) AttribI1i(x int) {
	C.glVertexAttribI1i(C.GLuint(indx), C.GLint(x))
}

func (indx AttribLocation) AttribI1iv(values *[1]int32) {
	C.glVertexAttribI1iv(C.GLuint(indx), (*C.GLint)(&values[0]))
}

func (indx AttribLocation) AttribI2i(x int, y int) {
	C.glVertexAttribI2i(C.GLuint(indx), C.GLint(x), C.GLint(y))
}

func (indx AttribLocation) AttribI2iv(values *[2]int32) {
	C.glVertexAttribI2iv(C.GLuint(indx), (*C.GLint)(&values[0]))
}

func (indx AttribLocation) AttribI3i(x int, y int, z int) {
	C.glVertexAttribI3i(C.GLuint(indx), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (indx AttribLocation) AttribI3iv(values *[3]int32) {
	C.glVertexAttribI3iv(C.GLuint(indx), (*C.GLint)(&values[0]))
}

func (indx AttribLocation) AttribI4i(x int, y int, z int, w int) {
	C.glVertexAttribI4i(C.GLuint(indx), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (indx AttribLocation) AttribI4iv(values *[4]int32) {
	C.glVertexAttribI4iv(C.GLuint(indx), (*C.GLint)(&values[0]))
}

func (indx AttribLocation) AttribI1ui(x uint) {
	C.glVertexAttribI1ui(C.GLuint(indx), C.GLuint(x))
}

func (indx AttribLocation) AttribI1uiv(values *[1]uint32) {
	C.glVertexAttribI1uiv(C.GLuint(indx), (*C.GLuint)(&values[0]))
}

func (indx AttribLocation) AttribI2ui(x uint, y uint) {
	C.glVertexAttribI2ui(C.GLuint(indx), C.GLuint(x), C.GLuint(y))
}

func (indx AttribLocation) AttribI2uiv(values *[2]uint32) {
	C.glVertexAttribI2uiv(C.GLuint(indx), (*C.GLuint)(&values[0]))
}

func (indx AttribLocation) AttribI3ui(x uint, y uint, z uint) {
	C.glVertexAttribI3ui(C.GLuint(indx), C.GLuint(x), C.GLuint(y), C.GLuint(z))
}

func (indx AttribLocation) AttribI3uiv(values *[3]uint32) {
	C.glVertexAttribI3uiv(C.GLuint(indx), (*C.GLuint)(&values[0]))
}

func (indx AttribLocation) AttribI4ui(x uint, y uint, z uint, w uint) {
	C.glVertexAttribI4ui(C.GLuint(indx), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w))
}

func (indx AttribLocation) AttribI4uiv(values *[4]uint32) {
	C.glVertexAttribI4uiv(C.GLuint(indx), (*C.GLuint)(&values[0]))
}

func (indx AttribLocation) AttribPointer(size uint, typ GLenum, normalized bool, stride int, pointer interface{}) {
	C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(typ),
		glBool(normalized), C.GLsizei(stride), ptr(pointer))
}

func (indx AttribLocation) EnableArray() {
	C.glEnableVertexAttribArray(C.GLuint(indx))
}

func (indx AttribLocation) DisableArray() {
	C.glDisableVertexAttribArray(C.GLuint(indx))
}

func (indx AttribLocation) AttribDivisor(divisor int) {
	C.glVertexAttribDivisor(C.GLuint(indx), C.GLuint(divisor))
}
