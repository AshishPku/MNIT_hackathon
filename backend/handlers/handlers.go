package handlers

import (
	"easyjobBackend/db"
	"easyjobBackend/types"

	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobHandler struct {
	JobStore db.JobStore
}

func NewJobHandler(jobStore db.JobStore) *JobHandler {
	return &JobHandler{
		JobStore: jobStore,
	}
}

func (p *JobHandler) HandleInsertJob(c *fiber.Ctx) error {

	companyID, err := primitive.ObjectIDFromHex((c.Locals("id")).(string))
	if err != nil {
		companyID = primitive.NewObjectID()
	}

	// Parse Salary Range
	salaryFrom, _ := strconv.Atoi(c.FormValue("salaryFrom"))
	salaryTo, _ := strconv.Atoi(c.FormValue("salaryTo"))
	salaryRange := types.SalaryRange{
		From:     int32(salaryFrom),
		To:       int32(salaryTo),
		Currency: c.FormValue("currency"),
	}

	// Parse Equity
	equityFrom, _ := strconv.ParseFloat(c.FormValue("equityFrom"), 64)
	equityTo, _ := strconv.ParseFloat(c.FormValue("equityTo"), 64)
	equity := types.Equity{
		From: equityFrom,
		To:   equityTo,
	}

	// Parse other fields from form data
	job := types.Job{
		ID:          primitive.NewObjectID(),
		CompanyID:   companyID,
		Position:    c.FormValue("position"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		URL:         c.FormValue("url"),
		Type:        c.FormValue("type"),
		Posted:      time.Now().Format("2006-01-02"), // sets posted date as current date
		Location:    c.FormValue("location"),
		Skills:      strings.Split(c.FormValue("skills"), ","), // assumes skills are comma-separated in form
		SalaryRange: salaryRange,
		Equity:      equity,
		Perks:       strings.Split(c.FormValue("perks"), ","), // assumes perks are comma-separated in form
		Apply:       c.FormValue("apply"),
	}

	err = p.JobStore.InsertJob(c.Context(), &job)
	if err != nil {
		return err
	}
	return c.JSON(job)

}

func (p *JobHandler) HandleGetAllJobs(c *fiber.Ctx) error {
	posts, err := p.JobStore.GetAllJobs(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(posts)
}

////----------------------------------- build it

// func (p *JobHandler) HandleSearchUser(c *fiber.Ctx) error {
// 	var filterData types.Job
// 	filter := bson.D{}
// 	err := c.BodyParser(&filterData)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// if len(filterData.City) > 0 {
// 	// 	filter = append(filter, bson.E{Key: "city", Value: filterData.City})
// 	// }
// 	// if filterData.Rent > 0 {
// 	// 	filter = append(filter, bson.E{Key: "rent", Value: filterData.Rent})
// 	// }
// 	// if len(filterData.State) > 0 {
// 	// 	filter = append(filter, bson.E{Key: "state", Value: filterData.State})
// 	// }
// 	// if len(filterData.Type) > 0 {
// 	// 	filter = append(filter, bson.E{Key: "type", Value: filterData.Type})
// 	// }
// 	res, err := p.JobStore.GetJobByFilter(c.Context(), &filter)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(res)
// }

func (p *JobHandler) HandleGetJobsByFilter(c *fiber.Ctx) error {

	// Initialize an empty filter
	filter := bson.D{}

	// Retrieve query parameters
	companyIDParam := c.Query("companyID")
	position := c.Query("position")
	title := c.Query("title")
	jobType := c.Query("type")
	location := c.Query("location")
	skills := c.Query("skills")         // expects comma-separated skills
	salaryFrom := c.Query("salaryFrom") // expects integer value as string
	salaryTo := c.Query("salaryTo")     // expects integer value as string

	// Build dynamic filter based on provided query parameters

	// CompanyID filter (if valid ObjectID provided)
	if companyIDParam != "" {
		companyID, err := primitive.ObjectIDFromHex(companyIDParam)
		if err == nil {
			filter = append(filter, bson.E{Key: "companyID", Value: companyID})
		}
	}

	// Position filter with partial match
	if len(position) > 0 {
		filter = append(filter, bson.E{Key: "position", Value: bson.M{"$regex": position, "$options": "i"}})
	}

	// Title filter with partial match
	if len(title) > 0 {
		filter = append(filter, bson.E{Key: "title", Value: bson.M{"$regex": title, "$options": "i"}})
	}

	// Job type filter
	if len(jobType) > 0 {
		filter = append(filter, bson.E{Key: "type", Value: jobType})
	}

	// Location filter
	if len(location) > 0 {
		filter = append(filter, bson.E{Key: "location", Value: location})
	}

	// Skills filter (uses $in to match any skill from a list)
	if len(skills) > 0 {
		// skillArray := parseCommaSeparated(skills)
		// filter = append(filter, bson.E{Key: "skills", Value: bson.M{"$in": skillArray}})
		skillArray := parseCommaSeparated(skills)

		// Create an array of regex filters for each skill
		var skillRegexFilters []bson.M
		for _, skill := range skillArray {
			skillRegexFilters = append(skillRegexFilters, bson.M{"skills": bson.M{"$regex": skill, "$options": "i"}})
		}

		// Use $or to match any skill partially
		filter = append(filter, bson.E{Key: "$or", Value: skillRegexFilters})
	}

	// Salary range filter
	if salaryFrom != "" || salaryTo != "" {
		salaryFilter := bson.M{}
		if salaryFrom != "" {
			from, err := strconv.Atoi(salaryFrom)
			if err == nil {
				salaryFilter["$gte"] = from
			}
		}
		if salaryTo != "" {
			to, err := strconv.Atoi(salaryTo)
			if err == nil {
				salaryFilter["$lte"] = to
			}
		}
		if len(salaryFilter) > 0 {
			filter = append(filter, bson.E{Key: "salaryRange.from", Value: salaryFilter})
		}
	}

	// Execute the query with the built filter
	res, err := p.JobStore.GetJobByFilter(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch jobs"})
	}

	// Return the response
	return c.JSON(res)
}

// Helper function to parse comma-separated values into a slice of strings
func parseCommaSeparated(input string) []string {
	// Split the input string by commas and trim any surrounding whitespace
	values := strings.Split(input, ",")
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}

type MentorHandler struct {
	MentorStore db.MentorStore
}

func NewMentorHandler(MentorStore db.MentorStore) *MentorHandler {
	return &MentorHandler{
		MentorStore: MentorStore,
	}
}
func (p *MentorHandler) HandleGetAllMentors(c *fiber.Ctx) error {
	posts, err := p.MentorStore.GetAllMentors(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(posts)
}
