# JellyRenamer

This is a small script I wrote to edit batches of video files for a jellyfin server. 
I had problems where Jellyfin would not understand where one seasons stops and the other starts or when a season starts with episode 25 and so on.
This script needs the following folder structure:

SeriesName
- Season
--Episode1
--Episode2


go run main.go "SeriesName/Season 1" or go run main.go "SeriesName"
