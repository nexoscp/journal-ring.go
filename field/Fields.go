package field

/*
	Fields: https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html
	MUST BE UPPERCASE!
*/
type Field string

const (
	Priority Field = "PRIORITY"
	Message  Field = "MESSAGE"
	/*
	   A 128-bit message identifier ID for recognizing certain message types, if this is desirable. This should contain
	   a 128-bit ID formatted as a lower-case hexadecimal string, without any separating dashes or suchlike. This is
	   recommended to be a UUID-compatible ID, but this is not enforced, and formatted differently. Developers can
	   generate a new ID for this purpose with systemd-id128 new.

	   Additional information could be linked with a catalog file.
	   See https://www.freedesktop.org/wiki/Software/systemd/catalog/
	*/
	MessageID Field = "MESSAGE_ID"
	/*
		also known as "tag"
	*/
	SyslogIdentifier Field = "SYSLOG_IDENTIFIER"
	/*
		A documentation URL with further information about the topic of the log message. Tools such as journalctl will
		include a hyperlink to an URL specified this way in their output. Should be a "http://", "https://", "file:/",
		"man:" or "info:" URL.
	*/
	Documentation Field = "DOCUMENTATION"
)
