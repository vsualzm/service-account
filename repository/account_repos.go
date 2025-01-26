package repository

import (
	"database/sql"
	"log"
	"service-account/model"
)

type AccountRepository interface {
	CreateAccount(account *model.Account) (*model.Account, error)
	GetAccountByEmail(email string) (*model.Account, error)
	GetAccountByNoRekening(noRekening string) (*model.Account, error)
	UpdateSaldo(account *model.Account) (*model.Account, error)
	SaveTransaction(transaction *model.Transaction) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) SaveTransaction(transaction *model.Transaction) error {
	log.Println("=== Save Transaction ===")

	query := `INSERT INTO transaction (account_id, no_rekening_to, code_transaction, total_amount, status, remark) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	// Log query dan parameter
	log.Printf("Executing query: %s\n", query)
	log.Printf("Parameters: AccountId=%v, NoRekeningTo=%v, CodeTransaction=%v, TotalAmount=%v, Status=%v, Remark=%v\n",
		transaction.AccountId, transaction.NoRekeningTo, transaction.CodeTrasaction, transaction.Total_amount, transaction.Status, transaction.Remark)

	// Eksekusi query
	_, err := r.db.Exec(query,
		transaction.AccountId, transaction.NoRekeningTo, transaction.CodeTrasaction,
		transaction.Total_amount, transaction.Status, transaction.Remark)

	if err != nil {
		// Log error jika query gagal
		log.Printf("Error executing query: %v\n", err)
		return err
	}

	log.Println("Query executed successfully")
	return nil
}

func (r *accountRepository) CreateAccount(account *model.Account) (*model.Account, error) {

	query := `INSERT INTO account (nama, email, password_hash, nik, no_hp, roles, saldo, no_rekening) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int
	err := r.db.QueryRow(query, account.Nama, account.Email, account.PasswordHash, account.Nik, account.NoHp, account.Roles, account.Saldo, account.NoRekening).Scan(&id)
	if err != nil {
		return nil, err
	}
	account.Id = id
	return account, nil
}

func (r *accountRepository) GetAccountByEmail(email string) (*model.Account, error) {
	query := `SELECT id, nama, email, password_hash, nik, no_hp, roles, saldo, no_rekening FROM account WHERE email = $1`
	account := &model.Account{}
	err := r.db.QueryRow(query, email).Scan(
		&account.Id,
		&account.Nama,
		&account.Email,
		&account.PasswordHash,
		&account.Nik,
		&account.NoHp,
		&account.Roles,
		&account.Saldo,
		&account.NoRekening,
	)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return nil, err
	}

	return account, nil
}

func (r *accountRepository) GetAccountByNoRekening(noRekening string) (*model.Account, error) {
	// cheking query

	log.Println("=== Get Account By No Rekening ===", noRekening)
	query := `SELECT id, nama, email, nik, no_hp, roles, saldo, no_rekening FROM account WHERE no_rekening = $1`
	account := &model.Account{}
	err := r.db.QueryRow(query, noRekening).Scan(&account.Id, &account.Nama, &account.Email, &account.Nik, &account.NoHp, &account.Roles, &account.Saldo, &account.NoRekening)
	if err != nil {
		return nil, err
	}

	log.Println("=== Get Account By No Rekening ===", account.Saldo)

	return account, nil
}

func (r *accountRepository) UpdateSaldo(account *model.Account) (*model.Account, error) {
	log.Println("=== Update Saldo ===", account.Saldo)
	query := `UPDATE account SET saldo = $1 WHERE no_rekening = $2`
	_, err := r.db.Exec(query, account.Saldo, account.NoRekening)
	if err != nil {
		return nil, err
	}
	return account, nil
}
