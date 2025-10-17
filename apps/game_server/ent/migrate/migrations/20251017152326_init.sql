-- Create "client_sessions" table
CREATE TABLE `client_sessions` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `token` text NOT NULL,
  `user_client_sessions` integer NOT NULL,
  CONSTRAINT `client_sessions_users_client_sessions` FOREIGN KEY (`user_client_sessions`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "client_sessions_token_key" to table: "client_sessions"
CREATE UNIQUE INDEX `client_sessions_token_key` ON `client_sessions` (`token`);
-- Create "transfer_codes" table
CREATE TABLE `transfer_codes` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `code` text NOT NULL,
  `user_transfer_code` integer NOT NULL,
  CONSTRAINT `transfer_codes_users_transfer_code` FOREIGN KEY (`user_transfer_code`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "transfer_codes_code_key" to table: "transfer_codes"
CREATE UNIQUE INDEX `transfer_codes_code_key` ON `transfer_codes` (`code`);
-- Create index "transfer_codes_user_transfer_code_key" to table: "transfer_codes"
CREATE UNIQUE INDEX `transfer_codes_user_transfer_code_key` ON `transfer_codes` (`user_transfer_code`);
-- Create "users" table
CREATE TABLE `users` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL
);
