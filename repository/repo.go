package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var res map[string]interface{}

func Load() error {
	file, err := ioutil.ReadFile((os.Getenv("FILE_JSON")))
	if err != nil {
		return err
	}

	errr := json.Unmarshal([]byte(file), &res)
	if errr != nil {
		return errr
	}

	return nil
}

func FindFollowers(username string) (int, error) {
	err := Load()
	if err != nil {
		return 0, err
	}

	for userid := range res {
		detail := res[userid].(map[string]interface{})
		for key, val := range detail {
			if key == "username" && val == username {
				return int(detail["followers"].(float64)), nil
			}
		}
	}
	return 0, nil
}

func FindUser(userid string) (string, int, error) {
	err := Load()
	if err != nil {
		return "", 0, err
	}

	for id := range res {
		if id == userid {
			detail := res[userid].(map[string]interface{})
			return detail["username"].(string), int(detail["followers"].(float64)), nil
		}
	}
	return "", 0, nil
}
