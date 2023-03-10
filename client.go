package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// var s string
	var devices []dbus.ObjectPath
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")
	fmt.Println(conn.Names())
	err = obj.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&devices)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call Foo function (is the server example running?):", err)
		os.Exit(1)
	}
	for _, v := range devices {
		println(v)
	}
	fmt.Println("Result from calling Foo function on interface:")
	// fmt.Println(s)

}

func workginExampleUpower() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// var s string
	var devices []dbus.ObjectPath
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")
	fmt.Println(conn.Names())
	err = obj.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&devices)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call Foo function (is the server example running?):", err)
		os.Exit(1)
	}
	for _, v := range devices {
		println(v)
	}
	fmt.Println("Result from calling Foo function on interface:")
	// fmt.Println(s)
}
func workingExample() {

	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.Names())
	// func (conn *Conn) Object(dest string, path ObjectPath) *Object
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	fmt.Println(conn.Names())
	// Interface from the specification:
	// UINT32 org.freedesktop.Notifications.Notify (STRING app_name, UINT32 replaces_id, STRING app_icon, STRING summary, STRING body, ARRAY actions, DICT hints, INT32 expire_timeout);

	// func (o *Object) Call(method string, flags Flags, args ...interface{}) *Call
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "c¼h", uint32(0), "", "Hallo Chaostreff!", "Ich begrüße euch herzlich zu meiner c¼h!", []string{}, map[string]dbus.Variant{}, int32(1000))
	if call.Err != nil {
		panic(call.Err)
	}
}
