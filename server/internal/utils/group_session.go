package utils

import "github.com/SwipEats/SwipEats/server/internal/types"

func CreateGroupSessionService() *types.GroupSessionService {
	return &types.GroupSessionService{
		Sessions: make(map[string]*types.GroupSession),
	}
}