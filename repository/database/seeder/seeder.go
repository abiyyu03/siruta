package seeder

import "fmt"

type SeederStruct struct{}

type SeederInterface interface {
	Execute() error
}

// Register the Seeders
func (s *SeederStruct) seedingProvider() []SeederInterface {
	return []SeederInterface{
		&AdminSeederStruct{},
		&RoleSeederStruct{},
		&MemberStatusSeederStruct{},
		&ReligionSeederStruct{},
		&LetterTypeSeeder{},
	}
}

func (s *SeederStruct) RunSeeders() error {
	seeders := s.seedingProvider()

	for _, seeder := range seeders {
		if err := seeder.Execute(); err != nil {
			return fmt.Errorf("failed to execute seeder: %w", err)
		}
	}
	fmt.Println("All seeders executed successfully!")

	return nil
}
