-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE profiles ADD `rel` text DEFAULT NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE profiles DROP `rel`;
