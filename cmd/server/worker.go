package main

func initWorkerQueue(nics []string) <-chan string {
	workerQueue := make(chan string)

	go func() {
		for _, nic := range nics {
			workerQueue <- nic
		}
	}()

	return workerQueue
}
