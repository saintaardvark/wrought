package qso

import (
	"fmt"
	"wrought/ham"
	"wrought/morsePlayer"
)

const (
	tnxBob = "TNX FOR CALL BT UR RST 599 599 HR"
	sldCpy = "TNX FOR RPT SLD CPY FB UR RST 599 599 BT"
	hwCpy  = "HW CPY?"
)

// A QSO is a set of exchanges between two hams
type QSO struct {
	Tx, Rx        *ham.Ham
	Transmissions []*Exchange
}

// NewQSO returns a new QSO struct
func NewQSO() *QSO {
	qso := QSO{}
	return &qso
}

// An Exchange is a sentence sent from one ham to another
type Exchange struct {
	Sender, Receiver *ham.Ham
	Sentence         string
}

// NewExchange returns a new Exchange struct
func NewExchange() *Exchange {
	exchange := Exchange{}
	return &exchange
}

// AppendExchange adds a new exchange to a QSO struct
func (qso *QSO) AppendExchange(sender, receiver *ham.Ham, s string) {
	exchange := Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: s,
	}
	qso.Transmissions = append(qso.Transmissions, &exchange)
}
}

func initialGreeting(caller, receiver *ham.Ham) string {
	callerRepeat := fmt.Sprintf("%s %s %s", caller.Callsign, caller.Callsign, caller.Callsign)
	receiverRepeat := fmt.Sprintf("%s %s %s", receiver.Callsign, receiver.Callsign, receiver.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", callerRepeat, caller.Callsign, receiverRepeat)
	return msg
}

func firstExchange(caller, receiver *ham.Ham) string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		de(receiver.Callsign, caller.Callsign),
		tnxBob,
		qth(caller.Location),
		name(caller.Name),
		hwCpy,
		kn(receiver.Callsign, caller.Callsign))
}

func secondExchange(caller, receiver *ham.Ham) string {
	return fmt.Sprintf("%s %s %s %s %s",
		de(caller.Callsign, receiver.Callsign),
		sldCpy,
		name(receiver.Name),
		qth(receiver.Location),
		kn(caller.Callsign, receiver.Callsign))
}

func gnightBob(caller, receiver *ham.Ham) string {
	msg := gnightBob1(caller, receiver)
	msg += "\n" + gnightBob2(receiver, caller)
	return msg
}

func gnightBob1(caller, receiver *ham.Ham) string {
	return fmt.Sprintf("%s %s %s %s %s",
		de(receiver.Callsign, caller.Callsign),
		"TNX FER FB QSO",
		receiver.Name,
		"HP CU AGN BT VY 73 TO U ES URS SK",
		de(receiver.Callsign, caller.Callsign))
}

func gnightBob2(caller, receiver *ham.Ham) string {
	return fmt.Sprintf("%s %s %s %s",
		de(receiver.Callsign, caller.Callsign),
		"TNX FER QSO "+receiver.Name,
		"BCNU BT VY 73 TO U ES URS SK",
		de(receiver.Callsign, caller.Callsign))
}

// BuildQSO creates an exchange between two hams
func BuildQSO(caller, receiver *ham.Ham, player *morsePlayer.MorsePlayer) {
	player.Exchange = append(player.Exchange, initialGreeting(caller, receiver))
	player.Exchange = append(player.Exchange, firstExchange(caller, receiver))
	player.Exchange = append(player.Exchange, secondExchange(caller, receiver))
	player.Exchange = append(player.Exchange, gnightBob(caller, receiver))

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
