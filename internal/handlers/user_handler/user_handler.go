package userhandler

import (
	userrequest "civi-id-app/internal/dto/request/user_request"
	userresponse "civi-id-app/internal/dto/response/user_response"
	userservice "civi-id-app/internal/services/user_service"
	errorresponse "civi-id-app/pkg/constant/error_response"
	"civi-id-app/pkg/constant/response"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService userservice.IUserService
}

func NewUserHandler(userService userservice.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req userrequest.RegisterUserRequest

	req.NIK = c.FormValue("nik")
	req.Name = c.FormValue("name")
	req.Email = c.FormValue("email")
	req.Password = c.FormValue("password")
	req.TempatLahir = c.FormValue("tempat_lahir")
	req.BirthDate = c.FormValue("birth_date")
	req.Agama = c.FormValue("agama")
	req.Address = c.FormValue("address")
	req.PhoneNumber = c.FormValue("phone_number")
	req.Status = c.FormValue("status")

	photoFile, err := c.FormFile("photo_file")
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Photo file is required", err.Error())
	}
	req.PhotoFile = photoFile

	err = h.userService.Register(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create user")
	}

	return response.Success(c, http.StatusCreated, "User Created Successfully", nil)
}

func (a *UserHandler) LoginUser(c echo.Context) error {
	var req userrequest.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	token, err := a.userService.Login(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "invalid login user")
	}

	return response.Success(c, http.StatusOK, "Login Successfully", map[string]interface{}{
		"access_token": token,
	})
}

func (a *UserHandler) GetProfileUser(c echo.Context) error {
	userToken := c.Get("user")
	if userToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	user, ok := userToken.(*jwt.Token)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	claims := user.Claims.(jwt.MapClaims)
	userIDStr := claims["user_id"].(string)
	userID, _ := strconv.Atoi(userIDStr)

	me, err := a.userService.GetProfile(c.Request().Context(), userID)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "invalid to get profile")
	}

	userResponse := userresponse.ToUserResponse(*me)
	return response.Success(c, http.StatusOK, "Get Profile Successfully", userResponse)
}

func (h *UserHandler) UpdateProfileUser(c echo.Context) error {
	userToken := c.Get("user")
	if userToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userIDStr := claims["user_id"].(string)
	userID, _ := strconv.Atoi(userIDStr)

	var req userrequest.UpdateUserRequest
	req.Email = c.FormValue("email")
	req.Address = c.FormValue("address")
	req.PhoneNumber = c.FormValue("phone_number")

	if err := h.userService.UpdateProfile(c.Request().Context(), userID, req); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "Failed to update profile")
	}

	return response.Success(c, http.StatusOK, "Profile updated successfully", nil)
}

func (a *UserHandler) GenerateQR(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["user_id"].(string))

	qrImage, err := a.userService.GenerateQR(c.Request().Context(), userID)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, 500, "Failed to generate QR", err.Error())
	}

	return response.Success(c, 200, "QR Generated", map[string]string{
		"qr_url": qrImage,
	})
}


func (a *UserHandler) LogoutUser(c echo.Context) error {
	userToken := c.Get("user")
	if userToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	user, ok := userToken.(*jwt.Token)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	claims := user.Claims.(jwt.MapClaims)
	userIDStr := claims["user_id"].(string)
	userID, _ := strconv.Atoi(userIDStr)

	if err := a.userService.Logout(c.Request().Context(), userID); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "Invalid to logout user")
	}

	return response.Success(c, http.StatusOK, "Logout Successfully", nil)
}
