package common

import "github.com/shopspring/decimal"

func Contains(str string, strs []string) bool {
	m := make(map[string]bool)
	for _, v := range strs {
		m[v] = true
	}
	if m[str] {
		return true
	}
	return false
}

func StringToDecimal(str string) decimal.Decimal {
	if str == "" {
		return decimal.Zero
	}
	dec, err := decimal.NewFromString(str)
	if err != nil {
		return decimal.Zero
	}
	return dec
}
