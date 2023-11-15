package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context[T any] struct {
	r            *http.Request
	w            http.ResponseWriter
	RequestParam T
}

func handlerCreateUser[T any](c Context[T]) error {
	userParams := c.RequestParam
	return JSON(http.StatusOK, userParams)
}

func JSON(code int, v any) error {
	return nil
}

type Handler[T any] func(Context[T]) error

func makeHttpHandler[T any](h Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqata T

		if err := json.NewDecoder(r.Body).Decode(&reqata); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		errs := Validate(reqata)
		if len(errs) > 0 {
			panic(errs)
		}

		h(Context[T]{
			r:            r,
			w:            w,
			RequestParam: reqata,
		})
	}
}

func POST[T any](router string, h Handler[T]) {
	http.HandleFunc(router, makeHttpHandler(h))
}

type CreateUserParams struct {
	Email    string
	Password string
}
type Validater interface {
	Validate() []error
}

func (cp CreateUserParams) Validate() []error {
	return []error{fmt.Errorf("Bad Request")}
}

func main_HandlerTest() {
	POST("/user", handlerCreateUser[CreateUserParams])
}

func Validate(data any) []error {
	if v, ok := data.(Validater); ok {
		return v.Validate()
	}
	return nil
}
