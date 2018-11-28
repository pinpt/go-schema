-- GENERATED SQL. DO NOT EDIT

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

--  ACLGrant is a mapping of resource to roles which provide a permission for a resource to a role
CREATE TABLE `acl_grant` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`resource_id`   VARCHAR(64) NOT NULL,
	`role_id`       VARCHAR(64) NOT NULL,
	`permission`    ENUM('read','readwrite','admin') NOT NULL,
	`admin_user_id` VARCHAR(64) NOT NULL,
	`customer_id`   VARCHAR(64) NOT NULL,
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	INDEX acl_grant_resource_id_index (`resource_id`),
	INDEX acl_grant_role_id_index (`role_id`),
	INDEX acl_grant_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  ACLResource is a defined resource in the system
CREATE TABLE `acl_resource` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`urn`           VARCHAR(255) NOT NULL,
	`description`   TEXT,
	`title`         TEXT,
	`public`        TINYINT(1) NOT NULL,
	`hidden`        TINYINT(1) NOT NULL,
	`admin`         TINYINT(1) NOT NULL,
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	INDEX acl_resource_urn_index (`urn`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  ACLRole is a role specific table defined by an admin
CREATE TABLE `acl_role` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`name`          VARCHAR(100) NOT NULL,
	`description`   TEXT,
	`admin_user_id` VARCHAR(64),
	`customer_id`   VARCHAR(64),
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	INDEX acl_role_name_index (`name`),
	INDEX acl_role_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  ACLUserRole maps users to roles
CREATE TABLE `acl_user_role` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`user_id`       VARCHAR(64) NOT NULL,
	`role_id`       VARCHAR(64) NOT NULL,
	`customer_id`   VARCHAR(64) NOT NULL,
	`admin_user_id` VARCHAR(64) NOT NULL,
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	INDEX acl_user_role_user_id_index (`user_id`),
	INDEX acl_user_role_role_id_index (`role_id`),
	INDEX acl_user_role_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  App is a higher level abstraction around a logical unit of code and systems
CREATE TABLE `app` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`name`          VARCHAR(255) NOT NULL,
	`description`   TEXT,
	`active`        BOOL NOT NULL,
	`repo_ids`      JSON,
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	`customer_id`   VARCHAR(64) NOT NULL,
	INDEX app_name_index (`name`),
	INDEX app_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Commit is a generic table which describes a single code commit
CREATE TABLE `commit` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`repo_id`           VARCHAR(64) NOT NULL,
	`sha`               VARCHAR(64) NOT NULL,
	`branch`            VARCHAR(255) DEFAULT "master",
	`message`           LONGTEXT,
	`mergecommit`       BOOL DEFAULT false,
	`excluded`          BOOL DEFAULT false,
	`parent`            VARCHAR(64),
	`parent_id`         VARCHAR(64),
	`date`              BIGINT UNSIGNED NOT NULL,
	`author_user_id`    VARCHAR(64) NOT NULL,
	`committer_user_id` VARCHAR(64) NOT NULL,
	`ordinal`           INT,
	`customer_id`       VARCHAR(64) NOT NULL,
	`ref_type`          VARCHAR(20) NOT NULL,
	`ref_id`            VARCHAR(64) NOT NULL,
	`metadata`          JSON,
	INDEX commit_repo_id_index (`repo_id`),
	INDEX commit_sha_index (`sha`),
	INDEX commit_author_user_id_index (`author_user_id`),
	INDEX commit_committer_user_id_index (`committer_user_id`),
	INDEX commit_customer_id_index (`customer_id`),
	INDEX commit_ref_type_index (`ref_type`),
	INDEX commit_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Commit Activity is a generic table which links users to loc by commit for a date
CREATE TABLE `commit_activity` (
	`id`                   VARCHAR(64) NOT NULL PRIMARY KEY,
	`date`                 BIGINT UNSIGNED NOT NULL,
	`sha`                  VARCHAR(64) NOT NULL,
	`user_id`              VARCHAR(64) NOT NULL,
	`repo_id`              VARCHAR(64) NOT NULL,
	`filename`             VARCHAR(700) NOT NULL,
	`language`             VARCHAR(100) DEFAULT "unknown",
	`ordinal`              BIGINT UNSIGNED NOT NULL,
	`loc`                  INT NOT NULL DEFAULT 0,
	`sloc`                 INT NOT NULL DEFAULT 0,
	`blanks`               INT NOT NULL DEFAULT 0,
	`comments`             INT NOT NULL DEFAULT 0,
	`complexity`           BIGINT NOT NULL DEFAULT 0,
	`weighted_complexity`  FLOAT NOT NULL DEFAULT 0,
	INDEX commit_activity_sha_index (`sha`),
	INDEX commit_activity_user_id_index (`user_id`),
	INDEX commit_activity_repo_id_index (`repo_id`),
	INDEX commit_activity_filename_index (`filename`),
	INDEX commit_activity_filename_repo_id_date_ordinal_index (`filename`,`repo_id`,`date`,`ordinal`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Commit File is a generic table which describes a single file in a specific commit
CREATE TABLE `commit_file` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`commit_id`         VARCHAR(64) NOT NULL,
	`repo_id`           VARCHAR(64) NOT NULL,
	`author_user_id`    VARCHAR(64) NOT NULL,
	`committer_user_id` VARCHAR(64) NOT NULL,
	`filename`          VARCHAR(1024) NOT NULL,
	`language`          VARCHAR(100) NOT NULL DEFAULT "unknown",
	`additions`         INT NOT NULL DEFAULT 0,
	`deletions`         INT NOT NULL DEFAULT 0,
	`size`              INT NOT NULL DEFAULT 0,
	`abinary`           BOOL NOT NULL DEFAULT false,
	`date`              BIGINT UNSIGNED NOT NULL,
	`branch`            VARCHAR(255) NOT NULL DEFAULT "master",
	`mergecommit`       BOOL NOT NULL DEFAULT false,
	`excluded`          BOOL NOT NULL DEFAULT false,
	`loc`               INT NOT NULL DEFAULT 0,
	`sloc`              INT NOT NULL DEFAULT 0,
	`comments`          INT NOT NULL DEFAULT 0,
	`blanks`            INT NOT NULL DEFAULT 0,
	`variance`          INT NOT NULL DEFAULT 0,
	`status`            ENUM('added','modified','deleted') NOT NULL,
	`renamed`           BOOL NOT NULL DEFAULT false,
	`renamed_from`      TEXT,
	`renamed_to`        TEXT,
	`customer_id`       VARCHAR(64) NOT NULL,
	`ref_type`          VARCHAR(20) NOT NULL,
	`ref_id`            VARCHAR(64) NOT NULL,
	`metadata`          JSON,
	INDEX commit_file_commit_id_index (`commit_id`),
	INDEX commit_file_customer_id_index (`customer_id`),
	INDEX commit_file_ref_type_index (`ref_type`),
	INDEX commit_file_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Commit Issue is a generic table which links one or more issues to a Commit
CREATE TABLE `commit_issue` (
	`id`             VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`       CHAR(64),
	`commit_id`      VARCHAR(64) NOT NULL,
	`branch`         VARCHAR(255) NOT NULL DEFAULT "master",
	`user_id`        VARCHAR(64) NOT NULL,
	`repo_id`        VARCHAR(64) NOT NULL,
	`issue_id`       VARCHAR(64) NOT NULL,
	`date`           BIGINT UNSIGNED NOT NULL,
	`customer_id`    VARCHAR(64) NOT NULL,
	`ref_type`       VARCHAR(20) NOT NULL,
	`ref_commit_id`  VARCHAR(64) NOT NULL,
	`ref_repo_id`    VARCHAR(64) NOT NULL,
	`ref_issue_type` VARCHAR(20) NOT NULL,
	`ref_issue_id`   VARCHAR(64) NOT NULL,
	`metadata`       JSON,
	INDEX commit_issue_commit_id_index (`commit_id`),
	INDEX commit_issue_customer_id_index (`customer_id`),
	INDEX commit_issue_ref_type_index (`ref_type`),
	INDEX commit_issue_ref_commit_id_index (`ref_commit_id`),
	INDEX commit_issue_ref_repo_id_index (`ref_repo_id`),
	INDEX commit_issue_ref_issue_type_index (`ref_issue_type`),
	INDEX commit_issue_ref_issue_id_index (`ref_issue_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `commit_summary` (
	`id`              VARCHAR(64) NOT NULL PRIMARY KEY,
	`commit_id`       VARCHAR(64) NOT NULL,
	`sha`             VARCHAR(64) NOT NULL,
	`author_user_id`  VARCHAR(64) NOT NULL,
	`customer_id`     VARCHAR(64) NOT NULL,
	`data_group_id`   VARCHAR(64),
	`repo_id`         VARCHAR(64) NOT NULL,
	`repo`            TEXT NOT NULL,
	`ref_type`        VARCHAR(20) NOT NULL,
	`additions`       INT NOT NULL DEFAULT 0,
	`deletions`       INT NOT NULL DEFAULT 0,
	`files_changed`   INT NOT NULL DEFAULT 0,
	`branch`          VARCHAR(255) DEFAULT "master",
	`language`        VARCHAR(500) NOT NULL DEFAULT "unknown",
	`date`            BIGINT UNSIGNED NOT NULL,
	`message`         LONGTEXT,
	`excluded`        BOOL NOT NULL DEFAULT false,
	`exclude_reason`  VARCHAR(500),
	`mergecommit`     BOOL DEFAULT false,
	INDEX commit_summary_commit_id_index (`commit_id`),
	INDEX commit_summary_sha_index (`sha`),
	INDEX commit_summary_author_user_id_index (`author_user_id`),
	INDEX commit_summary_customer_id_index (`customer_id`),
	INDEX commit_summary_data_group_id_index (`data_group_id`),
	INDEX commit_summary_repo_id_index (`repo_id`),
	INDEX commit_summary_ref_type_index (`ref_type`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Cost Center is a type of cost detail container that can be re-used for multiple users
CREATE TABLE `cost_center` (
	`id`             VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`       CHAR(64),
	`customer_id`    VARCHAR(64) NOT NULL,
	`name`           TEXT NOT NULL,
	`description`    TEXT,
	`identifier`     TEXT,
	`unit_cost`      FLOAT NOT NULL,
	`value_type`     ENUM('absolute','percentage') NOT NULL,
	`unit_type`      ENUM('salary','hourly','contractor','casual','other') NOT NULL,
	`unit_currency`  CHAR(3) NOT NULL,
	`allocation`     FLOAT NOT NULL,
	`created_at`     BIGINT UNSIGNED NOT NULL,
	`updated_at`     BIGINT UNSIGNED
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Customer is a generic table for holding customer specific details
CREATE TABLE `customer` (
	`id`          VARCHAR(64) NOT NULL PRIMARY KEY,
	`name`        TEXT NOT NULL,
	`active`      BOOL NOT NULL DEFAULT false,
	`created_at`  BIGINT UNSIGNED NOT NULL,
	`updated_at`  BIGINT UNSIGNED,
	`metadata`    JSON
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Data Group is a way to organize how the UI groups and filters data
CREATE TABLE `data_group` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`customer_id`       VARCHAR(64) NOT NULL,
	`name`              TEXT NOT NULL,
	`description`       TEXT,
	`parent_id`         VARCHAR(64),
	`issue_project_ids` JSON,
	`user_ids`          JSON,
	`repo_ids`          JSON,
	`created_at`        BIGINT UNSIGNED NOT NULL,
	`updated_at`        BIGINT UNSIGNED
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  ExclusionMapping is a table which has all the excluded ref_id
CREATE TABLE `exclusion_mapping` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`ref_id`       VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20) NOT NULL,
	`user_id`      VARCHAR(64),
	`date`         BIGINT UNSIGNED,
	`customer_id`  VARCHAR(64) NOT NULL,
	INDEX exclusion_mapping_ref_id_index (`ref_id`),
	INDEX exclusion_mapping_ref_type_index (`ref_type`),
	INDEX exclusion_mapping_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Issue is a generic table for an issue
CREATE TABLE `issue` (
	`id`               VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`         CHAR(64),
	`user_id`          VARCHAR(64),
	`project_id`       VARCHAR(64) NOT NULL,
	`title`            TEXT NOT NULL,
	`body`             LONGTEXT,
	`state`            ENUM('open','closed') NOT NULL,
	`type`             ENUM('bug','feature','enhancement','other') NOT NULL,
	`resolution`       ENUM('none','resolved','unresolved') NOT NULL,
	`created_at`       BIGINT UNSIGNED NOT NULL,
	`updated_at`       BIGINT UNSIGNED,
	`closed_at`        BIGINT UNSIGNED,
	`closedby_user_id` VARCHAR(64),
	`url`              VARCHAR(255) NOT NULL,
	`customer_id`      VARCHAR(64) NOT NULL,
	`ref_type`         VARCHAR(20) NOT NULL,
	`ref_id`           VARCHAR(64) NOT NULL,
	`metadata`         JSON,
	INDEX issue_customer_id_index (`customer_id`),
	INDEX issue_ref_type_index (`ref_type`),
	INDEX issue_ref_id_index (`ref_id`),
	INDEX issue_customer_id_user_id_project_id_index (`customer_id`,`user_id`,`project_id`),
	INDEX issue_state_created_at_customer_id_index (`state`,`created_at`,`customer_id`),
	INDEX issue_type_customer_id_project_id_index (`type`,`customer_id`,`project_id`),
	INDEX issue_ref_type_state_project_id_customer_id_index (`ref_type`,`state`,`project_id`,`customer_id`),
	INDEX issue_ref_type_state_closed_at_project_id_customer_id_index (`ref_type`,`state`,`closed_at`,`project_id`,`customer_id`),
	INDEX issue_ref_type_state_created_at_project_id_customer_id_index (`ref_type`,`state`,`created_at`,`project_id`,`customer_id`),
	INDEX issue_state_customer_id_project_id_index (`state`,`customer_id`,`project_id`),
	INDEX issue_state_ref_id_closed_at_index (`state`,`ref_id`,`closed_at`),
	INDEX issue_state_customer_id_closed_at_ref_id_type_index (`state`,`customer_id`,`closed_at`,`ref_id`,`type`),
	INDEX issue_customer_id_project_id_created_at_index (`customer_id`,`project_id`,`created_at`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Issue Comment is a generic table which describes a comment for a issue
CREATE TABLE `issue_comment` (
	`id`              VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`        CHAR(64),
	`issue_id`        VARCHAR(64) NOT NULL,
	`user_id`         VARCHAR(64),
	`project_id`      VARCHAR(64) NOT NULL,
	`body`            LONGTEXT NOT NULL,
	`deleted`         BOOL NOT NULL DEFAULT false,
	`deleted_user_id` VARCHAR(64),
	`created_at`      BIGINT UNSIGNED NOT NULL,
	`updated_at`      BIGINT UNSIGNED,
	`deleted_at`      BIGINT UNSIGNED,
	`url`             VARCHAR(255) NOT NULL,
	`customer_id`     VARCHAR(64) NOT NULL,
	`ref_type`        VARCHAR(20) NOT NULL,
	`ref_id`          VARCHAR(64) NOT NULL,
	`metadata`        JSON,
	INDEX issue_comment_issue_id_index (`issue_id`),
	INDEX issue_comment_user_id_index (`user_id`),
	INDEX issue_comment_customer_id_index (`customer_id`),
	INDEX issue_comment_ref_type_index (`ref_type`),
	INDEX issue_comment_ref_id_index (`ref_id`),
	INDEX issue_comment_customer_id_user_id_project_id_index (`customer_id`,`user_id`,`project_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Issue is a generic table for a project which contains issues
CREATE TABLE `issue_project` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`name`         TEXT NOT NULL,
	`url`          TEXT NOT NULL,
	`created_at`   BIGINT UNSIGNED NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	`metadata`     JSON,
	INDEX issue_project_customer_id_index (`customer_id`),
	INDEX issue_project_ref_type_index (`ref_type`),
	INDEX issue_project_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `issue_rework_summary` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`project_id`   VARCHAR(64) NOT NULL,
	`user_id`      VARCHAR(64),
	`issue_id`     VARCHAR(64) NOT NULL,
	`path`         VARCHAR(615) NOT NULL,
	`date`         BIGINT UNSIGNED,
	INDEX issue_rework_summary_customer_id_index (`customer_id`),
	INDEX issue_rework_summary_project_id_index (`project_id`),
	INDEX issue_rework_summary_user_id_index (`user_id`),
	INDEX issue_rework_summary_issue_id_index (`issue_id`),
	INDEX issue_rework_summary_customer_id_project_id_user_id_index (`customer_id`,`project_id`,`user_id`),
	INDEX issue_rework_summary_customer_id_user_id_path_index (`customer_id`,`user_id`,`path`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  IssueSummary is a summarization table with issue level summary data
CREATE TABLE `issue_summary` (
	`id`                          VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`                    CHAR(64),
	`issue_id`                    VARCHAR(64) NOT NULL,
	`total_issues`                INT UNSIGNED NOT NULL,
	`new30_days`                  INT UNSIGNED NOT NULL,
	`total_closed`                INT UNSIGNED NOT NULL,
	`closed30_days`               INT UNSIGNED NOT NULL,
	`estimated_work_months`       FLOAT NOT NULL,
	`estimated_work_months30_days` FLOAT NOT NULL,
	`title`                       TEXT NOT NULL,
	`url`                         TEXT,
	`priority`                    VARCHAR(100),
	`priority_id`                 VARCHAR(64),
	`status`                      VARCHAR(100),
	`status_id`                   VARCHAR(64),
	`issue_type`                  VARCHAR(100) NOT NULL,
	`issue_type_id`               VARCHAR(64),
	`resolution`                  VARCHAR(100),
	`resolution_id`               VARCHAR(64),
	`state`                       VARCHAR(10) NOT NULL,
	`custom_field_ids`            JSON,
	`teams`                       JSON,
	`parent_issue_id`             VARCHAR(64),
	`parents_issue_ids`           JSON,
	`metadata`                    JSON,
	`project_id`                  VARCHAR(64) NOT NULL,
	`sprints`                     JSON,
	`labels`                      JSON,
	`top_level`                   TINYINT UNSIGNED NOT NULL,
	`is_leaf`                     TINYINT UNSIGNED NOT NULL,
	`path`                        VARCHAR(1024) NOT NULL,
	`in_progress_duration`        BIGINT NOT NULL,
	`verification_duration`       BIGINT NOT NULL,
	`in_progress_count`           INT NOT NULL,
	`reopen_count`                INT NOT NULL,
	`mapped_type`                 VARCHAR(75) NOT NULL,
	`strategic_parent_id`         VARCHAR(64),
	`sprint_id`                   JSON,
	`issue_project_name`          VARCHAR(255) NOT NULL,
	`users`                       JSON,
	`initial_start_date`          BIGINT UNSIGNED NOT NULL,
	`total_duration`              BIGINT UNSIGNED NOT NULL,
	`created_at`                  BIGINT UNSIGNED,
	`closed_at`                   BIGINT UNSIGNED,
	`planned_end_date`            BIGINT UNSIGNED,
	`customer_id`                 VARCHAR(64) NOT NULL,
	`ref_type`                    VARCHAR(20) NOT NULL,
	`ref_id`                      VARCHAR(64) NOT NULL,
	`custom_field_ids_virtual`    TEXT,
	`release_duration`            BIGINT NOT NULL,
	`completed`                   BOOL NOT NULL,
	`completed_date`              BIGINT UNSIGNED,
	INDEX issue_summary_issue_id_index (`issue_id`),
	INDEX issue_summary_priority_id_index (`priority_id`),
	INDEX issue_summary_issue_type_id_index (`issue_type_id`),
	INDEX issue_summary_resolution_id_index (`resolution_id`),
	INDEX issue_summary_parent_issue_id_index (`parent_issue_id`),
	INDEX issue_summary_project_id_index (`project_id`),
	INDEX issue_summary_top_level_index (`top_level`),
	INDEX issue_summary_strategic_parent_id_index (`strategic_parent_id`),
	INDEX issue_summary_customer_id_index (`customer_id`),
	INDEX issue_summary_ref_id_index (`ref_id`),
	INDEX issue_summary_customer_id_parent_issue_id_index (`customer_id`,`parent_issue_id`),
	INDEX issue_summary_customer_id_top_level_index (`customer_id`,`top_level`),
	INDEX issue_summary_customer_id_top_level_issue_type_id_index (`customer_id`,`top_level`,`issue_type_id`),
	INDEX issue_summary_customer_id_top_level_issue_type_id_priority_id_in (`customer_id`,`top_level`,`issue_type_id`,`priority_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JiraCustomField is a table which contains the unique custom fields found across all jira projects
CREATE TABLE `jira_custom_field` (
	`id`               VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`         CHAR(64),
	`field_id`         VARCHAR(64) NOT NULL,
	`name`             VARCHAR(255) NOT NULL,
	`schema_type`      VARCHAR(100) NOT NULL,
	`schema_custom`    VARCHAR(255) NOT NULL,
	`schema_custom_id` BIGINT UNSIGNED,
	`customer_id`      VARCHAR(64) NOT NULL,
	INDEX jira_custom_field_name_index (`name`),
	INDEX jira_custom_field_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JiraCustomFieldValue is a table which contains the unique custom field values found across all jira projects
CREATE TABLE `jira_custom_field_value` (
	`id`              VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`        CHAR(64),
	`schema_type`     VARCHAR(100) NOT NULL,
	`value`           LONGTEXT NOT NULL,
	`custom_field_id` VARCHAR(64) NOT NULL,
	`customer_id`     VARCHAR(64) NOT NULL,
	INDEX jira_custom_field_value_custom_field_id_index (`custom_field_id`),
	INDEX jira_custom_field_value_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA issue is a specialized version of a Issue for JIRA
CREATE TABLE `jira_issue` (
	`id`                     VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`               CHAR(64),
	`issue_id`               VARCHAR(255) NOT NULL,
	`project_id`             VARCHAR(64) NOT NULL,
	`issue_type_id`          VARCHAR(64) NOT NULL,
	`user_id`                VARCHAR(64),
	`assignee_id`            VARCHAR(64),
	`priority_id`            VARCHAR(64),
	`status_id`              VARCHAR(64) NOT NULL,
	`resolution_id`          VARCHAR(64),
	`fix_version_ids`        JSON,
	`version_ids`            JSON,
	`environment`            TEXT,
	`component_ids`          JSON,
	`label_ids`              JSON,
	`duedate_at`             BIGINT UNSIGNED,
	`planned_start_at`       BIGINT UNSIGNED,
	`planned_end_at`         BIGINT UNSIGNED,
	`key`                    VARCHAR(20) NOT NULL,
	`custom_field_ids`       JSON,
	`sprint_id`              JSON,
	`epic_id`                VARCHAR(64),
	`parent_id`              VARCHAR(64),
	`strategic_parent_id`    VARCHAR(64),
	`in_progress_count`      INT NOT NULL,
	`reopen_count`           INT NOT NULL,
	`in_progress_duration`   BIGINT NOT NULL,
	`verification_duration`  BIGINT NOT NULL,
	`in_progress_start_at`   BIGINT,
	`customer_id`            VARCHAR(64) NOT NULL,
	`ref_id`                 VARCHAR(64) NOT NULL,
	`user_ref_id`            VARCHAR(64) NOT NULL,
	`cost`                   REAL NOT NULL DEFAULT 0,
	INDEX jira_issue_issue_id_index (`issue_id`),
	INDEX jira_issue_project_id_index (`project_id`),
	INDEX jira_issue_issue_type_id_index (`issue_type_id`),
	INDEX jira_issue_key_index (`key`),
	INDEX jira_issue_epic_id_index (`epic_id`),
	INDEX jira_issue_parent_id_index (`parent_id`),
	INDEX jira_issue_strategic_parent_id_index (`strategic_parent_id`),
	INDEX jira_issue_customer_id_index (`customer_id`),
	INDEX jira_issue_ref_id_index (`ref_id`),
	INDEX jira_issue_user_ref_id_index (`user_ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `jira_issue_change_log` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`user_id`      VARCHAR(64),
	`assignee_id`  VARCHAR(64),
	`created_at`   BIGINT UNSIGNED,
	`field`        VARCHAR(255) NOT NULL,
	`field_type`   TEXT,
	`from`         LONGTEXT,
	`from_string`  LONGTEXT,
	`to`           LONGTEXT,
	`to_string`    LONGTEXT,
	`ref_id`       VARCHAR(64) NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	`issue_id`     VARCHAR(64) NOT NULL,
	`project_id`   VARCHAR(64) NOT NULL,
	`ordinal`      INT NOT NULL,
	INDEX jira_issue_change_log_field_index (`field`),
	INDEX jira_issue_change_log_customer_id_index (`customer_id`),
	INDEX jira_issue_change_log_issue_id_index (`issue_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA Issue Comment is a specialized version of a Issue Comment for JIRA issue comments
CREATE TABLE `jira_issue_comment` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`comment_id`   VARCHAR(255) NOT NULL,
	`issue_id`     VARCHAR(64) NOT NULL,
	`project_id`   VARCHAR(64) NOT NULL,
	`user_id`      VARCHAR(64),
	`updated_at`   BIGINT UNSIGNED NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	INDEX jira_issue_comment_comment_id_index (`comment_id`),
	INDEX jira_issue_comment_issue_id_index (`issue_id`),
	INDEX jira_issue_comment_project_id_index (`project_id`),
	INDEX jira_issue_comment_user_id_index (`user_id`),
	INDEX jira_issue_comment_customer_id_index (`customer_id`),
	INDEX jira_issue_comment_ref_id_index (`ref_id`),
	INDEX jira_issue_comment_customer_id_project_id_user_id_updated_at_ind (`customer_id`,`project_id`,`user_id`,`updated_at`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA Issue Label is a unique label for a given project
CREATE TABLE `jira_issue_label` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`name`         VARCHAR(255) NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	INDEX jira_issue_label_name_index (`name`),
	INDEX jira_issue_label_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA Issue Priority is the priority table
CREATE TABLE `jira_issue_priority` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`issue_priority_id` VARCHAR(255) NOT NULL,
	`name`              TEXT NOT NULL,
	`icon_url`          TEXT NOT NULL,
	`customer_id`       VARCHAR(64) NOT NULL,
	INDEX jira_issue_priority_issue_priority_id_index (`issue_priority_id`),
	INDEX jira_issue_priority_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA Issue Progress tracks the progress for a given issue by user
CREATE TABLE `jira_issue_progress` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`user_id`      VARCHAR(64),
	`issue_id`     VARCHAR(64) NOT NULL,
	`start_date`   BIGINT UNSIGNED,
	`end_date`     BIGINT UNSIGNED,
	`duration`     BIGINT UNSIGNED,
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	INDEX jira_issue_progress_user_id_index (`user_id`),
	INDEX jira_issue_progress_issue_id_index (`issue_id`),
	INDEX jira_issue_progress_customer_id_index (`customer_id`),
	INDEX jira_issue_progress_ref_id_index (`ref_id`),
	INDEX jira_issue_progress_user_id_customer_id_start_date_index (`user_id`,`customer_id`,`start_date`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `jira_issue_type` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`issue_type_id` VARCHAR(64) NOT NULL,
	`name`          VARCHAR(255) NOT NULL,
	`description`   TEXT,
	`icon_url`      VARCHAR(255),
	`url`           VARCHAR(255),
	`subtask`       BOOL NOT NULL,
	`customer_id`   VARCHAR(64) NOT NULL,
	INDEX jira_issue_type_name_index (`name`),
	INDEX jira_issue_type_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA project
CREATE TABLE `jira_project` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`project_id`   VARCHAR(255) NOT NULL,
	`key`          VARCHAR(100) NOT NULL,
	`avatar_url`   VARCHAR(255) NOT NULL,
	`category_id`  VARCHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	INDEX jira_project_project_id_index (`project_id`),
	INDEX jira_project_customer_id_index (`customer_id`),
	INDEX jira_project_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `jira_project_resolution` (
	`id`             VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`       CHAR(64),
	`resolution_id`  VARCHAR(64) NOT NULL,
	`project_id`     VARCHAR(64) NOT NULL,
	`name`           VARCHAR(255) NOT NULL,
	`description`    TEXT,
	`customer_id`    VARCHAR(64) NOT NULL,
	INDEX jira_project_resolution_project_id_index (`project_id`),
	INDEX jira_project_resolution_name_index (`name`),
	INDEX jira_project_resolution_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `jira_project_sprint` (
	`id`                          VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`                    CHAR(64),
	`datagroup_id`                VARCHAR(64) NOT NULL,
	`project_ids`                 JSON NOT NULL,
	`sprint_id`                   TEXT NOT NULL,
	`name`                        TEXT NOT NULL,
	`customer_id`                 VARCHAR(64) NOT NULL,
	`complete_date`               BIGINT UNSIGNED,
	`end_date`                    BIGINT UNSIGNED,
	`start_date`                  BIGINT UNSIGNED,
	`state`                       TEXT,
	`goal`                        TEXT,
	`initial_issue_ids`           JSON,
	`final_issue_ids`             JSON,
	`final_closed_issue_ids`      JSON,
	`initial_issue_count`         INT NOT NULL,
	`final_issue_count`           INT NOT NULL,
	`closed_issue_count`          INT NOT NULL,
	`added_issue_count`           INT NOT NULL,
	`removed_issue_count`         INT NOT NULL,
	`initial_issues_closed`       INT NOT NULL,
	`initial_issues_in_final_count` INT NOT NULL,
	INDEX jira_project_sprint_datagroup_id_index (`datagroup_id`),
	INDEX jira_project_sprint_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `jira_project_status` (
	`id`              VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`        CHAR(64),
	`status_id`       VARCHAR(64) NOT NULL,
	`project_id`      VARCHAR(64) NOT NULL,
	`name`            VARCHAR(255) NOT NULL,
	`description`     TEXT,
	`icon_url`        TEXT,
	`category`        TEXT,
	`category_key`    TEXT,
	`category_color`  TEXT,
	`category_id`     VARCHAR(64),
	`issue_type_id`   VARCHAR(64) NOT NULL,
	`customer_id`     VARCHAR(64) NOT NULL,
	INDEX jira_project_status_project_id_index (`project_id`),
	INDEX jira_project_status_name_index (`name`),
	INDEX jira_project_status_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  JIRA user is a specific user
CREATE TABLE `jira_user` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`user_id`      VARCHAR(255) NOT NULL,
	`username`     TEXT NOT NULL,
	`name`         TEXT NOT NULL,
	`email`        VARCHAR(255) NOT NULL,
	`avatar_url`   VARCHAR(255) NOT NULL,
	`url`          VARCHAR(255) NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	INDEX jira_user_user_id_index (`user_id`),
	INDEX jira_user_customer_id_index (`customer_id`),
	INDEX jira_user_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `mockapi_application` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`ext_id`       BIGINT NOT NULL,
	`name`         TEXT NOT NULL,
	INDEX mockapi_application_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `mockapi_deployment` (
	`id`                      VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`                CHAR(64),
	`customer_id`             VARCHAR(64) NOT NULL,
	`application_ext_id_ext_id` BIGINT NOT NULL,
	`application_ext_id_id`   VARCHAR(64) NOT NULL,
	`ext_id`                  BIGINT NOT NULL,
	`name`                    TEXT NOT NULL,
	INDEX mockapi_deployment_customer_id_index (`customer_id`),
	INDEX mockapi_deployment_application_ext_id_ext_id_index (`application_ext_id_ext_id`),
	INDEX mockapi_deployment_application_ext_id_id_index (`application_ext_id_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Processing Error is used to track errors during processing
CREATE TABLE `processing_error` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`created_at`   BIGINT UNSIGNED NOT NULL,
	`type`         TEXT NOT NULL,
	`message`      MEDIUMTEXT NOT NULL,
	`fatal`        BOOL NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	INDEX processing_error_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Repo is a generic table which describes a code repository
CREATE TABLE `repo` (
	`id`            VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`      CHAR(64),
	`name`          VARCHAR(255) NOT NULL,
	`url`           VARCHAR(255) NOT NULL,
	`active`        BOOL NOT NULL DEFAULT true,
	`fork`          BOOL NOT NULL DEFAULT false,
	`parent_id`     VARCHAR(64),
	`description`   TEXT,
	`created_at`    BIGINT UNSIGNED NOT NULL,
	`updated_at`    BIGINT UNSIGNED,
	`customer_id`   VARCHAR(64) NOT NULL,
	`ref_type`      VARCHAR(20) NOT NULL,
	`ref_id`        VARCHAR(64) NOT NULL,
	`metadata`      JSON,
	INDEX repo_name_index (`name`),
	INDEX repo_customer_id_index (`customer_id`),
	INDEX repo_ref_type_index (`ref_type`),
	INDEX repo_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `repo_mapping` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`repo_id`      VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20) NOT NULL,
	INDEX repo_mapping_repo_id_index (`repo_id`),
	INDEX repo_mapping_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  RepoSummary table stores a summary for each repository
CREATE TABLE `repo_summary` (
	`id`                 VARCHAR(64) NOT NULL PRIMARY KEY,
	`name`               VARCHAR(255) NOT NULL,
	`description`        TEXT,
	`ref_type`           VARCHAR(20) NOT NULL,
	`data_group_id`      VARCHAR(20),
	`user_ids`           JSON NOT NULL,
	`commits`            BIGINT UNSIGNED NOT NULL,
	`additions`          BIGINT UNSIGNED NOT NULL,
	`deletions`          BIGINT UNSIGNED NOT NULL,
	`latest_commit_date` BIGINT UNSIGNED NOT NULL,
	`repo_id`            VARCHAR(64) NOT NULL,
	`customer_id`        VARCHAR(64) NOT NULL,
	INDEX repo_summary_ref_type_index (`ref_type`),
	INDEX repo_summary_data_group_id_index (`data_group_id`),
	INDEX repo_summary_repo_id_index (`repo_id`),
	INDEX repo_summary_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  Signal table is a generic set of signals that are calculated behind the scenes
CREATE TABLE `signal` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`name`         CHAR(60) NOT NULL,
	`value`        REAL NOT NULL DEFAULT 0,
	`timeunit`     INT NOT NULL,
	`date`         DATE NOT NULL,
	`metadata`     JSON,
	`customer_id`  VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20),
	`ref_id`       VARCHAR(64),
	INDEX signal_name_index (`name`),
	INDEX signal_customer_id_index (`customer_id`),
	INDEX signal_ref_type_index (`ref_type`),
	INDEX signal_ref_id_index (`ref_id`),
	INDEX signal_name_timeunit_customer_id_index (`name`,`timeunit`,`customer_id`),
	INDEX signal_name_timeunit_customer_id_date_index (`name`,`timeunit`,`customer_id`,`date`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `sonarqube_metric` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`customer_id`       VARCHAR(64) NOT NULL,
	`project_ext_id`    VARCHAR(255) NOT NULL,
	`project_id`        VARCHAR(64) NOT NULL,
	`project_key_ext_id` VARCHAR(255) NOT NULL,
	`ext_id`            VARCHAR(255) NOT NULL,
	`date`              BIGINT NOT NULL,
	`metric`            VARCHAR(64) NOT NULL,
	`value`             VARCHAR(64) NOT NULL,
	INDEX sonarqube_metric_customer_id_index (`customer_id`),
	INDEX sonarqube_metric_project_id_index (`project_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `sonarqube_project` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`ext_id`       TEXT NOT NULL,
	`key`          TEXT NOT NULL,
	`name`         TEXT NOT NULL,
	INDEX sonarqube_project_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  System is a mapping table which maps third-party integration to one or more repositories
CREATE TABLE `system` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`repo_id`      VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20) NOT NULL,
	`customer_id`  VARCHAR(64) NOT NULL,
	INDEX system_repo_id_index (`repo_id`),
	INDEX system_customer_id_index (`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  User is a generic table which maps a user into a specific customer organization
CREATE TABLE `user` (
	`id`                VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`          CHAR(64),
	`customer_id`       VARCHAR(64) NOT NULL,
	`email`             VARCHAR(255),
	`name`              TEXT,
	`username`          VARCHAR(100) NOT NULL,
	`active`            BOOL NOT NULL,
	`autoenrolled`      BOOL NOT NULL,
	`visible`           BOOL NOT NULL,
	`avatar_url`        VARCHAR(255) NOT NULL,
	`employee_id`       TEXT,
	`location`          TEXT,
	`region`            TEXT,
	`department`        TEXT,
	`reports_to_user_id` VARCHAR(64),
	`title`             TEXT,
	`role`              TEXT,
	`created_at`        BIGINT UNSIGNED NOT NULL,
	`updated_at`        BIGINT UNSIGNED,
	`metadata`          JSON,
	`ref_id`            VARCHAR(64),
	`ref_type`          VARCHAR(20),
	INDEX user_email_index (`email`),
	INDEX user_username_customer_id_index (`username`,`customer_id`),
	INDEX user_email_customer_id_index (`email`,`customer_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  User Cost Center is a specific record for a users cost details
CREATE TABLE `user_cost_center` (
	`id`             VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`       CHAR(64),
	`customer_id`    VARCHAR(64) NOT NULL,
	`user_id`        VARCHAR(64) NOT NULL,
	`cost_center_id` VARCHAR(64),
	`unit_cost`      FLOAT,
	`unit_type`      ENUM('salary','hourly','contractor','casual','other'),
	`unit_currency`  CHAR(3),
	`allocation`     FLOAT,
	`created_at`     BIGINT UNSIGNED NOT NULL,
	`updated_at`     BIGINT UNSIGNED
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--  User Login is a generic table for tracking user logins
CREATE TABLE `user_login` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`user_id`      VARCHAR(64) NOT NULL,
	`browser`      TEXT,
	`date`         BIGINT UNSIGNED NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_mapping` (
	`id`           VARCHAR(64) NOT NULL PRIMARY KEY,
	`checksum`     CHAR(64),
	`customer_id`  VARCHAR(64) NOT NULL,
	`user_id`      VARCHAR(64) NOT NULL,
	`ref_id`       VARCHAR(64) NOT NULL,
	`ref_type`     VARCHAR(20) NOT NULL,
	INDEX user_mapping_user_id_index (`user_id`),
	INDEX user_mapping_ref_id_index (`ref_id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS `acl_grant`;
DROP TABLE IF EXISTS `acl_resource`;
DROP TABLE IF EXISTS `acl_role`;
DROP TABLE IF EXISTS `acl_user_role`;
DROP TABLE IF EXISTS `app`;
DROP TABLE IF EXISTS `commit`;
DROP TABLE IF EXISTS `commit_activity`;
DROP TABLE IF EXISTS `commit_file`;
DROP TABLE IF EXISTS `commit_issue`;
DROP TABLE IF EXISTS `commit_summary`;
DROP TABLE IF EXISTS `cost_center`;
DROP TABLE IF EXISTS `customer`;
DROP TABLE IF EXISTS `data_group`;
DROP TABLE IF EXISTS `exclusion_mapping`;
DROP TABLE IF EXISTS `issue`;
DROP TABLE IF EXISTS `issue_comment`;
DROP TABLE IF EXISTS `issue_project`;
DROP TABLE IF EXISTS `issue_rework_summary`;
DROP TABLE IF EXISTS `issue_summary`;
DROP TABLE IF EXISTS `jira_custom_field`;
DROP TABLE IF EXISTS `jira_custom_field_value`;
DROP TABLE IF EXISTS `jira_issue`;
DROP TABLE IF EXISTS `jira_issue_change_log`;
DROP TABLE IF EXISTS `jira_issue_comment`;
DROP TABLE IF EXISTS `jira_issue_label`;
DROP TABLE IF EXISTS `jira_issue_priority`;
DROP TABLE IF EXISTS `jira_issue_progress`;
DROP TABLE IF EXISTS `jira_issue_type`;
DROP TABLE IF EXISTS `jira_project`;
DROP TABLE IF EXISTS `jira_project_resolution`;
DROP TABLE IF EXISTS `jira_project_sprint`;
DROP TABLE IF EXISTS `jira_project_status`;
DROP TABLE IF EXISTS `jira_user`;
DROP TABLE IF EXISTS `mockapi_application`;
DROP TABLE IF EXISTS `mockapi_deployment`;
DROP TABLE IF EXISTS `processing_error`;
DROP TABLE IF EXISTS `repo`;
DROP TABLE IF EXISTS `repo_mapping`;
DROP TABLE IF EXISTS `repo_summary`;
DROP TABLE IF EXISTS `signal`;
DROP TABLE IF EXISTS `sonarqube_metric`;
DROP TABLE IF EXISTS `sonarqube_project`;
DROP TABLE IF EXISTS `system`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `user_cost_center`;
DROP TABLE IF EXISTS `user_login`;
DROP TABLE IF EXISTS `user_mapping`;
