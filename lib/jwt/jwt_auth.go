package jwt

import (
    "crypto/md5"
    "crypto/rsa"
	"crypto/x509"
    "encoding/hex"
    "net/http"
    "strings"
    "time"
    "math/rand"
    "encoding/pem"
    "os"
    "bufio"

    jwtAuthenticate "github.com/dgrijalva/jwt-go"
    "go-api-jwt-v2/settings"
)

// const (
// 	tokenDuration = 72
// 	expireOffset  = 3600
// )

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

func Crackdependmaker(str string) string{
	strUno := RandomString(5)
	strDos := RandomString(30)
	finalstr :=  ")*&^%$#@!13423423CodeThenDEcoDe" +  str + "ToCodeOrNotToCodeHardCoder!@#$%^&*()_03746253" + strUno + strDos
	m := md5.New()
	m.Write([]byte(finalstr))
	return hex.EncodeToString(m.Sum(nil)) + "_" + strUno + "_" + strDos
}

func Crackdepend(str string,crypt string) bool{
   splitArray := strings.Split(crypt,"_")
   if len(splitArray) != 3 {
	   return false
   }
   finalstr :=  ")*&^%$#@!13423423CodeThenDEcoDe" +  str + "ToCodeOrNotToCodeHardCoder!@#$%^&*()_03746253" + splitArray[1] + splitArray[2]
   m := md5.New()
   m.Write([]byte(finalstr))
   encrypt  := hex.EncodeToString(m.Sum(nil)) + "_" + splitArray[1] + "_" + splitArray[2]
   if encrypt == crypt {
	   return true;
   }
   return false;
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

func RandomString(strlen int) string {
   rand.Seed(time.Now().UTC().UnixNano())
   const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
   result := make([]byte, strlen)
   for i := 0; i < strlen; i++ {
	   result[i] = chars[rand.Intn(len(chars))]
   }
   return string(result)
}