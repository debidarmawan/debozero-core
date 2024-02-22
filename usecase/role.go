package usecase

import (
	"regexp"
	"strings"

	"github.com/debidarmawan/debozero-core/repository"
)

type RoleUseCase interface {
	CanAccess(roleId string, path string, method string) (bool, string)
}

type roleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewRoleUseCase(roleRepo repository.RoleRepository) RoleUseCase {
	return &roleUseCase{
		roleRepo: roleRepo,
	}
}

// UUID v4
var re = regexp.MustCompile(`\/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

func (ru *roleUseCase) CanAccess(roleId string, path string, method string) (bool, string) {
	path = re.ReplaceAllString(path, `$1/{id}$2`)
	path = strings.TrimSuffix(path, "/")

	roleKeys, _ := ru.roleRepo.GetAllowedRoleKeyListByRoleId(roleId)
	canAccess := false

	for _, roleKey := range *roleKeys {
		if roleKey.Method == method && roleKey.Path == path {
			canAccess = true
			break
		}
	}

	return canAccess, path
}
