package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
	"math"
)

func main() {
	var coords [][]int
	var coord []int
	var points []Point
	var pnt Point

	fh, err := os.Open("sample.txt")
	defer fh.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	scn := bufio.NewScanner(fh)

	for scn.Scan() {
		coord, pnt = getCoords(scn.Text())
		coords = append(coords, coord)
		points = append(points, pnt)
	}

	check(scn.Err())

	result := findConvexHull(points)

	for _, pt := range result {
		for i := 0; i < len(coords); i += 1 {
			if coords[i][0] == pt.X && coords[i][1] == pt.Y {
				coords = append(coords[:i], coords[i+1:]...)
				i = 0
				break
			}
		}
	}

	// set := make(map[string]float64)

// 	for _, c := range coords {
// 	// 	fmt.Println(c[0], c[1])
//  //    set[fmt.Sprintf("%s_%s", string(c[0]), string(c[1]))] = 0
// for _, pt := range points {
// 	fmt.Println(c, pt)
// 	fmt.Println(math.Abs(float64(c[0] - pt.X)) + math.Abs(float64(c[1] - pt.Y)))
// 	fmt.Println("===")
//  //      set[fmt.Sprintf("%s_%s", string(c[0]), string(c[1]))] += (math.Abs(float64(c[0] - pt.X)) + math.Abs(float64(c[1] - pt.Y)))
//  //      // if coords[i][0] == pt.X && coords[i][1] == pt.Y {
//  //      //  // manhattanDistance()
//  //      //  // coords = append(coords[:i], coords[i+1:]...)
//  //      //  i = 0
//  //      //  break
//  //      // }
//     }
//   }


	fmt.Println(points)
	fmt.Println(result)
	fmt.Println(coords)
	// fmt.Println(set)
	// fmt.Println(result)
}

func getCoords(str string) ([]int, Point) {
	var c []int
	pts := strings.Split(str, ", ")

	x, err := strconv.Atoi(pts[0])
	check(err)
	c = append(c, x)

	y, err := strconv.Atoi(pts[1])
	check(err)
	c = append(c, y)

	return c, Point{x, y}
}

func manhattanDistance(a, b []float64) float64 {
	s := float64(0)
	for i := 0; i < len(a); i += 1 {
		s += math.Abs(b[i] - a[i])
	}

	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ConvexHull library

// Point is a struct that holds the X Y cooridinates
// of a specific point in the Ecliden plane or space.
type Point struct {
	X, Y int
}

// Points is a slice built up of Point structs.
type Points []Point

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) Len() int {
	return len(points)
}

// lets sort our Points by x and, if equal, by y
func (points Points) Less(i, j int) bool {
	if points[i].X == points[j].X {
		return points[i].Y < points[j].Y
	}
	return points[i].X < points[j].X
}

// returns the modulo (and sign) of the cross product between vetors OA and OB
func crossProduct(O, A, B Point) int {
	return (A.X-O.X)*(B.Y-O.Y) - (A.Y-O.Y)*(B.X-O.X)
}

// findConvexHull returns a slice of Point with a convex hull
// it is counterclockwise and starts and ends at the same point
// i.e. the same point is repeated at the beginning and at the end
func findConvexHull(points Points) Points {
	n := len(points)  // number of points to find convex hull
	var result Points // final result
	count := 0        // size of our convex hull (number of points added)

	// lets sort our points by x and if equal by y
	sort.Sort(points)

	if n == 0 {
		return result
	}

	// add the first element:
	result = append(result, points[0])
	count++

	// find the lower hull
	for i := 1; i < n; i++ {
		// remove points which are not part of the lower hull
		for count > 1 && crossProduct(result[count-2], result[count-1], points[i]) < 0 {
			count--
			result = result[:count]
		}

		// add a new better point than the removed ones
		result = append(result, points[i])
		count++
	}

	count0 := count // our base counter for the upper hull

	// find the upper hull
	for i := n - 2; i >= 0; i-- {
		// remove points which are not part of the upper hull
		for count-count0 > 0 && crossProduct(result[count-2], result[count-1], points[i]) < 0 {
			count--
			result = result[:count]
		}

		// add a new better point than the removed ones
		result = append(result, points[i])
		count++
	}

	return result
}
