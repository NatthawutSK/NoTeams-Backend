BEGIN;

TRUNCATE TABLE "User" CASCADE;
TRUNCATE TABLE "Oauth" CASCADE;
TRUNCATE TABLE "Team" CASCADE;
TRUNCATE TABLE "TeamMember" CASCADE;
TRUNCATE TABLE "Permission" CASCADE;
TRUNCATE TABLE "File" CASCADE;
TRUNCATE TABLE "Task" CASCADE;

COMMIT;
