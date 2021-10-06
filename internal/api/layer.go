package api

import (
	"docker/internal/types"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PullLayer(config *types.LayerConfig, auth *types.RegistryAuthResponse, client *http.Client) (*types.LayerResponse, error) {
	uri := fmt.Sprintf("%s/v2/library/%s/blobs/%s", config.RegistryAddress, config.ImageName, config.Digest)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.Token))

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

	layerResponse := &types.LayerResponse{
		Digest:  config.Digest,
		Content: body,
	}

	return layerResponse, nil
}
