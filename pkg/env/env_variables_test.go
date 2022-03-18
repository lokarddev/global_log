package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitEnvVariables(t *testing.T) {

	t.Run("Invalid. (no valid env provided)", func(t *testing.T) {
		err := InitEnvVariables()

		assert.Error(t, err)
	})
	t.Run(".env file", func(t *testing.T) {
		envFile := ".env"
		_, err := os.Create(envFile)
		assert.NoError(t, err)

		err = InitEnvVariables()

		assert.NoError(t, err)

		err = os.Remove(envFile)
		assert.NoError(t, err)
	})
	t.Run("OS env variables", func(t *testing.T) {
		err := os.Setenv("DEBUG", "true")
		assert.NoError(t, err)

		err = InitEnvVariables()

		assert.NoError(t, err)
		os.Clearenv()
	})
	t.Run(".env.dev file", func(t *testing.T) {
		envFile := ".env.dev"
		_, err := os.Create(envFile)
		assert.NoError(t, err)

		err = InitEnvVariables()

		assert.NoError(t, err)

		err = os.Remove(envFile)
		assert.NoError(t, err)
	})
}

func TestEnvFileLookup(t *testing.T) {
	envFile := ".env"
	t.Run("Valid", func(t *testing.T) {
		_, err := os.Create(envFile)
		assert.NoError(t, err)

		err = envFileLookup()

		assert.NoError(t, err)

		err = os.Remove(envFile)
		assert.NoError(t, err)
	})
	t.Run("No .env file", func(t *testing.T) {
		err := envFileLookup()

		assert.Error(t, err)
	})
}

func TestCheckOSDebug(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		err := os.Setenv("DEBUG", "true")
		assert.NoError(t, err)

		res, err := checkOSDebug()

		assert.NoError(t, err)
		assert.Equal(t, true, res)
		os.Clearenv()
	})
	t.Run("No OS env variable", func(t *testing.T) {
		res, err := checkOSDebug()

		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}

func TestLoadDefaultEnv(t *testing.T) {
	envFile := ".env.dev"
	t.Run("Valid", func(t *testing.T) {
		_, err := os.Create(envFile)
		assert.NoError(t, err)

		err = loadDefaultEnv()

		assert.NoError(t, err)

		err = os.Remove(envFile)
		assert.NoError(t, err)
	})
	t.Run("Invalid", func(t *testing.T) {
		err := loadDefaultEnv()

		assert.Error(t, err)
	})
}
