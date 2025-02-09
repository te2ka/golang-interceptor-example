package interceptor

import (
	"net/http"
)

type Interceptor func(http.HandlerFunc) http.HandlerFunc
