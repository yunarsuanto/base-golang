INSERT INTO "users" ("username", "password") VALUES
('superadmin', '$2a$10$7dBcJCNxbl12LXBhiBtJsO/cgQ6IEn.qdej8kJneSoqggEbE/YKNK'),
('admin',      '$2a$10$7dBcJCNxbl12LXBhiBtJsO/cgQ6IEn.qdej8kJneSoqggEbE/YKNK')
ON CONFLICT (username) DO NOTHING;
--------------------------------------------------
INSERT INTO "roles" ("name") VALUES
('admin'),
('user')
ON CONFLICT (name) DO NOTHING;
--------------------------------------------------
INSERT INTO "permissions" ("name") VALUES
('user_list'),
('user_detail'),
('user_create'),
('user_update'),
('user_delete'),
('role_list'),
('role_detail'),
('role_create'),
('role_update'),
('role_delete'),
('permission_list'),
('permission_detail'),
('permission_create'),
('permission_update'),
('permission_delete'),
('role_permission_upsert'),
('user_role_upsert'),
('category_lesson_list'),
('category_lesson_detail'),
('category_lesson_create'),
('category_lesson_update'),
('category_lesson_delete'),
('lesson_list'),
('lesson_detail'),
('lesson_create'),
('lesson_update'),
('lesson_delete'),
('lesson_item_list'),
('lesson_item_detail'),
('lesson_item_create'),
('lesson_item_update'),
('lesson_item_delete')
ON CONFLICT (name) DO NOTHING;
--------------------------------------------------
INSERT INTO "role_permission" ("role_id", "permission_id")
SELECT r.id, p.id
FROM roles r, permissions p
WHERE (r.name = 'admin' AND p.name IN (SELECT name FROM permissions))
OR (r.name = 'user' AND p.name IN (
    'category_lesson_list',
    'category_lesson_detail',
    'lesson_list',
    'lesson_detail',
    'lesson_item_list',
    'lesson_item_detail'
)) ON CONFLICT ("role_id", "permission_id") DO NOTHING;
