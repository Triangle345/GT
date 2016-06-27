package G3D

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type vertex struct {
	X, Y, Z float32
}

type vertex_normal struct {
	X, Y, Z float32
}

type face struct {
	v, vuv, vn []int
}

type Object struct {
	name  string
	vs    []vertex
	vns   []vertex_normal
	faces []face
}

func parseMtllib(i *int, dat []string) string {
	*i++
	return dat[*i]
}

func parseUseMtl(i *int, dat []string) string {
	*i++
	return dat[*i]
}

func parseObjectName(i *int, dat []string) string {
	*i++
	return dat[*i]
}

func parseVertex(i *int, dat []string) vertex {
	*i++
	v := vertex{}

	f, _ := strconv.ParseFloat(dat[*i], 32)

	v.X = float32(f)

	*i++
	f, _ = strconv.ParseFloat(dat[*i], 32)
	v.Y = float32(f)

	*i++
	f, _ = strconv.ParseFloat(dat[*i], 32)
	v.Z = float32(f)

	return v
}

func parseVertexNormal(i *int, dat []string) vertex_normal {
	*i++
	vn := vertex_normal{}

	f, _ := strconv.ParseFloat(dat[*i], 32)

	vn.X = float32(f)

	*i++
	f, _ = strconv.ParseFloat(dat[*i], 32)
	vn.Y = float32(f)

	*i++
	f, _ = strconv.ParseFloat(dat[*i], 32)
	vn.Z = float32(f)

	return vn
}

func parseFace(i *int, dat []string) face {
	*i++
	f := face{}
	fDat := strings.Split(dat[*i], "/")

	// first point of triangle
	v, _ := strconv.Atoi(fDat[0])
	vuv, _ := strconv.Atoi(fDat[1])
	vn, _ := strconv.Atoi(fDat[2])
	f.v = append(f.v, v)
	f.vuv = append(f.vuv, vuv)
	f.vn = append(f.vn, vn)

	*i++

	fDat = strings.Split(dat[*i], "/")

	// second point of triangle
	v, _ = strconv.Atoi(fDat[0])
	vuv, _ = strconv.Atoi(fDat[1])
	vn, _ = strconv.Atoi(fDat[2])
	f.v = append(f.v, v)
	f.vuv = append(f.vuv, vuv)
	f.vn = append(f.vn, vn)

	*i++

	fDat = strings.Split(dat[*i], "/")

	// third point of triangle
	v, _ = strconv.Atoi(fDat[0])
	vuv, _ = strconv.Atoi(fDat[1])
	vn, _ = strconv.Atoi(fDat[2])
	f.v = append(f.v, v)
	f.vuv = append(f.vuv, vuv)
	f.vn = append(f.vn, vn)

	return f
}

func ParseOBJ(location string) {
	dat, _ := ioutil.ReadFile(location)
	strDat := string(dat)
	re := regexp.MustCompile(`\r?\n`)
	strDat = re.ReplaceAllString(strDat, " ")
	strArray := strings.Split(strDat, " ")
	//fmt.Println("Strings: ", strDat)

	for i := 0; i < len(strArray); i++ {
		word := strArray[i]
		switch word {
		case "mtllib":
			m := parseMtllib(&i, strArray)
			fmt.Println("Parsed material: ", m)

		case "o":
			o := parseObjectName(&i, strArray)
			fmt.Println("Parsed object: ", o)

		case "v":
			v := parseVertex(&i, strArray)
			fmt.Println("Parsed Vertex: ", v)

		case "vn":
			vn := parseVertexNormal(&i, strArray)
			fmt.Println("Parsed Vertex normal: ", vn)

		case "usemtl":
			usemtl := parseUseMtl(&i, strArray)
			fmt.Println("Parsed usemtl: ", usemtl)

		case "f":
			f := parseFace(&i, strArray)
			fmt.Println("Parsed Face: ", f)

		}

	}

}
