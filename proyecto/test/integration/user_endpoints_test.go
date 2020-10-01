package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	server "example.com/mittaus/ddd-example/infraestructure/gin.server"
)

func TestGetEntries(t *testing.T) {
	// authHandler := jwt.New(viper.GetString("jwt.Salt"))

	// routerLogger := logger.NewLogger("TEST",
	// 	viper.GetString("log.level"),
	// 	viper.GetString("log.format"),
	// )

	// api := application.HandlerConstructor{
	// 	Logger:           routerLogger,
	// 	UserRW:           userRW.New(),
	// 	ArticleRW:        articleRW.New(),
	// 	UserValidator:    validator.New(),
	// 	AuthHandler:      authHandler,
	// 	Slugger:          slugger.New(),
	// 	ArticleValidator: articleValidator.New(),
	// 	TagsRW:           tagsRW.New(),
	// 	CommentRW:        commentRW.New(),
	// }.New()

	var jsonStr = []byte(`{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.RouterHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
