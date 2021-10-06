package api

import (
	"docker/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PullManifest(config *types.ManifestConfig, auth *types.RegistryAuthResponse, client *http.Client) (*types.ImageManifest, error) {
	uri := fmt.Sprintf("%s/v2/library/%s/manifests/%s", config.RegistryAddress, config.ImageName, config.ImageReference)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.Token))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var manifest *types.ImageManifest
	err = json.Unmarshal(body, &manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}
