alter table users add column business_unit_id INT not null DEFAULT 1;
alter table users add Constraint users_business_unit_id_fk foreign key (business_unit_id) references business_unit(id);