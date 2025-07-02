-- +goose Up
INSERT INTO public.courses (id, name, description, created_at, updated_at) VALUES (1, 'Go Basics', 'Introduction to Go', '2025-06-16 16:09:02.226261', '2025-06-16 16:09:02.226261');
INSERT INTO public.courses (id, name, description, created_at, updated_at) VALUES (2, 'Advanced Go', 'Deep dive into Go concurrency and memory management.', '2025-06-16 16:17:33.651741', '2025-06-16 16:17:33.651741');
-- +goose Down
-- no down