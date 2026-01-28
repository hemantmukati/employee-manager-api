package helpers

func CalculateSalary(country string, gross int64) (deduction int64, net int64) {
	var rate float64

	switch country {
	case "India":
		rate = 0.10
	case "United States":
		rate = 0.12
	default:
		rate = 0
	}

	deduction = int64(float64(gross) * rate)
	net = gross - deduction
	return
}
