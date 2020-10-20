package router

// BroadCastMenu is a method for broadcasting breakfast menu
func (r *Router) BroadCastMenu() error {
	if err := r.Bot.BroadcastFlex("早餐選單", r.Bot.NewMenu()); err != nil {
		return err
	}
	return nil
}
