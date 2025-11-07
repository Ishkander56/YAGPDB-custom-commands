{{/*
	When triggered by the "join message" timer, bans users if they lack the "first-time chatter" role. Replace 1431442192676950239 with your role ID.
	Users may leave the server to also avoid a ban, as voluntarily leaving the server is not scammer-like behavior.
	To remove this check and also ban leavers, remove line 12 and line 16.

	Trigger: None
	Author: Ishkander <https://github.com/ishkander56>
*/}}
{{$uid := toInt64 (index .ExecData "user")}}
{{$member := getMember $uid}}

{{if $member}}
	{{if not (targetHasRoleID $member 1431442192676950239)}}
		{{execAdmin (printf "ban %d You didn't post anything in the first three days, which seems like bot behavior." $uid )}}
	{{end}}
{{end}}
