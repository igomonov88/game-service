-- Version: 1.1
-- Description: create table results
CREATE TABLE results
(
    user_choice     INTEGER   NOT NULL,
    computer_choice INTEGER   NOT NULL,
    result          TEXT      NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP DEFAULT NULL
);
