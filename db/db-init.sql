create table if not exists "issues" (
                                        key text unique not null,
                                        createdTime date not null,
                                        closedTime  date,
                                        summary     text not null,
                                        type        text not null,
                                        priority    text not null,
                                        status      text not null,
                                        timeSpent   bigint,
                                        info json NOT NULL
);

create or replace function insertIssue(key_ text, createTime_ date, closedTime_ date,summary_ text, type_ text, priority_ text,
                                       status_  text, val json) returns void as $$
begin
Insert Into issues(key,createdTime,closedTime,summary, type,priority,status, info) values
    (key_,createTime_,closedTime_,summary_,type_,priority_,status_, val);
end;
$$ LANGUAGE plpgsql;

