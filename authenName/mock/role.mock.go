package mock

import (
	"authenName/model"
	repo "authenName/repository"
)

func MockRole(name string) {
	role := model.Role{}
	role.Name = name
	repo.CreateRole(role)
}
