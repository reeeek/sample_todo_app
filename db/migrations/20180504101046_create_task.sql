
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks(
id int unsigned NOT NULL AUTO_INCREMENT,
status tinyint NOT NULL,
title varchar(100) NOT NULL,
body varchar(512) NOT NULL,
created_at datetime NOT NULL,
updated_at datetime NOT NULL,
PRIMARY KEY (`id`)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tasks;
