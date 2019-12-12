DROP DATABASE IF EXISTS budgets;

CREATE DATABASE budgets;

USE budgets;

CREATE TABLE budgets (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    budget_name varchar(255),
    budget_price float
);

CREATE TABLE groups (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    group_name varchar(255),
    group_price float,
    budget_id int,
    people varchar(255),
    FOREIGN KEY (budget_id) REFERENCES budgets(id)
);



INSERT INTO budgets (id, budget_name, budget_price) VALUES (1, '2019', 500);
INSERT INTO groups (id, group_name, group_price, budget_id, people) VALUES (1,'InLaws - Hers', 0,1,"Josh, Char");
INSERT INTO groups (id, group_name, group_price, budget_id, people) VALUES (1,'InLaws - His', 0,1,"Michael, Rashell");
INSERT INTO groups (id, group_name, group_price, budget_id, people) VALUES (1,'Personal Family', 0,1,"Tori, Porter, Ammon");

