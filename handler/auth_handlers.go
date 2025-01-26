package handler

import (
	"log"
	"net/http"
	"service-account/model"
	"service-account/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{service}
}

func (h *AccountHandler) TestHealtAPI(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func (h *AccountHandler) CheckSaldo(c echo.Context) error {

	log.Println("HIT API CHECKSALDO")

	// Ambil klaim user dari context
	userClaims := c.Get("user").(jwt.MapClaims)
	userId := int(userClaims["id"].(float64)) // Konversi dari float64 ke int
	log.Println(userId, "chek id")

	noRekening := c.Param("norekening")
	account, err := h.service.CheckSaldo(noRekening, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "kesalahan terkait data yg dikirim"})
	}

	response := map[string]interface{}{
		"message": "Success",
		"Saldo":   account.Saldo,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) Tarik(c echo.Context) error {
	log.Println("HIT API TARIK")

	// Ambil klaim user dari context
	userClaims := c.Get("user").(jwt.MapClaims)
	userId := int(userClaims["id"].(float64)) // Konversi dari float64 ke int
	log.Println(userId, "chek id")

	var accountReq model.TransactionReq
	if err := c.Bind(&accountReq); err != nil {
		return err
	}

	// validate tarik
	if accountReq.Saldo <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Amount must be greater than 0"})
	}

	account, err := h.service.Tarik(&accountReq, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "kesalahan terkait data yg dikirim"})
	}

	response := map[string]interface{}{
		"message": "Success Tarik",
		"Saldo":   account.Saldo,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) Tabung(c echo.Context) error {
	log.Println("HIT API TABUNG")

	// Ambil klaim user dari context
	userClaims := c.Get("user").(jwt.MapClaims)
	userId := int(userClaims["id"].(float64)) // Konversi dari float64 ke int
	log.Println(userId, "chek id")

	var accountReq model.TransactionReq
	if err := c.Bind(&accountReq); err != nil {
		return err
	}

	account, err := h.service.Tabung(&accountReq, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "kesalahan terkait data yg dikirim"})
	}

	log.Println("checking", account.Saldo)

	response := map[string]interface{}{
		"message": "Success Tabung",
		"Saldo":   account.Saldo,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) DaftarAccount(c echo.Context) error {

	log.Println("HIT API DAFTAR ACCOUNT")
	var accountReq model.AccountReq
	if err := c.Bind(&accountReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "kesalahan terkait data yg dikirim"})
	}

	account, err := h.service.DaftarAccount(&accountReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "kesalahan terkait data yg dikirim"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Account registered successfully", "noRekening": account.NoRekening})
}

func (h *AccountHandler) LoginAccount(c echo.Context) error {

	var loginReq model.LoginReq
	err := c.Bind(&loginReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email and password are required"})
	}

	accountReq := model.AccountReq{
		Email:        loginReq.Email,
		PasswordHash: loginReq.Password,
	}

	token, err := h.service.LoginAccount(&accountReq)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success login", "token": token})

}
