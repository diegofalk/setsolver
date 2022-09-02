package cardfinder

import (
	"gocv.io/x/gocv"
	"setsolver/models"
)

const approxEpsilonPercent = 0.012
const minPerimeter = 1000
const minArea = 100000

func Run(src gocv.Mat) (cardContours []models.CardContour) {
	contours := gocv.FindContours(src, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	for i := 0; i < contours.Size(); i++ {
		var ct models.CardContour
		ct.Contour = contours.At(i)
		ct.Perimeter = gocv.ArcLength(ct.Contour, true)
		ct.Area = gocv.ContourArea(ct.Contour)
		ct.ApproxPoly = gocv.ApproxPolyDP(ct.Contour, approxEpsilonPercent*ct.Perimeter, true)
		if ct.ApproxPoly.Size() == 4 && ct.Perimeter > minPerimeter && ct.Area > minArea {
			cardContours = append(cardContours, ct)
		}
	}
	return
}
