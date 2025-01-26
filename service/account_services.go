package service

import (
	"errors"
	"log"
	"service-account/model"
	"service-account/repository"
	"service-account/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AccountService interface {
	DaftarAccount(accountReq *model.AccountReq) (*model.Account, error)
	LoginAccount(account *model.AccountReq) (string, error)
	Tabung(accountReq *model.TransactionReq, userId int) (*model.Account, error)
	Tarik(accountReq *model.TransactionReq, userId int) (*model.Account, error)
	CheckSaldo(account string, userId int) (*model.Account, error)
}

type accountService struct {
	repo      repository.AccountRepository
	jwtSecret string
}

func NewAccountService(repo repository.AccountRepository, jwtSecret string) AccountService {
	return &accountService{repo: repo, jwtSecret: jwtSecret}
}

func (s *accountService) CheckSaldo(account string, userId int) (*model.Account, error) {

	accountTransaction, err := s.repo.GetAccountByNoRekening(account)
	if err != nil {
		return nil, errors.New("failed get saldo account")
	}

	// validate user id
	if accountTransaction.Id != userId {
		return nil, errors.New("invalid user id")
	}

	return accountTransaction, nil
}

func (s *accountService) Tarik(accountReq *model.TransactionReq, userId int) (*model.Account, error) {

	accountTransaction, err := s.repo.GetAccountByNoRekening(accountReq.NoRekening)

	if err != nil {
		return nil, errors.New("failed get saldo account")
	}

	// validate user id
	if accountTransaction.Id != userId {
		return nil, errors.New("invalid user id")
	}

	if accountReq.Saldo > accountTransaction.Saldo {
		return nil, errors.New("insufficient balance")
	}

	log.Println("SALDO AWAL: ", accountTransaction.Saldo)
	log.Println("SALDO YANG DITARIK: ", accountReq.Saldo)

	// calculate amount
	accountTransaction.Saldo = accountTransaction.Saldo - accountReq.Saldo

	log.Println("SALDO AKHIR: ", accountReq.Saldo)

	accountTransaction, err = s.repo.UpdateSaldo(accountTransaction)
	if err != nil {
		return nil, errors.New("failed update saldo account")
	}

	transactionReq := &model.Transaction{
		AccountId:      accountTransaction.Id,
		NoRekeningTo:   accountReq.NoRekening,
		CodeTrasaction: utils.GenerateTransactionCode(),
		Total_amount:   accountReq.Saldo,
		Status:         "SUCCESS",
		Remark:         "Tarik Saldo",
	}

	// save transaction
	err = s.repo.SaveTransaction(transactionReq)
	if err != nil {
		return nil, errors.New("failed save transaction")
	}

	return accountTransaction, nil
}

func (s *accountService) Tabung(accountReq *model.TransactionReq, userId int) (*model.Account, error) {

	accountTransaction, err := s.repo.GetAccountByNoRekening(accountReq.NoRekening)

	if err != nil {
		return nil, errors.New("kesalahan terkait data yg dikirim.")
	}

	// validate user id
	if accountTransaction.Id != userId {
		return nil, errors.New("invalid user id")
	}

	// calculate amount

	log.Println("Saldo Awal: ", accountTransaction.Saldo)
	log.Println("Saldo Yang Ditabung: ", accountReq.Saldo)
	accountTransaction.Saldo = accountTransaction.Saldo + accountReq.Saldo

	log.Println("Saldo Akhir: ", accountTransaction.Saldo)

	accountTransaction, err = s.repo.UpdateSaldo(accountTransaction)
	if err != nil {
		return nil, errors.New("failed update saldo account")
	}

	generateCodes := utils.GenerateTransactionCode()
	log.Println("CODES:", generateCodes)

	transactionReq := &model.Transaction{
		AccountId:      accountTransaction.Id,
		NoRekeningTo:   accountReq.NoRekening,
		CodeTrasaction: generateCodes,
		Total_amount:   accountReq.Saldo,
		Status:         "SUCCESS",
		Remark:         "Tabung Saldo",
	}

	log.Println("masuk -------> 2", transactionReq)

	// save transaction
	err = s.repo.SaveTransaction(transactionReq)
	if err != nil {
		return nil, errors.New("failed save transaction")
	}

	return accountTransaction, nil
}

func (s *accountService) DaftarAccount(accountReq *model.AccountReq) (*model.Account, error) {

	account := &model.Account{}

	account.Nama = accountReq.Nama
	account.Email = accountReq.Email
	account.Nik = accountReq.Nik
	account.NoHp = accountReq.NoHp
	account.Roles = "ADMIN"
	account.Saldo = 0
	account.NoRekening = utils.Generate10DigitNumberRek()

	passWordHash, err := bcrypt.GenerateFromPassword([]byte(accountReq.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	account.PasswordHash = string(passWordHash)

	newAccount, err := s.repo.CreateAccount(account)

	if err != nil {
		return nil, err
	}
	return newAccount, nil

}

func (s *accountService) LoginAccount(accountReq *model.AccountReq) (string, error) {

	account, err := s.repo.GetAccountByEmail(accountReq.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(accountReq.PasswordHash))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    account.Id,
		"roles": account.Roles,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))

}
