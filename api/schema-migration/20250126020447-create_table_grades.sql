
-- +migrate Up
CREATE TABLE grades (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    assignment_id UUID NOT NULL,
    score DOUBLE PRECISION NOT NULL,
    feedback TEXT,
    teacher_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_assignment FOREIGN KEY (assignment_id) REFERENCES assignments (id) ON DELETE CASCADE
);
-- +migrate Down
