package adminhandler

import (
	adminrequest "civi-id-app/internal/dto/request/admin_request"
	adminresponse "civi-id-app/internal/dto/response/admin_response"
	adminservice "civi-id-app/internal/services/admin_service"
	errorresponse "civi-id-app/pkg/constant/error_response"
	"civi-id-app/pkg/constant/response"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService adminservice.IAdminService
}

func NewAdminHandler(adminService adminservice.IAdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (a *AdminHandler) RegisterAdmin(c echo.Context) error {
	var req adminrequest.RegisterAdminRequest
	req.Name = c.FormValue("name")
	req.Email = c.FormValue("email")
	req.Password = c.FormValue("password")

	err := a.adminService.Register(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create admin")
	}

	return response.Success(c, http.StatusCreated, "Admin Created Successfully", nil)

}

func (a *AdminHandler) LoginAdmin(c echo.Context) error {
	var req adminrequest.LoginAdminRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	token, err := a.adminService.Login(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "invalid login admin")
	}

	return response.Success(c, http.StatusOK, "Login Successfully", map[string]interface{}{
		"access_token": token,
	})
}

func (a *AdminHandler) GetProfileAdmin(c echo.Context) error {
	adminToken := c.Get("user")
	if adminToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	admin, ok := adminToken.(*jwt.Token)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	claims := admin.Claims.(jwt.MapClaims)
	adminIDStr := claims["user_id"].(string)
	adminID, _ := strconv.Atoi(adminIDStr)

	me, err := a.adminService.GetProfile(c.Request().Context(), adminID)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "invalid to get profile")
	}

	adminResponse := adminresponse.ToAdminResponse(*me)
	return response.Success(c, http.StatusOK, "Get Profile Successfully", adminResponse)
}

func (a *AdminHandler) UpdateProfileAdmin(c echo.Context) error {
	adminToken := c.Get("user")
	if adminToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	admin, ok := adminToken.(*jwt.Token)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	claims := admin.Claims.(jwt.MapClaims)
	adminIDStr := claims["user_id"].(string)
	adminID, _ := strconv.Atoi(adminIDStr)

	var req adminrequest.UpdateProfileRequest
	req.Name = c.FormValue("name")

	err := a.adminService.UpdateProfile(c.Request().Context(), adminID, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "Failed to update profile")
	}

	return response.Success(c, http.StatusOK, "Profile updated successfully", nil)
}

func (a *AdminHandler) LogoutAdmin(c echo.Context) error {
	adminToken := c.Get("user")
	if adminToken == nil {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	admin, ok := adminToken.(*jwt.Token)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized: token invalid or expired", nil)
	}

	claims := admin.Claims.(jwt.MapClaims)
	adminIDStr := claims["user_id"].(string)
	adminID, _ := strconv.Atoi(adminIDStr)

	if err := a.adminService.Logout(c.Request().Context(), adminID); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "Invalid to logout admin")
	}

	return response.Success(c, http.StatusOK, "Logout Successfully", nil)
}
