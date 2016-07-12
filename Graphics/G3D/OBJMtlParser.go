package G3D

import (
	"GT/Logging"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

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
			// fmt.Println("Parsed material name: ", mat)

			m.File = matLocation
			m.Name = mat
			mats[mat] = m

		case "Ka":
			a := parseColor(&i, strArray)
			// fmt.Println("Parsed material ambient: ", a)
			m.Ambient = a
		case "Kd":
			d := parseColor(&i, strArray)
			// fmt.Println("Parsed material diffuse: ", d)
			m.Diffuse = d

		case "Ks":
			s := parseColor(&i, strArray)
			// fmt.Println("Parsed material specular: ", s)
			m.Specular = s
		case "Ke":
			e := parseColor(&i, strArray)
			// fmt.Println("Parsed material emission: ", e)
			m.Emission = e

		case "map_Kd":
			t := parseTexture(&i, strArray)
			// fmt.Println("Parsed material texture diffuse: ", t)
			m.DiffuseTex = t
		}

	}

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
