package main

import "fmt"

func main() {
	// // Define Map
	// emails := make(map[string]string)

	// // Assign Key Values
	// emails["Haha"] = "haha@domainname.domainsuffix"
	// emails["Lol"] = "lol@domainname.domainsuffix"
	// emails["Wow"] = "wow@domainname.domainsuffix"

	// Declare map and assign Key Values
	emails := map[string]string{"Haha": "haha@domainname.domainsuffix", "Lol": "lol@domainname.domainsuffix"}

	// Add Values
	emails["Wow"] = "wow@domainname.domainsuffix"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Lol"])

	// Delete from map
	delete(emails, "Lol")

	fmt.Println(emails)
}
