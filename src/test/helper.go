package test

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
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

func RawToDtoPermission(perm *model.Permission) *proto.Permission {
	return &proto.Permission{
		Id:   uint32(perm.ID),
		Name: perm.Name,
		Code: perm.Code,
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
