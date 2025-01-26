
-- +migrate Up
CREATE TABLE assignments (
    id UUID PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    tittle VARCHAR(255) NOT NULL,
    student_id UUID NOT NULL,
    content TEXT,
    status VARCHAR(20) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate Down
