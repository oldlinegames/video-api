package api

import (
	// "net/http"
	// "io/ioutil"
	"fmt"
	"os"
	"sync"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/oldlinegames/video-api/download"
	"github.com/oldlinegames/video-api/structs"
)

var VideoQueue *structs.VideoQueue

func init() {
	if _, err := os.Stat("./saved_videos"); os.IsNotExist(err) {
		os.Mkdir("saved_videos", 0777)
	}
}

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
		return c.SendString(err.Error())
	}
	var wg sync.WaitGroup
	for _, videoURL := range v.Films {
		wg.Add(1)
		go download.DownloadVideo(videoURL, &wg, VideoQueue)
	}
	wg.Wait()
	return c.SendString("Videos uploaded succesfully")
}

func dlHandler(c *fiber.Ctx) error {
	VideoQueue.Mux.Lock()

	if len(VideoQueue.Queue) < 1 {
		return c.SendString("No videos loaded in queue")
	}

	videoTitle := VideoQueue.Queue[0]
	VideoQueue.Queue = VideoQueue.Queue[1:]
	VideoQueue.Mux.Unlock()
	c.Append("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, videoTitle))
	data, err := os.ReadFile(fmt.Sprintf("./saved_videos/%s", videoTitle))

	if err != nil {
		return c.SendString(err.Error())
	}

	err = os.Remove(fmt.Sprintf("./saved_videos/%s", videoTitle))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Send(data)
}
