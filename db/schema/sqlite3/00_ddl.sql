CREATE TABLE IF NOT EXISTS `tasks` (
  `id` varchar(36) NOT NULL,
  `name` varchar(140) NOT NULL,
  `status` int NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `due_date` varchar(16) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY kEY (`id`)
);

