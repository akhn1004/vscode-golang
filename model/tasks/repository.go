package tasks

type Repository interface {
	Add(Task) int
	List() []*Task
	Done(id int) error
}
