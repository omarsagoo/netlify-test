package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/upload", func(c echo.Context) error {
		// Multipart form
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}

		src, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()

		r, err := zip.NewReader(src, file.Size)
		if err != nil {
			log.Fatal(err)
		}
		for _, zip_file := range r.File {
			filepaths := strings.Split(zip_file.Name, "/")
			if filepaths[0] == "__MACOSX" {
				continue
			}
			path := strings.Join(filepaths[:len(filepaths)-1], "/")
			err := os.MkdirAll("builds/"+path, 0777)
			if err != nil {
				log.Fatal(err)
			}

			if strings.Contains(zip_file.Name, ".") {
				dst, err := os.Create("builds/" + zip_file.Name)
				if err != nil {
					log.Fatal(err)
				}
				defer dst.Close()

				src, err := zip_file.Open()
				if err != nil {
					log.Fatal(err)
				}
				defer src.Close()

				io.Copy(dst, src)
			}
		}

		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", r.Comment))
	})

	e.Static("/", "builds/roll-a-ball")

	e.Logger.Fatal(e.Start(":8080"))
}
