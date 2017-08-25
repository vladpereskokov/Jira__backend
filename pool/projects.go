package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"github.com/DVI-GI-2017/Jira__backend/utils"
)

func init() {
	resolvers["Project"] = projectsResolver
}

const (
	ProjectCreate     = Action("ProjectCreate")
	ProjectExists     = Action("ProjectExists")
	ProjectsAll       = Action("ProjectsAll")
	ProjectFindById   = Action("ProjectFindById")
	ProjectAllUsers   = Action("ProjectAllUsers")
	ProjectAllTasks   = Action("ProjectAllTasks")
	ProjectAddUser    = Action("ProjectAddUser")
	ProjectDeleteUser = Action("ProjectDeleteUser")
	ProjectUserExists = Action("ProjectUserExists")
)

func projectsResolver(action Action) (ServiceFunc, error) {
	switch action {
	case ProjectCreate:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if project, ok := data.(models.Project); ok {
				return projects.CreateProject(source, project)
			}
			return models.Project{}, utils.CastFailsMsg(data, models.Project{})
		}, nil

	case ProjectExists:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if project, ok := data.(models.Project); ok {
				return projects.CheckProjectExists(source, project)
			}
			return false, utils.CastFailsMsg(data, models.Project{})
		}, nil

	case ProjectsAll:
		return func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}, nil

	case ProjectFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.FindProjectById(source, id)
			}
			return models.Project{}, utils.CastFailsMsg(data, models.RequiredId{})
		}, nil

	case ProjectAllUsers:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.AllUsersInProject(source, id)
			}
			return models.UsersList{}, utils.CastFailsMsg(data, models.RequiredId{})
		}, nil

	case ProjectAllTasks:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.AllTasksInProject(source, id)
			}
			return models.TasksList{}, utils.CastFailsMsg(data, models.RequiredId{})
		}, nil

	case ProjectAddUser:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.AddUserToProject(source, ids.ProjectId, ids.UserId)
			}
			return models.UsersList{}, utils.CastFailsMsg(data, models.ProjectUser{})
		}, nil

	case ProjectDeleteUser:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.DeleteUserFromProject(source, ids.ProjectId, ids.UserId)
			}
			return models.UsersList{}, utils.CastFailsMsg(data, models.ProjectUser{})
		}, nil

	case ProjectUserExists:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.CheckUserInProject(source, ids.UserId, ids.ProjectId)
			}
			return false, utils.CastFailsMsg(data, models.ProjectUser{})
		}, nil

	default:
		return nil, fmt.Errorf("can not find resolver with action: %v, in projects resolvers", action)

	}
}
