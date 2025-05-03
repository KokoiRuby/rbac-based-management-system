-- admin/admin
INSERT INTO rbac_mgmt_system.users (username, password, is_admin, created_at, updated_at)
VALUES
    (
        'admin',
        '$2a$10$rj0NtCIyPIbhAWwrB4YVWeoaixI/rfijezWmaSlBVl/ih85tnqiLe',
        true,
        NOW(),
        NOW()
    );