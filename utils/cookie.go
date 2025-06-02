package utils

import (
	"fmt"
	"net/http"
)

func SetCookie(writter http.ResponseWriter, request *http.Request) {
	Cookie := new(http.Cookie)
	Cookie.Name = "Todo Cookie"
	Cookie.Value = request.URL.Query().Get("VALUE")
	Cookie.Path = "/"

	http.SetCookie(writter, Cookie)
	fmt.Fprint(writter, "Cookie has been set successfully")
}
