package main

import (
	"fmt"
	"log"

	"os"
	"path/filepath"

	"github.com/huynhminhtruong/go-instagram-downloader/internal/config"
)

func main() {
	cookiePath, err := getCookiePath("config/cookies.yaml")
	if err != nil {
		log.Fatal("❌ Get cookie path failed:", err)
	}

	cfg, err := config.LoadCookieConfig(cookiePath)
	if err != nil {
		log.Fatal("❌ Load cookie YAML failed:", err)
	}

	fmt.Println("✅ Cookies loaded successfully")
	fmt.Print(cfg.Cookies)

	// postURL := "https://www.instagram.com/p/DK_KmfZp5LQ/"
	// images, err := chromedphelper.StartSessionWithCookies(postURL, cfg.Cookies)
	// if err != nil {
	// 	log.Fatal("❌ Chromedp session failed:", err)
	// }

	// fmt.Println("✅ Image URLs found:")
	// for i, img := range images {
	// 	fmt.Printf("[%d] %s\n", i+1, img)
	// }
}

func getCookiePath(fname string) (string, error) {
	// Get the path of the executable file
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, fname), nil
}
