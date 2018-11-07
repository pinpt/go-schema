-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("34c41b9c8565aade","d48b5fab5ea9d63e","urn:webapp:/admin","","",false,true,true,1541599525701);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("40e07d8fe3f69e9b","f080518c934677b5","urn:webapp:/admin/cost-center","Create and manage cost centers","Cost Centers",false,false,true,1541599525701);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3a73baba427edf09","90de81dd5b05791f","urn:webapp:/admin/mapping","Associate projects and code repositories with teams","Team Mapping",false,false,true,1541599525701);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0e670c147cedd51a","75dca0c20fc8a322","urn:webapp:/admin/people","Manage your users","People",false,false,true,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("cc30a417700eb3bc","c4d771629ca8f1ff","urn:webapp:/admin/roles","Create and manage roles","Roles",false,false,true,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("c63df0d00578b6dd","8ba6f4b8651d9e78","urn:webapp:/commit/:id","","Commit Detail",false,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("ef46db3751d8e999","42e8e62f88393464","urn:feature:cost/all","Disabling salary data will remove salary and cost data from all views, including 2x2 cost oriented matrixes","Salaries",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("298dd2d0c37f9078","3d6b631ec1fa9b6e","urn:webapp:/cost-center/:id","","Cost Center Detail",false,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("a25ca86fead82e11","d6d78a68472a9c02","urn:webapp:/data/commits","","Data - Commits",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("330667884b05e787","fd0daa3d2540490a","urn:webapp:/data/issues","","Data - Issues",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3333c1aa2fd45297","576e60f92224ed7f","urn:webapp:/data/locations","","Data - Locations",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("19d46d2f7501a9b2","ad811c57b2b861a7","urn:webapp:/data/people","","Data - People",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("bec1bebaacc0e936","5c0d0fd77f7db091","urn:webapp:/data/projects","","Data - Projects",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("34d48343ea38eba3","7a08177ff67554a0","urn:webapp:/data/repositories","","Data - Repositories",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("dbf0617630089e7a","0962b8cbd9cb4678","urn:webapp:/data/sprints","","Data - Sprints",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("c6da1f7bb1295a51","0f005f7059fb873c","urn:webapp:/data/teams","","Data - Teams",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("d36114c721917300","4d72fb65bd95fda0","urn:webapp:/error","","Error",true,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("13110b47edd72e37","10c6c74052b983a2","urn:webapp:/file/:id","","File Detail",false,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("6029394193209d74","03af75f68226a933","urn:webapp:/issue/:id","","Issue Detail",false,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("74559779b75eeb96","31677f49fc0810e5","urn:webapp:/issues/forecast","Summary of forecasts and delivery status for larger open issues","Work - Forecast",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("fd6afe817ce547e6","8ffd7854e5d5fb1c","urn:webapp:/issues/workflow/:team/:issueType/:interval","Analysis of typical issue workflow paths","Work - Issue Workflow",false,false,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("ba1fd6ced759206e","c063f3913ec217f0","urn:webapp:/language/:id","","Language Detail",false,true,false,1541599525702);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("52a6c63ca841ad5f","ad3d422a24ff02af","urn:webapp:/location/:id","","Location Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("c9cdbd81906a0b73","6972ed7c49c797af","urn:webapp:/locations/performance","Compare locations across a range of signals","Location - Performance Summary",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("e89cd67289eddaea","3e58dbb06c0db070","urn:webapp:/","","",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0f41068501c9af03","4754d1f22bda445f","urn:webapp:/person/:id","","Person Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("f48ef8eec7ae8c01","3015db514f68689c","urn:webapp:/signal/changes-per-commit/person/:id","","People - Changes Per Commit",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3a1252027a07f6c8","dfc9afbaee87c5be","urn:webapp:/signal/code-ownership/person/:id","","People - Code Ownership",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3ec6f3c826e1fb2b","a14f0a84549d6e72","urn:webapp:/signal/commits/person/:id","","People - Commits",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("8f4cb2d1d5b57025","bbabdeb09c0708f9","urn:webapp:/signal/cycle-time/person/:id","","People - Cycle Time",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("cd7384eeb3d5804a","3b53d0c57863aa00","urn:webapp:/signal/issues-worked/person/:id","","People - Issues Worked",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("757dbe7a3f28a531","c835136a73086025","urn:webapp:/people/performance","Compare employees across a range of signals","People - Performance Summary",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("80ec82b56b24f003","d1fd1d352debedf1","urn:webapp:/signal/rework-rate/person/:id","","People - Rework Rate",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("2184738413722241","74d804d51fecc642","urn:webapp:/signal/traceability/person/:id","","People - Traceability",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("3285ac4d4adf482c","5e9dc70dedfd52c4","urn:webapp:/project/:id","","Project Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("6ca3232c13d02ae7","3e165f4be8cf8726","urn:webapp:/repository/:id","","Repository Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("113050129c759805","a36f25c3179859cb","urn:webapp:/sprint/:id","","Sprint Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("e711a6ac576cc326","b06eac9ac28400fe","urn:webapp:/team/:id","","Team Detail",false,true,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("f521c0a35aae5dfb","e07837c4f8f62831","urn:webapp:/signal/backlog-change/team/:id","","Teams - Backlog Change",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0a5b4a25a61980cb","6fed9ca0c05e2e63","urn:webapp:/signal/cost/team/:id","","Teams - Cost",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("9f637e34a1c99420","b61e02c1d1625edf","urn:webapp:/signal/cycle-time/team/:id","","Teams - Cycle Time",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("25db0e044f487e8e","c16135d4f9551d92","urn:webapp:/signal/defects-density/team/:id","","Teams - Defect Density",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("cf22f78a0079ed93","72c70150f929d89a","urn:webapp:/signal/defects-rate/team/:id","","Teams - Defect Rate",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("18f633bff893dea7","b38fd4f69d83cdc1","urn:webapp:/signal/delivered-vs-committed/team/:id","","Teams - Delivered vs. Planned",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0e494e5d37db0e0d","f74d914e2be4510c","urn:webapp:/signal/issues-completed/team/:id","","Teams - Completed Issues",false,false,false,1541599525703);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("fbaf66fd9463735b","72f05962b048d165","urn:webapp:/signal/innovation-rate/team/:id","","Teams - Innovation Rate",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("88643c94b2d7ffc1","5280a4f6f504eac9","urn:webapp:/signal/on-time-delivery/team/:id","","Teams - On-time Delivery",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("e184024d5e546adb","0ac14527a14b8ee0","urn:webapp:/teams/performance","Compare teams across a range of signals","Teams - Performance Summary",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("0e39bfa412e78db0","4debac0fb670a61d","urn:webapp:/signal/rework-rate/team/:id","","Teams - Rework Rate",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("8009c6c9688e430c","45cfa45924d86273","urn:webapp:/signal/scheduled-rate/team/:id","","Teams - Planned Issues",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("52c9e0e1ac1f01d6","2ce1f663855362b2","urn:webapp:/signal/sprint-health/team/:id","","Teams - Sprint Health",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("2181e7028f08a3cb","60031e9344813d9c","urn:webapp:/signal/sprint-volatility/team/:id","","Teams - Sprint Volatility",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("fcd92eebe06fc77f","21ac4d4ddcf05614","urn:webapp:/signal/initiative-issues/team/:id","","Teams - Strategic Issues",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("15553ff67e31ca4c","cc87731dc72581e2","urn:webapp:/signal/throughput/team/:id","","Teams - Throughput",false,false,false,1541599525704);
INSERT INTO `acl_resource` (`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`) VALUES ("67aa3b4fdeb14ba2","b7e2ee974458e903","urn:webapp:/welcome","","Welcome",true,true,false,1541599525704);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("1f15a4c9d4798230","admin","Administrative role",1541002044312);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("3ae149a8dc444bce","executive","Executive role",1541002044312);
INSERT INTO `acl_role` (`id`,`name`,`description`,`created_at`) VALUES ("455511f8e362deb9","manager","Manager role",1541002044312);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DELETE FROM `acl_resource` WHERE `id` = "34c41b9c8565aade";
DELETE FROM `acl_resource` WHERE `id` = "40e07d8fe3f69e9b";
DELETE FROM `acl_resource` WHERE `id` = "3a73baba427edf09";
DELETE FROM `acl_resource` WHERE `id` = "0e670c147cedd51a";
DELETE FROM `acl_resource` WHERE `id` = "cc30a417700eb3bc";
DELETE FROM `acl_resource` WHERE `id` = "c63df0d00578b6dd";
DELETE FROM `acl_resource` WHERE `id` = "ef46db3751d8e999";
DELETE FROM `acl_resource` WHERE `id` = "298dd2d0c37f9078";
DELETE FROM `acl_resource` WHERE `id` = "a25ca86fead82e11";
DELETE FROM `acl_resource` WHERE `id` = "330667884b05e787";
DELETE FROM `acl_resource` WHERE `id` = "3333c1aa2fd45297";
DELETE FROM `acl_resource` WHERE `id` = "19d46d2f7501a9b2";
DELETE FROM `acl_resource` WHERE `id` = "bec1bebaacc0e936";
DELETE FROM `acl_resource` WHERE `id` = "34d48343ea38eba3";
DELETE FROM `acl_resource` WHERE `id` = "dbf0617630089e7a";
DELETE FROM `acl_resource` WHERE `id` = "c6da1f7bb1295a51";
DELETE FROM `acl_resource` WHERE `id` = "d36114c721917300";
DELETE FROM `acl_resource` WHERE `id` = "13110b47edd72e37";
DELETE FROM `acl_resource` WHERE `id` = "6029394193209d74";
DELETE FROM `acl_resource` WHERE `id` = "74559779b75eeb96";
DELETE FROM `acl_resource` WHERE `id` = "fd6afe817ce547e6";
DELETE FROM `acl_resource` WHERE `id` = "ba1fd6ced759206e";
DELETE FROM `acl_resource` WHERE `id` = "52a6c63ca841ad5f";
DELETE FROM `acl_resource` WHERE `id` = "c9cdbd81906a0b73";
DELETE FROM `acl_resource` WHERE `id` = "e89cd67289eddaea";
DELETE FROM `acl_resource` WHERE `id` = "0f41068501c9af03";
DELETE FROM `acl_resource` WHERE `id` = "f48ef8eec7ae8c01";
DELETE FROM `acl_resource` WHERE `id` = "3a1252027a07f6c8";
DELETE FROM `acl_resource` WHERE `id` = "3ec6f3c826e1fb2b";
DELETE FROM `acl_resource` WHERE `id` = "8f4cb2d1d5b57025";
DELETE FROM `acl_resource` WHERE `id` = "cd7384eeb3d5804a";
DELETE FROM `acl_resource` WHERE `id` = "757dbe7a3f28a531";
DELETE FROM `acl_resource` WHERE `id` = "80ec82b56b24f003";
DELETE FROM `acl_resource` WHERE `id` = "2184738413722241";
DELETE FROM `acl_resource` WHERE `id` = "3285ac4d4adf482c";
DELETE FROM `acl_resource` WHERE `id` = "6ca3232c13d02ae7";
DELETE FROM `acl_resource` WHERE `id` = "113050129c759805";
DELETE FROM `acl_resource` WHERE `id` = "e711a6ac576cc326";
DELETE FROM `acl_resource` WHERE `id` = "f521c0a35aae5dfb";
DELETE FROM `acl_resource` WHERE `id` = "0a5b4a25a61980cb";
DELETE FROM `acl_resource` WHERE `id` = "9f637e34a1c99420";
DELETE FROM `acl_resource` WHERE `id` = "25db0e044f487e8e";
DELETE FROM `acl_resource` WHERE `id` = "cf22f78a0079ed93";
DELETE FROM `acl_resource` WHERE `id` = "18f633bff893dea7";
DELETE FROM `acl_resource` WHERE `id` = "0e494e5d37db0e0d";
DELETE FROM `acl_resource` WHERE `id` = "fbaf66fd9463735b";
DELETE FROM `acl_resource` WHERE `id` = "88643c94b2d7ffc1";
DELETE FROM `acl_resource` WHERE `id` = "e184024d5e546adb";
DELETE FROM `acl_resource` WHERE `id` = "0e39bfa412e78db0";
DELETE FROM `acl_resource` WHERE `id` = "8009c6c9688e430c";
DELETE FROM `acl_resource` WHERE `id` = "52c9e0e1ac1f01d6";
DELETE FROM `acl_resource` WHERE `id` = "2181e7028f08a3cb";
DELETE FROM `acl_resource` WHERE `id` = "fcd92eebe06fc77f";
DELETE FROM `acl_resource` WHERE `id` = "15553ff67e31ca4c";
DELETE FROM `acl_resource` WHERE `id` = "67aa3b4fdeb14ba2";
DELETE FROM `acl_role` WHERE `id` = "1f15a4c9d4798230";
DELETE FROM `acl_role` WHERE `id` = "3ae149a8dc444bce";
DELETE FROM `acl_role` WHERE `id` = "455511f8e362deb9";
