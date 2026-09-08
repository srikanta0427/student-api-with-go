package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/srikanta0427/student-api/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var st types.Student

		err:= json.NewDecoder(r.Body).Decode(&st)
		if errors.Is(err, io.EOF){
			
		}
		slog.Info("creating student")

		w.Write([]byte("Welcome to student api"))
	}
}
