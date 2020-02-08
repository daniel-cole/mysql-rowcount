DROP DATABASE IF EXISTS db1_test;

CREATE DATABASE db1_test;
CREATE TABLE db1_test.test1(
id INT PRIMARY KEY,
    description varchar(255)
);
INSERT INTO db1_test.test1(id, description) values(0, 'foo');
INSERT INTO db1_test.test1(id, description) values(1, 'bar');
INSERT INTO db1_test.test1(id, description) values(2, 'hello');
INSERT INTO db1_test.test1(id, description) values(3, 'world');
INSERT INTO db1_test.test1(id, description) values(4, 'foooobar');


CREATE TABLE db1_test.test2(
    id INT PRIMARY KEY,
    description varchar(255)
);
INSERT INTO db1_test.test2(id, description) values(1, 'foo');
INSERT INTO db1_test.test2(id, description) values(2, 'bar');
