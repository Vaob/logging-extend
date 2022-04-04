
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