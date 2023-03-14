package main

/*
#cgo CFLAGS: -DPNG_DEBUG=1
#cgo amd64 386 CFLAGS: -DX86=1
#cgo LDFLAGS: -lstdc++ -lm
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
*/
import (
	"C"
	"fmt"
	"os"

	conn "github.com/alessandrovaprio/dbus/connFactory"
	helpers "github.com/alessandrovaprio/dbus/helpers"
	"github.com/godbus/dbus/v5"
)
import "strings"

func main() {
	// tmp := CallDBus("System", "org.freedesktop.UPower", "org.freedesktop.UPower.EnumerateDevices")
	// fmt.Println(tmp)

}

//export CallDBus
func CallDBus(busType *C.char, destination *C.char, operation *C.char) *C.char {
	// , Output **C.char
	goBusType := C.GoString(busType)
	goDestination := C.GoString(destination)
	goOperation := C.GoString(operation)
	conn, err := conn.GetBusConnection(goBusType)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to "+goBusType+":", err)
		os.Exit(1)
	}
	defer conn.Close()
	path, dbusPathErr := helpers.GetDbusPath(goDestination)
	if dbusPathErr != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect "+dbusPathErr.Error()+":", err)
		os.Exit(1)
	}
	obj := conn.Object(goDestination, path)
	var busReturn []dbus.ObjectPath
	err = obj.Call(goOperation, 0).Store(&busReturn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call "+goOperation+" function (is the server example running?):", err)
		os.Exit(1)
	}
	var values []string
	for _, v := range busReturn {
		tmp := fmt.Sprintf("%v", v)
		values = append(values, tmp)
	}
	retVal := strings.Join(values, "****")
	tmp := C.CString(fmt.Sprintf("%s", retVal))
	// fmt.Println(C.GoString(retVal))
	// *Output = C.CString(fmt.Sprintf("%s", retVal))
	// return int32(len(values))
	// defer C.free(unsafe.Pointer(retVal))
	return tmp
}

//export WorkingExampleUpower
func WorkingExampleUpower() {
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
