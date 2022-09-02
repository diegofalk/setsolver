package alligner

import (
	"gocv.io/x/gocv"
	"image"
	"setsolver/models"
)

const (
	cardWidth  = 200
	cardHeight = 300
)

func Run(src gocv.Mat, contour models.CardContour) gocv.Mat {
	srcPoints := orderPoints(contour.ApproxPoly)
	rect := gocv.BoundingRect(contour.ApproxPoly)

	dstPoints := gocv.NewPointVector()
	dstPoints.Append(image.Point{0, 0})
	if rect.Dy() > rect.Dx() {
		dstPoints.Append(image.Point{0, cardHeight})
		dstPoints.Append(image.Point{cardWidth, cardHeight})
		dstPoints.Append(image.Point{cardWidth, 0})
	} else {
		dstPoints.Append(image.Point{cardWidth, 0})
		dstPoints.Append(image.Point{cardWidth, cardHeight})
		dstPoints.Append(image.Point{0, cardHeight})
	}

	t := gocv.GetPerspectiveTransform(srcPoints, dstPoints)
	r := gocv.NewMat()
	gocv.WarpPerspective(src, &r, t, image.Point{cardWidth, cardHeight})
	return r
}

func orderPoints(pts gocv.PointVector) gocv.PointVector {
	sums := []int{}
	diffs := []int{}

	for i := 0; i < pts.Size(); i++ {
		sums = append(sums, pts.At(i).X+pts.At(i).Y)
		diffs = append(diffs, pts.At(i).X-pts.At(i).Y)
	}

	p2, p0 := maxAndMinIdx(sums)
	p3, p1 := maxAndMinIdx(diffs)

	return gocv.NewPointVectorFromPoints([]image.Point{
		pts.At(p0),
		pts.At(p1),
		pts.At(p2),
		pts.At(p3),
	})
}

func maxAndMinIdx(vals []int) (int, int) {
	max := vals[0]
	min := vals[0]
	maxIdx := 0
	minIdx := 0
	for i, s := range vals {
		if s > max {
			max = s
			maxIdx = i
		}
		if s < min {
			min = s
			minIdx = i
		}
	}
	return maxIdx, minIdx
}
