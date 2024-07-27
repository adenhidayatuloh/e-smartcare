package service

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	siswarepository "esmartcare/repository/siswaRepository"
	"net/http"
	"strconv"

	riwayattanyajawabrepository "esmartcare/repository/riwayatTanyaJawabRepository"
	"esmartcare/repository/userrepository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	PakarAuthorization() gin.HandlerFunc
	SiswaAuthorization() gin.HandlerFunc
	AdminAndPakarAuthorization() gin.HandlerFunc
	RiwayatAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo    userrepository.UserRepository
	siswaRepo   siswarepository.SiswaRepository
	riwayatRepo riwayattanyajawabrepository.RiwayatTanyaJawabRepository
}

func NewAuthService(userRepo userrepository.UserRepository, siswaRepo siswarepository.SiswaRepository, riwayatRepo riwayattanyajawabrepository.RiwayatTanyaJawabRepository) AuthService {
	return &authService{userRepo: userRepo, siswaRepo: siswaRepo, riwayatRepo: riwayatRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userData", result)
		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.JenisAkun != "1" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) PakarAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.JenisAkun != "2" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) SiswaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.JenisAkun != "3" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) AdminAndPakarAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.JenisAkun != "1" && userData.JenisAkun != "2" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) RiwayatAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		taskID := ctx.Param("id")
		taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Riwayat id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		_ = taskIDUint

		task, err2 := a.riwayatRepo.FindById(uint(taskIDUint))
		if err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err2)
			return
		}

		if task.Email == "" {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, "data not found")
			return

		}

		if task.Email != userData.Email {
			newError := errs.NewUnauthorized("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}
