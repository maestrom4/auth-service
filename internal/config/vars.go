package config

import (
	t "auth-service/internal/types"
)

const (
	ResolverKey t.ContextKey = "resolver"
	UserIDKey   t.ContextKey = "userID"
)
