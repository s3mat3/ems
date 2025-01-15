CREATE DATABASE test_db

\c test_db;

DROP TABLE IF EXISTS status;
CREATE TABLE status (
       id integer NOT NULL PRIMARY KEY,
       name varchar(16) NOT NULL
);

DROP TABLE IF EXISTS todo;
CREATE TABLE todo (
       id integer NOT NULL PRIMARY KEY,
       subject varchar(100) NOT NULL,
       status integer NOT NULL,
       CONSTRAINT todo_status_fk FOREIGN KEY (status) REFERENCES status (id) ON DELETE CASCADE
);

CREATE INDEX todo_status_idx ON todo (status);

INSERT INTO status (id, name) VALUES(-10, 'reject');
INSERT INTO status (id, name) VALUES(-1, 'done');
INSERT INTO status (id, name) VALUES(0, 'new');
INSERT INTO status (id, name) VALUES(1, 'open');
INSERT INTO status (id, name) VALUES(2, 'in progress');

INSERT INTO todo (id, status, subject) VALUES( 1, 0, 'hogehoge1');
INSERT INTO todo (id, status, subject) VALUES( 2, 0, 'hogehoge2');
INSERT INTO todo (id, status, subject) VALUES( 3, 0, 'hogehoge3');
INSERT INTO todo (id, status, subject) VALUES( 4, 1, 'hogehoge4');
INSERT INTO todo (id, status, subject) VALUES( 5, -10, 'hogehoge5');
INSERT INTO todo (id, status, subject) VALUES( 6, -1, 'hogehoge6');
INSERT INTO todo (id, status, subject) VALUES( 7, -1, 'hogehoge7');
INSERT INTO todo (id, status, subject) VALUES( 8, 1, 'hogehoge8');
INSERT INTO todo (id, status, subject) VALUES( 9, 2, 'hogehoge9');
INSERT INTO todo (id, status, subject) VALUES(10, 1, 'hogehoge10');
INSERT INTO todo (id, status, subject) VALUES(11, 2, 'hogehoge11');
INSERT INTO todo (id, status, subject) VALUES(12, 0, 'hogehoge12');
INSERT INTO todo (id, status, subject) VALUES(13, 1, 'hogehoge13');
INSERT INTO todo (id, status, subject) VALUES(14, 0, 'hogehoge14');
INSERT INTO todo (id, status, subject) VALUES(15, -10, 'hogehoge15');
INSERT INTO todo (id, status, subject) VALUES(16, 1, 'hogehoge16');
INSERT INTO todo (id, status, subject) VALUES(17, 2, 'hogehoge17');
INSERT INTO todo (id, status, subject) VALUES(18, 2, 'hogehoge18');
INSERT INTO todo (id, status, subject) VALUES(19, -1, 'hogehoge19');
INSERT INTO todo (id, status, subject) VALUES(20, -10, 'hogehoge20');
INSERT INTO todo (id, status, subject) VALUES(21, 1, 'hogehoge21');
