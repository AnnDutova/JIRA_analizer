create table if not exists "project" (
    id serial primary key,
    title text not null unique
);

create table if not exists "author" (
    id serial primary key,
    name text not null unique
);

create table if not exists "issues" (
    id serial primary key,
    projectId int not null,
    authorId int not null,
    assigneeId int,
    key text not null,
    createdTime timestamp not null,
    updatedTime timestamp not null,
    closedTime  timestamp,
    summary     text not null,
    description text,
    type        text not null,
    priority    text not null,
    status      text not null,
    timeSpent   bigint,
    info json,
    constraint "fk_issues_project" foreign key (projectId) references project(id) MATCH FULL,
    constraint "fk_issues_author" foreign key (authorId) references author(id) MATCH FULL,
    constraint "fk_issues_assignee" foreign key (assigneeId) references author(id) MATCH FULL
    );

create table if not exists "statusChange" (
    authorId int not null,
    issueId int not null,
    changeTime timestamp not null,
    fromStatus text,
    toStatus text not null,
    constraint "fk_statusChange_author" foreign key (authorId) references author(id) MATCH FULL
);

create table if not exists "comments" (
    authorId int not null,
    issueId int not null,
    changeTime timestamp not null,
    content text not null,
    constraint "fk_comments_author" foreign key (authorId) references author(id) MATCH FULL,
    constraint "fk_comments_issues" foreign key (issueId) references issues(id) MATCH FULL
);
