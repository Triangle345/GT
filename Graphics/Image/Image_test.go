package Image

import (
	"fmt"
	"image"
	"testing"
)

func TestBuddyAggregator(t *testing.T) {

	ba := NewBuddyAggregator(100)
	part := ba.Insert("test1", 5, 5)

	if ok := part.Bounds().Eq(image.Rect(0, 0, 6, 6)); !ok {
		t.Error("Q1 test failed")
	}

	fmt.Println("Partition insert", part)

	part = ba.Insert("test2", 15, 15)

	if ok := part.Bounds().Eq(image.Rect(25, 0, 50, 25)); !ok {
		t.Error("Q2 test failed")
	}

	fmt.Println("Partition insert", part)

	part = ba.Insert("test3", 12, 12)

	if ok := part.Bounds().Eq(image.Rect(12, 0, 25, 12)); !ok {
		t.Error("Q2 again test failed")
	}

	fmt.Println("Partition insert", part)

	part = ba.Insert("test4", 13, 13)

	if ok := part.Bounds().Eq(image.Rect(12, 12, 25, 25)); !ok {
		t.Error("Q2 again test failed")
	}

	fmt.Println("Partition insert", part)

	part = ba.Insert("test5", 12, 12)

	if ok := part.Bounds().Eq(image.Rect(0, 12, 12, 25)); !ok {
		t.Error("Q2 again test failed")
	}

	fmt.Println("Partition insert", part)

	part = ba.Insert("test6", 20, 20)

	fmt.Println("Partition insert", part)

	part = ba.Insert("test7", 45, 20)
	if ok := part.Bounds().Eq(image.Rect(50, 0, 100, 50)); !ok {
		t.Error("Failed large bounds non uniform")
	}
	fmt.Println("Partition insert", part)

	part = ba.Insert("test8", 105, 45)
	if ok := part; ok != nil {
		t.Error("Failed to handle out of bounds")
	}
	fmt.Println("Partition insert", part)

}
