package service

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"gorm.io/gorm"
)

func RawToDtoContact(cont *model.Contact) *proto.Contact {
	return &proto.Contact{
		Id:        uint32(cont.ID),
		Facebook:  cont.Facebook,
		Instagram: cont.Instagram,
		Twitter:   cont.Twitter,
		Linkedin:  cont.Linkedin,
	}
}

func DtoToRawContact(cont *proto.Contact) *model.Contact {
	return &model.Contact{
		Model:     gorm.Model{ID: uint(cont.Id)},
		Facebook:  cont.Facebook,
		Instagram: cont.Instagram,
		Twitter:   cont.Twitter,
		Linkedin:  cont.Linkedin,
	}
}

func RawToDtoLocation(loc *model.Location) *proto.Location {
	return &proto.Location{
		Id:       uint32(loc.ID),
		Address:  loc.Address,
		District: loc.District,
		Province: loc.Province,
		Country:  loc.Country,
		Zipcode:  loc.ZipCode,
	}
}

func DtoToRawLocation(loc *proto.Location) *model.Location {
	return &model.Location{
		Model:    gorm.Model{ID: uint(loc.Id)},
		Address:  loc.Address,
		District: loc.District,
		Province: loc.Province,
		Country:  loc.Country,
		ZipCode:  loc.Zipcode,
	}
}

func RawToDtoPermission(perm *model.Permission) *proto.Permission {
	return &proto.Permission{
		Id:   uint32(perm.ID),
		Name: perm.Name,
		Code: perm.Code,
	}
}

func DtoToRawPermission(permission *proto.Permission) *model.Permission {
	return &model.Permission{
		Model: gorm.Model{
			ID: uint(permission.Id),
		},
		Name: permission.Name,
		Code: permission.Code,
	}
}

func DtoToRawRole(role *proto.Role) *model.Role {
	var perms []*model.Permission
	for _, perm := range role.Permissions {
		perms = append(perms, DtoToRawPermission(perm))
	}

	return &model.Role{
		Model:       gorm.Model{ID: uint(role.Id)},
		Name:        role.Name,
		Description: role.Description,
		Permissions: perms,
	}
}

func RawToDtoRole(role *model.Role) *proto.Role {
	var permissions []*proto.Permission
	for _, permission := range role.Permissions {
		rolePerm := RawToDtoPermission(permission)
		permissions = append(permissions, rolePerm)
	}
	return &proto.Role{
		Id:          uint32(role.ID),
		Name:        role.Name,
		Description: role.Description,
		Permissions: permissions,
	}
}

func RawToDtoUser(user *model.User) *proto.User {
	var orgs []*proto.Organization
	for _, org := range user.Organizations {
		orgs = append(orgs, RawToDtoOrganization(org))
	}

	var teams []*proto.Team
	for _, team := range user.Teams {
		teams = append(teams, RawToDtoTeam(team))
	}

	return &proto.User{
		Id:            uint32(user.ID),
		Firstname:     user.Firstname,
		Lastname:      user.Lastname,
		DisplayName:   user.DisplayName,
		ImageUrl:      user.ImageUrl,
		Teams:         teams,
		Organizations: orgs,
		Contact:       RawToDtoContact(&user.Contact),
		Address:       RawToDtoLocation(&user.Location),
	}
}

func DtoToRawUser(user *proto.User) *model.User {
	return &model.User{
		Model:       gorm.Model{ID: uint(user.Id)},
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		DisplayName: user.DisplayName,
		ImageUrl:    user.ImageUrl,
	}
}

func RawToDtoSubTeams(team []*model.Team) []*proto.Team {
	var subTeams []*proto.Team
	for _, t := range team {
		subTeam := proto.Team{
			Id:          uint32(t.ID),
			Name:        t.Name,
			Description: t.Description,
			SubTeams:    RawToDtoSubTeams(t.SubTeams),
		}
		subTeams = append(subTeams, &subTeam)
	}
	return subTeams
}

func RawToDtoTeam(team *model.Team) *proto.Team {
	var members []*proto.User
	for _, member := range team.Members {
		members = append(members, RawToDtoUser(member))
	}

	return &proto.Team{
		Id:          uint32(team.ID),
		Name:        team.Name,
		Description: team.Description,
		SubTeams:    RawToDtoSubTeams(team.SubTeams),
		Members:     members,
	}
}

func RawToDtoOrganization(org *model.Organization) *proto.Organization {
	var roles []*proto.Role
	for _, role := range org.Roles {
		roles = append(roles, &proto.Role{Id: uint32(role.ID)})
	}

	var members []*proto.User
	for _, member := range org.Members {
		members = append(members, RawToDtoUser(member))
	}

	return &proto.Organization{
		Id:          uint32(org.ID),
		Name:        org.Name,
		Email:       org.Email,
		Description: org.Description,
		Teams:       RawToDtoSubTeams(org.Teams),
		Roles:       roles,
		Members:     members,
		Contact:     RawToDtoContact(&org.Contact),
		Location:    RawToDtoLocation(&org.Location),
	}
}

func DtoToRawSubTeams(team []*proto.Team) []*model.Team {
	var subTeams []*model.Team
	for _, t := range team {
		subTeam := model.Team{
			Model:       gorm.Model{ID: uint(t.Id)},
			Name:        t.Name,
			Description: t.Description,
			SubTeams:    DtoToRawSubTeams(t.SubTeams),
		}
		subTeams = append(subTeams, &subTeam)
	}
	return subTeams
}

func DtoToRawTeam(team *proto.Team) *model.Team {
	return &model.Team{
		Model:       gorm.Model{ID: uint(team.Id)},
		Name:        team.Name,
		Description: team.Description,
	}
}

func DtoToRawOrganization(org *proto.Organization) *model.Organization {
	var roles []*model.Role
	for _, role := range org.Roles {
		roles = append(roles, &model.Role{
			Model:       gorm.Model{ID: uint(role.Id)},
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return &model.Organization{
		Model:       gorm.Model{ID: uint(org.Id)},
		Name:        org.Name,
		Email:       org.Email,
		Description: org.Description,
		Teams:       DtoToRawSubTeams(org.Teams),
		Roles:       roles,
	}
}
