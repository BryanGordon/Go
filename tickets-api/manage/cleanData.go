
func CleanDataFile(ctx context.Context) {
	timer := time.NewTicker(5 * time.Minute)

	go func() {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				deleteData()
			case <-ctx.Done():
				return
			}
		}
	}()
}
