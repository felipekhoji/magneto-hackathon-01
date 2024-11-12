package datasource

// Getter defines an interface for retrieving data from a source.
type Getter interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
}
