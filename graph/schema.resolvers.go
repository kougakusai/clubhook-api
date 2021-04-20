package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/kougakusaiHPTeam/clubhook-api/graph/generated"
	"github.com/kougakusaiHPTeam/clubhook-api/graph/model"
	"github.com/kougakusaiHPTeam/clubhook-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user := models.User{
		Name:    input.Name,
		Email:   input.Email,
		Grade:   input.Grade,
		GroupID: input.GroupID,
	}
	res := r.DB.Create(&user)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.EditUser) (*models.User, error) {
	var user models.User
	res := r.DB.First(&user, input.UserID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&user).Updates(models.User{
		Name:    *input.Name,
		Email:   *input.Email,
		Grade:   *input.Grade,
		GroupID: *input.GroupID,
	})
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).First(&user, input.UserID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	res := r.DB.First(&user, userID)
	if err := res.Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&user).Association("Events").Clear(); err != nil {
		return nil, err
	}
	res = r.DB.Delete(&user, userID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (*models.Group, error) {
	group := models.Group{
		Name: input.Name,
	}
	res := r.DB.Create(&group)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).First(&group)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, input model.EditGroup) (*models.Group, error) {
	var group models.Group
	res := r.DB.First(&group, input.GroupID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&group).Update("Name", input.Name)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).First(&group, input.GroupID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, groupID uint) (*models.Group, error) {
	var group models.Group
	res := r.DB.Delete(&group, groupID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*models.Event, error) {
	users := make([]*models.User, 0)
	for _, userID := range input.UsersID {
		users = append(users, &models.User{Model: gorm.Model{ID: userID}})
	}
	event := models.Event{
		Name:         input.Name,
		Place:        input.Place,
		OwnerID:      input.OwnerID,
		CalenderDate: time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, time.UTC),
		Users:        users,
	}
	res := r.DB.Create(&event)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&event)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, input model.EditEvent) (*models.Event, error) {
	var event models.Event
	res := r.DB.First(&event, input.ID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&event).Updates(models.Event{
		Name:         input.Name,
		Place:        input.Place,
		OwnerID:      input.OwnerID,
		CalenderDate: time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, time.UTC),
	})
	if err := res.Error; err != nil {
		return nil, err
	}
	users := make([]*models.User, 0)
	for _, userID := range input.UsersID {
		users = append(users, &models.User{Model: gorm.Model{ID: userID}})
	}
	if err := r.DB.Model(&event).Association("Users").Replace(users); err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&event)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, eventID uint) (*models.Event, error) {
	var event models.Event
	res := r.DB.First(&event, eventID)
	if err := res.Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&event).Association("Users").Clear(); err != nil {
		return nil, err
	}
	res = r.DB.Delete(&event, eventID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *mutationResolver) CreateCalender(ctx context.Context, input model.NewCalender) (*models.Calender, error) {
	calender := models.Calender{
		Date:    time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, time.UTC),
		Weekday: int(input.Date.Weekday()),
		Holiday: input.Holiday,
	}
	res := r.DB.Create(&calender)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload("Events").First(&calender)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &calender, nil
}

func (r *mutationResolver) UpdateCalender(ctx context.Context, input model.EditCalender) (*models.Calender, error) {
	var calender models.Calender
	res := r.DB.First(&calender, "date = ?", time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, time.UTC))
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&calender).Update("Holiday", input.Holiday)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload("Events").First(&calender, "date = ?", time.Date(input.Date.Year(), input.Date.Month(), input.Date.Day(), 0, 0, 0, 0, time.UTC))
	if err := res.Error; err != nil {
		return nil, err
	}
	return &calender, nil
}

func (r *mutationResolver) DeleteCalender(ctx context.Context, calenderDate time.Time) (*models.Calender, error) {
	var calender models.Calender
	res := r.DB.Select("Events").Where("date = ?", time.Date(calenderDate.Year(), calenderDate.Month(), calenderDate.Day(), 0, 0, 0, 0, time.UTC)).Delete(&calender)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &calender, nil
}

