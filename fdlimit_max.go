package fdlimit

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func MaxIt() error {

	if err := os.Setenv("GODEBUG", "netdns=go"); err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	limit, err := Maximum()
	if err != nil {
		return fmt.Errorf("failed to retrieve file descriptor allowance:%s", err)
	}
	_, err = Raise(uint64(limit))
	if err != nil {
		return fmt.Errorf("failed to raise file descriptor allowance:%s", err)
	}
	return nil
}
