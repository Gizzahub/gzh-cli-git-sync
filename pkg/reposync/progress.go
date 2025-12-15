package reposync

// ProgressSink receives progress events from the executor.
type ProgressSink interface {
	OnStart(action Action)
	OnProgress(action Action, message string, progress float64)
	OnComplete(result ActionResult)
}
