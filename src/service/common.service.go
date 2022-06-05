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
	return &proto.User{
		Id:          uint32(user.ID),
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		DisplayName: user.DisplayName,
		ImageUrl:    user.ImageUrl,
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
