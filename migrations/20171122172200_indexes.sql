-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- the Partition below needs to remove the primary key since it's not useful anyway
ALTER TABLE `signal` DROP PRIMARY KEY, ADD PRIMARY KEY(id, date, name);

-- the Partition will be by YEAR and then subpartitioned by signal
ALTER TABLE `signal` PARTITION BY RANGE(YEAR(date)) 
  SUBPARTITION BY KEY(name)
  SUBPARTITIONS 12 (
		PARTITION p0 VALUES LESS THAN (2010),
		PARTITION p1 VALUES LESS THAN (2011),
		PARTITION p2 VALUES LESS THAN (2012),
		PARTITION p3 VALUES LESS THAN (2013),
		PARTITION p4 VALUES LESS THAN (2014),
		PARTITION p5 VALUES LESS THAN (2015),
		PARTITION p6 VALUES LESS THAN (2016),
		PARTITION p7 VALUES LESS THAN (2017),
		PARTITION p8 VALUES LESS THAN (2018),
		PARTITION p9 VALUES LESS THAN (2019),
		PARTITION p10 VALUES LESS THAN (2020),
		PARTITION p11 VALUES LESS THAN (2021),
		PARTITION p99 VALUES LESS THAN MAXVALUE
);

CREATE INDEX issue_state_customer_id_project_id_closed_at_ref_id_type_index ON issue(state, customer_id, project_id, closed_at, ref_id, type);
CREATE INDEX jira_issue_change_log_issue_id_field_index ON jira_issue_change_log(issue_id, field);
CREATE INDEX issue_summary_customer_id_planned_closed_state_total_issues on issue_summary(customer_id, planned_end_date, closed_at, state, total_issues);
CREATE INDEX issue_summary_customer_id_planned_end_date_state_total_issues on issue_summary(customer_id, planned_end_date, state, total_issues);

-- create an index for the generated virtual field
ALTER TABLE issue_summary ADD FULLTEXT INDEX `custom_fields_full_idx`(`custom_field_ids_virtual`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

