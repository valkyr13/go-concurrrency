package workerpool

type job struct {
	ID int
}

type result struct {
	job    job
	output string
}
