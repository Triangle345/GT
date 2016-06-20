package Image

import "image"

type quadrant int

const ( // iota is reset to 0
	tlQ quadrant = iota //1
	trQ          = iota //2
	blQ          = iota //3
	brQ          = iota //4
)

// var count int = 0

type partition struct {
	leaf, filled   bool
	tl, tr, bl, br *partition
	bounds         image.Rectangle
	tag            string
	quad           quadrant
}

func (this *partition) topLeft() image.Rectangle {
	//curMin, curMin, curMax/2, curMax/2
	return image.Rect(this.bounds.Min.X,
		this.bounds.Min.Y,
		this.bounds.Min.X+this.bounds.Dx()/2,
		this.bounds.Min.Y+this.bounds.Dy()/2)
}

func (this *partition) topRight() image.Rectangle {
	//curMax/2, curMin, curMax, curMax/2
	return image.Rect(this.bounds.Min.X+this.bounds.Dx()/2,
		this.bounds.Min.Y,
		this.bounds.Max.X,
		this.bounds.Min.Y+this.bounds.Dy()/2)
}

func (this *partition) bottomLeft() image.Rectangle {
	//curMin, curMax/2, curMax/2, curMax
	return image.Rect(this.bounds.Min.X,
		this.bounds.Min.Y+this.bounds.Dy()/2,
		this.bounds.Min.X+this.bounds.Dx()/2,
		this.bounds.Max.Y)
}

func (this *partition) bottomRight() image.Rectangle {
	//curMax/2, curMax/2, curMax, curMax
	return image.Rect(this.bounds.Min.X+this.bounds.Dx()/2,
		this.bounds.Min.Y+this.bounds.Dy()/2,
		this.bounds.Max.X,
		this.bounds.Max.Y)
}

// Insert - Inserts tag into an open top most left most partition, returns that partition as reference
func (this *partition) Insert(tag string, width, height int) *partition {

	// count += 1

	// if count > 30 {
	// 	os.Exit(-1)
	// }
	// fmt.Println("Attempting insert in :", this.Bounds())

	if (width > this.bounds.Dx() || height > this.bounds.Dy()) || this.filled {
		// fmt.Println("W: ", width, "H:", height, " Doesnt FIT in bounds: ", this)
		return nil
	}

	// until any children actually get filled, we are still a leaf
	if this.leaf {

		tl, tr, bl, br := this.split()
		this.tl = tl
		this.tr = tr
		this.bl = bl
		this.br = br
	}

	// now attempt to insert
	// if any one of these get filled, we are not a leaf any more
	if p := this.tl.Insert(tag, width, height); p != nil {
		this.leaf = false
		return p
	}

	if p := this.tr.Insert(tag, width, height); p != nil {
		this.leaf = false
		return p
	}
	if p := this.bl.Insert(tag, width, height); p != nil {
		this.leaf = false
		return p
	}
	if p := this.br.Insert(tag, width, height); p != nil {
		this.leaf = false
		return p
	}

	// however; if nothing is filled it means we are an acceptable choice, and we are still a leaf
	if this.leaf {

		this.filled = true
		this.tag = tag
		return this
	}

	return nil

}

func (this *partition) split() (tl, tr, bl, br *partition) {
	tl = &partition{
		leaf:   true,
		filled: false,
		quad:   tlQ,
		bounds: this.topLeft()}

	tr = &partition{
		leaf:   true,
		filled: false,
		quad:   trQ,
		bounds: this.topRight()}

	bl = &partition{
		leaf:   true,
		filled: false,
		quad:   blQ,
		bounds: this.bottomLeft()}

	br = &partition{
		leaf:   true,
		filled: false,
		quad:   brQ,
		bounds: this.bottomRight()}

	return tl, tr, bl, br
}

// Bounds - returns bounds of partition
func (this *partition) Bounds() image.Rectangle {
	return this.bounds
}

func newpartition(bounds image.Rectangle) *partition {

	//fmt.Println("Creating partition from: Min", curMin, "Max: ", curMax)

	part := partition{leaf: true,
		filled: false,
		bounds: bounds}

	tl, tr, bl, br := part.split()

	part.tl = tl
	part.tr = tr
	part.bl = bl
	part.br = br

	return &part
}

// NewBuddyAggregator creates a new partition which recursively finds free space in any aggregate image
func NewBuddyAggregator(maxBounds int) *partition {

	return newpartition(image.Rect(0, 0, maxBounds, maxBounds))
}
