{{/*
	A simple template for creating different ban messages for permanent and timed bans.
  If you have an appeal link for your server, consider only sending the appeal link for permanent bans.
  Ban duration is measured in nanoseconds, and a ban that is 0 nanoseconds long is permanent.

	Trigger: None
	Author: Ishkander <https://github.com/ishkander56>
*/}}
You have been {{.ModAction}}
**Reason:** {{.Reason}}
{{if (eq .Duration 0)}}
This ban is permanent. Appeals are available at (somewhere.)
{{else}}
This ban is temporary. Appeals will not be available.
{{end}}
