package api

import (
	"docker/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Authenticate(config *types.RegistryAuthConfig, client *http.Client) (*types.RegistryAuthResponse, error) {
	uri := fmt.Sprintf("%s/token?service=%s&scope=%s", config.ServerAddress, config.Service, config.Scope)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var auth *types.RegistryAuthResponse
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
