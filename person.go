package gofakeit

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// PersonInfo is a struct of person information
type PersonInfo struct {
	FirstName  string          `json:"first_name" xml:"first_name"`
	LastName   string          `json:"last_name" xml:"last_name"`
	Gender     string          `json:"gender" xml:"gender"`
	SSN        string          `json:"ssn" xml:"ssn"`
	EAMS       string          `json:"eams" xml:"eams"`
	ADJNumber  string          `json:"adj_number" xml:"adj_number"`
	Image      string          `json:"image" xml:"image"`
	Hobby      string          `json:"hobby" xml:"hobby"`
	Job        *JobInfo        `json:"job" xml:"job"`
	Address    *AddressInfo    `json:"address" xml:"address"`
	Contact    *ContactInfo    `json:"contact" xml:"contact"`
	CreditCard *CreditCardInfo `json:"credit_card" xml:"credit_card"`
}

// Person will generate a struct with person information
func Person() *PersonInfo { return person(globalFaker.Rand) }

// Person will generate a struct with person information
func (f *Faker) Person() *PersonInfo { return person(f.Rand) }

func person(r *rand.Rand) *PersonInfo {
	return &PersonInfo{
		FirstName:  firstName(r),
		LastName:   lastName(r),
		Gender:     gender(r),
		SSN:        ssn(r),
		Image:      imageURL(r, number(r, 100, 500), number(r, 100, 500)),
		Hobby:      hobby(r),
		Job:        job(r),
		Address:    address(r),
		Contact:    contact(r),
		CreditCard: creditCard(r),
	}
}

// Name will generate a random First and Last Name
func Name() string { return name(globalFaker.Rand) }

// Name will generate a random First and Last Name
func (f *Faker) Name() string { return name(f.Rand) }

func name(r *rand.Rand) string {
	return getRandValue(r, []string{"person", "first"}) + " " + getRandValue(r, []string{"person", "last"})
}

// FirstName will generate a random first name
func FirstName() string { return firstName(globalFaker.Rand) }

// FirstName will generate a random first name
func (f *Faker) FirstName() string { return firstName(f.Rand) }

func firstName(r *rand.Rand) string { return getRandValue(r, []string{"person", "first"}) }

// LastName will generate a random last name
func LastName() string { return lastName(globalFaker.Rand) }

// LastName will generate a random last name
func (f *Faker) LastName() string { return lastName(f.Rand) }

func lastName(r *rand.Rand) string { return getRandValue(r, []string{"person", "last"}) }

// NamePrefix will generate a random name prefix
func NamePrefix() string { return namePrefix(globalFaker.Rand) }

// NamePrefix will generate a random name prefix
func (f *Faker) NamePrefix() string { return namePrefix(f.Rand) }

func namePrefix(r *rand.Rand) string { return getRandValue(r, []string{"person", "prefix"}) }

// NameSuffix will generate a random name suffix
func NameSuffix() string { return nameSuffix(globalFaker.Rand) }

// NameSuffix will generate a random name suffix
func (f *Faker) NameSuffix() string { return nameSuffix(f.Rand) }

func nameSuffix(r *rand.Rand) string { return getRandValue(r, []string{"person", "suffix"}) }

// SSN will generate a random Social Security Number
func SSN() string { return ssn(globalFaker.Rand) }

// EAMS will generate a random eams Number
func EAMS() string { return eams(globalFaker.Rand) }

// ADJNumber will generate a random adjust number
func ADJNumber() string { return adjNumber(globalFaker.Rand) }

// SSN will generate a random Social Security Number
func (f *Faker) SSN() string { return ssn(f.Rand) }

// EAMS will generate a random eams Number
func (f *Faker) EAMS() string { return eams(f.Rand) }

// ADJNumber will generate a random adjust number
func (f *Faker) ADJNumber() string { return adjNumber(f.Rand) }

