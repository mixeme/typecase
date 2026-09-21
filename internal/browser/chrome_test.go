package browser

import "testing"

// parseProcessList must find the browser that still holds a harness profile
// and nothing else — in particular not the shell that ran the query, whose
// own command line names the same profile directory.
func TestParseProcessList(t *testing.T) {
	const marker = "typecase-browser-123456"
	out := "" +
		"4236\tbash.exe\tgrep typecase-browser-123456\n" +
		"13420\tpowershell.exe\tGet-CimInstance Win32_Process | ... typecase-browser-123456\n" +
		"9672\tmsedge.exe\tC:\\msedge.exe --headless --user-data-dir=C:\\Temp\\typecase-browser-123456\n" +
		"9680\tchrome.exe\tchrome --user-data-dir=/tmp/typecase-browser-999999\n" +
		"\n" +
		"notapid\tmsedge.exe\t--user-data-dir=typecase-browser-123456\n"
	got := parseProcessList(out, marker, []string{"chrome.exe", "msedge.exe"})
	if len(got) != 1 || got[0] != 9672 {
		t.Fatalf("got %v, want [9672]", got)
	}
}

func TestParseProcessListPS(t *testing.T) {
	// ps -A -o pid=,comm=,args= pads the pid column and separates with spaces.
	const marker = "typecase-browser-777"
	out := "" +
		" 1201 chromium /usr/bin/chromium --headless --user-data-dir=/tmp/typecase-browser-777\n" +
		" 1202 bash     pgrep -f typecase-browser-777\n"
	got := parseProcessList(out, marker, browserNames())
	if len(got) != 1 || got[0] != 1201 {
		t.Fatalf("got %v, want [1201]", got)
	}
}
