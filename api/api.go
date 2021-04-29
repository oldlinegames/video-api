package api

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
)

func HostAPI() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Getting to that bag Sheeesh")
    })

	app.Get("/download", dlHandler)
	app.Post("/upload", upHandler)

    app.Listen(":3000")
}

func upHandler(c *fiber.Ctx) error {
	return nil
}

func dlHandler(c *fiber.Ctx) error {
	return nil
}

func DownloadVideo(url string) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("user-agent", "Prometheus2.24.1")
	req.Header.Add("accept", "application/openmetrics-text; version=0.0.1,text/plain;version=0.0.4;q=0.5,*/*;q=0.1")
	req.Header.Add("X-Prometheus-Scrape-Timeout-Seconds", "10")
	// req.Header.Add("Accept-Encoding", "gzip")

	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return err
	}

	err = os.WriteFile("test.mp4", body, 0666)
	if err != nil {
		return err
	}
	return nil
}