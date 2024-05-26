package repository

import (
	"mobile/internal/app/models"
	"sort"
	"time"

	"gorm.io/gorm"
)

type ScheduleRespository struct {
	DB *gorm.DB
}

func NewScheduleRespository(db *gorm.DB) *ScheduleRespository {
	return &ScheduleRespository{DB: db}
}

func (r *ScheduleRespository) GetPrepodsCount(prepods models.Prepod) int {
	var count int64
	r.DB.Model(prepods).Count(&count)
	return int(count)
}

func (r *ScheduleRespository) GetGroupsCount(group models.Group) int {
	var count int64
	r.DB.Model(group).Count(&count)
	return int(count)
}

func (r *ScheduleRespository) GetUrokiCount(urok models.Urok) int {
	var count int64
	r.DB.Model(urok).Count(&count)
	return int(count)
}

func (r *ScheduleRespository) GetSubjectsCount(subjects models.Subject) int {
	var count int64
	r.DB.Model(subjects).Count(&count)
	return int(count)
}

func (r *ScheduleRespository) GetPrepodsList() map[uint]models.Prepod {
	prepods := []models.Prepod{}
	r.DB.Order("id").Find(&prepods)
	prepodsMap := make(map[uint]models.Prepod)
	for _, prepod := range prepods {
		prepodsMap[prepod.ID] = prepod
	}
	return prepodsMap
}

func (r *ScheduleRespository) GetChangeByID(id int) (models.Change, error) {
	var change models.Change
	result := r.DB.Where("id = ?", id).First(&change)
	if result.Error != nil {
		return models.Change{}, result.Error
	}
	prepod := models.Prepod{}
	r.DB.First(&prepod, change.PrepodID)
	subject := models.Subject{}
	r.DB.First(&subject, change.SubjectID)
	output := models.Change{
		ID:        change.ID,
		IsChange:  true,
		Date:      change.Date,
		GroupID:   change.GroupID,
		Urok:      change.Urok,
		Classroom: change.Classroom,
		Prepod:    prepod,
		Subject:   subject,
		Comment:   change.Comment,
	}
	return output, nil
}

func (r *ScheduleRespository) GetUrokByID(id int) (models.Urok, error) {
	var urok models.Urok
	result := r.DB.Where("id = ?", id).First(&urok)
	if result.Error != nil {
		return models.Urok{}, result.Error
	}
	prepod := models.Prepod{}
	r.DB.First(&prepod, urok.PrepodID)
	subject := models.Subject{}
	r.DB.First(&subject, urok.SubjectID)
	output := models.Urok{
		ID:        urok.ID,
		IsChange:  false,
		Date:      urok.Date,
		GroupID:   urok.GroupID,
		Number:    urok.Number,
		Classroom: urok.Classroom,
		Prepod:    prepod,
		Subject:   subject,
	}
	return output, nil
}

func (r *ScheduleRespository) GetChangesByDay(date time.Time) ([]models.Change, error) {
	var changes []models.Change
	result := r.DB.Where("date = ?", date.Format("2006-01-02")).Order("GroupId").Find(&changes)
	if result.Error != nil {
		return []models.Change{}, result.Error
	}

	for i, change := range changes {
		var prepodRes models.Prepod
		r.DB.Model(&models.Prepod{}).Select("Surname, Firstname, SecondName").Where("id = ?", change.PrepodID).Row().Scan(&prepodRes)

		subject := models.Subject{}
		r.DB.First(&subject, change.SubjectID)

		changes[i].Prepod = prepodRes
		changes[i].Subject = subject
	}

	return changes, nil
}

func (r *ScheduleRespository) GetAllSubjects() ([]string, error) {
	var subjects []models.Subject
	result := r.DB.Find(&subjects)
	if result.Error != nil {
		return []string{}, result.Error
	}

	var subjectNames []string
	for _, subject := range subjects {
		subjectNames = append(subjectNames, subject.NAME)
	}
	sort.Strings(subjectNames)

	return subjectNames, nil
}

func (r *ScheduleRespository) GetAllPrepods() ([]string, error) {
	var prepods []models.Prepod
	result := r.DB.Order("id").Find(&prepods)
	if result.Error != nil {
		return []string{}, result.Error
	}

	var prepodNames []string
	for _, prepod := range prepods {
		prepodNames = append(prepodNames, prepod.Surname+" "+prepod.FirstName+" "+prepod.SecondName)
	}

	return prepodNames, nil
}

func (r *ScheduleRespository) GetAllGroups() ([]string, error) {
	var groups []models.Group
	result := r.DB.Order("Name").Find(&groups)
	if result.Error != nil {
		return []string{}, result.Error
	}

	var groupNames []string
	for _, group := range groups {
		groupNames = append(groupNames, group.NAME)
	}
	sort.Strings(groupNames)

	return groupNames, nil
}

func (r *ScheduleRespository) GetSubjectByID(id int) (string, error) {
	var subject models.Subject
	result := r.DB.Where("id = ?", id).First(&subject)
	if result.Error != nil {
		return "", result.Error
	}

	return subject.NAME, nil
}

