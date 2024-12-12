package seeder

import "fmt"

type SeederStruct struct{}

type SeederInterface interface {
	Execute() error
}

// Register the Seeders
func (s *SeederStruct) Seeding() []SeederInterface {
	return []SeederInterface{
		&RoleSeederStruct{},
		&MemberStatusSeederStruct{},
		&ReligionSeederStruct{},
	}
}

func (s *SeederStruct) RunSeeders() error {
	seeders := s.Seeding()

	for _, seeder := range seeders {
		if err := seeder.Execute(); err != nil {
			return fmt.Errorf("failed to execute seeder: %w", err)
		}
	}

	return nil
}
