package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/karchx/realword-nx/conduit"
	"github.com/karchx/realword-nx/mock"
)

func Test_createUser(t *testing.T) {
	userStore := &mock.UserService{}
	srv := testServer()
	srv.userService = userStore

	input := `{
    "user": {
      "email": "test@mail.com",
      "username": "test",
      "password": "password"
    }
  }`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(input))
	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()

	var user conduit.User
	userStore.CreateUserFn = func(u conduit.User) error {
		user = u
		return nil
	}

	srv.router.ServeHTTP(w, req)
	expectedResp := userResponse(&user)
	gotResp := M{}
	extractResponseBody(w.Body, &gotResp)

	if code := w.Code; code != http.StatusCreated {
		t.Errorf("expected status code of 201, but got %d ", code)
	}

	if !reflect.DeepEqual(expectedResp, gotResp) {
		t.Errorf("expected response %v, but got %v ", expectedResp, gotResp)
	}
}

func testServer() *Server {
	srv := &Server{
		router: mux.NewRouter(),
	}
	srv.routes()
	return srv
}

func extractResponseBody(body io.Reader, v interface{}) {
	mm := M{}
	_ = readJson[any](body, &mm)
	byt, err := json.Marshal(mm["user"])
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byt, err)
}
