package utils

import "github.com/KokoiRuby/rbac-based-management-system/backend/global"

func IsReady() bool {
	flag := true
	for _, ready := range global.Readiness {
		if !ready {
			flag = false
			break
		}
	}
	return flag
}
