package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AccountDetail struct {
	Userid    string `json:""`
	Username  string `json:"username"`
	Followers string `json:"followers"`
}

func Load() error {
	acc := AccountDetail{}
	file, err := ioutil.ReadFile("./hw.json")
	if err != nil {
		return err
	}
	fmt.Println([]byte(file))
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
