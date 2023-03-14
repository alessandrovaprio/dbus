package connFactory

import (
	"C"

	"github.com/godbus/dbus/v5"
)
import (
	"errors"
	"strings"
)

var allowed_buses = []string{"System", "Session"}
var allowed_destinations = []string{"org.freedesktop.UPower"}

//GetConnection
func GetBusConnection(busType string) (*dbus.Conn, error) {
	if contains(allowed_buses, busType) {
		switch busType {
		case "Session":
			return dbus.SessionBus()
		case "System":
			return dbus.SystemBus()
		default:
			return nil, errors.New("Specified Bus Not allowed IN " + strings.Join(allowed_buses, ","))
		}
	} else {
		return nil, errors.New("Specified Bus Not allowed IN " + strings.Join(allowed_buses, ","))
	}
}

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
