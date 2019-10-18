package main

import (
	_sessionRepo "2019_2_Covenant/pkg/session/repository"
	_sessionUsecase "2019_2_Covenant/pkg/session/usecase"
	"2019_2_Covenant/pkg/user/delivery"
	_userRepo "2019_2_Covenant/pkg/user/repository"
	_userUsecase "2019_2_Covenant/pkg/user/usecase"
	"github.com/labstack/echo"
	"log"
)

func main() {
	e := echo.New()

	userStorage := _userRepo.NewUserStorage()
	userUsecase := _userUsecase.NewUserUsecase(userStorage)
	sessionStorage := _sessionRepo.NewSessionStorage()
	sessionUsecase := _sessionUsecase.NewSessionUsecase(sessionStorage)

	delivery.NewUserHandler(e, userUsecase, sessionUsecase)

	log.Fatal(e.Start(":8000"))
}
