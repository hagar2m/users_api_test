package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadViber() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./configs")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func GetDatabaseUrl() string {
	return viper.GetString("database.url")
}
func GetPort() string {
	return viper.GetString("server.port")
}