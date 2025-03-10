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

// ScanPort of website with
// website IP address given.
func ScanPort(target string, workers int, timeout int, debug bool) {
	const totalPorts = 1024
	const batchSize = 100

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var openPorts []int
	var completedPorts int32
	limit := make(chan struct{}, workers)
	progressChan := make(chan int, totalPorts)

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
	batches := totalPorts / batchSize

	if totalPorts%batchSize != 0 {
		batches++
	}

	go func() {
		for progress := range progressChan {
			progressBar.Set(progress)
		}
	}()

	for batch := 0; batch < batches; batch++ {
		wg.Add(1)
		go func(batch int) {
			defer wg.Done()

			startPort := batch*batchSize + 1
			endPort := startPort + batchSize - 1

			if endPort > totalPorts {
				endPort = totalPorts
			}

			var portWG sync.WaitGroup
			currentTimeout := time.Duration(timeout) * time.Millisecond
			for port := startPort; port <= endPort; port++ {
				portWG.Add(1)

				go func(port int) {
					limit <- struct{}{}
					defer portWG.Done()
					defer func() { <-limit }()

					address := net.JoinHostPort(target, fmt.Sprint(port))
					timeNow := time.Now()
					connect, err := net.DialTimeout("tcp", address, currentTimeout)
					latency := time.Since(timeNow)

					if err != nil && debug {
						utils.Log("[DEBUG] %s", err)
					}

					if err == nil {
						mutex.Lock()
						openPorts = append(openPorts, port)
						mutex.Unlock()
						connect.Close()

						if latency < currentTimeout {
							currentTimeout = min(latency*2, time.Duration(timeout)*time.Millisecond)
						}
					}

					if !debug {
						atomic.AddInt32(&completedPorts, 1)
						progressChan <- int(atomic.LoadInt32(&completedPorts))
					}

				}(port)
			}
			portWG.Wait()
		}(batch)
	}
	wg.Wait()
	close(limit)

	if !debug {
		progressBar.Finish()
		close(progressChan)
	}

	for _, port := range openPorts {
		utils.Log("[+] Open Port: %d", port)
	}
}
