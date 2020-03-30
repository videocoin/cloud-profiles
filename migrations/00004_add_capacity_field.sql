-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE profiles ADD `capacity` json DEFAULT NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE profiles DROP `capacity`;
