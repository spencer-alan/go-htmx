package common

import "github.com/BurntSushi/toml"

type Navbar struct {
	Links []Link `toml:"links"`
}

type AppSettings struct {
	DatabaseAddress  string `toml:"database_address"`
	DatabasePort     int    `toml:"database_port"`
	DatabaseUser     string `toml:"database_user"`
	DatabasePassword string `toml:"database_password"`
	DatabaseName     string `toml:"databse_name"`
	WebserverPort    int    `toml:"webserver_port"`
	AdminPort        int    `toml:"admin_port"` // NOTE: probably get rid of this at some point? Cool idea but I do not think I will nee it
	ImageDirectory   string `toml:"image_dir"`
	CacheEnabled     bool   `toml:"cache_bool"`
	AppNavbar        Navbar `toml:"navbar"`
}

func ReadConfigToml(filepath string) (AppSettings, error) {
	var config AppSettings
	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		return AppSettings{}, err
	}

	return config, nil
}
