CREATE TABLE items
(
    id   BIGINT NOT NULL AUTO_INCREMENT,
    name varchar(250) NOT NULL,
    views BIGINT  NOT NULL default 0,
    PRIMARY KEY (id)
);