package social

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/style"
)

type Platform int

const (
    Instagram Platform = iota + 1
    TikTok
    GitHub
    Snapchat
    Twitch
    YouTube
    X
    LinkedIn
    Reddit
    Pinterest
    Medium
    Vimeo
    DeviantArt
    SoundCloud
    Flickr
    Dribbble
    Behance
)

var Platforms = []Platform{
    Instagram, TikTok, GitHub, Snapchat, Twitch, YouTube,
    X, LinkedIn, Reddit, Pinterest, Medium, Vimeo,
    DeviantArt, SoundCloud, Flickr, Dribbble, Behance,
}

var PlatformStrings = map[Platform]string{
    Instagram:  "Instagram",
    TikTok:     "TikTok",
    GitHub:     "GitHub",
    Snapchat:   "Snapchat",
    Twitch:     "Twitch",
    YouTube:    "YouTube",
    X:    			"X",
    LinkedIn:   "LinkedIn",
    Reddit:     "Reddit",
    Pinterest:  "Pinterest",
    Medium:     "Medium",
    Vimeo:      "Vimeo",
    DeviantArt: "DeviantArt",
    SoundCloud: "SoundCloud",
    Flickr:     "Flickr",
    Dribbble:   "Dribbble",
    Behance:    "Behance",
}

var PlatformBaseUrls = map[Platform]string{
    Instagram:  "https://instagram.com/",
    TikTok:     "https://us.tiktok.com/@",
    GitHub:     "https://github.com/",
    Snapchat:   "https://www.snapchat.com/add/",
    Twitch:     "https://www.twitch.tv/",
    YouTube:    "https://youtube.com/@",
    X:    			"https://x.com/",
    LinkedIn:   "https://www.linkedin.com/in/",
    Reddit:     "https://www.reddit.com/user/",
    Pinterest:  "https://www.pinterest.com/",
    Medium:     "https://medium.com/@",
    Vimeo:      "https://vimeo.com/",
    DeviantArt: "https://www.deviantart.com/",
    SoundCloud: "https://soundcloud.com/",
    Flickr:     "https://www.flickr.com/people/",
    Dribbble:   "https://dribbble.com/",
    Behance:    "https://www.behance.net/",
}

func (p Platform) String() string {
	return PlatformStrings[p]
}

func (p Platform) BaseUrl() string {
	return PlatformBaseUrls[p]
}

func (p Platform) Spacer() int {
	return style.MaxCharWidth - len(p.String())
}

type Status int

const (
	Unavailable Status = iota - 1
	Unknown
	Available
)

var StatusMessages = map[Status]string{
	Unavailable: style.Red.Colorize("✖️"),
	Unknown:     style.Red.Colorize("?"),
	Available:   style.Green.Colorize("✔️"),
}

func (s Status) String() string {
	return StatusMessages[s]
}

func (p Platform) GetAvailability(username string) Status {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := p.BaseUrl() + username
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Unknown
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Unknown
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNotFound:
		return Available
	case http.StatusNotAcceptable:
		if p == GitHub {
			return Available
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Unknown
	}

	switch p {
	case Instagram:
		if strings.Contains(string(body), "Sorry, this page isn't available.") {
			return Available
		}
	case TikTok:
		if strings.Contains(string(body), "Couldn't find this account") {
			return Available
		}
	case Snapchat:
		if strings.Contains(string(body), "This content was not found") {
			return Available
		}
	case Twitch:
		if !strings.Contains(string(body), username) {
			return Available
		}
	case YouTube:
		if strings.Contains(string(body), "<title>404 Not Found</title>") {
			return Available
		}
	case X:
		if !strings.Contains(string(body), username) {
			return Available
		}
	case LinkedIn:
		if !strings.Contains(string(body), username) {
			return Available
		}
	case Reddit:
		if strings.Contains(string(body), "Sorry, nobody on Reddit goes by that name.") {
			return Available
		}
	case Pinterest:
		if strings.Contains(string(body), "<title></title>") {
			return Available
		}
	case Medium:
		if strings.Contains(string(body), "PAGE NOT FOUND") {
			return Available
		}
	case Vimeo:
		if strings.Contains(string(body), "Sorry, we couldn’t find that page") {
			return Available
		}
	case DeviantArt:
		if strings.Contains(string(body), "<title>DeviantArt: 404</title>") {
			return Available
		}
	case SoundCloud:
		if strings.Contains(string(body), "We can’t find that user.") {
			return Available
		}
	case Flickr:
		if strings.Contains(string(body), "This is not the page you’re looking for.") {
			return Available
		}
	case Dribbble:
		if strings.Contains(string(body), "<title>Uh-oh!</title>") {
			return Available
		}
	case Behance:
		if strings.Contains(string(body), "Oops! We can’t find that page.") {
			return Available
		}
	}

	return Unavailable
}
