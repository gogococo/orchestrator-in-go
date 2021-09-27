package scheduler

type Scheduler interface {
	SelectCandidateNotes()
	Score()
	Pick()
}
