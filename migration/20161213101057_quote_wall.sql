
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Quote_Wall(
    "id" serial PRIMARY key NOT NULL,
    "quote" text NOT NULL,
    "name" text NOT NULL
);


-- +goose Down
drop table quote_wall
-- SQL section 'Down' is executed when this migration is rolled back

