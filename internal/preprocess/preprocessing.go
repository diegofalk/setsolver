package preprocess

import (
	"gocv.io/x/gocv"
	"image"
)

const (
	cardThreshold = 90
	blurSize      = 5
	MorphOpenSize = 5
)

func Run(src gocv.Mat, dst *gocv.Mat) {
	gocv.CvtColor(src, dst, gocv.ColorBGRToGray)
	gocv.GaussianBlur(*dst, dst, image.Point{blurSize, blurSize}, 0, 0, gocv.BorderConstant)
	gocv.Threshold(*dst, dst, cardThreshold, 255, gocv.ThresholdBinary)
	gocv.MorphologyEx(*dst, dst, gocv.MorphOpen, gocv.GetStructuringElement(gocv.MorphRect, image.Point{MorphOpenSize, MorphOpenSize}))
}
