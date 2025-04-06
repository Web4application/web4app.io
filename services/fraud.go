package services

import "math/rand"

func IsFraudulentTransaction(amount float64, ip string) bool {
    // Placeholder logic for fraud detection
    return rand.Float64() > 0.95 // Simulate fraud detection with randomness
}
