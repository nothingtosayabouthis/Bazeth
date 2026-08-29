package dns

import "fmt"

func PrintSection(name string, values []string) {
	if len(values) == 0 {
		return
	}

	fmt.Printf("\n[%s]\n", name)

	for _, value := range values {
		fmt.Printf("[*] %s\n", value)
	}
}
