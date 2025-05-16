package resume

type Resume struct {
	Basics       Basics        `yaml:"basics"`
	Awards       []Award       `yaml:"awards"`
	Work         []Work        `yaml:"work"`
	Volunteer    []Volunteer   `yaml:"volunteer"`
	Education    []Education   `yaml:"education"`
	Certificates []Certificate `yaml:"certificates"`
	Publications []Publication `yaml:"publications"`
	Skills       []Skill       `yaml:"skills"`
	Languages    []Language    `yaml:"languages"`
	Interests    []Interest    `yaml:"interests"`
	References   []Reference   `yaml:"references"`
	Projects     []Project     `yaml:"projects"`
}

// Basics //
type Basics struct {
	Name     string    `yaml:"name"`
	Label    string    `yaml:"label"`
	Image    string    `yaml:"image"`
	Email    string    `yaml:"email"`
	Phone    string    `yaml:"phone"`
	Url      string    `yaml:"url"`
	Summary  string    `yaml:"summary"`
	Location Location  `yaml:"location"`
	Profiles []Profile `yaml:"profiles"`
}

type Location struct {
	Address     string `yaml:"address"`
	PostalCode  string `yaml:"postalCode"`
	City        string `yaml:"city"`
	CountryCode string `yaml:"countryCode"`
	Region      string `yaml:"region"`
}

type Profile struct {
	Network  string `yaml:"network"`
	Username string `yaml:"username"`
	Url      string `yaml:"url"`
}

// Work //
type Work struct {
	Name       string   `yaml:"name"`
	Position   string   `yaml:"position"`
	Url        string   `yaml:"url"`
	StartDate  string   `yaml:"startDate"`
	EndDate    string   `yaml:"endDate"`
	Summary    string   `yaml:"summary"`
	Highlights []string `yaml:"highlights"`
}

// Volunteer //
type Volunteer struct {
	Organization string   `yaml:"organization"`
	Position     string   `yaml:"position"`
	Url          string   `yaml:"url"`
	StartDate    string   `yaml:"startDate"`
	EndDate      string   `yaml:"endDate"`
	Summary      string   `yaml:"summary"`
	Highlights   []string `yaml:"highlights"`
}

// Education //
type Education struct {
	Institution string   `yaml:"institution"`
	Url         string   `yaml:"url"`
	Area        string   `yaml:"area"`
	StudyType   string   `yaml:"studyType"`
	StartDate   string   `yaml:"startDate"`
	EndDate     string   `yaml:"endDate"`
	Score       string   `yaml:"score"`
	Courses     []string `yaml:"courses"`
}

// Awards //
type Award struct {
	Title   string `yaml:"title"`
	Awarder string `yaml:"awarder"`
	Date    string `yaml:"date"`
	Summary string `yaml:"summary"`
}

// Certificates //
type Certificate struct {
	Name   string `yaml:"name"`
	Date   string `yaml:"date"`
	Issuer string `yaml:"issuer"`
	Url    string `yaml:"url"`
}

// Publications //
type Publication struct {
	Name        string `yaml:"name"`
	Publisher   string `yaml:"publisher"`
	Releasedate string `yaml:"releaseDate"`
	Url         string `yaml:"url"`
	Summary     string `yaml:"summary"`
}

// Skill //
type Skill struct {
	Name     string   `yaml:"name"`
	Level    string   `yaml:"level"`
	Keywords []string `yaml:"keywords"`
}

// Language //
type Language struct {
	Language string `yaml:"language"`
	Fluency  string `yaml:"fluency"`
}

// Interests //
type Interest struct {
	Name     string   `yaml:"name"`
	Keywords []string `yaml:"keywords"`
}

// References //
type Reference struct {
	Name      string `yaml:"name"`
	Reference string `yaml:"reference"`
}

// Projects //
type Project struct {
	Name        string   `yaml:"name"`
	StartDate   string   `yaml:"startDate"`
	EndDate     string   `yaml:"endDate"`
	Description string   `yaml:"description"`
	Highlights  []string `yaml:"highlights"`
	Url         string   `yaml:"url"`
}
