package strategies

type Strategy interface {
	Analyze() (string, error) //strategies implement this interface
	Set(candlesFile string, candlesFileVolu