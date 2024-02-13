package web

import (
	"embed"
	"github.com/bjut-tech/auth-server/internal/config"
	"io/fs"
	"net/http"
	"os"
)

//go:embed dist
var EmbedFs embed.FS

func HttpFs() http.FileSystem {
	if config.Production {
		fsys, err := fs.Sub(EmbedFs, "dist")
		if err != nil {
			panic(err)
		}

		return http.FS(fsys)
	}

	return http.FS(os.DirFS("web/dist"))
}
