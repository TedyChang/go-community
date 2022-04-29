package service

import (
	"net/http"
)

func UserPage(w http.ResponseWriter, req *http.Request) {
	err := Render.JSON(w, http.StatusOK, map[string]string{"page": "user"})
	if err != nil {
		_ = Render.JSON(w, http.StatusBadGateway, map[string]string{"code": "fail"})
	}
}
