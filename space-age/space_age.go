package space

// Age takes the amount of seconds passed on a plannet and returns the
// amount of years that would be on that plannet
func Age(seconds float64, planet string) float64 {
	earthYear := 31557600.0
	multiplier := map[string]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / earthYear / multiplier[planet]
}
