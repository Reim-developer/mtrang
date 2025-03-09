package core

import (
	"fmt"
	"mtrang/utils"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Scan port of website with
// website IP address given.
func ScanPort(target string, workers int, timeout int, debug bool) {
	var wg sync.WaitGroup
	limit := make(chan struct{}, workers)
	var mu sync.Mutex

	progressBar := progressbar.NewOptions(
		1024,
		progressbar.OptionSetDescription("Scanning Ports"),
		progressbar.OptionFullWidth(),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerHead:    "█",
			SaucerPadding: "─",
			BarStart:      "|",
			BarEnd:        "|",
		}),
	)

	resultChan := make(chan int, 1024)
	debugInfoChan := make(chan string, 1024)

	go func() {
		for info := range debugInfoChan {
			if debug {
				utils.Log("%s", info)
			}
		}
	}()

	var completedPorts int32
	for port := 1; port <= 1024; port++ {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()
			limit <- struct{}{}

			address := net.JoinHostPort(target, fmt.Sprintf("%d", port))
			connect, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Millisecond)

			if err != nil {
				debugInfoChan <- fmt.Sprintf("[DEBUG] %s\n", err)
			}

			if err == nil {
				mu.Lock()
				resultChan <- port
				mu.Unlock()
				connect.Close()
			}

			if !debug {
				atomic.AddInt32(&completedPorts, 1)
				progressBar.Set(int(atomic.LoadInt32(&completedPorts)))
			}

			<-limit
		}(port)
	}

	wg.Wait()
	close(resultChan)
	close(debugInfoChan)

	for port := range resultChan {
		utils.Log("[+] Open Port: %d", port)
	}
}
