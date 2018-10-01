-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("1f15a4c9d4798230","admin","Administrative role",1538410596280);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("3ae149a8dc444bce","executive","Executive role",1538410596280);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("455511f8e362deb9","manager","Manager role",1538410596280);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DELETE FROM `acl_role` WHERE `id` = "1f15a4c9d4798230";
DELETE FROM `acl_role` WHERE `id` = "3ae149a8dc444bce";
DELETE FROM `acl_role` WHERE `id` = "455511f8e362deb9";
