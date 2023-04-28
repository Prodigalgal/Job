CREATE TABLE `classify` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `author` varchar(255) DEFAULT NULL,
                            `update_time` datetime NOT NULL,
                            `creat_time` datetime DEFAULT NULL,
                            `name` varchar(255) NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `dangerous` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) NOT NULL,
                             `create_time` datetime NOT NULL,
                             `update_time` datetime NOT NULL,
                             `author` varchar(255) DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `job` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `name` varchar(255) NOT NULL,
                       `create_time` datetime NOT NULL,
                       `update_time` datetime NOT NULL,
                       `path` varchar(255) DEFAULT NULL,
                       `author` varchar(255) DEFAULT NULL,
                       `dangerous` int NOT NULL,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `job_classify_union` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `job_id` int NOT NULL,
                                      `classify_id` int NOT NULL,
                                      `create_time` datetime NOT NULL,
                                      `update_time` datetime DEFAULT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `username` varchar(255) NOT NULL,
                        `password` varchar(255) NOT NULL,
                        `create_time` datetime NOT NULL,
                        `update_time` datetime NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

