// Sample app for the ground-truth tests. It pins a vulnerable dependency
// version ON PURPOSE. Never deploy it.
//
// GO-2023-2048 only affects Windows. This app is kept on its own so the
// ground truth for that policy question lives in one place.
package main

import (
	"fmt"

	securejoin "github.com/cyphar/filepath-securejoin"
)

func main() {
	path, err := securejoin.SecureJoin("/srv/data", "../etc/passwd") // ref:GO-2023-2048
	fmt.Println(path, err)
}
