SELECT
	o.site_id, s.URL,
	d.recorded AS started, d.err AS err, d.msg AS msg,
	u.recorded AS ended
FROM incidents AS o
	LEFT JOIN sources AS s on s.site_id=o.site_id
	LEFT JOIN probes AS u ON o.up_probe_id=u.probe_id
	LEFT JOIN probes AS d ON o.down_probe_id=d.probe_id
WHERE o.site_id=?
LIMIT 1