func adjNumber(r *rand.Rand) string {
	return fmt.Sprintf("ADJ%d", randIntRange(r, 100000000, 999999999))
}

func eams(r *rand.Rand) string { return strconv.Itoa(randIntRange(r, 100000000, 999999999)) }

func ssn(r *rand.Rand) string { return strconv.Itoa(randIntRange(r, 100000000, 999999999)) }

// Gender will generate a random gender string
func Gender() string { return gender(globalFaker.Rand) }

// Gender will generate a random gender string
func (f *Faker) Gender() string { return gender(f.Rand) }

func gender(r *rand.Rand) string {
	if boolFunc(r) {
		return "male"
	}

	return "female"
}

// Hobby will generate a random hobby string
func Hobby() string { return hobby(globalFaker.Rand) }

// Hobby will generate a random hobby string
func (f *Faker) Hobby() string { return hobby(f.Rand) }

func hobby(r *rand.Rand) string { return getRandValue(r, []string{"person", "hobby"}) }

// ContactInfo struct full of contact info
type ContactInfo struct {
	Phone string `json:"phone" xml:"phone"`
	Email string `json:"email" xml:"email"`
}

// Contact will generate a struct with information randomly populated contact information
func Contact() *ContactInfo { return contact(globalFaker.Rand) }

// Contact will generate a struct with information randomly populated contact information
func (f *Faker) Contact() *ContactInfo { return contact(f.Rand) }

func contact(r *rand.Rand) *ContactInfo {
	return &ContactInfo{
		Phone: phone(r),
		Email: email(r),
	}
}

// Phone will generate a random phone number string
func Phone() string { return phone(globalFaker.Rand) }

// Phone will generate a random phone number string
func (f *Faker) Phone() string { return phone(f.Rand) }

func phone(r *rand.Rand) string { return replaceWithNumbers(r, "##########") }

// PhoneFormatted will generate a random phone number string
func PhoneFormatted() string { return phoneFormatted(globalFaker.Rand) }

// PhoneFormatted will generate a random phone number string
func (f *Faker) PhoneFormatted() string { return phoneFormatted(f.Rand) }

func phoneFormatted(r *rand.Rand) string {
	return replaceWithNumbers(r, getRandValue(r, []string{"person", "phone"}))
}

// Email will generate a random email string
func Email() string { return email(globalFaker.Rand) }

// Email will generate a random email string
func (f *Faker) Email() string { return email(f.Rand) }

