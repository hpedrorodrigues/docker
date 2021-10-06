package types

type ImageLayer struct {
	BlobSum string `json:"blobSum"`
}

type ImageManifest struct {
	Name     string       `json:"name"`
	Tag      string       `json:"tag"`
	FsLayers []ImageLayer `json:"fsLayers"`
}

type ManifestConfig struct {
	RegistryAddress string
	ImageName       string
	ImageReference  string
}

type LayerConfig struct {
	RegistryAddress string
	ImageName       string
	Digest          string
}

type LayerResponse struct {
	Digest  string
	Content []byte
}
