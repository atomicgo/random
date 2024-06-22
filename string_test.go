package random_test

import (
	"fmt"

	"fmt"

	"atomicgo.dev/random"
)

func ExampleString() {
	random.Seed(1337) // Set seed for deterministic output, not required

	fmt.Print(random.String(5, random.StringSetLowercase+random.StringSetUppercase))

	// Output: YsFVm
}

func ExampleStringSlice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.StringSlice(5, random.StringSetLowercase+random.StringSetUppercase, 5)
	fmt.Println(s)

	// Output: [YsFVm KAfdn AIdJS kxApO apmsT]
}
