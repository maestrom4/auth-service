package utils

import (
	"context"
	"fmt"
)

func GetStringFromContext(ctx context.Context, key string) (string, error) {
	value := ctx.Value(key)
	if value == nil {
		return "", fmt.Errorf("key '%s' not found in context", key)
	}
	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("value stored under key '%s' is not a string", key)
	}
	return strValue, nil
}

func GetBoolFromContext(ctx context.Context, key string) (bool, error) {
	value := ctx.Value(key)
	if value == nil {
		return false, fmt.Errorf("key '%s' not found in context", key)
	}
	boolValue, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("value stored under key '%s' is not a boolean", key)
	}
	return boolValue, nil
}
