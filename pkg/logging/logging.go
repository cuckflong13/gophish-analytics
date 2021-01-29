package logging

import "fmt"

const ERROR = 0
const DEBUG = 1
const SUCCESS = 2

var Debug bool

func Log(status int, message string) {
	switch status {
	case ERROR:
		fmt.Printf("[-] %s\n", message)
	case DEBUG:
		if Debug {
			fmt.Printf("[!] %s\n", message)
		}
	case SUCCESS:
		fmt.Printf("[+] %s\n", message)
	}
}
