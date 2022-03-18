alter table event_type add constraint eventUnique unique (code);
alter table log_level add constraint logUnique unique (code);
alter table source add constraint sourceUnique unique (code);