func email(r *rand.Rand) string {
	email := getRandValue(r, []string{"person", "first"}) + getRandValue(r, []string{"person", "last"})
	email += "@"
	email += getRandValue(r, []string{"person", "last"}) + "." + getRandValue(r, []string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}

// Teams takes in an array of people and team names and randomly places the people into teams as evenly as possible
func Teams(peopleArray []string, teamsArray []string) map[string][]string {
	return teams(globalFaker.Rand, peopleArray, teamsArray)
}

// Teams takes in an array of people and team names and randomly places the people into teams as evenly as possible
func (f *Faker) Teams(peopleArray []string, teamsArray []string) map[string][]string {
	return teams(f.Rand, peopleArray, teamsArray)
}

func teams(r *rand.Rand, people []string, teams []string) map[string][]string {
	// Shuffle the people if more than 1
	if len(people) > 1 {
		shuffleStrings(r, people)
	}

	peopleIndex := 0
	teamsOutput := make(map[string][]string)
	numPer := math.Ceil(float64(len(people)) / float64(len(teams)))
	for _, team := range teams {
		teamsOutput[team] = []string{}
		for i := 0.00; i < numPer; i++ {
			if peopleIndex < len(people) {
				teamsOutput[team] = append(teamsOutput[team], people[peopleIndex])
				peopleIndex++
			}
		}
	}

	return teamsOutput
}

func addPersonLookup() {
	AddFuncLookup("person", Info{
		Display:     "Person",
		Category:    "person",
		Description: "Random set of person info",
		Example: `{
			first_name: "Markus",
			last_name: "Moen",
			gender: "male",
			ssn: "420776036",
			image: "https://picsum.photos/300/300/people",
			hobby: "Swimming",
			job: {
				company: "Lockman and Sons",
				title: "Developer",
				descriptor: "Global",
				level: "Brand"
			}, 
			address: {
				address: "5369 Streamville, Rossieview, Hawaii 42591",
				street: "5369 Streamville",
				city: "Rossieview",
				state: "Hawaii",
				zip: "42591",
				country: "Burkina Faso",
				latitude: "-6.662594491850811",
				longitude: "23.921575244414612"
			},
			contact: {
				phone: "6136459948",
				email: "carolecarroll@bosco.com"
			},
			credit_card: {
				type: "Visa",
				number: "6536459948995369",
				exp: "03/27",
				cvv: "353"
			}
		}`,
		Output:      "map[string]interface",
		ContentType: "application/json",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return person(r), nil
		},
	})

	AddFuncLookup("name", Info{
		Display:     "Name",
		Category:    "person",
		Description: "Random name",
		Example:     "Markus Moen",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return name(r), nil
		},
	})

	AddFuncLookup("nameprefix", Info{
		Display:     "Name Prefix",
		Category:    "person",
		Description: "Random name prefix",
		Example:     "Mr.",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return namePrefix(r), nil
		},
	})

	AddFuncLookup("namesuffix", Info{
		Display:     "Name Suffix",
		Category:    "person",
		Description: "Random name suffix",
		Example:     "Jr.",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return nameSuffix(r), nil
		},
	})

	AddFuncLookup("firstname", Info{
		Display:     "First Name",
		Category:    "person",
		Description: "Random first name",
		Example:     "Markus",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return firstName(r), nil
		},
	})

	AddFuncLookup("lastname", Info{
		Display:     "Last Name",
		Category:    "person",
		Description: "Random last name",
		Example:     "Daniel",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return lastName(r), nil
		},
	})

	AddFuncLookup("gender", Info{
		Display:     "Gender",
		Category:    "person",
		Description: "Random gender",
		Example:     "male",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return gender(r), nil
		},
	})

	AddFuncLookup("ssn", Info{
		Display:     "SSN",
		Category:    "person",
		Description: "Random social security number",
		Example:     "296446360",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return ssn(r), nil
		},
	})

	AddFuncLookup("hobby", Info{
		Display:     "Hobby",
		Category:    "person",
		Description: "Random hobby activity",
		Example:     "Swimming",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return hobby(r), nil
		},
	})

	AddFuncLookup("email", Info{
		Display:     "Email",
		Category:    "person",
		Description: "Random email",
		Example:     "markusmoen@pagac.net",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return email(r), nil
		},
	})

	AddFuncLookup("phone", Info{
		Display:     "Phone",
		Category:    "person",
		Description: "Random phone number",
		Example:     "6136459948",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return phone(r), nil
		},
	})

	AddFuncLookup("phoneformatted", Info{
		Display:     "Phone Formatted",
		Category:    "person",
		Description: "Random formatted phone number",
		Example:     "136-459-9489",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return phoneFormatted(r), nil
		},
	})

	AddFuncLookup("teams", Info{
		Display:     "Teams",
		Category:    "person",
		Description: "Randomly split people into teams",
		Example:     `{"Team 1": ["Sharon","Jeff"], "Team 2": ["Billy","Connor"]}`,
		Output:      "map[string][]string",
		Params: []Param{
			{Field: "people", Display: "Strings", Type: "[]string", Description: "Array of people"},
			{Field: "teams", Display: "Strings", Type: "[]string", Description: "Array of teams"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			people, err := info.GetStringArray(m, "people")
			if err != nil {
				return nil, err
			}

			teamsArray, err := info.GetStringArray(m, "teams")
			if err != nil {
				return nil, err
			}

			return teams(r, people, teamsArray), nil
		},
	})
}
