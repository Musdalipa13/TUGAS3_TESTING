package controller

import (
	"bufio"
	"context"
	"fmt"
	"golang-database-user/model"
	"golang-database-user/service"
	"os"
	"strings"
)

func DefaultChoose() {
	fmt.Println("Incorrect Number")
}

func CreateUser(userService service.UserService) {

	ctx := context.Background()

	// buat inputan

	var Name, Email, Password, PhoneNumber string

	fmt.Print("Masukkan Name : ")
	fmt.Scan(&Name)

	fmt.Print("Masukkan Email : ")
	fmt.Scan(&Email)

	fmt.Print("Masukkan Password : ")
	fmt.Scan(&Password)

	fmt.Print("Masukkan PhoneNumber : ")
	fmt.Scan(&PhoneNumber)

	user := model.MstUser{
		Name:        Name,
		Email:       Email,
		Password:    Password,
		PhoneNumber: PhoneNumber,
	}

	mstUser := userService.CreateUser(ctx, user)

	fmt.Println(mstUser)
}

func UpdateUser(userService service.UserService) {
	reader := bufio.NewReader(os.Stdin)

	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan Id User yang ingin di update: ")
	fmt.Scanln(&userId)

	fmt.Print("Masukkan Name: ")
	Name, _ := reader.ReadString('\n')
	Name = strings.TrimSpace(Name)

	fmt.Print("Masukkan Email: ")
	Email, _ := reader.ReadString('\n')
	Email = strings.TrimSpace(Email)

	fmt.Print("Masukkan password: ")
	Password, _ := reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	fmt.Print("Masukkan PhoneNumber: ")
	PhoneNumber, _ := reader.ReadString('\n')
	PhoneNumber = strings.TrimSpace(PhoneNumber)

	user := model.MstUser{
		Name:        Name,
		Email:       Email,
		Password:    Password,
		PhoneNumber: PhoneNumber,
	}

	mstUser := userService.UpdateUser(ctx, user, userId)

	fmt.Println(mstUser)
}

func ReadUser(userService service.UserService) {
	ctx := context.Background()

	users, err := userService.ReadUser(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nAll Users:")
	for _, mstUser := range users {
		fmt.Println("Id : ", mstUser.IdUser, "\nName : ", mstUser.Name, "\nEmail : ", mstUser.Email, "\nPhoneNumber : ", mstUser.PhoneNumber)
		fmt.Println()
	}
}

func DeleteUser(userService service.UserService) {
	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan id user yang ingin di hapus: ")
	fmt.Scanln(&userId)

	err := userService.DeleteUser(ctx, userId)
	if err != nil {
		fmt.Println("Gagal menghapus user:", err)
	} else {
		fmt.Println("User berhasil dihapus")
	}
}
