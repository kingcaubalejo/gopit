package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func EncryptPlainText(str string) string {
	if str == "" {
		return ""
	}

	finalstr :=  ")*&^%$#@!13423423CodeThenDEcoDe" +  str + "ToCodeOrNotToCodeHardCoder!@#$%^&*()_03746253"
	m := md5.New()
	m.Write([]byte(finalstr))
	return hex.EncodeToString(m.Sum(nil))
}

func CompareEncryptedPlainText(plainText string, cipherText string) bool {
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