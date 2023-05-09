SELECT site_id, recorded, response_time, err, msg FROM probes
	WHERE site_id=?
	AND recorded >= DATE_SUB(NOW(), INTERVAL ? SECOND)
	LIMIT ?;
