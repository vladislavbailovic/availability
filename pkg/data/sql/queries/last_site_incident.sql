SELECT
	o.site_id, IFNULL(o.down_probe_id, 0) AS down_probe_id, IFNULL(o.up_probe_id, 0) AS up_probe_id -- ,
	-- u.recorded AS up, d.recorded AS down, d.err AS err, d.msg AS msg
FROM incidents AS o
	LEFT JOIN probes AS u ON o.up_probe_id=u.probe_id
	LEFT JOIN probes AS d ON o.down_probe_id=d.probe_id
WHERE o.site_id=?
ORDER BY d.recorded DESC LIMIT 1
