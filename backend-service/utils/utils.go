package utils

import "os"

func GetEnvOrFallback(env, fallback string) string {
	envVar := os.Getenv(env)
	if envVar != "" {
		return envVar
	}
	return fallback
}
