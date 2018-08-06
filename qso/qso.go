package qso

import (
	"bufio"
	"fmt"
	"os"
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
	// Tx is the Ham calling CQ
	// RX is the Ham responding
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

// PlayCW plays a QSO's exchange as Morse code
func (qso *QSO) PlayCW(player *morsePlayer.MorsePlayer) {
	var sentences []*string
	for _, exch := range qso.Transmissions {
		sentences = append(sentences, &exch.Sentence)
	}
	player.PlayCW(sentences)
}

// PlayRemoteHalf plays the remote half of the Exchange
func (qso *QSO) PlayRemoteHalf(player *morsePlayer.MorsePlayer) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Send your CQ and hit [enter] when ready to continue...")
	reader.ReadString('\n')
	for _, exch := range qso.Transmissions {
		if exch.Sender.Callsign == qso.Rx.Callsign {
			player.PlayCW([]*string{&exch.Sentence})

			fmt.Print("Your turn! Hit [enter] when ready to continue...")
			reader.ReadString('\n')
		}
	}
}

// PrintText prints the plain text of the Exchange
func (qso *QSO) PrintText() {
	for _, s := range qso.Transmissions {
		fmt.Println(s.Sentence)
	}
}

// BuildQSO creates a QSO between two hams
// Tx is the initiator (the one calling CQ); Rx is the one who replies.
func BuildQSO(Tx, Rx *ham.Ham, player *morsePlayer.MorsePlayer) *QSO {
	qso := NewQSO()
	qso.Tx = Tx
	qso.Rx = Rx
	qso.AppendExchange(initialCQ(Tx))
	qso.AppendExchange(initialReply(Rx, Tx))
	qso.AppendExchange(firstExchange(Tx, Rx))
	qso.AppendExchange(secondExchange(Rx, Tx))
	qso.AppendExchange(gnightBob1(Tx, Rx))
	qso.AppendExchange(gnightBob2(Rx, Tx))
	return qso
}

// initialCQ returns an Exchange containing a CQ call
func initialCQ(sender *ham.Ham) *Exchange {
	senderRepeat := fmt.Sprintf("%s %s %s", sender.Callsign, sender.Callsign, sender.Callsign)
	msg := fmt.Sprintf("CQ CQ CQ DE %s K", senderRepeat)
	return &Exchange{
		Sender:   sender,
		Sentence: msg,
	}
}

// initialReply returns an Exchange replying to a CQ call
func initialReply(sender, receiver *ham.Ham) *Exchange {
	replyRepeat := fmt.Sprintf("%s %s %s", sender.Callsign, sender.Callsign, sender.Callsign)
	msg := fmt.Sprintf("%s DE %s KN", receiver.Callsign, replyRepeat)
	return &Exchange{
		Sender:   sender,
		Sentence: msg,
	}
}

func firstExchange(sender, receiver *ham.Ham) *Exchange {
	msg := fmt.Sprintf("%s %s %s %s %s %s",
		de(receiver.Callsign, sender.Callsign),
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
		de(receiver.Callsign, sender.Callsign),
		sldCpy,
		name(sender.Name),
		qth(sender.Location),
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
