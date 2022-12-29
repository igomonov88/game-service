INSERT INTO results (user_choice, computer_choice, result, created_at)
VALUES (1, 2, 'win', '2022-01-01 00:00:00') ON CONFLICT DO NOTHING;