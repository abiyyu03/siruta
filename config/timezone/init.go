package timezone

import "os"

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
}
