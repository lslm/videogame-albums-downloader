package main

import (
  "fmt"
  "log"
  "net/http"
  "errors"
  "os"
  "io"
  "sync"
  "github.com/SouUmLucas/videogame-albums-downloader/model"
  "github.com/PuerkitoBio/goquery"
)

var (
  err       error
  tracks    []model.Track
  outputDir string
  albumUrl  string
  orq       sync.WaitGroup
)


func main() {
  err = setArgs()
  if err != nil {
    log.Fatal(err)
  }

  err = fetchURL(albumUrl)
  if err != nil {
    log.Fatal(err)
  }

  orq.Add(len(tracks))

  for _, track := range tracks {
    go downloadSoundtrack(track)
  }

  fmt.Printf("Please wait until the program is finished\n\n")
  orq.Wait()
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

      output, err := os.Create(outputDir + track.SongName + ".mp3")
      if err != nil {
        log.Fatal(err)
        orq.Done()
        return
      }

      defer output.Close()

      fmt.Printf("Dowloading soundtrack #%s: %s (%s)\n", track.Track, track.SongName, track.Size)
      response, err := http.Get(downloadUrl)
      if err != nil {
        log.Fatal(err)
        orq.Done()
        return
      }

      defer response.Body.Close()

      _, err = io.Copy(output, response.Body)

      if err != nil {
        log.Fatal(err)
        orq.Done()
        return
      }
    })
  }

  orq.Done()
  
  return
}

func setArgs() (err error) {
  if (len(os.Args) < 3) {
    err = errors.New("Invalid arguments: you must specify the URL of the album and the output directory")
    return
  } else {
    albumUrl = os.Args[1]
    outputDir = os.Args[2]
  }

  return
}