package internal

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/aqyuki/soreppoi/internal/strings"
)

func randomGenerator(content []string) string {
	return content[rand.Intn((len(content) - 1))]
}

func printLoop(ctx context.Context) {
	content := strings.Outputs

LOOP:
	for {
		fmt.Printf("%s\n", randomGenerator(content))
		select {
		case <-ctx.Done():
			return
		default:
			continue LOOP
		}
	}
}

// PrintSorrepoiStrings show sorrepoi output
func PrintSorrepoiStrings(seconds int32) error {

	if seconds <= 0 {
		return fmt.Errorf("specified value is too small, please specify 0 or more")
	}

	ctx, cancel := context.WithCancel(context.Background())
	t := time.After(time.Duration(seconds) * time.Second)
	go printLoop(ctx)
	<-t
	cancel()
	fmt.Println("Happy Hacking!!")

	return nil
}
