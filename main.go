package main

import (
	"github.com/oldlinegames/video-api/api"
)

func main() {
	err := api.DownloadVideo(`https://v1.musicallydown.com/dl?url=aHR0cHM6Ly92MzktZXUudGlrdG9rY2RuLmNvbS85ZWViMjc4NDgxZmU1NjBlYWNmMTdkNDMyNzlmYWIzNy82MDhhODM0ZS92aWRlby90b3MvdXNlYXN0MmEvdG9zLXVzZWFzdDJhLXZlLTAwNjhjMDA0LzY3ZDU3MjA0OGI3ZDRiZWVhZjExYTc4YTJjNmMyN2JkLz9hPTEyMzMmYnI9MTE5MCZidD01OTUmY2Q9MCU3QzAlN0MwJmNoPTAmY3I9MCZjcz0wJmN2PTEmZHI9MCZkcz02JmVyPSZsPTIwMjEwNDI5MDM1ODI4MDEwMDk5MDgxMDA0MzgwMDAwMDMmbHI9YWxsJm1pbWVfdHlwZT12aWRlb19tcDQmbmV0PTAmcGw9MCZxcz0wJnJjPWFqdHNiSGw0T1RoeU5ETXpOemN6TTBBcE9qZHBORFExWjJRMk4yWTBaanc4WldkamEyNHVOR1EyYUhGZ0xTMWtNVFp6Y3pSallqTmhZVjR2WWw5aFlWNWlZalU2WXclM0QlM0Qmdmw9JnZyPQ==&name=ZGFtaWVuLndvb2Q1MDU=&type=mp4`)
	if err != nil {
		panic(err)
	}
}