-- +goose Up
INSERT INTO public.lessons (id, module_id, title, content, created_at, updated_at) VALUES (1, 1, 'Introduction to Go', 'Learn the basics of Go: syntax, structure, and Hello World.', '2025-06-30 18:36:54.723056', '2025-06-30 18:36:54.723056');
INSERT INTO public.lessons (id, module_id, title, content, created_at, updated_at) VALUES (2, 1, 'Variables and Data Types', 'Understand var, short declaration :=, and basic data types.', '2025-06-30 18:37:02.698089', '2025-06-30 18:37:02.698089');
INSERT INTO public.lessons (id, module_id, title, content, created_at, updated_at) VALUES (6, 3, 'Goroutines', 'Understanding goroutines and their role in concurrency.', '2025-06-30 18:39:15.109421', '2025-06-30 18:39:15.109421');
INSERT INTO public.lessons (id, module_id, title, content, created_at, updated_at) VALUES (8, 4, 'Using Mutex in Go', 'Learn how to synchronize access to shared resources using sync.Mutex.', '2025-06-30 18:40:22.598820', '2025-06-30 18:40:22.598820');
INSERT INTO public.lessons (id, module_id, title, content, created_at, updated_at) VALUES (9, 5, 'Escape Analysis in Go', 'Understand memory allocation and escape analysis for optimization.', '2025-06-30 18:40:30.748267', '2025-06-30 18:40:30.748267');
-- +goose Down
-- no down