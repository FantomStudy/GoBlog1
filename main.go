package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//https://youtu.be/WF9sjPg448A?si=9w3YDQXXTGOwJb8j
//xampp - для запуска сервера
//команда \xampp\xampp_start.exe для запуска
//команда \xampp\xampp_stop.exe для остановки
//https://gorm.io/docs/belongs_to.html - для бд
//https://www.apachefriends.org/ru/download.html - сервер
//https://gowebexamples.com/password-hashing/ - шифрование паролей
//https://github.com/joho/godotenv
//https://pkg.go.dev/github.com/golang-jwt/jwt - JWT пакет
//https://docs.gofiber.io/ - GO FIBER
//http://localhost/phpmyadmin/ - там бд после постановки сервера
//http://localhost/phpmyadmin/index.php?route=/sql&pos=0&db=goblog&table=blogs - рандомные пикчи для постов
//go mod tidy - восстановление зависимостей

func main() {
	database.Connect()
	err:=godotenv.Load()
	if err!= nil{
		log.Fatal("Error load .env files")
	}
	port:=os.Getenv("PORT")
	app:=fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)
}