# Todoist

## Database
sqlx

- sql query
- migrate
- testing
- go:embed

```sql
CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(36) NOT NULL,
  `identity` varchar(128) NOT NULL,
  `passwd`  varchar(128) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
) PRIMARY kEY (`id`);

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

CREATE INDEX IF NOT EXISTS `idx_tasks_user_id` ON tasks(`user_id`);
```

## APIs

- create a task
- gets a task
- gets all tasks
- update complete task
- delete complete task

using protobuf request/response data with go:generate command generate protobuf source
```sh
protoc --go_out=. go_opt paths=source_relative
````





