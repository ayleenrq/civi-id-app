package rolehandler

import (
	rolerequest "civi-id-app/internal/dto/request/role_request"
	roleresponse "civi-id-app/internal/dto/response/role_response"
	roleservice "civi-id-app/internal/services/role_service"
	errorresponse "civi-id-app/pkg/constant/error_response"
	"civi-id-app/pkg/constant/response"
	"civi-id-app/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService roleservice.IRoleService
}

func NewRoleHandler(roleService roleservice.IRoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (r *RoleHandler) CreateRole(c echo.Context) error {
	var req rolerequest.CreateRoleRequest
	req.Name = c.FormValue("name")

	err := r.roleService.CreateRole(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create role")
	}

	return response.Success(c, http.StatusOK, "Role Created Successfully", nil)
}

func (r *RoleHandler) GetAllRole(c echo.Context) error {
	pageInt, limitInt := utils.ParsePaginationParams(c, 10)
	search := c.QueryParam("search")

	roles, total, err := r.roleService.GetAllRole(c.Request().Context(), pageInt, limitInt, search)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get role")
	}

	meta := utils.BuildPaginationMeta(c, pageInt, limitInt, total)
	data := make([]roleresponse.RoleResponse, len(roles))
	for i, s := range roles {
		data[i] = roleresponse.ToRoleResponse(*s)
	}

	return response.PaginatedSuccess(c, http.StatusOK, "Get All Role Successfully", data, meta)
}

func (r *RoleHandler) GetByIdRole(c echo.Context) error {
	roleIdStr := c.Param("roleId")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid role ID", err.Error())
	}

	role, err := r.roleService.GetByIdRole(c.Request().Context(), roleId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get role")
	}

	res := roleresponse.ToRoleResponse(*role)

	return response.Success(c, http.StatusOK, "Get Role Successfully", res)
}

func (r *RoleHandler) UpdateRole(c echo.Context) error {
	roleIdStr := c.Param("roleId")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid role ID", err.Error())
	}

	var req rolerequest.UpdateRoleRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = r.roleService.UpdateRole(c.Request().Context(), roleId, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to update role")
	}

	return response.Success(c, http.StatusOK, "Role Updated Successfully", nil)
}

func (r *RoleHandler) DeleteRole(c echo.Context) error {
	roleIdStr := c.Param("roleId")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid role ID", err.Error())
	}

	if err := r.roleService.DeleteRole(c.Request().Context(), roleId); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to delete role")
	}

	return response.Success(c, http.StatusOK, "Role Deleted Successfully", nil)
}
