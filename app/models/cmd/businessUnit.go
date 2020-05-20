package cmd

import "github.com/getfider/fider/app/models"

type AddBusinessUnit struct {
	Name string

	Result *models.BusinessUnit
}

type GetBusinessUnitByName struct {
	Name string

	Result *models.BusinessUnit
}

type GetBusinessUnitByID struct {
	Id string

	Result *models.BusinessUnit
}
