package G3D

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

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

func parseVertexNormal(i *int, dat []string) vertexNormal {
	*i++
	vn := vertexNormal{}

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

	// NOTE: subtract one because we want index and obj files do not do index
	// first point of triangle
	//TODO replace Atoi with parseuint for bigger values
	v, _ := strconv.Atoi(fDat[0])
	uv, _ := strconv.Atoi(fDat[1])
	vn, _ := strconv.Atoi(fDat[2])
	f.V = append(f.V, v-1)
	f.UV = append(f.UV, uv-1)
	f.VN = append(f.VN, vn-1)

	*i++

	fDat = strings.Split(dat[*i], "/")

	// second point of triangle
	v, _ = strconv.Atoi(fDat[0])
	uv, _ = strconv.Atoi(fDat[1])
	vn, _ = strconv.Atoi(fDat[2])
	f.V = append(f.V, v-1)
	f.UV = append(f.UV, uv-1)
	f.VN = append(f.VN, vn-1)

	*i++

	fDat = strings.Split(dat[*i], "/")

	// third point of triangle
	v, _ = strconv.Atoi(fDat[0])
	uv, _ = strconv.Atoi(fDat[1])
	vn, _ = strconv.Atoi(fDat[2])
	f.V = append(f.V, v-1)
	f.UV = append(f.UV, uv-1)
	f.VN = append(f.VN, vn-1)

	return f
}

func parseMat(matLocation string) (*Material, error) {
	dat, _ := ioutil.ReadFile(matLocation)
	strDat := string(dat)
	re := regexp.MustCompile(`\r?\n`)
	strDat = re.ReplaceAllString(strDat, " ")
	strArray := strings.Split(strDat, " ")

	m := Material{}

	for i := 0; i < len(strArray); i++ {
		word := strArray[i]
		switch word {
		case "mtllib":

		}
	}

	return &m, nil
}

func ParseOBJ(objLocation, matLocation string) (*Mesh, error) {
	dat, _ := ioutil.ReadFile(objLocation)
	strDat := string(dat)
	re := regexp.MustCompile(`\r?\n`)
	strDat = re.ReplaceAllString(strDat, " ")
	strArray := strings.Split(strDat, " ")

	m := Mesh{}

	m.File = objLocation

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
			m.Vs = append(m.Vs, v)

		case "vn":
			vn := parseVertexNormal(&i, strArray)
			fmt.Println("Parsed Vertex normal: ", vn)
			m.VNs = append(m.VNs, vn)

		case "usemtl":
			usemtl := parseUseMtl(&i, strArray)
			fmt.Println("Parsed usemtl: ", usemtl)

		case "f":
			f := parseFace(&i, strArray)
			fmt.Println("Parsed Face: ", f)
			m.Faces = append(m.Faces, f)

		}

	}
	return &m, nil

}
