{{/*
	This is part of the required setup for banning lurkers.
	This part of the command schedules Custom Command #3 after 72 hours (both of these values may be changed), passing a reference to the user who just joined.

	Trigger: Add this to your server's join message. Enable join messages if they're not available already.
	Author: Ishkander <https://github.com/ishkander56>
*/}}
{{execCC 3 nil 259200 (sdict "user" .User.ID)}}
