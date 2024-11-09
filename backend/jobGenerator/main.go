package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// JobPost represents the structure of a job posting
type JobPost struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Position    string   `json:"position"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Type        string   `json:"type"`
	Posted      string   `json:"posted"`
	Location    string   `json:"location"`
	Skills      []string `json:"skills"`
	SalaryRange struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
	} `json:"salaryRange"`
	Equity struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"equity"`
	Perks []string `json:"perks"`
	Apply string   `json:"apply"`
}

// JobCategory represents a job category with associated positions, titles, and descriptions
type JobCategory struct {
	Position    []string
	Title       []string
	Description []string
	Skills      []string
	Perks       []string
}

// InitializeJobCategories initializes the job categories with their respective data
func InitializeJobCategories() map[string]JobCategory {
	return map[string]JobCategory{
		"Engineering": {
			Position: []string{
				"Software Engineer", "Data Scientist", "DevOps Engineer", "Machine Learning Engineer",
				"Backend Developer", "Frontend Developer", "Full Stack Developer", "Cloud Architect",
				"Database Administrator", "Cybersecurity Analyst", "Network Engineer",
			},
			Title: []string{
				"Senior Software Engineer", "Junior Data Scientist", "Lead DevOps Engineer",
				"Machine Learning Specialist", "Backend Developer", "Frontend Developer",
				"Full Stack Developer", "Cloud Solutions Architect", "Database Administrator",
				"Cybersecurity Analyst", "Network Infrastructure Engineer",
			},
			Description: []string{
				"Develop and maintain scalable web applications using modern frameworks.",
				"Analyze large datasets to extract actionable insights and drive decision-making.",
				"Implement and manage CI/CD pipelines to ensure efficient deployment processes.",
				"Design and deploy machine learning models to solve complex business problems.",
				"Build and optimize server-side components to ensure high performance and responsiveness.",
				"Create user-friendly interfaces and improve overall user experience.",
				"Handle both frontend and backend development tasks to deliver comprehensive solutions.",
				"Architect cloud-based solutions to enhance scalability and reliability.",
				"Manage and optimize database systems to ensure data integrity and availability.",
				"Ensure the security of systems by implementing robust cybersecurity measures.",
				"Design and maintain network infrastructure to support organizational operations.",
			},
			Skills: []string{
				"Go", "Python", "Java", "C++", "JavaScript", "SQL", "Cloud Computing",
				"DevOps", "Machine Learning", "Cybersecurity", "Docker", "Kubernetes",
			},
			Perks: []string{
				"Health Insurance", "401(k) Matching", "Stock Options", "Remote Work Options",
				"Flexible Schedule", "Gym Membership", "Professional Development",
			},
		},
		"Marketing": {
			Position: []string{
				"Marketing Manager", "Digital Marketer", "SEO Specialist", "Content Strategist",
				"Brand Manager", "Social Media Manager", "Public Relations Specialist",
			},
			Title: []string{
				"Senior Marketing Manager", "Digital Marketing Specialist", "SEO Analyst",
				"Content Marketing Manager", "Brand Strategist", "Social Media Coordinator",
				"Public Relations Coordinator",
			},
			Description: []string{
				"Develop and execute comprehensive marketing strategies to increase brand awareness.",
				"Manage digital marketing campaigns across various online platforms.",
				"Optimize website content to improve search engine rankings and drive organic traffic.",
				"Create and manage content strategies to engage target audiences.",
				"Oversee brand positioning and ensure consistency across all channels.",
				"Manage social media accounts and engage with the online community.",
				"Handle public relations activities to maintain a positive company image.",
			},
			Skills: []string{
				"Digital Marketing", "SEO", "Content Creation", "Social Media Management",
				"Brand Management", "Google Analytics", "Copywriting", "Email Marketing",
			},
			Perks: []string{
				"Health Insurance", "Paid Time Off", "Flexible Schedule", "Remote Work Options",
				"Professional Development", "Company Events",
			},
		},
		"Human Resources": {
			Position: []string{
				"HR Specialist", "Recruiter", "Training Coordinator", "HR Manager",
			},
			Title: []string{
				"Senior HR Specialist", "Technical Recruiter", "Employee Training Coordinator",
				"HR Operations Manager",
			},
			Description: []string{
				"Manage employee relations and ensure compliance with labor laws.",
				"Source, interview, and hire top talent to meet organizational needs.",
				"Coordinate and develop training programs to enhance employee skills.",
				"Oversee HR operations and implement effective human resource strategies.",
			},
			Skills: []string{
				"Recruitment", "Employee Relations", "Training & Development",
				"HR Management", "Conflict Resolution", "Performance Management",
			},
			Perks: []string{
				"Health Insurance", "Paid Time Off", "Retirement Plan", "Flexible Schedule",
				"Professional Development",
			},
		},
		"Finance": {
			Position: []string{
				"Financial Analyst", "Investment Banker", "Accountant", "Auditor",
				"Financial Planner", "Portfolio Manager",
			},
			Title: []string{
				"Senior Financial Analyst", "Investment Banking Associate", "Certified Public Accountant",
				"Internal Auditor", "Financial Planning Advisor", "Portfolio Manager",
			},
			Description: []string{
				"Analyze financial data to support strategic business decisions.",
				"Assist clients with investment strategies and financial planning.",
				"Manage financial records and ensure accurate reporting.",
				"Conduct audits to ensure compliance with financial regulations.",
				"Provide financial advice and planning services to clients.",
				"Manage investment portfolios to maximize returns for clients.",
			},
			Skills: []string{
				"Financial Modeling", "Budgeting", "Accounting", "Auditing",
				"Investment Strategies", "Risk Management", "Excel", "Financial Reporting",
			},
			Perks: []string{
				"Health Insurance", "401(k) Matching", "Bonuses", "Paid Time Off",
				"Professional Development",
			},
		},
		"Customer Service": {
			Position: []string{
				"Customer Service Representative", "Customer Support Specialist",
				"Call Center Agent", "Customer Success Manager",
			},
			Title: []string{
				"Senior Customer Service Representative", "Technical Support Specialist",
				"Call Center Supervisor", "Customer Success Manager",
			},
			Description: []string{
				"Assist customers with inquiries and resolve issues promptly.",
				"Provide technical support and troubleshoot customer problems.",
				"Manage call center operations and supervise customer service agents.",
				"Ensure customer satisfaction and foster long-term relationships with clients.",
			},
			Skills: []string{
				"Communication", "Problem Solving", "Customer Service", "Technical Support",
				"Conflict Resolution", "CRM Software",
			},
			Perks: []string{
				"Health Insurance", "Paid Time Off", "Flexible Schedule", "Remote Work Options",
				"Professional Development",
			},
		},
		// Add more categories as needed
	}
}

