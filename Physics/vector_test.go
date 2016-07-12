package Physics

import "testing"

func TestVector3(t *testing.T) {

	v1 := NewVector3(10, 20, 30)
	v2 := NewVector3(5, 5, 5)

	if v1.X != 10 && v1.Y != 20 && v1.Z != 30 {
		t.Error("v1 is not valid: ", v1)
	}

	res := v1.Add(v2)

	if res.X != 15 && res.Y != 25 && res.Z != 35 {
		t.Error("res is not valid: ", res)
	}

	res = v1.Sub(v2)

	if res.X != 5 && res.Y != 15 && res.Z != 25 {
		t.Error("res is not valid: ", res)
	}

	res = v1.Lerp(.5)

	if res.X != 5 && res.Y != 10 && res.Z != 15 {
		t.Error("res is not valid: ", res)
	}

	res2 := v1.Dot(v2)

	if res2 != float64(300.00) {
		t.Error("Dot product failed:", res2)
	}

	res = v1.Cross(v2)
	if res.X != -50 && res.Y != 100 && res.Z != -50 {
		t.Error("res is not valid: ", res)
	}

	res2 = v1.Mag()
	if res2 != float64(37.416573867739416) {
		t.Error("Dot product failed:", res2)
	}

	res = v1.Normalize()
	if res.X != 0.26726124191242434 && res.Y != 0.5345224838248487 && res.Z != -0.5345224838248487 {
		t.Error("res is not valid: ", res)
	}
}
