{{/*
	A lightweight implementation of XKCD's ROBOT9000. (https://blog.xkcd.com/2008/01/14/robot9000-and-xkcd-signal-attacking-noise-in-chat/)
	Deletes triggers if a message with a similar prefix has been posted within the last 100 messages. Customizable.

	Trigger: Regex | ".*"
	Channels: Anywhere people post repetitive messages.
	Roles: Anybody who you don't want posting repetitive messages.
	Author: Ishkander <https://github.com/ishkander56>
*/}}

{{/* Configurable prefix size and database size. Smaller prefixes are more strict (if the prefix is 5, "hello everyone" and "hello world" are treated as similar!), smaller databases forget messages sooner. */}}
{{$r9kPrefixSize := 100}}
{{$r9kDatabaseSize := 100}}

{{/* Actual processing begins here. */}}
{{$triggerMention := .User.Mention}}
{{$prefix := .Message.Content}}
{{if gt (len $prefix) $r9kPrefixSize}}
	{{$prefix = (slice $prefix 0 $r9kPrefixSize)}}
{{end}}
{{$key := (printf "robot8999-%d" .Channel.ID)}}
{{$raw := dbGet 0 $key}}
{{$list := cslice}}

{{if $raw}}
	{{$list = $raw.Value}}
{{end}}

{{if true}}
	{{- range $i, $v := $list}}
		{{if eq $v $prefix}}
			{{deleteTrigger}}
			{{$myReply := sendMessageRetID nil (printf "%s attempted to post a message that was already posted recently." $triggerMention)}}
			{{deleteMessage nil $myReply 10}}
			{{return}}
		{{end}}
	{{end}}
{{end}}

{{$list = $list.Append $prefix}}

{{if gt (len $list) $r9kDatabaseSize}}
	{{$list = (slice $list (sub (len $list) $r9kDatabaseSize))}}
{{end}}

{{dbSet 0 $key ($list)}}
