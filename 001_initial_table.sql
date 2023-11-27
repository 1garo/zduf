-- This is a sample migration.

create table teacher(
    id serial primary key,
    first_name varchar not null,
    last_name varchar not null,
    age smallint not null
);

create table student(
    id serial primary key,
    first_name varchar not null,
    last_name varchar not null,
    age smallint not null
);

create table class(
    id serial primary key,
    name varchar not null,
    time timestamp not null,

    teacher_id int not null,
    CONSTRAINT fk_teacher FOREIGN KEY(teacher_id) REFERENCES teacher(id)
);

create table student_class(
    student_id int not null,
    class_id int not null,

    PRIMARY KEY (student_id, class_id),
    CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES student(id),
    CONSTRAINT fk_class FOREIGN KEY(class_id) REFERENCES class(id)
);

---- create above / drop below ----

drop table student_class;
drop table teacher;
drop table student;
drop table class;
