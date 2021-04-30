package download

import(
	"net/http"
	"fmt"
	"github.com/google/uuid"
	"os"
	"io/ioutil"
	"sync"
	"github.com/oldlinegames/video-api/structs"
)

func DownloadVideo(url string, wg *sync.WaitGroup, videoQueue *structs.VideoQueue) {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	tiktokName := fmt.Sprintf("tiktok%s.mp4", uuid.NewString())
	err = os.WriteFile(fmt.Sprintf("./saved_videos/%s", tiktokName), body, 0666)
	if err != nil {
		return
	}

	videoQueue.Mux.Lock()
	videoQueue.Queue = append(videoQueue.Queue, tiktokName)
	videoQueue.Mux.Unlock()

	return
}