func (r *mutationResolver) CreateVote(ctx context.Context, input model.NewVote) (*models.Vote, error) {
	vote := models.Vote{
		Name:    input.Name,
		EventID: input.EventID,
	}
	res := r.DB.Create(&vote)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&vote)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *mutationResolver) UpdateVote(ctx context.Context, input model.EditVote) (*models.Vote, error) {
	var vote models.Vote
	res := r.DB.First(&vote, input.ID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&vote).Updates(models.Vote{
		Name:    input.Name,
		EventID: input.EventID,
	})
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&vote)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *mutationResolver) DeleteVote(ctx context.Context, voteID uint) (*models.Vote, error) {
	var vote models.Vote
	res := r.DB.Delete(&vote, voteID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *mutationResolver) CreateOption(ctx context.Context, input model.NewOption) (*models.Option, error) {
	option := models.Option{
		Name:   input.Name,
		VoteID: input.VoteID,
	}
	res := r.DB.Create(&option)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&option)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *mutationResolver) UpdateOption(ctx context.Context, input model.EditOption) (*models.Option, error) {
	var option models.Option
	res := r.DB.First(&option, input.ID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&option).Update("name", input.Name)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&option)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *mutationResolver) DeleteOption(ctx context.Context, optionID uint) (*models.Option, error) {
	var option models.Option
	res := r.DB.Delete(&option, optionID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *mutationResolver) CreateCast(ctx context.Context, input model.NewCast) (*models.Cast, error) {
	cast := models.Cast{
		OptionID: input.OptionID,
		UserID:   input.UserID,
	}
	res := r.DB.Create(&cast)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&cast)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &cast, nil
}

func (r *mutationResolver) UpdateCast(ctx context.Context, input model.EditCast) (*models.Cast, error) {
	var cast models.Cast
	res := r.DB.First(&cast, input.ID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Model(&cast).Update("option_id", input.OptionID)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Preload(clause.Associations).Find(&cast)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &cast, nil
}

func (r *mutationResolver) DeleteCast(ctx context.Context, castID uint) (*models.Cast, error) {
	var cast models.Cast
	res := r.DB.Delete(&cast, castID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &cast, nil
}

func (r *queryResolver) User(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	res := r.DB.Preload(clause.Associations).First(&user, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Group(ctx context.Context, id uint) (*models.Group, error) {
	var group models.Group
	res := r.DB.Preload("Users").First(&group, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *queryResolver) Event(ctx context.Context, id uint) (*models.Event, error) {
	var event models.Event
	res := r.DB.Preload(clause.Associations).First(&event, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *queryResolver) Calender(ctx context.Context, date time.Time) (*models.Calender, error) {
	var calender models.Calender
	res := r.DB.Preload("Events").First(&calender, "date = ?", time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
	if err := res.Error; err != nil {
		return nil, err
	}
	return &calender, nil
}

func (r *queryResolver) Vote(ctx context.Context, id uint) (*models.Vote, error) {
	var vote models.Vote
	res := r.DB.Preload(clause.Associations).First(&vote, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *queryResolver) Option(ctx context.Context, id uint) (*models.Option, error) {
	var option models.Option
	res := r.DB.Preload(clause.Associations).First(&option, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *queryResolver) Cast(ctx context.Context, id uint) (*models.Cast, error) {
	var cast models.Cast
	res := r.DB.Preload(clause.Associations).First(&cast, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &cast, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	res := r.DB.Preload(clause.Associations).Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) AllGroups(ctx context.Context) ([]*models.Group, error) {
	var groups []*models.Group
	res := r.DB.Preload("Users").Find(&groups)
	if err := res.Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *queryResolver) AllEvents(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	res := r.DB.Preload(clause.Associations).Find(&events)
	if err := res.Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *queryResolver) MonthCalenders(ctx context.Context, month time.Time) ([]*models.Calender, error) {
	firstDay := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	var calenders []*models.Calender
	res := r.DB.Preload("Events").Where("date BETWEEN ? AND ?", firstDay, lastDay).Find(&calenders)
	if err := res.Error; err != nil {
		return nil, err
	}
	return calenders, nil
}

func (r *queryResolver) AllOptions(ctx context.Context, voteID uint) ([]*models.Option, error) {
	var vote models.Vote
	res := r.DB.First(&vote, voteID)
	if err := res.Error; err != nil {
		return nil, err
	}
	var options []*models.Option
	if err := r.DB.Preload(clause.Associations).Model(&vote).Association("Options").Find(&options); err != nil {
		return nil, err
	}
	return options, nil
}

func (r *queryResolver) AllCasts(ctx context.Context, optionID uint) ([]*models.Cast, error) {
	var option models.Option
	res := r.DB.First(&option, optionID)
	if err := res.Error; err != nil {
		return nil, err
	}
	var casts []*models.Cast
	if err := r.DB.Preload(clause.Associations).Model(&option).Association("Casts").Find(&casts); err != nil {
		return nil, err
	}
	return casts, nil
}

func (r *voteResolver) Event(ctx context.Context, obj *models.Vote) (*models.Event, error) {
	var event models.Event
	res := r.DB.Preload(clause.Associations).First(&event, obj.EventID)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Vote returns generated.VoteResolver implementation.
func (r *Resolver) Vote() generated.VoteResolver { return &voteResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type voteResolver struct{ *Resolver }
