package main

import (
	"fmt"
  "log"
	"net/http"
	"os"
	"io"
  "github.com/SouUmLucas/videogame-albums-downloader/model"
	"github.com/PuerkitoBio/goquery"
)

var (
  tracks []model.Track
)


func main() {
	err := fetchURL("https://downloads.khinsider.com/game-soundtracks/album/metal-gear-solid-3-snake-eater-the-complete-soundtrack-flac-gamerip")
  if err != nil {
    log.Fatal(err)
  }

  for _, track := range tracks {
    downloadSoundtrack(track)
  }
}

func fetchURL(url string) (err error) {
  tracks = make([]model.Track, 0)

  doc, err := goquery.NewDocument(url)
  if err != nil {
    return
  }

  doc.Find("#songlist").Each(func(i int, s *goquery.Selection) {
    s.Find("tr").Each(func(j int, q *goquery.Selection) {

      track := model.Track{}

      q.Find("td").Each(func (k int, p *goquery.Selection) {
        col := p.Text()

        switch k {
        case 1:
          track.Track = col
        case 2:
          track.SongName = col

          p.Find("a").Each(func (b int, a *goquery.Selection) {
            track.Url = a.Nodes[0].Attr[0].Val
          })

        case 3:
          track.Length = col
        case 4:
          track.Size = col
        }
      })

      tracks = append(tracks, track)
    })
  })

  return
}

func downloadSoundtrack(track model.Track) (err error) {
  if track.Url != "" {
    url := "https://downloads.khinsider.com" + track.Url
    doc, _ := goquery.NewDocument(url)

    doc.Find("audio").Each(func (i int, s *goquery.Selection) {
      downloadUrl := s.Nodes[0].Attr[2].Val

      output, err := os.Create("downloadedSoundtracks/" + track.SongName + ".mp3")
      if err != nil {
        log.Fatal(err)
        return
      }

      defer output.Close()

      fmt.Printf("Dowloading soundtrack #%s: %s (%s)\n", track.Track, track.SongName, track.Size)
      response, err := http.Get(downloadUrl)
      if err != nil {
        log.Fatal(err)
        return
      }

      defer response.Body.Close()

      _, err = io.Copy(output, response.Body)

      if err != nil {
        log.Fatal(err)
        return
      }
    })
  }
  
  return
}
