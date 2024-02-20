package symwalker

import (
	"fmt"
)

func walkError(e WalkErr) string {
	fmt.Println("Error: ", e)
	return e.Error()
}
