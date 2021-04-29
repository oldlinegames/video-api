package api

import (
	// "net/http"
	// "io/ioutil"
	// "os"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/oldlinegames/video-api/download"
	"github.com/oldlinegames/video-api/structs"
)

var VideoQueue *structs.VideoQueue

func HostAPI() {
	VideoQueue = &structs.VideoQueue{
		Queue: []string{},
	}

	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Getting to that bag Sheeesh")
    })

	app.Get("/download", dlHandler)
	app.Post("/upload", upHandler)

    app.Listen(":3000")
}


func upHandler(c *fiber.Ctx) error {
	v := new(structs.VideoUpload)
	if err := c.BodyParser(v); err != nil {
		return err
	}

	for _, videoURL := range v.Videos {
		title, err := download.DownloadVideo(videoURL)
		if err != nil {
			return err
		}
		VideoQueue.Mux.Lock()
		VideoQueue.Queue = append(VideoQueue.Queue, title)
		VideoQueue.Mux.Unlock()
	}
	fmt.Println(VideoQueue.Queue)
	return nil
}

func dlHandler(c *fiber.Ctx) error {
	return nil
}