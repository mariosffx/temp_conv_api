package main

import "fmt"

func convert(value float64, from, to string) (float64, error) {
	// First convert to Celsius as a common base.
	var celsius float64
	switch from {
	case "celsius":
		celsius = value
	case "fahrenheit":
		celsius = (value - 32) * 5 / 9
	case "kelvin":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("unknown source unit: %s (use celsius, fahrenheit, or kelvin)", from)
	}

	// Then convert from Celsius to the target unit.
	switch to {
	case "celsius":
		return celsius, nil
	case "fahrenheit":
		return celsius*9/5 + 32, nil
	case "kelvin":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("unknown target unit: %s (use celsius, fahrenheit, or kelvin)", to)
	}
}
