package bench

import (
	"fmt"

	"github.com/pions/webrtc"
	log "github.com/sirupsen/logrus"
)

// Useful for unit tests
func (s *Session) onNewDataChannelHelper(name string, channelID uint16, d *webrtc.DataChannel) {
	log.Tracef("New DataChannel %s (id: %x)\n", name, channelID)

	switch channelID {
	case s.downloadChannelID():
		log.Traceln("Created Download data channel")
		d.OnClose(s.onCloseHandlerDownload())
		go s.onOpenHandlerDownload(d)()

	case s.uploadChannelID():
		log.Traceln("Created Upload data channel")

	default:
		log.Warningln("Created unknown data channel")
	}
}

func (s *Session) onNewDataChannel() func(d *webrtc.DataChannel) {
	return func(d *webrtc.DataChannel) {
		if d == nil || d.ID() == nil {
			return
		}
		s.onNewDataChannelHelper(d.Label(), *d.ID(), d)
	}
}

func (s *Session) createMasterSession() error {
	if err := s.sess.CreateOffer(); err != nil {
		log.Errorln(err)
		return err
	}

	fmt.Println("Please, paste the remote SDP:")
	if err := s.sess.ReadSDP(); err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func (s *Session) createSlaveSession() error {
	fmt.Println("Please, paste the remote SDP:")
	if err := s.sess.ReadSDP(); err != nil {
		log.Errorln(err)
		return err
	}

	fmt.Println("SDP:")
	if err := s.sess.CreateAnswer(); err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}