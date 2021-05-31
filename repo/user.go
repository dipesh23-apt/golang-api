package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/dipesh23-apt/golang_api/models"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func encode_sha1(x string) string {
	h := sha1.New()
	h.Write([]byte(x))
	pass_hash := hex.EncodeToString(h.Sum(nil))
	return pass_hash
}
func validFields(us models.User) error {
	validate = validator.New()
	err := validate.Struct(us)
	if err != nil {
		return (err)
	}
	return nil
}

func GetUserfromDB(id string) (res models.User, err error) {
	result := DB.First(&res, id)
	return res, result.Error
}
func GetallUsersfromDB(x []uint) (c []models.User, err error) {
	result := DB.Where(x).Find(&c)
	return c, result.Error
}

func CreateUserinDB(us models.User) (d models.User, err error) {
	var b models.User
	err = validFields(us)
	if err != nil {
		return models.User{}, err
	}
	DB.Table("user").Select("Id").Last(&b)
	us.Id = b.Id + 1
	us.Password = encode_sha1(us.Password)
	result := DB.Create(&us)
	return us, result.Error

}
func DeleteUserfromDB(id string) (err error) {
	var a models.User
	DB.Where("id= ?", id).Find(&a)
	result := DB.Delete(&a)
	return result.Error
}
