package strategies

type Strategy interface {
	Analyze() (string, error) //strategies implement this interf