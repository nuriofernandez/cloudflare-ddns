package scheduler

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Schedule(executableTask func()) {
	repeatInterval := parseMinutes()
	if repeatInterval != 0 {
		fmt.Printf("[DDNS] [Scheduler] Running every %d minutes...\n", repeatInterval)
	}

	// Execute task (both once & scheduled)
	executableTask()

	// Setup scheduler
	if repeatInterval != 0 {
		ticker := time.NewTicker(time.Duration(repeatInterval) * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			executableTask()
			fmt.Printf("[DDNS] [Scheduler] Next update will be in %d minutes...\n", repeatInterval)
		}
	}
}

func parseMinutes() int {
	var intervalMinutes = os.Getenv("DDNS_REFRESH_INTERVAL_MINUTES")
	if intervalMinutes == "" {
		return 0
	}

	// Parse string to minutes
	minutes, err := strconv.Atoi(intervalMinutes)
	if err != nil || minutes <= 0 {
		fmt.Printf("[DDNS] [Scheduler] Invalid interval '%s'.\n", intervalMinutes)
		return 0
	}

	return minutes
}
