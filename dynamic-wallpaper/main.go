package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const configFile = "config.json"

const (
	SPI_SETDESKWALLPAPER = 20
	SPIF_UPDATEINIFILE   = 0x01
	SPIF_SENDCHANGE      = 0x02
)

type Wallpaper struct {
	Time string `json:"time"`
	File string `json:"file"`
}

type Config struct {
	Wallpapers []Wallpaper `json:"wallpapers"`
}

var (
	user32                = windows.NewLazySystemDLL("user32.dll")
	systemParametersInfoW = user32.NewProc("SystemParametersInfoW")
)

func main() {

	fmt.Println("Dynamic Wallpaper started")

	last_wallpaper := ""
	for {
		config, err := loadConfig()
		if err != nil {
			fmt.Println("Failed to load config:", err)
			return
		}
		now := time.Now()

		for _, wallpaper := range config.Wallpapers {
			// fmt.Printf("\nChecking Time(%s) and Wallpaper (%s , %s)", now.Format("15:04"), wallpaper.Time, wallpaper.File)
			if now.Format("15:04") == wallpaper.Time && last_wallpaper != wallpaper.File {
				if err := setWallpaper(wallpaper.File); err != nil {
					fmt.Println("Failed to set wallpaper:", err)
				} else {
					last_wallpaper = wallpaper.File
					fmt.Println("Wallpaper changed:", wallpaper.File)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func loadConfig() (Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func setWallpaper(file string) error {
	absolutePath, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	// Make sure the file actually exists.
	if _, err := os.Stat(absolutePath); err != nil {
		return fmt.Errorf("wallpaper not found: %s", absolutePath)
	}

	fmt.Println("Setting wallpaper:", absolutePath)

	path, err := windows.UTF16PtrFromString(absolutePath)
	if err != nil {
		return err
	}

	ret, _, callErr := systemParametersInfoW.Call(
		uintptr(SPI_SETDESKWALLPAPER),
		0,
		uintptr(unsafe.Pointer(path)),
		uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE),
	)

	if ret == 0 {
		return fmt.Errorf("SystemParametersInfoW failed: %w", callErr)
	}

	return nil
}
