package static

import (
	"net/http"
	"path/filepath"
	"strings"
)

type Server struct {
	staticFilePath string
}

func New(staticFilePath string) *Server {
	return &Server{
		staticFilePath: staticFilePath,
	}
}

func (h *Server) Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, `/`)

		filesPath := h.staticFilePath

		switch {
		case len(pathSplit) > 1 && pathSplit[1] == `api`:
			// если запрос проишел в апи, продолжаем выполнение
			next.ServeHTTP(w, r)

		case len(pathSplit) > 1 && pathSplit[1] != `api`:
			qwe := strings.Join(pathSplit[1:len(pathSplit)], "/")
			filepath.Join(filesPath, qwe)

			_, err := http.Dir(filesPath).Open(qwe)
			if err != nil { // todo проверить что ошибка == file not found
				//  в противном случае отдаём index
				http.ServeFile(w, r, filepath.Join(filesPath, "index.html"))
				return
			}

			// если удаётся найти файл, то сервим его
			fileServer := http.FileServer(http.Dir(filesPath))
			http.StripPrefix("/", fileServer).ServeHTTP(w, r)

		}

		defer func() {}()
	}

	return http.HandlerFunc(fn)
}
