package main

import (
	"log"
	"time"

	"observerdemo/state" // Import the ObservableBool type
)

// Create a shared observable boolean value named isNodeActive.
// This will be watched by multiple components like the scheduler and cronjob.
var isNodeActive = state.NewObservableBool(false)

func main() {
	// Sart the background subscribers.
	// each will lissten for changes to isNodeActive via their own channel.
	startScheduler()
	launchCronA()

	// Simulate a delay before activating the node.
	// This mimics a real-world condition change (e.g., service becomes healthy).
	log.Println("Simulate node becomes active: Waiting 3 seconds before activating node...")
	time.Sleep(3 * time.Second)
	isNodeActive.Set(true) // This triggers all subscribers with "true"

	// Give subscribers time to react before changing state again
	time.Sleep(2 * time.Second)

	// Now simulate the node going inactive
	log.Println(" Simulate node deactivates: Deactivating node...")
	isNodeActive.Set(false) // Subscribers will now be notified with "false"

	// Final sleep to let reactions log before main exits
	time.Sleep(2 * time.Second)
}

// startScheduler creates a goroutine that listens for state changes to isNodeActive.
// It reacts by starting or stopping scheduled jobs.
func startScheduler() {
	ch := isNodeActive.Subscribe() // Get a unique channel for this subscriber

	go func() {
		for val := range ch {
			if val {
				log.Println(" Scheduler: Node is active — starting jobs")
				startScheduledJobs()
			} else {
				log.Println("Scheduler: Node is inactive — stopping jobs")
				stopScheduledJobs()
			}
		}
	}()
}

// Simulated job-start logic
func startScheduledJobs() {
	log.Println("[Scheduler] Jobs started.")
}

// Simulated job-stop logic
func stopScheduledJobs() {
	log.Println("[Scheduler] Jobs stopped.")
}

// launchCronA is anothR subscriber, representing a different reacting component.
// It only reacts when the node becomes active.
func launchCronA() {
	ch := isNodeActive.Subscribe() // Get a unique channel for this subscriber

	go func() {
		for val := range ch {
			if val {
				log.Println("CronA: Node is active — running playbook")
				runMyPlaybook()
			}
			// No action needed on false
		}
	}()
}

// Simulated cronjob logic
func runMyPlaybook() {
	log.Println("[CronA] Playbook A running...")
}
