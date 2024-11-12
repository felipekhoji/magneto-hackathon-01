package datasource

// Poster defines an interface for storing or persisting data to a source.
type Poster interface {
	AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error
}
