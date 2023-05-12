UPDATE sources SET
	active=?,
	changed=NOW()
WHERE site_id=? LIMIT 1;
