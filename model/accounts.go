package model

import "time"

type Account struct {
	Id           int
	Nama         string
	Email        string
	PasswordHash string
	Roles        string
	Nik          string
	NoHp         string
	Saldo        float64
	NoRekening   string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

type AccountReq struct {
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Nik          string `json:"nik"`
	NoHp         string `json:"no_hp"`
	// Roles        string `json:"roles"`
	// Saldo        float64 `json:"saldo"`
	// NoRekening   string  `json:"no_rekening"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TransactionReq struct {
	NoRekening string  `json:"no_rekening"`
	Saldo      float64 `json:"saldo"`
}

type Transaction struct {
	Id             int
	AccountId      int
	NoRekeningTo   string
	CodeTrasaction string
	Total_amount   float64
	Status         string
	Remark         string
	CreatedAt      time.Time
	UpdateAt       time.Time
}

type TransactionUpdate struct {
	Id             int     `json:"id"`
	AccountId      int     `json:"account_id"`
	NoRekeningTo   string  `json:"no_rekening_to"`
	CodeTrasaction string  `json:"code_transaction"`
	Total_amount   float64 `json:"total_amount"`
	Status         string  `json:"status"`
	Remark         string  `json:"remark"`
}
