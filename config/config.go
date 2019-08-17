package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/benemenda/klocctl/kw"
)

var port string
var host string
var protocol string
var user string
var ltoken string
var url string
var urlProm string

type klocworkConfig struct {
	host     string
	port     string
	protocol string
	user     string
	ltoken   string
}

type prometheusConfig struct {
	host string
	port string
}

type config struct {
	Klocwork   klocworkConfig
	Prometheus prometheusConfig
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
			port:     viper.GetString("klocctl.port"),
			protocol: viper.GetString("klocctl.protocol"),
			user:     viper.GetString("klocctl.user"),
			ltoken:   viper.GetString("klocctl.ltoken"),
		},
		Prometheus: prometheusConfig{
			host: viper.GetString("klocctl.prometheus.host"),
		},
	}
}

func Print() {
	cfg := values()
	fmt.Println("KW Configuration:")
	fmt.Printf("%v \n", cfg.Klocwork)
	fmt.Printf("%v \n", cfg.Prometheus)
	kw.ReceiveRequest("get", "builds", nil)

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
	if viper.Get("klocctl.prometheus.host") == nil {
		viper.Set("klocctl.prometheus.host", "localhost")
	}
	if viper.Get("klocctl.prometheus.port") == nil {
		viper.Set("klocctl.prometheus.port", "9090")
	}
}

func Config() {
	host = viper.GetString("klocctl.host")
	port = viper.GetString("klocctl.port")
	protocol = viper.GetString("klocctl.protocol")
	user = viper.GetString("klocctl.user")
	ltoken = viper.GetString("klocctl.ltoken")
	url = fmtURL(host, port, protocol)
	urlProm = fmtURL(viper.GetString("klocctl.prometheus.host"), viper.GetString("klocctl.prometheus.port"), protocol)
}
