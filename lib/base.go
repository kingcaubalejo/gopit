package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"math/rand"
	"time"
)

func EncryptPlainText(str string) string{
	strUno := RandomString(5)
	strDos := RandomString(30)
	finalstr :=  ")*&^%$#@!13423423CodeThenDEcoDe" +  str + "ToCodeOrNotToCodeHardCoder!@#$%^&*()_03746253" + strUno + strDos
	m := md5.New()
	m.Write([]byte(finalstr))
	return hex.EncodeToString(m.Sum(nil)) + "_" + strUno + "_" + strDos
}

func CompareEncryptedPlainText(plainText string, cipherText string) bool{
   splitArray := strings.Split(cipherText, "_")
   if len(splitArray) != 3 {
	   return false
   }
   finalstr :=  ")*&^%$#@!13423423CodeThenDEcoDe" +  plainText + "ToCodeOrNotToCodeHardCoder!@#$%^&*()_03746253" + splitArray[1] + splitArray[2]
   m := md5.New()
   m.Write([]byte(finalstr))
   encrypted  := hex.EncodeToString(m.Sum(nil)) + "_" + splitArray[1] + "_" + splitArray[2]
   if encrypted == cipherText {
	   return true;
   }
   return false;
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