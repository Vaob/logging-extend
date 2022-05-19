
package traderutils

import (
	"exmo-trading/app/data"
	"math"
	"time"
)

func GetMA(array []float64, frame int) []float64 {
	length := len(array)
	ma := make([]float64, length)
	for i := frame; i < length; i++ {
		sum := 0.00
		for j := i - frame; j < i; j++ {
			sum = sum + array[j]
		}
		ma[i] = sum / float64(frame)
	}
	return ma
}

func GetSD(priceArray, ma []float64, frame int) []float64 {
	length := len(priceArray)
	sd := make([]float64, length)
	for i := frame; i < length; i++ {
		sum := 0.00
		for j := i - frame; j < i; j++ {
			sum = sum + (priceArray[j]-ma[i])*(priceArray[j]-ma[i])
		}
		sd[i] = math.Sqrt(sum / float64(frame))
	}
	return sd