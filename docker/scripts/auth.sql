CREATE SCHEMA profile;

CREATE TABLE profile.userdata (
    tg_id BIGINT PRIMARY KEY,
    profile_name VARCHAR(100) NOT NULL
);

INSERT INTO profile.userdata (tg_id, profile_name) VALUES 
(616547839, 'user1'),
(616547830, 'user2');
