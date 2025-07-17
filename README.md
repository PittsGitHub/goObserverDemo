# goObserverDemo

quick example of the observer pattern

go run main.go

output:
2025/07/17 16:07:45 Simulate node becomes active: Waiting 3 seconds before activating node...
2025/07/17 16:07:48 CronA: Node is active — running playbook
2025/07/17 16:07:48 [CronA] Playbook A running...
2025/07/17 16:07:48 Scheduler: Node is active — starting jobs
2025/07/17 16:07:48 [Scheduler] Jobs started.
2025/07/17 16:07:50 Simulate node deactivates: Deactivating node...
2025/07/17 16:07:50 Scheduler: Node is inactive — stopping jobs
2025/07/17 16:07:50 [Scheduler] Jobs stopped.
