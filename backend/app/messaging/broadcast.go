package messaging

import "breakfaster/mybot"

// BroadCastMenu is a method for broadcasting breakfast menu
func (m *MessageControllerImpl) BroadCastMenu() error {
	if err := m.pusher.BroadcastFlex("早餐選單", mybot.NewMenu()); err != nil {
		return err
	}
	return nil
}
