CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE modules (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    module_id INT NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
