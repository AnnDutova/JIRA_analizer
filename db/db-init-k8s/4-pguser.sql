Create user pguser with password 'pgpwd';

GRANT CONNECT on database "testdb" to pguser;

GRANT USAGE, SELECT ON SEQUENCE project_id_seq TO pguser;
GRANT USAGE, SELECT ON SEQUENCE author_id_seq TO pguser;
GRANT USAGE, SELECT ON SEQUENCE issues_id_seq TO pguser;

GRANT SELECT, INSERT, UPDATE on table "project", "issues","author", "statusChange", "comments"  to pguser;

