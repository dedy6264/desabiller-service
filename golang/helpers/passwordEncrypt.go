package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"desabiller/configs"
	"desabiller/models"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

//	type JwtCustClaim struct {
//		SnDevice   string `json:"snDevice"`
//		UserId     string `json:"userId"`
//		OutletId   string `json:"outletId"`
//		MerchantId string `json:"merchantId"`
//		jwt.RegisteredClaims
//	}
func TokenJWTDecode(ctx echo.Context) (data models.DataToken) {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	data.MerchantId = int(claims["merchantId"].(float64))
	data.MerchantOutletId = int(claims["outletId"].(float64))
	data.MerchantOutletUsername = claims["outletUsername"].(string)
	return data
}
func TokenJwtGenerate(mID, oID int, oUsername string) (tkn string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// claims["snDevice"] = snDev
	claims["outletUsername"] = oUsername
	claims["outletId"] = oID
	claims["merchantId"] = mID
	// claims["clientId"] = cID
	//claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	t, err := token.SignedString([]byte(configs.KEY))
	if err != nil {
		return tkn, err
	}
	return t, nil
}

// func TokenJwtGenerateDashboard(uID int) (tkn string, err error) {
// 	// claim := &JwtCustClaim{
// 	// 	SnDevice:   "qqqq",
// 	// 	UserId:     "wwww",
// 	// 	OutletId:   "eee",
// 	// 	MerchantId: "rrrrr",
// 	// 	RegisteredClaims: jwt.RegisteredClaims{
// 	// 		ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 72)},
// 	// 	},
// 	// }
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	// claims["snDevice"] = snDev
// 	claims["userDashboardId"] = uID
// 	// claims["outletId"] = oID
// 	// claims["merchantId"] = mID
// 	// claims["clientId"] = cID
// 	//claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
// 	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

//		t, err := token.SignedString([]byte(configs.KEY))
//		fmt.Println("PPPPPPPP", err)
//		if err != nil {
//			return tkn, err
//		}
//		return t, nil
//	}
func PassEncrypt(pswrd string) (result string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pswrd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return result, err
	}
	return string(hashedPassword), nil
}
func PassCheck(reqpswrd string, pssword string) {
	err := bcrypt.CompareHashAndPassword([]byte(reqpswrd), []byte(pssword))
	fmt.Println("::::::CHECK", err)
}

func PswEnc(word string) (enc string, err error) {
	chipBlock, err := aes.NewCipher([]byte(configs.KEY))
	if err != nil {
		log.Println("ERROR ENCRYP 1", err.Error())
		return "", err
	}
	stream, err := cipher.NewGCM(chipBlock)
	if err != nil {
		log.Println("ERROR ENCRYP 1", err.Error())
		return "", err
	}

	cp := make([]byte, stream.NonceSize())
	cipherText := stream.Seal(cp, cp, []byte(word), nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil

}
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func Decrypt(datain string) (string, string, string) {
	dataindec, _ := base64.StdEncoding.DecodeString(datain)
	data := []byte(dataindec)
	key := []byte(createHash(configs.KEY))
	block, errAes := aes.NewCipher(key)
	if errAes != nil {
		return "", "81", "Decrypt : " + errAes.Error()
	}
	gcm, errCipher := cipher.NewGCM(block)
	if errCipher != nil {
		return "", "81", "Decrypt : " + errCipher.Error()
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, errGcm := gcm.Open(nil, nonce, ciphertext, nil)
	if errGcm != nil {
		return "", "81", "Decrypt : " + errGcm.Error()
	}
	output := string(plaintext)
	return output, "00", "Success"
}

func PswDec(word string) (decr string, err error) {
	dataindec, _ := base64.StdEncoding.DecodeString(word)
	data := []byte(dataindec)
	key := []byte(createHash(configs.KEY))
	block, errAes := aes.NewCipher(key)
	if errAes != nil {
		return "", errAes
	}
	gcm, errCipher := cipher.NewGCM(block)
	if errCipher != nil {
		return "", errCipher
	}
	nonceSize := gcm.NonceSize()
	fmt.Println("::", nonceSize)
	fmt.Println("::", data[:nonceSize])
	fmt.Println("::", data[nonceSize:])
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, errGcm := gcm.Open(nil, nonce, ciphertext, nil)
	if errGcm != nil {
		return "", errGcm
	}
	output := string(plaintext)
	return output, err
}
