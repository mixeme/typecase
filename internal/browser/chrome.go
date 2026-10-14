package browser

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var chromeExec string

func init() {
	chromeExec, _ = locateChrome()
}

// ChromePath returns the browser executable used by chromedp tests.
func ChromePath() string {
	return chromeExec
}

// ChromeAvailable reports whether a Chromium-based browser was found.
func ChromeAvailable() bool {
	return chromeExec != ""
}

func locateChrome() (string, error) {
	if p := os.Getenv("CHROME_PATH"); p != "" {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	if p, err := exec.LookPath("google-chrome"); err == nil {
		return p, nil
	}
	if p, err := exec.LookPath("chromium"); err == nil {
		return p, nil
	}
	if p, err := exec.LookPath("chromium-browser"); err == nil {
		return p, nil
	}
	if runtime.GOOS == "windows" {
		for _, p := range []string{
			`C:\Program Files\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files\Microsoft\Edge\Application\msedge.exe`,
			`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
		} {
			if _, err := os.Stat(p); err == nil {
				return p, nil
			}
		}
	}
	return "", exec.ErrNotFound
}

// instance is a headless browser started by the harness.
//
// The harness starts the process itself instead of letting chromedp's exec
// allocator do it: the allocator learns the DevTools address by reading the
// "DevTools listening on" line from stderr, and a browser built as a Windows
// GUI subsystem binary (Edge, notably) writes nothing to an inherited pipe.
// The DevToolsActivePort file inside the profile directory carries the same
// address on every platform and every Chromium build.
type instance struct {
	cmd      *exec.Cmd
	dir      string
	endpoint string
	stopped  bool
	mu       sync.Mutex
}

const portFileTimeout = 30 * time.Second

// profilePrefix names the harness profile directories. killProfileHolders
// matches command lines against it, so nothing else may use it.
const profilePrefix = "typecase-browser-"

func launchChrome() (*instance, error) {
	if chromeExec == "" {
		return nil, exec.ErrNotFound
	}
	dir, err := os.MkdirTemp("", profilePrefix)
	if err != nil {
		return nil, fmt.Errorf("profile dir: %w", err)
	}
	cmd := exec.Command(chromeExec,
		"--headless",
		"--disable-gpu",
		"--disable-lcd-text",
		"--force-color-profile=srgb",
		"--no-sandbox",
		"--no-first-run",
		"--no-default-browser-check",
		"--disable-extensions",
		"--disable-background-networking",
		"--disable-sync",
		// what chromedp gives a browser it starts itself: a tab in the
		// background is neither slowed nor frozen, and a busy page is not
		// taken for a hung one — each check opens its own tab
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-renderer-backgrounding",
		"--disable-hang-monitor",
		"--disable-ipc-flooding-protection",
		"--disable-popup-blocking",
		"--disable-prompt-on-repost",
		"--disable-breakpad",
		"--remote-allow-origins=*",
		"--user-data-dir="+dir,
		"--remote-debugging-port=0",
		"about:blank",
	)
	if err := cmd.Start(); err != nil {
		os.RemoveAll(dir)
		return nil, fmt.Errorf("start %s: %w", chromeExec, err)
	}
	in := &instance{cmd: cmd, dir: dir}
	port, err := waitPortFile(filepath.Join(dir, "DevToolsActivePort"))
	if err != nil {
		in.stop()
		return nil, err
	}
	in.endpoint = "http://127.0.0.1:" + strconv.Itoa(port) + "/"
	return in, nil
}

func waitPortFile(path string) (int, error) {
	deadline := time.Now().Add(portFileTimeout)
	for {
		port, err := readPortFile(path)
		if err == nil {
			return port, nil
		}
		if time.Now().After(deadline) {
			return 0, fmt.Errorf("browser did not publish %s within %s: %w", filepath.Base(path), portFileTimeout, err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func readPortFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	if !scan.Scan() {
		return 0, fmt.Errorf("empty port file")
	}
	port, err := strconv.Atoi(strings.TrimSpace(scan.Text()))
	if err != nil {
		return 0, fmt.Errorf("bad port file: %w", err)
	}
	if port <= 0 {
		return 0, fmt.Errorf("bad port %d", port)
	}
	return port, nil
}

func (in *instance) stop() {
	in.mu.Lock()
	defer in.mu.Unlock()
	if in.stopped {
		return
	}
	in.stopped = true
	if in.cmd != nil && in.cmd.Process != nil {
		pid := in.cmd.Process.Pid
		if runtime.GOOS == "windows" {
			_ = exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(pid)).Run()
		} else {
			_ = in.cmd.Process.Kill()
		}
		_, _ = in.cmd.Process.Wait()
	}
	killProfileHolders(in.dir)
	if in.dir != "" {
		_ = os.RemoveAll(in.dir)
	}
}

// killProfileHolders kills browser processes that still name dir as their
// --user-data-dir.
//
// stop kills the tree of the process the harness started. A browser that
// re-execs itself leaves that tree, so the tree kill misses it; the orphan
// keeps the profile locked and the next frame's screenshot hangs until its
// deadline. The profile directory name is unique per harness, so it is a safe
// thing to match command lines against — as long as only browsers are
// considered, or the query finds the shell that runs it.
func killProfileHolders(dir string) {
	if dir == "" || !strings.HasPrefix(filepath.Base(dir), profilePrefix) {
		return
	}
	self := os.Getpid()
	for _, pid := range profileHolders(dir) {
		if pid <= 0 || pid == self {
			continue
		}
		if runtime.GOOS == "windows" {
			_ = exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(pid)).Run()
			continue
		}
		if p, err := os.FindProcess(pid); err == nil {
			_ = p.Kill()
		}
	}
}

// browserNames are the process names a Chromium-based browser runs under: the
// one the harness launched, plus the family, because a re-exec may pick a
// sibling binary.
func browserNames() []string {
	names := []string{"chrome.exe", "msedge.exe", "chromium.exe", "chrome", "chromium", "chromium-browser", "msedge"}
	if chromeExec != "" {
		names = append([]string{filepath.Base(chromeExec)}, names...)
	}
	return names
}

// processListPS prints "pid<TAB>name<TAB>command line" for every process.
var processListPS = "Get-CimInstance Win32_Process | ForEach-Object { \"$($_.ProcessId)\t$($_.Name)\t$($_.CommandLine)\" }"

// profileHolders returns the pids of browser processes holding dir.
func profileHolders(dir string) []int {
	marker := filepath.Base(dir)
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", processListPS).Output()
	} else {
		out, err = exec.Command("ps", "-A", "-o", "pid=,comm=,args=").Output()
	}
	if err != nil && len(out) == 0 {
		return nil
	}
	return parseProcessList(string(out), marker, browserNames())
}

// parseProcessList picks the pids whose command line names marker and whose
// process name is one of names. A line is "pid<sep>name<sep>command line",
// separated by tabs on Windows and by spaces from ps elsewhere.
func parseProcessList(out, marker string, names []string) []int {
	isBrowser := make(map[string]bool, len(names))
	for _, n := range names {
		isBrowser[strings.ToLower(n)] = true
	}
	var pids []int
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimRight(line, "\r")
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == '\t' || r == ' ' })
		if len(fields) < 3 || !strings.Contains(line, marker) {
			continue
		}
		pid, convErr := strconv.Atoi(fields[0])
		if convErr != nil {
			continue
		}
		if !isBrowser[strings.ToLower(filepath.Base(fields[1]))] {
			continue
		}
		pids = append(pids, pid)
	}
	return pids
}
