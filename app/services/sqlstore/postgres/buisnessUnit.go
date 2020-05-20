package postgres

import (
	"context"

	"github.com/getfider/fider/app/models"
	"github.com/getfider/fider/app/models/cmd"
	"github.com/getfider/fider/app/pkg/dbx"
	"github.com/getfider/fider/app/pkg/errors"
)

type dbBusinessUnit struct {
	ID     int       `db:"id"`
	Name   string    `db:"name"`
	Tenant *dbTenant `db:"tenant"`
}

func (c *dbBusinessUnit) toModel() *models.BusinessUnit {
	businessUnit := &models.BusinessUnit{
		ID:     c.ID,
		Name:   c.Name,
		Tenant: c.Tenant.toModel(),
	}
	return businessUnit
}

func addBusinessUnit(ctx context.Context, c *cmd.AddBusinessUnit) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {

		var id int
		err := trx.Get(&id,
			`INSERT INTO business_unit (name, tenant_id) VALUES($1,$2)`, c.Name, tenant.ID)
		if err != nil {
			return err
		}
		c.Result = &models.BusinessUnit{
			ID:     id,
			Name:   c.Name,
			Tenant: tenant,
		}
		return nil
	})
}

func getBusinessUnitByName(ctx context.Context, c *cmd.GetBusinessUnitByName) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		businessUnit := dbBusinessUnit{}

		err := trx.Get(&businessUnit, "SELECT id, name FROM business_unit WHERE tenant_id = $1 AND name = $2", tenant.ID, c.Name)
		if err != nil {
			return errors.Wrap(err, "failed to get buisness unit with name '%s'", c.Name)
		}
		c.Result = businessUnit.toModel()
		return nil
	})
}
