//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package dashboard

import "testing"

func TestToValidFileName(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{"NGINX - Overview", "NGINX_-_Overview"},
		{"AWS EC2 (reception)", "AWS_EC2_(reception)"},
		{"Re:Invent - Coffee House", "Re:Invent_-_Coffee_House"},
		{"MySQL Disk I/O", "MySQL_Disk_IO"},
	}

	for _, tt := range tests {
		if got := toValidFileName(tt.text); got != tt.want {
			t.Fatalf("\n got=%s\nwant=%s", got, tt.want)
		}
	}
}
