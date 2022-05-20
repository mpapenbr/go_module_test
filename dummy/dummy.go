package dummy

import "fmt"

func GetVersion() string {
	return version;
}

func PrintVersion() {
	fmt.Printf("Version: %s\n", version)
}