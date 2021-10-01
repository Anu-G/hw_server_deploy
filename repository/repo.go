package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AccountDetail struct {
	Userid    string `json:""`
	Username  string `json:"username"`
	Followers string `json:"followers"`
}

func Load() error {
	acc := AccountDetail{}
	file, err := ioutil.ReadFile((os.Getenv("FILE_JSON")))
	if err != nil {
		return err
	}

	errr := json.Unmarshal([]byte(file), &acc)
	if errr != nil {
		return errr
	}

	fmt.Println(acc)
	return nil
}

func FindFollowers(username string) (int, error) {

	return 0, nil
}

func FindUser(userid string) (string, int, error) {
	return "", 0, nil
}
