package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate10DigitNumberRek() string {
	// rand.Seed(time.Now().UnixNano())
	number := rand.Int63n(1e10) // Angka maksimum 10^10 - 1

	// Konversi ke string
	noRekening := fmt.Sprintf("%010d", number)
	return noRekening
}

func GenerateRandomCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomCode := make([]byte, 5)
	for i := range randomCode {
		randomCode[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomCode)
}

func GenerateTransactionCode() string {
	// Dapatkan random code
	randomCode := GenerateRandomCode()

	// Dapatkan tanggal saat ini
	currentTime := time.Now()

	// Format tanggal menjadi "DDMMYY"
	formattedDate := currentTime.Format("020106")

	// Gabungkan format TRX - RANDOMCODE - DDMMYY
	transactionCode := fmt.Sprintf("TRX-%s-%s", randomCode, formattedDate)

	return transactionCode
}
