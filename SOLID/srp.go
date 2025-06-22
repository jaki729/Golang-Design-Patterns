/*
Solid principles - Single Responsibility Principle (SRP)
The Single Responsibility Principle states that a class should have only one reason to change,
should have only one job or responsibility. 
This principle helps in reducing the complexity of the code
and makes it easier to maintain and understand. 
In Go, we can apply SRP by ensuring that each struct or package
has a single responsibility and does not mix different functionalities.
*/

package solid

type Survey struct {
	Title       string
	Questions  []string
}

func (s *Survey) GetTitle() string {
 return s.Title
}

func (s *Survey) Validate() bool {
	return len(s.Questions) > 0
}

// ========= SurveyService =========
// The SurveyService is responsible for managing surveys. 
// This will be in  a separate package, e.g., "surveyservice". Or in diffrent file.
type Repository interface {
	Save(survey *Survey) error	
}

type InMeoryRepository struct {
	surveys []*Survey
}

func (r *InMeoryRepository) Save(survey *Survey) error {
 r.surveys = append(r.surveys, survey)
 return nil
}

func SaveSurvey(survey *Survey, repo Repository) error {
 if !survey.Validate() {
  return nil // or return an error
 }
 return repo.Save(survey)
}