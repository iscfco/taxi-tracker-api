package jwttasks

import (
	"gbmchallenge/api/config"
	"os"
)

const (
	// Values are in minutes
	accessTokenDuration  = 5
	refreshTokenDuration = (10 * 1) * 1 // = (mins * hours) * days
)

func getJwtPrivateKeyPath() string {
	switch config.ApiEnv {
	case config.Production:
		return "api/security/jwttasks/keys/private_key"
	case config.Local:
		return os.Getenv("GOPATH") + "/src/gbmchallenge/api/security/jwttasks/keys/private_key"
	}
	return ""
}

func getJwtPublicKeyPath() string {
	switch config.ApiEnv {
	case config.Production:
		return "api/security/jwttasks/keys/public_key.pub"
	case config.Local:
		return os.Getenv("GOPATH") + "/src/gbmchallenge/api/security/jwttasks/keys/public_key.pub"
	}
	return ""
}
