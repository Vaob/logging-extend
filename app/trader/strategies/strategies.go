package strategies

type Strategy interface {
	Analyze() (string, error) //strateg