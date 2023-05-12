INSERT INTO sources
	(site_id, url, created, changed, active)
VALUES
	(?,       ?,   NOW(),   NOW(),   1);
