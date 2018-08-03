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
func (qso *QSO) AppendExchange(exchange *Exchange) {
	qso.Transmissions = append(qso.Transmissions, exchange)
}

	}
	qso.Transmissions = append(qso.Transmissions, &exchange)

// PlayRemoteHalf plays the remote half of the Exchange
func (qso *QSO) PlayRemoteHalf() {
	return
}

// PrintText prints the plain text of the Exchange
func (qso *QSO) PrintText() {
	for _, s := range qso.Transmissions {
		fmt.Println(s.Sentence)
	}
}

// BuildQSO creates a QSO between two hams
func BuildQSO(caller, receiver *ham.Ham, player *morsePlayer.MorsePlayer) *QSO {
	qso := NewQSO()
	qso.AppendExchange(initialGreeting(caller, receiver))
	qso.AppendExchange(firstExchange(receiver, caller))
	qso.AppendExchange(secondExchange(caller, receiver))
	qso.AppendExchange(gnightBob1(receiver, caller))
	qso.AppendExchange(gnightBob2(caller, receiver))
	return qso
}

func initialGreeting(sender, receiver *ham.Ham) *Exchange {
	senderRepeat := fmt.Sprintf("%s %s %s", sender.Callsign, sender.Callsign, sender.Callsign)
	receiverRepeat := fmt.Sprintf("%s %s %s", receiver.Callsign, receiver.Callsign, receiver.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K\n%s DE %s KN\n", senderRepeat, sender.Callsign, receiverRepeat)
	return &Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: msg,
	}
}

func firstExchange(sender, receiver *ham.Ham) *Exchange {
	msg := fmt.Sprintf("%s %s %s %s %s %s",
		de(sender.Callsign, receiver.Callsign),
		tnxBob,
		qth(sender.Location),
		name(sender.Name),
		hwCpy,
		kn(sender.Callsign, receiver.Callsign))
	return &Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: msg,
	}
}

func secondExchange(sender, receiver *ham.Ham) *Exchange {
	msg := fmt.Sprintf("%s %s %s %s %s",
		de(sender.Callsign, receiver.Callsign),
		sldCpy,
		name(receiver.Name),
		qth(receiver.Location),
		kn(sender.Callsign, receiver.Callsign))
	return &Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: msg,
	}
}

// func gnightBob(caller, receiver *ham.Ham) *Exchange {
// 	msg := gnightBob1(caller, receiver)
// 	msg += "\n" + gnightBob2(receiver, caller)
// 	return msg
// }

func gnightBob1(sender, receiver *ham.Ham) *Exchange {
	msg := fmt.Sprintf("%s %s %s %s %s",
		de(receiver.Callsign, sender.Callsign),
		"TNX FER FB QSO",
		receiver.Name,
		"HP CU AGN BT VY 73 TO U ES URS SK",
		de(receiver.Callsign, sender.Callsign))
	return &Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: msg,
	}
}

func gnightBob2(sender, receiver *ham.Ham) *Exchange {
	msg := fmt.Sprintf("%s %s %s %s",
		de(receiver.Callsign, sender.Callsign),
		"TNX FER QSO "+receiver.Name,
		"BCNU BT VY 73 TO U ES URS SK",
		de(receiver.Callsign, sender.Callsign))
	return &Exchange{
		Sender:   sender,
		Receiver: receiver,
		Sentence: msg,
	}
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
