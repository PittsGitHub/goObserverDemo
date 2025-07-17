# goObserverDemo

Quick example of the observer pattern in Go using channels.
The idea is to create an object (ObservableBool) that holds:
- a boolean value,
- a list of subscriber channels,
- and a mutex to ensure thread-safe access.

Any number of functions or components can subscribe to this object.
When the value changes via .Set(true/false) functions, then all subscribers are notified on their individual channels 

go run main.go
output:
```
2025/07/17 16:07:45 Simulate node becomes active: Waiting 3 seconds before activating node...
2025/07/17 16:07:48 CronA: Node is active — running playbook
2025/07/17 16:07:48 [CronA] Playbook A running...
2025/07/17 16:07:48 Scheduler: Node is active — starting jobs
2025/07/17 16:07:48 [Scheduler] Jobs started.
2025/07/17 16:07:50 Simulate node deactivates: Deactivating node...
2025/07/17 16:07:50 Scheduler: Node is inactive — stopping jobs
2025/07/17 16:07:50 [Scheduler] Jobs stopped.
``
