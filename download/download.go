package download

import(
	"net/http"
	"fmt"
	"github.com/google/uuid"
	"os"
	"io/ioutil"
)

func DownloadVideo(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	tiktokName := fmt.Sprintf("tiktok%s.mp4", uuid.NewString())
	err = os.WriteFile(fmt.Sprintf("./saved_videos/%s", tiktokName), body, 0666)
	if err != nil {
		return "", err
	}
	return tiktokName, nil
}