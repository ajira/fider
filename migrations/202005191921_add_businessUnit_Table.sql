create table if not exists business_unit (
  id           serial primary key,
  name         varchar(60) not null,   
  tenant_id    int not null,
  foreign key (tenant_id) references tenants(id),
  unique (id, tenant_id)
);