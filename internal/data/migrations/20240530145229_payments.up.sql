 CREATE TABLE payments (
    id SERIAL NOT NULL,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    payment_status VARCHAR(25) NOT NULL DEFAULT 'pending',
    intent_id VARCHAR(255) UNIQUE NOT NULL,
    check_in_status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPtz NOT NULL DEFAULT 'now()',
    updated_at TIMESTAMPtz NOT NULL DEFAULT 'now'
);

ALTER TABLE payments ADD PRIMARY KEY(id);
ALTER TABLE payments ADD FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE payments ADD FOREIGN KEY(event_id) REFERENCES events(id);