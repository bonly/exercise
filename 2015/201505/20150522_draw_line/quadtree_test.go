package quadtree

import (
	"math/rand"
	"testing"
	"time"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"golang.org/x/image/colornames"
)

func TestQuadtreeCreation(t *testing.T) {
	//x, y, width, height
	qt := setupQuadtree(0, 0, 640, 480)
	if qt.Bounds.Width != 640 && qt.Bounds.Height != 480 {
		t.Errorf("Quadtree was not created correctly")
	}
}

func TestSplit(t *testing.T) {

	//x, y, width, height
	qt := setupQuadtree(0, 0, 640, 480)
	qt.split()
	if len(qt.Nodes) != 4 {
		t.Error("Quadtree did not split correctly, expected 4 nodes got", len(qt.Nodes))
	}

	qt.split()
	if len(qt.Nodes) != 4 {
		t.Error("Quadtree should not split itself more than once", len(qt.Nodes))
	}

}

func TestTotalSubnodes(t *testing.T) {

	//x, y, width, height
	qt := setupQuadtree(0, 0, 640, 480)
	qt.split()
	for i := 0; i < len(qt.Nodes); i++ {
		qt.Nodes[i].split()
	}

	total := qt.TotalNodes()
	if total != 20 {
		t.Error("Quadtree did not split correctly, expected 20 nodes got", total)
	}

}

func TestQuadtreeInsert(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for i := 0; i < numObjects; i++ {

		x := randMinMax(0, gridh) * grid
		y := randMinMax(0, gridv) * grid

		randomObject = Bounds{
			X:      x,
			Y:      y,
			Width:  randMinMax(1, 4) * grid,
			Height: randMinMax(1, 4) * grid,
		}

		index := qt.getIndex(randomObject)
		if index < -1 || index > 3 {
			t.Errorf("The index should be -1 or between 0 and 3, got %d \n", index)
		}

		qt.Insert(randomObject)

	}

	if qt.Total != numObjects {
		t.Errorf("Error: Should have totalled %d, got %d \n", numObjects, qt.Total)
	} else {
		t.Logf("Success: Total objects in the Quadtree is %d (as expected) \n", qt.Total)
	}

}

func TestCorrectQuad(t *testing.T) {

	qt := setupQuadtree(0, 0, 100, 100)

	var index int
	pass := true

	topRight := Bounds{
		X:      99,
		Y:      99,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topRight)
	index = qt.getIndex(topRight)
	if index == 0 {
		t.Errorf("The index should be 0, got %d \n", index)
		pass = false
	}

	topLeft := Bounds{
		X:      99,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topLeft)
	index = qt.getIndex(topLeft)
	if index == 1 {
		t.Errorf("The index should be 1, got %d \n", index)
		pass = false
	}

	bottomLeft := Bounds{
		X:      1,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomLeft)
	index = qt.getIndex(bottomLeft)
	if index == 2 {
		t.Errorf("The index should be 2, got %d \n", index)
		pass = false
	}

	bottomRight := Bounds{
		X:      1,
		Y:      51,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomRight)
	index = qt.getIndex(bottomRight)
	if index == 3 {
		t.Errorf("The index should be 3, got %d \n", index)
		pass = false
	}

	if pass == true {
		t.Log("Success: The points were inserted into the correct quadrants")
	}

}

func TestQuadtreeRetrieval(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	var randomObject Bounds
	numObjects := 100

	for i := 0; i < numObjects; i++ {

		randomObject = Bounds{
			X:      float64(i),
			Y:      float64(i),
			Width:  0,
			Height: 0,
		}

		qt.Insert(randomObject)

	}

	for j := 0; j < numObjects; j++ {

		Cursor := Bounds{
			X:      float64(j),
			Y:      float64(j),
			Width:  0,
			Height: 0,
		}

		objects := qt.Retrieve(Cursor)

		found := false

		if len(objects) >= numObjects {
			t.Error("Objects should not be equal to or bigger than the number of retrieved objects")
		}

		for o := 0; o < len(objects); o++ {
			if objects[o].X == float64(j) && objects[o].Y == float64(j) {
				found = true
			}
		}
		if found != true {
			t.Error("Error finding the correct point")
		}

	}

}

func TestQuadtreeRandomPointRetrieval(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	numObjects := 1000

	for i := 1; i < numObjects+1; i++ {

		randomObject := Bounds{
			X:      float64(i),
			Y:      float64(i),
			Width:  0,
			Height: 0,
		}

		qt.Insert(randomObject)

	}

	failure := false
	iterations := 20
	for j := 1; j < iterations+1; j++ {

		Cursor := Bounds{
			X:      float64(j),
			Y:      float64(j),
			Width:  0,
			Height: 0,
		}

		point := qt.RetrievePoints(Cursor)

		for k := 0; k < len(point); k++ {
			if point[k].X == 0 {
				failure = true
			}
			if point[k].Y == 0 {
				failure = true
			}
			if failure {
				t.Error("Point was incorrectly retrieved", point)
			}
			if point[k].IsPoint() == false {
				t.Error("Point should have width and height of 0")
			}
		}

	}

	if failure == false {
		t.Logf("Success: All the points were retrieved correctly", iterations, numObjects)
	}

}

