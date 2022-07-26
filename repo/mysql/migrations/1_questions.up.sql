-- CREATE TABLE IF NOT EXISTS questions (
--     `id` int NOT NULL AUTO_INCREMENT,
--     `no` int NOT NULL,
--     `question` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
--     `answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
--     PRIMARY KEY (`id`) USING BTREE
-- );

CREATE TABLE IF NOT EXISTS questions (
    `id` int NOT NULL AUTO_INCREMENT,
    `no` int NOT NULL,
    `question` longtext  NULL,
    `answer` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
);