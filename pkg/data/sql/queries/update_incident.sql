UPDATE incidents SET up_probe_id=?
	WHERE site_id=?
	AND down_probe_id=?
LIMIT 1
