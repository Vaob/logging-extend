
package query

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type User struct {
	Id        string
	PublicKey string
	SecretKey string
}

type Query interface {
	Do() (*http.Response, error)
}

type PostQuery struct {
	UserParams     User
	Sign           string
	Method         string
	Params         map[string]string
	PreparedParams string
}

type GetQuery struct {
	Method string
}

func (q *PostQuery) PrepareParams() {
	postParams := url.Values{}
	if q.Params != nil {
		for key, value := range q.Params {
			postParams.Add(key, value)
		}
	}