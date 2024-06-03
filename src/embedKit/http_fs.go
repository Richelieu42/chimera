package embedKit

import (
	"embed"
	"net/http"
)

func ToHttpFileSystem(embedFs embed.FS) http.FileSystem {
	return http.FS(embedFs)
}
