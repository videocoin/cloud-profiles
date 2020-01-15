-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE profiles 
    ADD `deposit` varchar(255) DEFAULT NULL;
    ADD `reward`  varchar(255) DEFAULT NULL, 

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE profiles DROP `deposit`, DROP `reward`;
