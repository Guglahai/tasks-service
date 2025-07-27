package task

type Service interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasks() ([]*Task, error)
	GetTaskByID(id uint) (*Task, error)
	GetTasksByUserID(userID uint) ([]*Task, error)
	UpdateTask(task *Task) (Task, error)
	DeleteTask(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTask(task *Task) (*Task, error) {
	return s.repo.CreateTask(task)
}

func (s *service) GetAllTasks() ([]*Task, error) {
	return s.repo.GetAllTasks()
}

func (s *service) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *service) GetTasksByUserID(userID uint) ([]*Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *service) UpdateTask(task *Task) (Task, error) {
	return s.repo.UpdateTask(task)
}

func (s *service) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
