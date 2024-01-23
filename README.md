# Vbox7 scraper in Go

## Overview
This project is an open-source tool developed in Golang for download the video from vbox7 page
As many know, the page vbox7 is been taken down and they give one month to download the videos from the page
https://www.reddit.com/r/DataHoarder/comments/19d8omx/this_videosharing_website_is_going_to_delete_99/

## Features
- Full search support
- Extracts detailed product information from Airbnb
- Implemented in Go for performance and efficiency
- Easy to integrate with existing Go projects
- The code is optimize to work on this format: ```https://www.vbox7.com/play:[page id] or https://www.vbox7.com/article:[page id]```

## Examples

### Getting the page ID
     https://www.vbox7.com/article:d7e76484cf    id is: d7e76484cf
     https://www.vbox7.com/play:527d6b347b       id is: 527d6b347b       
     https://www.vbox7.com/article:d9000f5f42    id is: d9000f5f42
     https://www.vbox7.com/play:9b5b912a3f       id is: 9b5b912a3f

### Quick testing
```Go
    package main

    import (
        "log"
        "os"
        "github.com/johnbalvin/vbox7"
    )
    func main(){
        pageURL := "https://www.vbox7.com/article:d9000f5f42"
        videoURL, err := vbox7.GetFinalVideoURLFromPageURL(pageURL, nil)
        if err != nil {
            log.Println(err)
            return
        }
        video, contentType, err := vbox7.DowloadVideoFromFinalVideoURL(videoURL, nil)
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("contentType: *%s*\n", contentType)
        if err := os.WriteFile("./video.mp4", video, 0644); err != nil {
            log.Println(err)
            return
        }    
    }
```
```Go
    package main

    import (
        "log"
        "os"
        "github.com/johnbalvin/vbox7"
    )
    func main(){
        pageID:="9b5b912a3f"
        videoURL, err := vbox7.vbox7GetFinalVideoURLFromPageID(pageID, nil)
        if err != nil {
            log.Println(err)
            return
        }
        video, contentType, err := vbox7.DowloadVideoFromFinalVideoURL(videoURL, nil)
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("contentType: *%s*\n", contentType)
        if err := os.WriteFile("./video.mp4", video, 0644); err != nil {
            log.Println(err)
            return
        }    
    }
```
