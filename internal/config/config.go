package config

import (
	"flag" //flag is the package that contains the flag package
	"log"
	"os" //os is the package that contains all the os functions

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct { //we added this struct to hold all the config values
	Addr string  `yaml:"address" env-required:"true"`
}

// env-default:"production"

type Config struct { //we added this struct to hold all the config values
	Env         string `yaml:"env" env:"ENV" env-required:"true" `
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config { //we added this function to load the config file
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file") // "config" is the name of the flag and "" is the default value for the flag if it is not passed
		flag.Parse()  //parse means that it will look for the flag in the command line
			configPath = *flags    //pointer to flag

			if configPath == "" {
				log.Fatal("config path is not setup")
			}
	}

	if _,err := os.Stat(configPath); os.IsNotExist(err) {   // if _, err := os.Stat(configPath); os.IsNotExist(err) is a condition that checks if the file exists
		log.Fatalf("config file does not exist: %s", configPath)  // if the file does not exist, it will log the error , we used fatal because we want the program to stop
	}

		var cfg Config //we added this var to hold the config values	

	err := cleanenv.ReadConfig(configPath, &cfg) //ReadConfig is a function that reads the config file and fills the struct
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	return &cfg
}