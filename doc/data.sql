BEGIN;
INSERT INTO groups (name) VALUES ('admin');
INSERT INTO principals (name, ssh_key, provisioned) VALUES ('admin', NULL, 't');

COPY acl_types (acl_type_id, name) FROM stdin;
1	read
2	write
3	manage
4	discover
5	group_manage
6	enrol
7	principal_manage
\.

INSERT INTO group_membership (group_id, principal_id) VALUES ((SELECT group_id FROM groups WHERE name = 'admin'), (SELECT principal_id FROM principals WHERE name = 'admin'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'principal_manage'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'group_manage'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'manage'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'write'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'read'));
INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'enrol'));

INSERT INTO acls (secret_id, group_id, acl_type_id) VALUES (NULL, (SELECT group_id FROM groups WHERE name = 'admin'), (SELECT acl_type_id FROM acl_types WHERE name = 'enrol'));

COMMIT;
