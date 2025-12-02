package response

import (
	"civi-id-app/internal/models"
	"civi-id-app/pkg/utils"
)

type AdminResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToAdminResponse(admin models.User) AdminResponse {
	return AdminResponse{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Role:      admin.Role.Name,
		CreatedAt: utils.FormatDate(admin.CreatedAt),
		UpdatedAt: utils.FormatDate(admin.UpdatedAt),
	}
}
