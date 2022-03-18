create table source
(
    id serial primary key,
    code varchar(20),
    value varchar(20),
    description text
);

create table event_type
(
    id serial primary key,
    code varchar(20),
    value varchar(20),
    description text
);

create table metrics (
                         id serial primary key,
                         code varchar(20),
                         value varchar(20),
                         description text
);

create table log_level (
                           id serial primary key,
                           code varchar(20),
                           value varchar(20),
                           description text
);

create table event
(
    id bigserial primary key,
    created_at timestamptz not null default now(),
    updated_at timestamp not null default now(),
    event_type_id integer references event_type,
    source_id integer references source
);

create table logs (
    id bigserial primary key,
    created_at timestamptz not null default now(),
    updated_at timestamp not null default now(),
    log_level_id integer not null references log_level,
    payload text,
    event_id bigint references event
);

create table analytics (
    id bigserial primary key,
    created_at timestamptz not null default now(),
    updated_at timestamp not null default now(),
    metrics_type_id integer references metrics,
    payload jsonb,
    event_id bigint references event
);
