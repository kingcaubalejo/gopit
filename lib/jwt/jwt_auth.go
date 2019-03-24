package jwt

import (
    "crypto/rsa"
	"crypto/x509"
    "net/http"
    "time"
    "encoding/pem"
    "os"
    "bufio"

    jwtAuthenticate "github.com/dgrijalva/jwt-go"
    "go-api-jwt-v2/settings"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

func SignToken(crypt string, tokenType string) (int,string) {
    token, err := GenerateToken(crypt, tokenType)
    if err != nil {
        return http.StatusInternalServerError, ""
    } else {
        return http.StatusOK, token
    }
}

func GenerateToken(crypt string, tokenType string) (string,interface{}) {
    claims := GenerateClaims(crypt, tokenType)
    token  := jwtAuthenticate.NewWithClaims(jwtAuthenticate.SigningMethodRS512,claims)
    tokenString, err := token.SignedString(getPrivateKey())
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func getPrivateKey() *rsa.PrivateKey {
	filePath := settings.GetExecDirectory()
	privateKeyFile, err := os.Open(filePath + "/" + settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	filePath := settings.GetExecDirectory()
	publicKeyFile, err := os.Open(filePath + "/" + settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}

func GenerateClaims(crypt string, tokenType string) jwtAuthenticate.MapClaims {
   var duration = 0
   if duration = settings.Get().JWTRefreshTokenExpiration; tokenType == "acess_token" {
	   duration = settings.Get().JWTAccessTokenExpiration
   }
   
   claims := make(jwtAuthenticate.MapClaims)
   claims["exp"] = time.Now().Add(time.Hour * time.Duration(duration)).Unix()
   claims["sub"] = crypt
   return claims
}