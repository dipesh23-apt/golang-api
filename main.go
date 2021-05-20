package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id         uint `gorm:"primary key" json:"id"`
	CreatedAt  int64
	UpdatedAt  int64
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	City       string  `json:"city"`
	Phone      string  `json:"phone"`
	Height     float32 `json:"height"`
	Gender     string  `json:"gender"`
	Password   string  `json:"password"`
	Married    bool    `json:"married"`
}
type result struct {
	Id         uint    `gorm:"primary key" json:"id"`
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	City       string  `json:"city"`
	Phone      string  `json:"phone"`
	Height     float32 `json:"height"`
	Gender     string  `json:"gender"`
	Password   string  `json:"password"`
	Married    bool    `json:"married"`
}

var users []User
var DB *gorm.DB
var res result
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var IsDigit = regexp.MustCompile(`^[0-9]*$`).MatchString

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	param := mux.Vars(r)
	b := DB.Find(&users).RowsAffected
	i, _ := strconv.ParseInt(param["id"], 10, 64)
	if i > b {
		fmt.Fprintf(w, "Invalid id.Does not Exist !!")
		return
	}
	DB.Table("user").Select("id", "first_name", "last_name", "city", "phone", "height", "gender", "password", "married").Where("id = ?", param["id"]).Scan(&res)

	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}
func getallUsers(w http.ResponseWriter, r *http.Request) {

	//var users []User
	//result := DB.Find(&users)
	var ro []int
	DB.Table("user").Select("id").Scan(&ro)
	fmt.Println(ro)
	if len(ro) != 0 {
		b := make(map[string][]int)
		b["ids"] = ro
		json.NewEncoder(w).Encode(b)
	} else {
		fmt.Fprintf(w, "No data Present in the Database !!")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("CreateUser() called")
	var us User
	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !IsLetter(us.First_name) || !IsLetter(us.Last_name) || !IsLetter(us.City) || !IsLetter(us.Gender) {
		fmt.Fprintf(w, "Invalid Entry in one of the fields(should not contain digit)")
		return
	}
	if !IsDigit(us.Phone) {
		fmt.Fprintf(w, "Invalid Entry in one of the fields(Only Digits allowed)")
		return
	}
	if len(us.Phone) != 10 {
		fmt.Fprintf(w, "Phone Number should consist of 10 digits !!")
		return
	}
	b := DB.Find(&users).RowsAffected
	us.Id = uint(b + 1)
	us.CreatedAt = time.Now().Unix()
	us.UpdatedAt = time.Now().Unix()

	//Encrypting the password and saving
	h := sha1.New()
	h.Write([]byte(us.Password))
	pass_hash := hex.EncodeToString(h.Sum(nil))
	us.Password = pass_hash
	fmt.Println(us)

	//Commiting Into Database
	result := DB.Create(&us)
	//fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	json.NewEncoder(w).Encode(us)

}
func (User) TableName() string {
	return "user"
}

func InitialMigration() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	//db.Migrator().HasTable(&User{})
	//db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	return db, nil
}
func main() {
	x, err := InitialMigration()
	DB = x
	if err != nil {
		panic("Could not connect Database")
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/v1/user/fetch", getallUsers).Methods("POST")
	r.HandleFunc("/api/v1/user/create", createUser).Methods("POST")
	//r.HandleFunc("/api/v1/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/v1/user/{id}", deleteUser).Methods("DELETE")
	http.Handle("/", r)
	fmt.Println("Running on server 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	var a User
	DB.Where("id= ?", param["id"]).Find(&a)
	DB.Delete(&a)
	fmt.Fprintf(w, "Successfully deleted the row with id= %s", param["id"])
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	//Will Update accordingly,as of now it isTemporary
	
	/*w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	b := DB.Find(&users).RowsAffected
	i, _ := strconv.ParseInt(param["id"], 10, 64)
	if i > b {
		fmt.Fprintf(w, "Invalid id.Does not Exist !!")
		return
	}
	var res2 result
	DB.Table("user").Select("id", "first_name", "last_name", "city", "phone", "height", "gender", "password", "married").Where("id = ?", param["id"]).Scan(&res2)
	fmt.Println(res2)
	var a User
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(a)
	*/
}
