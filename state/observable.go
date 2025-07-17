package state

import "sync"

// ObservableBool represents a boolean value that notifies multiple components
// when its state changes. Perfect for scenarios like `isNodeActive`, where
// multiple goroutines (e.g., scheduler, cronjobs) need to react to changes
// in one shared state — without polling.
type ObservableBool struct {
	value       bool        // the current true/false state
	subscribers []chan bool // each component gets its own channel for updates
	mu          sync.Mutex  // ensures thread-safe access
}

// NewObservableBool creates a new observable boolean with an initial state.
// For example: isNodeActive := NewObservableBool(false)
func NewObservableBool(initial bool) *ObservableBool {
	return &ObservableBool{
		value:       initial,
		subscribers: make([]chan bool, 0),
	}
}

// Subscribe returns a new channel that will receive updates
// whenever the value changes (from false→true or true→false).
// Each cronjob or service should call this to get notified.
func (o *ObservableBool) Subscribe() <-chan bool {
	o.mu.Lock()
	defer o.mu.Unlock()

	ch := make(chan bool, 1) // buffer avoids blocking if subscriber is slow
	o.subscribers = append(o.subscribers, ch)
	return ch
}

// Set updates the observable value and notifies all subscribers,
// but only if the value actually changed (avoids redundant signals).
// This is how you activate or deactivate the shared state — e.g.,
// isNodeActive.Set(true) to notify cronjobs that the node is now active.
func (o *ObservableBool) Set(newVal bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.value == newVal {
		return // no state change → no need to notify
		//but we could notify on no state change probably don't want that
	}

	o.value = newVal
	for _, ch := range o.subscribers {
		select {
		case ch <- newVal:
			// notify the subscriber of the change
			// anyone that subscribed will trigger
		default:
			// skip if their channel is full — avoids blocking the loop
		}
	}
}

// Get returns the current value (true/false).
// just a useful handle if we want to check it outside of
// the reactive workflow we're aiming for
func (o *ObservableBool) Get() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.value
}
