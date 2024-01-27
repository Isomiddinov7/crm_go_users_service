CREATE TYPE LEVEL_TYPE AS ENUM('BEGINER', 'ELEMENTARY', 'INTERMEDIATE', 'IELTS');

CREATE TABLE IF NOT EXISTS "role"(
    "id" UUID NOT NULL PRIMARY KEY,
    "role" VARCHAR NOT NULL
);


CREATE TABLE IF NOT EXISTS "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "supper_admin"(
    "id" VARCHAR PRIMARY KEY,
    "full_name" VARCHAR,
    "password" VARCHAR,
    "login" VARCHAR UNIQUE,
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "manager"(
    "id" VARCHAR PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "salary" NUMERIC,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "teacher"(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "phone" VARCHAR(13),
    "salary" VARCHAR,
    "password" VARCHAR,
    "login" VARCHAR UNIQUE,
    "branch_id" UUID REFERENCES "branch_id";
    "role_id" UUID REFERENCES "role"("id"),
    "months_worked" VARCHAR,
    "total_sum" VARCHAR,
    "ielts_score" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "support_teacher"(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "phone" VARCHAR(13),
    "salary" VARCHAR,
    "password" VARCHAR,
    "login" VARCHAR UNIQUE,
    "branch_id" UUID REFERENCES "branch_id";
    "role_id" UUID REFERENCES "role"("id"),
    "months_worked" VARCHAR,
    "total_sum" VARCHAR,
    "ielts_score" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "administrator"(
    "id" VARCHAR PRIMARY KEY,
    "full_name" VARCHAR,
    "phone" VARCHAR UNIQUE,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "salary" NUMERIC,
    "ielts_score" VARCHAR,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"), 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "group"(
    "id" UUID PRIMARY KEY,
    "uniqueID" VARCHAR NOT NULL UNIQUE,
    "branch_id" UUID REFERENCES "branch"("id"),
    "type" LEVEL_TYPE,
    "teacher_id" VARCHAR REFERENCES "teacher"("id"),
    "support_teacher_id" VARCHAR REFERENCES "support_teacher"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "student"(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "phone" VARCHAR(13),
    "paid_sum" VARCHAR,
    "cource_count" INT,
    "total_sum" VARCHAR,
    "group_id" UUID REFERENCES "group"("id"),
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "event"(
    "id" UUID NOT NULL PRIMARY KEY,
    "topic" VARCHAR,
    "start_time" TIME NOT NULL, 
    "day" VARCHAR,
    "branch_id" UUID REFERENCES "branch"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "assign_student"(
    "id" UUID PRIMARY KEY,
    "event_id" UUID REFERENCES "event"("id"),
    "student_id" VARCHAR REFERENCES "student"("id") UNIQUE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "journal"(
    "id" UUID PRIMARY KEY,
    "group_id" UUID REFERENCES "group"("id"),
    "from" VARCHAR,
    "to" VARCHAR,
    "student_id" VARCHAR REFERENCES "student"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "schedule"(
    "id" UUID NOT NULL PRIMARY KEY,
    "start_time"    VARCHAR, 
    "end_time"  VARCHAR, 
    "journal_id" UUID REFERENCES "journal"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "lesson"(
    "id" UUID NOT NULL PRIMARY KEY,
    "schedule_id" UUID REFERENCES "schedule"("id") UNIQUE,
    "lesson" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "task"(
    "id" UUID NOT NULL PRIMARY KEY,
    "lesson_id" UUID REFERENCES "lesson"("id"),
    "label" VARCHAR,
    "deadline" DATE NOT NULL,
    "score" INT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "payment"(
    "id" UUID PRIMARY KEY,
    "student_id" VARCHAR REFERENCES "student"("id"),
    "branch_id" UUID REFERENCES "branch"("id"), 
    "paid_sum" NUMERIC,
    "total_sum" NUMERIC,
    "course_count" INTEGER,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "score"(
    "id" UUID PRIMARY KEY,
    "task_id" UUID REFERENCES "task"("id"),
    "student_id" VARCHAR REFERENCES "student"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

