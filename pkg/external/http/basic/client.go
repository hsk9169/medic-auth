package basic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	dto "github.com/medic-basic/auth/pkg/dto/http"
)

var BasicImpl Basic

const (
	User  = "user"
	Medic = "medic"
    Host  = "{{host_alb_url}}"
)

type Basic struct {
	url    string
	scheme string
}

func InitBasic() {
	BasicImpl = Basic{
        url:    "{{host_alb_url}}"
		scheme: "http",
	}
}

func GetAccountInfo(reqData dto.GetAccountInfoReqToBasic) (any, error) {
	var url string
	if reqData.Role == User {
		url = Host + "/" + User + "/" + reqData.UserID
	} else if reqData.Role == Medic {
		url = Host + "/" + Medic + "/" + reqData.UserID
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var medicData dto.MedicDataFromBasicApi
	if err = json.Unmarshal(respBody, &medicData); err != nil {
		fmt.Println("Failed to unmarshal response json body")
		return nil, errors.New("failed to parse data from basic api")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200")
	}

	return medicData, nil
}
