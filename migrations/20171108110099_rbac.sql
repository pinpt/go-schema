-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("34c41b9c8565aade","d48b5fab5ea9d63e","urn:webapp:/admin","","",false,true,true,1542928548480);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("243b7977676a7a37","040ee90cc9764f64","urn:webapp:/admin/activate","","Admin Activation",true,true,false,1542928548480);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("40e07d8fe3f69e9b","c511e7aef7f67497","urn:webapp:/admin/cost-center","Manage cost centers and assign them to employees","Cost Centers",false,false,true,1542928548480);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3a73baba427edf09","b8c77b8ff5e8b38a","urn:webapp:/admin/mapping","Associate projects, repositories and people with teams","Team Mapping",false,false,true,1542928548480);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("cc30a417700eb3bc","a49b75671938da8e","urn:webapp:/admin/roles","Manage user roles and assign them to employees","Roles",false,false,true,1542928548480);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("ef46db3751d8e999","42e8e62f88393464","urn:feature:cost/all","Disabling salary data will remove salary and cost data from all views, including 2x2 cost oriented matrixes","Salaries",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("d36114c721917300","4d72fb65bd95fda0","urn:webapp:/error","","Error",true,true,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("32780b49037be639","8330c23183767aa5","urn:webapp:/jira/:team","Usage and adoption data for all Jira projects","Technology - Jira",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("03f3eeaa722d8d19","34727d53b42af770","urn:webapp:/languages","Contributions by programming language for people and projects","Technology - Languages",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("e89cd67289eddaea","3e58dbb06c0db070","urn:webapp:/","","",false,true,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("f720f0775840fe45","65cad1178bc3c2b2","urn:webapp:/signal/person/changes-per-commit/:id","","People - Changes Per Commit",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("bb2043b9f4db8d0f","9f8c1dfcc7245f64","urn:webapp:/signal/person/code-ownership/:id","","People - Code Ownership",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("be202bc02517c6bc","03030037e893460b","urn:webapp:/signal/person/commits/:id","","People - Commits",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("07b8e659030155b0","ccb2fcd3fbb0253b","urn:webapp:/signal/person/cycle-time/:id","","People - Cycle Time",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("bd8d3ecdfdf449cf","0d9696f256d93666","urn:webapp:/signal/person/issues-worked/:id","","People - Issues Worked",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("f5074c3084e4a0a9","0f905da1618bc2ce","urn:webapp:/performance/people/:filter/:id","Compare employees across a range of signals","People - Performance Summary",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("deb0f3669064471b","224bb562acbe8042","urn:webapp:/performance-detail/people/:person","Detailed performance signals for a specific person","People - Performance Detail",false,true,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("269063b0706f2c8f","3d4136ee86584490","urn:webapp:/signal/person/rework-rate/:id","","People - Rework Rate",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("6658efa46325d3aa","e355c2d4e26be42c","urn:webapp:/signal/person/traceability/:id","","People - Commit Traceability",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3f5e913a8b0bec3b","2384cbf73740c2c5","urn:webapp:/signal/team/backlog-change/:id","","Team - Backlog Change",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0835c75f474868da","b3262a101bfb7af2","urn:webapp:/signal/team/cost/:id","","Team - Cost",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("77da3e502145129c","b7dac9608654ad3f","urn:webapp:/signal/team/cycle-time/:id","","Team - Cycle Time",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("fc105dad5a73e677","6a8fabb5f313e980","urn:webapp:/signal/team/defects-density/:id","","Team - Defects Density",false,false,false,1542928548481);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("b6bb43d71557263f","8a5440a7e894f357","urn:webapp:/signal/team/defects-rate/:id","","Team - Defects Rate",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("5ab91df03468bda4","22dad4d16df90310","urn:webapp:/signal/team/delivered-vs-committed/:id","","Team - Delivered vs. Committed",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("ccc522ac53513655","67627f72da7b7140","urn:webapp:/signal/team/issues-closed/:id","","Team - Issues Closed",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("1dfd81e333fcb5f5","bfb99dbbe810fa9c","urn:webapp:/signal/team/innovation-rate/:id","","Team - Innovation Rate",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3020ad6b9c6cb452","0c74dd85982b5b41","urn:webapp:/performance/team/:team","Compare teams across a range of signals","Team - Performance Summary",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("485e9c90439126cf","fa93cf64687d92f2","urn:webapp:/performance-detail/team/:team","Detailed performance signals for a specific team","Team - Performance Detail",false,true,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("5c9d5357537e97f6","e3ac3d8f5b17be05","urn:webapp:/signal/team/rework-rate/:id","","Team - Rework Rate",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0094d897e70b3318","08430c0dd8cbc2ff","urn:webapp:/signal/team/scheduled-rate/:id","","Team - Scheduled Rate",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("d10a16ac5d42bcba","a44f96df2e8e1dff","urn:webapp:/signal/team/sprint-volatility/:id","","Team - Sprint Volatility",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("6fbbf822c566bf77","780f26c8bb4270bc","urn:webapp:/signal/team/initiative-issues/:id","","Team - Initiative Issues",false,false,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("67aa3b4fdeb14ba2","b7e2ee974458e903","urn:webapp:/welcome","","Welcome",true,true,false,1542928548482);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("8e56607a319a7c79","74df32cd953e911a","urn:webapp:/work-forecast/:team/:issueType/:priority/:id/:customField","Status and estimates for open issues","Initiative - Work Forecast",false,false,false,1542928548482);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("1f15a4c9d4798230","admin","Administrative role",1541002044312);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("3ae149a8dc444bce","executive","Executive role",1541002044312);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("455511f8e362deb9","manager","Manager role",1541002044312);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DELETE FROM `acl_resource` WHERE `id` = "34c41b9c8565aade";
DELETE FROM `acl_resource` WHERE `id` = "243b7977676a7a37";
DELETE FROM `acl_resource` WHERE `id` = "40e07d8fe3f69e9b";
DELETE FROM `acl_resource` WHERE `id` = "3a73baba427edf09";
DELETE FROM `acl_resource` WHERE `id` = "cc30a417700eb3bc";
DELETE FROM `acl_resource` WHERE `id` = "ef46db3751d8e999";
DELETE FROM `acl_resource` WHERE `id` = "d36114c721917300";
DELETE FROM `acl_resource` WHERE `id` = "32780b49037be639";
DELETE FROM `acl_resource` WHERE `id` = "03f3eeaa722d8d19";
DELETE FROM `acl_resource` WHERE `id` = "e89cd67289eddaea";
DELETE FROM `acl_resource` WHERE `id` = "f720f0775840fe45";
DELETE FROM `acl_resource` WHERE `id` = "bb2043b9f4db8d0f";
DELETE FROM `acl_resource` WHERE `id` = "be202bc02517c6bc";
DELETE FROM `acl_resource` WHERE `id` = "07b8e659030155b0";
DELETE FROM `acl_resource` WHERE `id` = "bd8d3ecdfdf449cf";
DELETE FROM `acl_resource` WHERE `id` = "f5074c3084e4a0a9";
DELETE FROM `acl_resource` WHERE `id` = "deb0f3669064471b";
DELETE FROM `acl_resource` WHERE `id` = "269063b0706f2c8f";
DELETE FROM `acl_resource` WHERE `id` = "6658efa46325d3aa";
DELETE FROM `acl_resource` WHERE `id` = "3f5e913a8b0bec3b";
DELETE FROM `acl_resource` WHERE `id` = "0835c75f474868da";
DELETE FROM `acl_resource` WHERE `id` = "77da3e502145129c";
DELETE FROM `acl_resource` WHERE `id` = "fc105dad5a73e677";
DELETE FROM `acl_resource` WHERE `id` = "b6bb43d71557263f";
DELETE FROM `acl_resource` WHERE `id` = "5ab91df03468bda4";
DELETE FROM `acl_resource` WHERE `id` = "ccc522ac53513655";
DELETE FROM `acl_resource` WHERE `id` = "1dfd81e333fcb5f5";
DELETE FROM `acl_resource` WHERE `id` = "3020ad6b9c6cb452";
DELETE FROM `acl_resource` WHERE `id` = "485e9c90439126cf";
DELETE FROM `acl_resource` WHERE `id` = "5c9d5357537e97f6";
DELETE FROM `acl_resource` WHERE `id` = "0094d897e70b3318";
DELETE FROM `acl_resource` WHERE `id` = "d10a16ac5d42bcba";
DELETE FROM `acl_resource` WHERE `id` = "6fbbf822c566bf77";
DELETE FROM `acl_resource` WHERE `id` = "67aa3b4fdeb14ba2";
DELETE FROM `acl_resource` WHERE `id` = "8e56607a319a7c79";
DELETE FROM `acl_role` WHERE `id` = "1f15a4c9d4798230";
DELETE FROM `acl_role` WHERE `id` = "3ae149a8dc444bce";
DELETE FROM `acl_role` WHERE `id` = "455511f8e362deb9";
