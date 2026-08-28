package model

type Job struct {
	ID      int
	Numbers []int
}

type Result struct {
	JobID    int
	WorkerID int 
	Sum      int
}