package cmd

import "auth-service/internal/initializer"

func init() {
	initializer.LoadEnv()
}

func main() {

	config := initializer.LoadConfig()
}