func main() {
	// Seed the math/rand package
	rand.Seed(time.Now().UnixNano())

	// Initialize job categories
	jobCategories := InitializeJobCategories()

	// Define other general data
	companies := []string{
		"Google", "Apple", "Microsoft", "Amazon", "Facebook", "IBM", "Oracle", "Intel",
		"Cisco", "Samsung", "Dell", "HP", "Adobe", "Salesforce", "Uber", "Airbnb",
		"Netflix", "Spotify", "Twitter", "LinkedIn", "Snapchat", "Pinterest", "PayPal",
		"eBay", "Shopify", "Slack", "Zoom", "Dropbox", "Reddit", "Square", "Nvidia",
		"Tesla", "SpaceX", "Boeing", "Nike", "Coca-Cola", "PepsiCo", "McDonald's",
		"Starbucks", "Walmart", "Target", "Disney", "Sony", "LG", "Philips", "Panasonic",
	}

	types := []string{"full-time", "part-time", "contract", "internship"}

	locations := []string{
		"US", "UK", "Canada", "Australia", "Germany", "France", "India", "China",
		"Japan", "Remote", "Brazil", "South Africa", "Mexico", "Spain", "Italy",
		"Netherlands", "Sweden", "Switzerland", "Singapore", "New Zealand",
	}

	currencies := []string{"USD", "EUR", "GBP", "CAD", "AUD", "JPY", "INR", "CNY"}

	// Initialize maps to ensure uniqueness
	urlMap := make(map[string]bool)
	idMap := make(map[string]bool)

	var jobPosts []JobPost

	// Get list of categories
	categoryKeys := make([]string, 0, len(jobCategories))
	for category := range jobCategories {
		categoryKeys = append(categoryKeys, category)
	}

	for i := 0; i < 1000; i++ {
		// Select a random category
		category := categoryKeys[rand.Intn(len(categoryKeys))]
		jobCategory := jobCategories[category]

		// Select random position, title, and description within the category
		position := jobCategory.Position[rand.Intn(len(jobCategory.Position))]
		title := jobCategory.Title[rand.Intn(len(jobCategory.Title))]
		description := jobCategory.Description[rand.Intn(len(jobCategory.Description))]

		// Generate a unique ObjectId
		var oid string
		for {
			oid = generateObjectId()
			if !idMap[oid] {
				idMap[oid] = true
				break
			}
		}

		// Select random company
		company := companies[rand.Intn(len(companies))]

		// Generate a unique URL
		var url string
		for {
			url = fmt.Sprintf("https://careers.%s.com/jobs/%s-%d", strings.ToLower(company), slugify(title), rand.Intn(100000))
			if !urlMap[url] {
				urlMap[url] = true
				break
			}
		}

		// Generate random apply URL
		applyURL := url + "/apply"

		// Create JobPost instance
		job := JobPost{}

		// Set ID
		job.ID.Oid = oid

		// Assign fields
		job.Position = position
		job.Title = title
		job.Description = description
		job.URL = url
		job.Type = types[rand.Intn(len(types))]
		job.Posted = randomDate()
		job.Location = locations[rand.Intn(len(locations))]

		// Assign skills based on category
		job.Skills = randomSelection(jobCategory.Skills, 3)

		// Assign salary range
		job.SalaryRange.From = rand.Intn(50000) + 50000              // From $50k to $100k
		job.SalaryRange.To = job.SalaryRange.From + rand.Intn(50000) // To $100k to $150k
		job.SalaryRange.Currency = currencies[rand.Intn(len(currencies))]

		// Assign equity
		job.Equity.From = roundFloat(rand.Float64() * 0.05)               // 0% to 5%
		job.Equity.To = roundFloat(job.Equity.From + rand.Float64()*0.05) // From equity to equity + 5%

		// Assign perks based on category
		job.Perks = randomSelection(jobCategory.Perks, 2)

		// Assign apply URL
		job.Apply = applyURL

		// Append to jobPosts slice
		jobPosts = append(jobPosts, job)
	}

	// Write jobPosts to JSON file
	file, err := os.Create("job_posts.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jobPosts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated 10,000 job posts and saved to job_posts.json")
}

// generateObjectId generates a valid 24-character hexadecimal MongoDB ObjectId
func generateObjectId() string {
	b := make([]byte, 12)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

// slugify converts a string to a URL-friendly slug
func slugify(text string) string {
	// Convert to lowercase
	slug := strings.ToLower(text)
	// Replace spaces and underscores with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	// Remove all non-alphanumeric and non-hyphen characters
	reg, _ := regexp.Compile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")
	// Remove multiple consecutive hyphens
	regHyphens, _ := regexp.Compile("-+")
	slug = regHyphens.ReplaceAllString(slug, "-")
	// Trim hyphens from the start and end
	slug = strings.Trim(slug, "-")
	return slug
}

// randomDate generates a random date between Jan 1, 2020 and today
func randomDate() string {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Now().Unix()
	sec := rand.Int63n(end-start) + start
	return time.Unix(sec, 0).Format("2006-01-02")
}

// randomSelection selects 'count' unique random elements from a list
func randomSelection(list []string, count int) []string {
	if count > len(list) {
		count = len(list)
	}
	shuffled := make([]string, len(list))
	copy(shuffled, list)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled[:count]
}

// roundFloat rounds a float to three decimal places
func roundFloat(f float64) float64 {
	return math.Round(f*1000) / 1000
}
