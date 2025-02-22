package util

const (
	USD = "USD"
	EUR = "EUR"
	UAH = "UAH"
)

func IsSupportedCurrence(currency string) bool {
	switch currency {
	case USD, EUR, UAH:
		return true
	}
	return false
}
