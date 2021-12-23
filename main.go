package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
	"github.com/shkh/lastfm-go/lastfm"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func end(err error) {
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Press any key to close window...")
	fmt.Scanln()
}

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		end(err)
		return
	}

	trim := strings.TrimSpace
	token := trim(cfg.Section("discord").Key("token").String())
	apiKey := trim(cfg.Section("lastfm").Key("api_key").String())
	username := trim(cfg.Section("lastfm").Key("username").String())
	title := trim(cfg.Section("app").Key("title").String())
	endlessMode, err := strconv.ParseBool(cfg.Section("app").Key("endless_mode").String())

	if err != nil {
		end(err)
		return
	}
	configInterval, err := cfg.Section("lastfm").Key("check_interval").Int()

	if err != nil {
		end(err)
		return
	}

	api := lastfm.New(apiKey, "")

	log.Println("Settings loaded: config.ini")

	if endlessMode {
		log.Println("Endless mode! Press `Ctrl+C` to exit")
	}
	dg, err := discordgo.New(token)
	if err != nil {
		end(err)
		return
	}

	log.Println("Authorized to Discord")

	if err := dg.Open(); err != nil {
		end(err)
		return
	}

	defer dg.Close()
	log.Println("Connected to Discord")

	interval := time.Duration(configInterval*1000) * time.Millisecond
	ticker := time.NewTicker(interval)

	var deathChan = make(chan os.Signal, 0)
	signal.Notify(deathChan, os.Interrupt, syscall.SIGTERM)
	go func(dg *discordgo.Session, deatchChan chan os.Signal) {
		<-deathChan
		statusData := discordgo.UpdateStatusData{Game: nil}
		if err := dg.UpdateStatusComplex(statusData); err != nil {
			log.Println("Error during deleting status:", err)
			end(nil)
			return
		}
		log.Println("Deleting status... (press Ctrl+C again to permament stop)")
		go func(dchan chan os.Signal) {
			<-dchan
			os.Exit(0)
		}(deathChan)
		time.Sleep(5 * time.Second)
		log.Println("Deleted status!")
		end(nil)
		os.Exit(0)
	}(dg, deathChan)
	var (
		prevTrack string
		// fTitle is f(ull) title
		fTitle string
	)
	for {
		select {
		case <-ticker.C:
			result, err := api.User.GetRecentTracks(lastfm.P{"limit": "1", "user": username})
			if err != nil {
				log.Println("LastFM error: ", err)
				if !endlessMode {
					end(nil)
					return
				}
			} else {
				if len(result.Tracks) > 0 {
					currentTrack := result.Tracks[0]
					isNowPlaying, _ := strconv.ParseBool(currentTrack.NowPlaying)
					trackName := currentTrack.Artist.Name + " - " + currentTrack.Name
					if isNowPlaying {
						fTitle = strings.Replace(title, "{{name}}", currentTrack.Name, -1)
						fTitle = strings.Replace(fTitle, "{{artist}}", currentTrack.Artist.Name, -1)
						fTitle = strings.Replace(fTitle, "{{album}}", currentTrack.Album.Name, -1)
						statusData := discordgo.UpdateStatusData{
							Game: &discordgo.Game{
								Name:    fTitle,
								Type:    discordgo.GameTypeListening,
								Details: currentTrack.Name,
								State:   currentTrack.Artist.Name,
							},
							AFK:    false,
							Status: "online",
						}
						if err := dg.UpdateStatusComplex(statusData); err != nil {
							log.Println("Discord error: ", err)
							if !endlessMode {
								end(nil)
								break
							}
						}
						if prevTrack != trackName {
							log.Println("Now playing: " + trackName)
							prevTrack = trackName
						}
					} else if !isNowPlaying {
						statusData := discordgo.UpdateStatusData{
							Game:   nil,
							Status: "offline",
						}
						if err := dg.UpdateStatusComplex(statusData); err != nil {
							log.Println("Discord error: ", err)
							if !endlessMode {
								end(nil)
								break
							}
						}
					}
				}
			}
		}
	}
}
