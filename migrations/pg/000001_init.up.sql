CREATE TABLE users (
    id                  serial primary key,
    email			    varchar(255) UNIQUE not null,
    first_name          varchar(255) not null,
    last_name           varchar(255) not null,
    hashed_password 	varchar(255) not null
);

CREATE TABLE statuses (
    id                  serial primary key,
    status_name         varchar(30) unique not null
);

INSERT INTO statuses (status_name) VALUES
    ('open'),
    ('in progress'),
    ('completed');

CREATE TABLE priorities (
    id                  serial primary key,
    priority_name       varchar(30) unique not null
);

INSERT INTO priorities (priority_name) VALUES
    ('low'),
    ('medium'),
    ('high');

CREATE TABLE tasks (
    id                  serial primary key,
    user_id             integer references users(id) not null,
    created_date        timestamp not null,
    task_name 		    varchar(255) not NULL,
    description         text not null,
    status_id	        integer references statuses(id),
    deadline_from	    timestamp,
    deadline_to		    timestamp,
    priority_id		    integer references priorities(id),
    parent_task_id		integer references tasks(id)
);

CREATE TABLE roles (
    id                  serial primary key,
    role_name           varchar(30) unique not null
);

INSERT INTO roles (role_name) VALUES
    ('owner'),
    ('member'),
    ('superuser');

CREATE TABLE members_tasks (
    user_id             integer references users(id) not null,
    task_id             integer references tasks(id) not null,
    role_id			    integer references roles(id) not null
);

CREATE TABLE actions (
    id                  serial primary key,
    action_name         varchar(30) unique not null
);

INSERT INTO actions (action_name) VALUES
    ('create'),
    ('read'),
    ('update'),
    ('delete');

CREATE TABLE events (
    id                  serial primary key,
    user_id             integer references users(id) not null,
    task_id             integer references tasks(id) not null,
    action_id			integer references actions(id) not null,
    created_date		timestamp not null
);

CREATE TABLE task_logs (
    id                  serial primary key,
    user_id             integer references users(id) not null,
    task_id             integer references tasks(id) not null,
    status_id 		    integer references statuses(id) not null,
    begin_date			timestamp,
    end_date			timestamp,
    total_hours			integer,
    description		    text
);

CREATE TABLE comments (
    id                  serial primary key,
    user_id             integer references users(id) not null,
    task_id 		    integer references tasks(id) not null,
    created_date		varchar(255) not null,
    comment			    text not null,
    parent_comment_id	integer references comments(id)
);