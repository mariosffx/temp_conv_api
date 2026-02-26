package main

// ConvertRequest represents the incoming conversion request.
type ConvertRequest struct {
	Value float64 `json:"value"`
	From  string  `json:"from"` // "celsius", "fahrenheit", "kelvin"
	To    string  `json:"to"`   // "celsius", "fahrenheit", "kelvin"
}

// ConvertResponse represents the conversion result.
type ConvertResponse struct {
	OriginalValue float64 `json:"original_value"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	Result        float64 `json:"result"`
}

// ErrorResponse represents an error message.
type ErrorResponse struct {
	Error string `json:"error"`
}