func (r *ScheduleRespository) GetChangesForPrepodByDate(date time.Time, prepodID int) map[int]interface{} {
	var changes []models.Change
	r.DB.Where("date = ? AND prepod_id = ?", date.Format("2006-01-02"), prepodID).Find(&changes)

	var uroki []models.Urok
	r.DB.Where("date = ? AND prepod_id = ?", date.Format("2006-01-02"), prepodID).Order("Number").Find(&uroki)

	output := make(map[int]interface{})

	for _, change := range changes {
		changeOutput := map[string]interface{}{
			"IsChange":  true,
			"Date":      change.Date,
			"Group":     change.GroupID,
			"Urok":      change.Urok,
			"Classroom": change.Classroom,
			"Prepod":    change.Prepod,  // Assuming that Prepod is a field in models.Change
			"Subject":   change.Subject, // Assuming that Subject is a field in models.Change
			"Comment":   change.Comment,
		}
		output[change.Number] = changeOutput
	}

	for _, urok := range uroki {
		urokOutput := map[string]interface{}{
			"IsChange":  false,
			"Subject":   urok.Subject, // Assuming that Subject is a field in models.Urok
			"Prepod":    urok.Prepod,  // Assuming that Prepod is a field in models.Urok
			"Group":     urok.GroupID,
			"Date":      urok.Date,
			"Number":    urok.Number,
			"Classroom": urok.Classroom,
		}
		output[urok.Number] = urokOutput
	}

	for _, change := range changes {
		if change.DeleteUrok != nil {
			delete(output, *change.DeleteUrok)
		}
	}

	for _, urok := range uroki {
		if urokMap, ok := output[urok.Number].(map[string]interface{}); ok {
			if urok.Subject != urokMap["Subject"] {
				delete(output, urok.Number)
			}
		}
	}

	return output
}

func (r *ScheduleRespository) GetUrokByDateNumberAndGroup(date time.Time, group string, number int) (models.Urok, error) {
	var urok models.Urok
	result := r.DB.Where("Date = ? AND Group = ? AND Number = ?", date.Format("2006-01-02"), group, number).First(&urok)
	if result.Error != nil {
		return models.Urok{}, result.Error
	}
	return urok, nil
}

func (r *ScheduleRespository) GetChangesForGroupByDate(date time.Time, group string) ([]models.Change, error) {
	var changes []models.Change
	result := r.DB.Where("Date = ? AND Group = ?", date.Format("2006-01-02"), group).Order("Group").Find(&changes)
	if result.Error != nil {
		return []models.Change{}, result.Error
	}
	return changes, nil
}

func (r *ScheduleRespository) GetScheduleWithChangesDay(date time.Time, group string) map[int][]interface{} {
	var changes []models.Change
	r.DB.Where("Date = ? AND Group = ?", date.Format("2006-01-02"), group).Find(&changes)

	var uroki []models.Urok
	r.DB.Where("Date = ? AND Group = ?", date.Format("2006-01-02"), group).Order("Number").Find(&uroki)

	output := make(map[int][]interface{})

	for _, change := range changes {
		changeOutput := map[string]interface{}{
			"Date":      change.Date,
			"Group":     change.GroupID,
			"Number":    change.Number,
			"Classroom": change.Classroom,
			"Prepod":    change.Prepod,  // Assuming that Prepod is a field in models.Change
			"Subject":   change.Subject, // Assuming that Subject is a field in models.Change
			"IsChange":  true,
		}
		output[change.Number] = append(output[change.Number], changeOutput)
	}

	for _, urok := range uroki {
		urokOutput := map[string]interface{}{
			"Subject":   urok.Subject, // Assuming that Subject is a field in models.Urok
			"Prepod":    urok.Prepod,  // Assuming that Prepod is a field in models.Urok
			"Group":     urok.GroupID,
			"Date":      urok.Date,
			"Number":    urok.Number,
			"Classroom": urok.Classroom,
			"IsChange":  false,
		}
		output[urok.Number] = append(output[urok.Number], urokOutput)
	}

	return output
}

func (r *ScheduleRespository) RemoveOldSchedule() error {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	result := r.DB.Where("Date < ?", sevenDaysAgo).Delete(&models.Urok{})
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Where("Date < ?", sevenDaysAgo).Delete(&models.Change{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ScheduleRespository) AddChange(subjectID int, prepodID int, group string, date time.Time, number int, classroom, comment string, delete *int) error {
	change := models.Change{
		SubjectID:  uint(subjectID),
		PrepodID:   uint(prepodID),
		GroupID:    group,
		Date:       date,
		Number:     number,
		Classroom:  classroom,
		Comment:    comment,
		DeleteUrok: delete,
	}
	result := r.DB.Create(&change)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScheduleRespository) AddUrok(subjectID int, prepodID int, group string, date time.Time, number int, classroom string) error {
	urok := models.Urok{
		SubjectID: uint(subjectID),
		PrepodID:  uint(prepodID),
		GroupID:   group,
		Date:      date,
		Number:    number,
		Classroom: classroom,
	}
	result := r.DB.Create(&urok)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScheduleRespository) GetAllClassrooms() ([]string, error) {
	var classrooms []string
	result := r.DB.Model(&models.Urok{}).Select("DISTINCT Classroom").Pluck("Classroom", &classrooms)
	if result.Error != nil {
		return []string{}, result.Error
	}
	sort.Strings(classrooms)
	return classrooms, nil
}

func (r *ScheduleRespository) DeleteChange(id int) error {
	result := r.DB.Delete(&models.Change{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScheduleRespository) AddSubject(name string) error {
	subject := models.Subject{
		NAME: name,
	}
	result := r.DB.Create(&subject)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScheduleRespository) AddPrepod(surname, firstName, secondName string) error {
	prepod := models.Prepod{
		Surname:    surname,
		FirstName:  firstName,
		SecondName: secondName,
	}
	result := r.DB.Create(&prepod)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScheduleRespository) AddGroup(group string) error {
	groupM := models.Group{
		NAME: group,
	}

	result := r.DB.Create(&groupM)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
