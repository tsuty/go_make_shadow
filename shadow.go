package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/md5_crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
	_ "github.com/GehirnInc/crypt/sha512_crypt"

	"github.com/rickb777/date"
)

const (
	saltSize = 16

	MD5ID    = 1
	SHA256ID = 5
	SHA512ID = 6
)

var saltBytes = []byte{
	// A-Z
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a,
	// a-z
	0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a,
	// 0-9
	0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	// ./
	0x2e, 0x2f,
}
var saltBytesLen = len(saltBytes)

type shadow struct {
	LoginName        string
	Salt             []byte
	MinPasswordAge   *uint   `long:"min" description:"The minimum password age" value-name:"days"`
	MaxPasswordAge   *uint   `long:"max" description:"The maximum password age" value-name:"days"`
	WarningPeriod    *uint   `long:"warning" description:"The number of days before a password is going to expire" value-name:"days"`
	InactivityPeriod *uint   `long:"inactivity" description:"The number of days after a password has expired" value-name:"days"`
	ExpirationDate   *uint   `long:"expiration" description:"The date of expiration of the account, expressed as the number of days since Jan 1, 1970" value-name:"days"`
	ReservedField    *string `long:"reserved" description:"This field is reserved for future use" hidden:"1"`
	SaltString       string  `long:"salt" description:"This encryption salt" hidden:"1"`
	MD5              bool    `long:"md5" description:"MD5"`
	SHA256           bool    `long:"sha256" description:"SHA-256"`
	SHA512           bool    `long:"sha512" description:"SHA-512 (default)"`
	OnlyEncrypt      bool    `long:"only-encrypt" description:"Only encrypt password"`
	Help             bool    `long:"help" short:"h" description:"Show this help"`
}

func (s *shadow) id() uint {
	if s.SHA512 {
		return SHA512ID
	} else if s.SHA256 {
		return SHA256ID
	} else {
		return MD5ID
	}
}

func (s *shadow) crypter() crypt.Crypter {
	if s.SHA512 {
		return crypt.SHA512.New()
	} else if s.SHA256 {
		return crypt.SHA256.New()
	} else if s.MD5 {
		return crypt.MD5.New()
	} else {
		log.Fatal("known crypter.")
		return nil
	}
}

func (s *shadow) encryptedPassword(password []byte) string {
	if len(s.SaltString) > 0 {
		s.Salt = []byte(s.SaltString)
	}
	if len(s.Salt) == 0 {
		s.Salt = s.makeSalt()
	}

	c := s.crypter()
	salt := []byte(fmt.Sprintf("$%d$", s.id()))
	salt = append(salt, s.Salt...)
	ret, err := c.Generate(password, salt)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func (s *shadow) makeSalt() []byte {
	var salt []byte
	for i := 0; i < saltSize; i++ {
		rand.Seed(time.Now().UnixNano())
		p := rand.Intn(saltBytesLen)
		salt = append(salt, saltBytes[p])
	}
	return salt
}

func (s *shadow) lastPasswordChangeDays() string {
	d := date.Today()
	return fmt.Sprintf("%d", d.DaysSinceEpoch())
}

func (s *shadow) minPasswordAge() string {
	if s.MinPasswordAge == nil {
		return ""
	}
	return fmt.Sprintf("%d", s.MinPasswordAge)
}

func (s *shadow) maxPasswordAge() string {
	if s.MaxPasswordAge == nil {
		return ""
	}
	return fmt.Sprintf("%d", s.MaxPasswordAge)
}

func (s *shadow) warningPeriod() string {
	if s.WarningPeriod == nil {
		return ""
	}
	return fmt.Sprintf("%d", s.WarningPeriod)
}

func (s *shadow) inactivityPeriod() string {
	if s.InactivityPeriod == nil {
		return ""
	}
	return fmt.Sprintf("%d", s.InactivityPeriod)
}

func (s *shadow) expirationDate() string {
	if s.ExpirationDate == nil {
		return ""
	}
	return fmt.Sprintf("%d", s.ExpirationDate)
}

func (s *shadow) reservedField() string {
	if s.ReservedField == nil {
		return ""
	}
	return fmt.Sprintf("%v", s.ReservedField)
}

func (s *shadow) make(password []byte) string {
	fields := make([]string, 9)
	fields[0] = s.LoginName
	fields[1] = s.encryptedPassword(password)
	fields[2] = s.lastPasswordChangeDays()
	fields[3] = s.minPasswordAge()
	fields[4] = s.maxPasswordAge()
	fields[5] = s.warningPeriod()
	fields[6] = s.inactivityPeriod()
	fields[7] = s.expirationDate()
	fields[8] = s.reservedField()

	return strings.Join(fields, ":")
}
