package cfsservice

type RuntimeConfig struct {
	Port               uint64
	DBConnectionString string
	JWTPublicKeyPath   string
}
