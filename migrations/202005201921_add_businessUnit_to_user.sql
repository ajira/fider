alter table users add column businessUnit_id INT not null DEFAULT 1;
alter table users add Constraint business_unit_id_fk foreign key (businessUnit_id) references business_unit(id);