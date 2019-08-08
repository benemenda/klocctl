package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"klocctl/kwservertool"
)

var port string
var host string
var protocol string
var url string

type klocworkConfig struct {
	host     string
	port     string
	protocol string
}

type config struct {
	Klocwork klocworkConfig
}

func fmtURL(host, port, protocol string) string {
	var u string = ""
	if port != "0" {
		u = protocol + ":" + "//" + host + ":" + port
	}
	// if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
	// 	u = "http://" + u
	// }
	fmt.Println("Configured URL: " + u)
	return u
}

func values() config {
	return config{
		Klocwork: klocworkConfig{
			host:     viper.GetString("klocctl.host"),
			port:     "8080",
			protocol: "http",
		},
	}
}

func Print() {
	cfg := values()
	fmt.Println("KW Configuration:")
	fmt.Printf("%v \n", cfg.Klocwork)
	kwservertool.Request()

}

func Init(cfgFile string) {
	viper.SetEnvPrefix("klocctl")
	viper.SetConfigName("klocctl")        // name of config file (without extension)
	viper.AddConfigPath("$HOME/.klocctl") // adding home directory as first search path
	viper.AddConfigPath(".")              // adding home directory as first search path
	viper.AutomaticEnv()                  // read in environment variables that match

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		_, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if viper.Get("klocctl.port") == nil {
		viper.Set("klocctl.port", "8080")
	}
	if viper.Get("klocctl.host") == nil {
		viper.Set("klocctl.host", "localhost")
	}
	if viper.Get("klocctl.protocol") == nil {
		viper.Set("klocctl.protocol", "http")
	}
}

func Config() {
	host = viper.GetString("klocctl.host")
	port = viper.GetString("klocctl.port")
	protocol = viper.GetString("klocctl.protocol")
	url = fmtURL(host, port, protocol)
}
