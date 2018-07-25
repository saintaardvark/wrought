package main

import (
	"fmt"

	"wrought/ham"
)

func initialGreeting(caller, receiver *ham.Ham) string {
	callerRepeat := fmt.Sprintf("%s %s %s", caller.Callsign, caller.Callsign, caller.Callsign)
	receiverRepeat := fmt.Sprintf("%s %s %s", receiver.Callsign, receiver.Callsign, receiver.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", callerRepeat, caller.Callsign, receiverRepeat)
	return msg
}

func firstExchange(caller, receiver *ham.Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FOR CALL BT UR RST 599 599 HR "
	msg = msg + qth(caller.Location) + " "
	msg = msg + name(caller.Name) + " "
	msg = msg + "HW CPY? "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func secondExchange(caller, receiver *ham.Ham) string {
	msg := de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FOR RPT SLD CPY FB UR RST 599 599 BT" + " "
	msg = msg + name(receiver.Name) + " "
	msg = msg + qth(receiver.Location) + " "
	msg = msg + kn(caller.Callsign, receiver.Callsign)
	return msg
}

func gnightBob(caller, receiver *ham.Ham) string {
	msg := de(receiver.Callsign, caller.Callsign) + " "
	msg = msg + "TNX FER FB QSO " + receiver.Name + " "
	msg = msg + "HP CU AGN BT VY 73 TO U ES URS SK" + " "
	msg = msg + de(receiver.Callsign, caller.Callsign) + "\n"
	// And now the reply
	msg = msg + de(caller.Callsign, receiver.Callsign) + " "
	msg = msg + "TNX FER QSO " + caller.Name + " "
	msg = msg + "BCNU BT VY 73 TO U ES URS SK" + " "
	msg = msg + de(caller.Callsign, receiver.Callsign)
	return msg
}

func name(name string) string {
	return fmt.Sprintf("NAME %s %s", name, name)
}

func qth(location string) string {
	return fmt.Sprintf("QTH %s %s", location, location)
}

func de(cx, rx string) string {
	return fmt.Sprintf("%s DE %s", cx, rx)
}

func kn(cx, rx string) string {
	return fmt.Sprintf("%s KN", de(cx, rx))
}
