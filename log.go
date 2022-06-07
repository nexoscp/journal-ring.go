package journal_ring

import (
	"bytes"
	"golang.org/x/sys/unix"
	"journal-ring/field"
	"journal-ring/priority"
	"strconv"
)

// https://lists.freedesktop.org/archives/systemd-devel/2012-November/007359.html
// https://github.com/systemd/systemd/blob/master/src/journal/journal-send.c function sd_journal_sendv line 213
type Journal struct {
	//used for SYSLOG_IDENTIFIER
	//see also https://linux.die.net/man/3/program_invocation_short_name
	tag      string
	socketFD int
}

/*
	generate messageID with "systemd-id128 new" on command line
	see https://www.freedesktop.org/wiki/Software/systemd/catalog/
*/
type MessageID string

func Open(tag string) (*Journal, error) {
	if socketFD, err := unix.Socket(unix.AF_UNIX, unix.SOCK_DGRAM, 0); err != nil {
		return nil, err
	} else if err := unix.Connect(socketFD, &unix.SockaddrUnix{Name: "/run/systemd/journal/socket"}); err != nil {
		return nil, err
	} else {
		return &Journal{tag: tag, socketFD: socketFD}, nil
	}
}

/*
else if ring, err := gouring.New(256, nil); err != nil {
		return err
	} else {
		defer ring.Close()
	}
*/

func (journal *Journal) l(priority priority.SysLogLevel, message string) {
	journal.log(CreateRecord(priority, message, ""))
}

func (journal *Journal) log(record *Record) {
	// TODO from systemd src/basic/util.c implement string_is_safe(..)

	// https://systemd.io/JOURNAL_NATIVE_PROTOCOL/
	var buffer bytes.Buffer
	write(&buffer, field.SyslogIdentifier, journal.tag)
	write(&buffer, field.Priority, strconv.Itoa(int(record.priority)))
	write(&buffer, field.Message, record.message)
	write(&buffer, field.MessageID, string(record.messageId))
	if record.fields != nil {
		for key, value := range record.fields {
			write(&buffer, key, value)
		}
	}
	_, err := unix.Write(journal.socketFD, buffer.Bytes())
	if err != nil {
		panic("Error logging " + err.Error()) //TODO better error handling? returning the error pollutes business code *argl*
	}
}

func write(buffer *bytes.Buffer, fieldName field.Field, value string) {
	if len(value) != 0 {
		buffer.WriteString(string(fieldName))
		buffer.WriteString("=")
		buffer.WriteString(value)
		buffer.WriteString("\n")
	}
}

type Record struct {
	priority  priority.SysLogLevel
	message   string
	messageId MessageID
	fields    map[field.Field]string
}

func CreateRecord(priority priority.SysLogLevel, message string, messageId MessageID) *Record {
	return &Record{priority: priority, message: message, messageId: messageId}
}

// TODO add method for map[field.Field]string and map[string]string and maybe some interface to prevent copy
func (record *Record) With(key field.Field, value string) *Record {
	if record.fields == nil {
		record.fields = make(map[field.Field]string, 4)
	}
	record.fields[key] = value
	return record
}
