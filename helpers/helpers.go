package helpers

import (
	"errors"
	"strings"

	"github.com/godbus/dbus/v5"
)

// GetDbusPath
func GetDbusPath(dest string) (dbus.ObjectPath, error) {
	tmp := strings.Replace(dest, ".", "/", -1)
	retVal := "/" + tmp
	if dbus.ObjectPath(retVal).IsValid() {
		return dbus.ObjectPath(retVal), nil
	}
	return dbus.ObjectPath(""), errors.New(retVal + " is not a valid dbus path")
}
