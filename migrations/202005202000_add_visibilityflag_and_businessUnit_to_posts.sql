
alter table posts add column is_public boolean not null DEFAULT false;
alter table posts add column business_unit_id INT not null DEFAULT 1;
alter table posts add Constraint posts_business_unit_id_fk foreign key (business_unit_id) references business_unit(id);
alter table posts add column estimated_date_for_completion date;