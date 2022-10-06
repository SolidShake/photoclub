package app

import (
	"net/http"

	"github.com/SolidShake/photoclub/internal/handlers"
)

func Run() {
	http.ListenAndServe(":8080", handlers.InitHandlers())
}
