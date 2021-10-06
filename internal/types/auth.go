package types

import "time"

type RegistryAuthConfig struct {
	ServerAddress string
	Service       string
	Scope         string
}

type RegistryAuthResponse struct {
	Token       string    `json:"token"`
	AccessToken string    `json:"access_token"`
	ExpiresIn   int32     `json:"expires_in"`
	IssuedAt    time.Time `json:"issued_at"`
}
