-- 000001_init_schema.up.sql
CREATE TABLE themes (
    name VARCHAR(50) PRIMARY KEY,
    floating_shapes_theme VARCHAR(50) NOT NULL
);

CREATE TABLE thank_you_messages (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    subtitle TEXT,
    icon VARCHAR(10),
    button_text VARCHAR(100),
    button_url VARCHAR(200),
    button_new_tab BOOLEAN DEFAULT false
);

CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    theme_name VARCHAR(50) REFERENCES themes(name),
    thank_you_message_id INTEGER REFERENCES thank_you_messages(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    form_id INTEGER REFERENCES forms(id) ON DELETE CASCADE,
    question_order INTEGER NOT NULL,
    type VARCHAR(50) NOT NULL,
    question TEXT NOT NULL,
    subtext TEXT,
    image VARCHAR(200),
    gradient VARCHAR(100),
    bg_color VARCHAR(20),
    input_type VARCHAR(50),
    placeholder TEXT,
    validation TEXT,
    max_selections INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE question_options (
    id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
    text VARCHAR(200) NOT NULL,
    icon VARCHAR(10),
    color VARCHAR(20)
);

CREATE TABLE next_question_conditions (
    id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
    answer TEXT,
    next_question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
    is_default BOOLEAN DEFAULT false
);

CREATE TABLE form_submissions (
    id SERIAL PRIMARY KEY,
    form_id INTEGER REFERENCES forms(id) ON DELETE CASCADE,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE submission_answers (
    id SERIAL PRIMARY KEY,
    submission_id INTEGER REFERENCES form_submissions(id) ON DELETE CASCADE,
    question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
    answer TEXT NOT NULL
);
