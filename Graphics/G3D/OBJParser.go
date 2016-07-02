package G3D

import (
	"GT/Logging"
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

func parseVertexTexture(i *int, dat []string) vertexTexture {
	*i++
	vt := vertexTexture{}

	f, _ := strconv.ParseFloat(dat[*i], 32)

	vt.U = float32(f)

	*i++
	f, _ = strconv.ParseFloat(dat[*i], 32)
	vt.V = float32(f)

	return vt
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

func populateFaceTriData(dat string, f *face) {
	fDat := strings.Split(dat, "/")

	// NOTE: subtract one because we want index and obj files do not do index
	// first point of triangle
	//TODO replace Atoi with parseuint for bigger values
	if v, errV := strconv.Atoi(fDat[0]); errV == nil {
		f.V = append(f.V, v-1)
	}
	if uv, errUV := strconv.Atoi(fDat[1]); errUV == nil {
		f.UV = append(f.UV, uv-1)
	}

	if vn, errVN := strconv.Atoi(fDat[2]); errVN == nil {
		f.VN = append(f.VN, vn-1)
	}
}

func parseFace(i *int, dat []string) face {

	f := face{}

	// populate for three vertices - triangle

	*i++

	populateFaceTriData(dat[*i], &f)

	*i++
	populateFaceTriData(dat[*i], &f)

	*i++

	populateFaceTriData(dat[*i], &f)

	return f
}

func ParseOBJ(objLocation, matLocation string) (*Mesh, error) {
	dat, _ := ioutil.ReadFile(objLocation)
	strDat := string(dat)
	re := regexp.MustCompile(`\r?\n`)
	strDat = re.ReplaceAllString(strDat, " ")
	strArray := strings.Split(strDat, " ")

	var err error = nil

	m := Mesh{}

	mats, e := parseMat(matLocation)

	if e != nil {
		err = e
	}

	m.Materials = mats

	m.File = objLocation

	currentMaterial := ""

	for i := 0; i < len(strArray); i++ {
		word := strArray[i]
		switch word {
		case "mtllib":
			m := parseMtllib(&i, strArray)
			fmt.Println("Parsed object material: ", m)

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

		case "vt":
			vt := parseVertexTexture(&i, strArray)
			fmt.Println("Parsed Vertex texture: ", vt)
			m.VTs = append(m.VTs, vt)

		case "usemtl":
			usemtl := parseUseMtl(&i, strArray)
			fmt.Println("Parsed usemtl: ", usemtl)
			currentMaterial = usemtl

		case "f":
			f := parseFace(&i, strArray)
			fmt.Println("Parsed Face: ", f)
			f.Material = currentMaterial
			m.Faces = append(m.Faces, f)

		}

	}
	return &m, err

}

///////////////////////////////////////////////////////////////////////

func parseMat(matLocation string) (map[string]*Material, error) {
	dat, _ := ioutil.ReadFile(matLocation)
	strDat := string(dat)
	re := regexp.MustCompile(`\r?\n`)
	strDat = re.ReplaceAllString(strDat, " ")
	strArray := strings.Split(strDat, " ")

	mats := map[string]*Material{}

	var m *Material

	for i := 0; i < len(strArray); i++ {

		word := strArray[i]
		switch word {
		case "newmtl":
			m = &Material{}

			mat := parseMtl(&i, strArray)
			fmt.Println("Parsed material name: ", mat)

			m.File = matLocation
			m.Name = mat
			mats[mat] = m

		case "Ka":
			a := parseColor(&i, strArray)
			fmt.Println("Parsed material ambient: ", a)
			m.Ambient = a
		case "Kd":
			d := parseColor(&i, strArray)
			fmt.Println("Parsed material diffuse: ", d)
			m.Diffuse = d

		case "Ks":
			s := parseColor(&i, strArray)
			fmt.Println("Parsed material specular: ", s)
			m.Specular = s
		case "Ke":
			e := parseColor(&i, strArray)
			fmt.Println("Parsed material emission: ", e)
			m.Emission = e

		case "map_Kd":
			t := parseTexture(&i, strArray)
			fmt.Println("Parsed material texture diffuse: ", t)
			m.DiffuseTex = t
		}

	}

	fmt.Println("Materisls 001 :", mats["Material.001"])
	fmt.Println("Materisls 002 :", mats["Material.002"])
	return mats, nil
}

func parseTexture(i *int, dat []string) string {
	*i++
	return dat[*i]
}

func parseMtl(i *int, dat []string) string {
	*i++
	return dat[*i]
}

func parseColor(i *int, dat []string) Color {

	errString := "Failed to convert Color data: " + dat[*i]
	*i++

	var r, g, b float64
	var err error

	if r, err = strconv.ParseFloat(dat[*i], 32); err != nil {
		Logging.Info(errString, err)
	}

	*i++

	if g, err = strconv.ParseFloat(dat[*i], 32); err != nil {
		Logging.Info(errString, err)
	}

	*i++

	if b, err = strconv.ParseFloat(dat[*i], 32); err != nil {
		Logging.Info(errString, err)
	}

	return Color{R: float32(r), G: float32(g), B: float32(b)}

}
