package services

import "math/rand"

// list of user agents
const (
	// Chrome
	ChromeMacOSXUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"
	ChromeWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"
	ChromeLinuxUserAgent = "Mozilla/5.0 (X11; Linux x86_64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"

	// Firefox
	FirefoxMacOSXUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:88.0) " +
		"Gecko/20100101 Firefox/88.0"
	FirefoxWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) " +
		"Gecko/20100101 Firefox/88.0"
	FirefoxLinuxUserAgent = "Mozilla/5.0 (X11; Linux i686; rv:88.0) " +
		"Gecko/20100101 Firefox/88.0"

	// Safari
	SafariMacOSXUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15"
	SafariWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15"
	SafariLinuxUserAgent = "Mozilla/5.0 (X11; Linux x86_64) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15"

	// Edge
	EdgeMacOSXUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.66"
	EdgeWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.66"
	EdgeLinuxUserAgent = "Mozilla/5.0 (X11; Linux x86_64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.66"

	// Opera
	OperaMacOSXUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15 OPR/76.0.4017.177"
	OperaWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15 OPR/76.0.4017.177"
	OperaLinuxUserAgent = "Mozilla/5.0 (X11; Linux x86_64) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15 OPR/76.0.4017.177"

	// Internet Explorer
	InternetExplorerWindowsUserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) " +
		"like Gecko"
	InternetExplorerLinuxUserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) " +
		"like Gecko"

	// Android
	AndroidUserAgent = "Mozilla/5.0 (Linux; Android 11; SM-G991B) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.210 Mobile Safari/537.36"

	// iOS
	IOSUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 14_5_1 like Mac OS X) " +

		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Mobile/15E148 Safari/604.1"

	// Windows Phone
	WindowsPhoneUserAgent = "Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; " +
		"Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/70.0.3538.110 Mobile Safari/537.36 Edge/15.14977"

	// Blackberry
	BlackberryUserAgent = "Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en) " +
		"AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.1.0.346 Mobile Safari/534.11+"

	// Linux
	LinuxUserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:88.0) " +

		"Gecko/20100101 Firefox/88.0"

	// Chrome OS
	ChromeOSUserAgent = "Mozilla/5.0 (X11; CrOS x86_64 13816.34.0) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.218 Safari/537.36"

	// Playstation
	PlaystationUserAgent = "Mozilla/5.0 (PlayStation 4 3.11) AppleWebKit/537.73 " +
		"(KHTML, like Gecko)"

	// Xbox
	XboxUserAgent = "Mozilla/5.0 (Xbox One 6.2.12998.0) AppleWebKit/537.73 " +
		"(KHTML, like Gecko) Version/10.0 Safari/537.73"

	// Nintendo
	NintendoUserAgent = "Mozilla/5.0 (Nintendo Switch; Firmware 9.2.0; " +
		"NX) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/9.2.0 Safari/605.1.15"
)

// list of user agents
var UserAgents = []string{
	// Chrome
	ChromeMacOSXUserAgent,
	ChromeWindowsUserAgent,
	ChromeLinuxUserAgent,

	// Firefox
	FirefoxMacOSXUserAgent,
	FirefoxWindowsUserAgent,
	FirefoxLinuxUserAgent,

	// Safari
	SafariMacOSXUserAgent,
	SafariWindowsUserAgent,
	SafariLinuxUserAgent,

	// Edge
	EdgeMacOSXUserAgent,
	EdgeWindowsUserAgent,
	EdgeLinuxUserAgent,

	// Opera
	OperaMacOSXUserAgent,
	OperaWindowsUserAgent,
	OperaLinuxUserAgent,

	// Internet Explorer
	InternetExplorerWindowsUserAgent,
	InternetExplorerLinuxUserAgent,

	// Android
	AndroidUserAgent,

	// iOS
	IOSUserAgent,

	// Windows Phone
	WindowsPhoneUserAgent,

	// Blackberry
	BlackberryUserAgent,

	// Linux
	LinuxUserAgent,

	// Chrome OS
	ChromeOSUserAgent,

	// Playstation
	PlaystationUserAgent,

	// Xbox
	XboxUserAgent,

	// Nintendo
	NintendoUserAgent,
}

func GetRandomUserAgent() string {
	return UserAgents[getRandomUserAgentIndex()]
}

func getRandomUserAgentIndex() int {
	return getRandomInt(0, len(UserAgents)-1)
}

func getRandomInt(min, max int) int {
	if max < min {
		return 0
	}
	if max == min {
		return max
	}
	return min + rand.Intn(max-min)
}
