package api

// WeightService contains the methods of the weight service
type WeightService interface {
	New(request NewWeightRequest) error
	CalculateBMR(height, age, weight int, sex string) (int, error)
	DailyIntake(BMR, activityLevel int, weightGoal string) (int, error)
}

// WeightRepository is what lets our service do db operations without knowing anything about the implementation
type WeightRepository interface {
	CreateWeightEntry(w Weight) error
	GetUser(userID int) (User, error)
}

type weightService struct {
	storage WeightRepository
}

func NewWeightService(weightRepo WeightRepository) WeightService {
	return &weightService{
		storage: weightRepo,
	}
}

// add these below NewWeightService
func (w *weightService) CalculateBMR(height, age, weight int, sex string) (int, error) {
	panic("implement me")
}

func (w *weightService) DailyIntake(BMR, activityLevel int, weightGoal string) (int, error) {
	panic("implement me")
}

func (w *weightService) New(request NewWeightRequest) error {
	panic("implement me")
}