func TestIntersectionRetrieval(t *testing.T) {
	qt := setupQuadtree(0, 0, 640, 480)
	qt.Insert(Bounds{
		X:      1,
		Y:      1,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      5,
		Y:      5,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      10,
		Y:      10,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      15,
		Y:      15,
		Width:  10,
		Height: 10,
	})
	inter := qt.RetrieveIntersections(Bounds{
		X:      5,
		Y:      5,
		Width:  2.5,
		Height: 2.5,
	})
	if len(inter) != 2 {
		t.Error("Should have two intersections")
	}
}

func TestQuadtreeClear(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for i := 0; i < numObjects; i++ {

		x := randMinMax(0, gridh) * grid
		y := randMinMax(0, gridv) * grid

		randomObject = Bounds{
			X:      x,
			Y:      y,
			Width:  randMinMax(1, 4) * grid,
			Height: randMinMax(1, 4) * grid,
		}

		index := qt.getIndex(randomObject)
		if index < -1 || index > 3 {
			t.Errorf("The index should be -1 or between 0 and 3, got %d \n", index)
		}

		qt.Insert(randomObject)

	}

	qt.Clear()

	if qt.Total != 0 {
		t.Errorf("Error: The Quadtree should be cleared")
	} else {
		t.Logf("Success: The Quadtree was cleared correctly")
	}

}

// Benchmarks

func BenchmarkInsertOneThousand(b *testing.B) {

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject Bounds
	numObjects := 1000

	for n := 0; n < b.N; n++ {
		for i := 0; i < numObjects; i++ {

			x := randMinMax(0, gridh) * grid
			y := randMinMax(0, gridv) * grid

			randomObject = Bounds{
				X:      x,
				Y:      y,
				Width:  randMinMax(1, 4) * grid,
				Height: randMinMax(1, 4) * grid,
			}

			qt.Insert(randomObject)

		}
	}

}

// Convenience Functions

func setupQuadtree(x float64, y float64, width float64, height float64) *Quadtree {

	return &Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		MaxObjects: 4,
		MaxLevels:  8,
		Level:      0,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0),
	}

}

func randMinMax(min float64, max float64) float64 {
	val := min + (rand.Float64() * (max - min))
	return val
}

func Test_Bonly1(ts *testing.T){
	qt := setupQuadtree(0, 0, 2, 2);
	obj := Bounds{1, 1, 1, 1};
	qt.Insert(obj);
	obj = Bounds{1, 1, 0, 0};
	// qt.Insert(obj);

	bd := qt.RetrieveIntersections(Bounds{0, 0, 1, 1});
	if len(bd) >=1 {
		ts.Logf("贴边碰撞！%#v\n", bd);
	}

	bd = qt.RetrieveIntersections(Bounds{0, 0, 2, 1});
	if len(bd) >=1 {
		ts.Logf("x相交碰撞！%#v\n", bd);
	}	

	bd = qt.RetrieveIntersections(Bounds{0.9, 0, 0.1, 1.1});
	if len(bd) >=1 {
		ts.Logf("y相交碰撞！%#v\n", bd);
	}		

	bd = qt.RetrieveIntersections(Bounds{0.1, .1, 0.8, 5});
	if len(bd) < 1 {
		ts.Logf("x不够不会碰撞！%#v\n", bd);
	}	

	bd = qt.RetrieveIntersections(Bounds{2, .2, 1, .4});
	if len(bd) <1 {
		ts.Logf("y不够不会碰撞！%#v\n", bd);
	}	

	bd = qt.RetrievePoints(Bounds{1, 1, 1, 1});
	if len(bd) >= 1{
		ts.Logf("找到点！%#v\n", bd);
	}

	// bd = qt.Retrieve(Bounds{1, 1, 1, 1});
	// if len(bd) >= 1{
	// 	ts.Logf("找到物件！%#v\n", bd);
	// }
}

func Test_Bonly2(ts *testing.T){
	const width = 500;
	const height = 500;

	im := image.NewRGBA(image.Rect(0, 0, width, height));
    蓝 := color.RGBA{0, 0, 255, 255};
    白 := color.RGBA{255, 255, 255, 255};
	draw.Draw(im, im.Bounds(), &image.Uniform{白}, image.ZP, draw.Src)
	im.Set(50, 50, 蓝);
	
	for idx := 0; idx < width; idx+=10{
		drawline(idx, 0, idx, height, func(x, y int){
			im.Set(x, y, colornames.Gray);
		});
	}

	for idx := 0; idx < width; idx+=10{
		drawline(0, idx, width, idx, func(x, y int){
			im.Set(x, y, colornames.Gray);
		});
	}	

	w, _ := os.Create("pic.png")
	defer w.Close();
	png.Encode(w, im);
}

// Putpixel describes a function expected to draw a point on a bitmap at (x, y) coordinates.
type Putpixel func(x, y int)
// 求绝对值
func abs(x int) int {
    if x >= 0 {
    return x
    }
    return -x
}
// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// https://github.com/akavel/polyclip-go/blob/9b07bdd6e0a784f7e5d9321bff03425ab3a98beb/polyutil/draw.go
// TODO: handle int overflow etc.
func drawline(x0, y0, x1, y1 int, brush Putpixel) {
    dx := abs(x1 - x0)
    dy := abs(y1 - y0)
    sx, sy := 1, 1
    if x0 >= x1 {
    sx = -1
    }
    if y0 >= y1 {
    sy = -1
    }
    err := dx - dy
    for {
    brush(x0, y0)
    if x0 == x1 && y0 == y1 {
        return
    }
    e2 := err * 2
    if e2 > -dy {
        err -= dy
        x0 += sx
    }
    if e2 < dx {
        err += dx
        y0 += sy
    }
    }
}