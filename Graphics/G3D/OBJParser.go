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
	x, y, z float32
}

type face struct {
	v, vn, vuv int
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

		}

	}

}
