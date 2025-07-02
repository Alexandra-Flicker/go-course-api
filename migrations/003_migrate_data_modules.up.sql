-- +goose Up
INSERT INTO public.modules (id, course_id, title, created_at, updated_at) VALUES (1, 1, 'Installing Go', '2025-06-24 12:56:43.618929', '2025-06-24 12:56:43.618929');
INSERT INTO public.modules (id, course_id, title, created_at, updated_at) VALUES (3, 1, 'Basic Syntax and Variables', '2025-06-24 12:58:14.920043', '2025-06-24 12:58:14.920043');
INSERT INTO public.modules (id, course_id, title, created_at, updated_at) VALUES (4, 2, 'Understanding Goroutines', '2025-06-24 12:58:23.766668', '2025-06-24 12:58:23.766668');
INSERT INTO public.modules (id, course_id, title, created_at, updated_at) VALUES (5, 2, 'Working with Mutex and Channels', '2025-06-24 12:58:30.439257', '2025-06-24 12:58:30.439257');
INSERT INTO public.modules (id, course_id, title, created_at, updated_at) VALUES (6, 2, 'Memory Management and Escape Analysis', '2025-06-24 12:58:37.798224', '2025-06-24 12:58:37.798224');

-- +goose Down


