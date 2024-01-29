# JellyRenamer

This is a small script I wrote to edit batches of video files for a jellyfin server. 
I had problems where Jellyfin would not understand where one seasons stops and the other starts or when a season starts with episode 25 and so on.

The script uses the parent folder of the video files and the folder above the parent folder to name rename all of the video files in the order they are already in. 
It will rename it to a structure that jellyfin prefers.
The best results can be achieved by having the following structure to your media folder.

`SeriesName/Season X/Episode X.mp4`

The episode name does not matter, only the original order the episodes are listed as.

The script can be run by using the following command:


`go run main.go "SeriesName/Season 1"`

or if the series does not have any season (it will assume only one season):

`go run main.go "SeriesName"`

Using the relative path works just fine.

## Make sure not so run this anywhere else! It could brick your PC

This script will blindly rename the contents of a folder. It can and will break your PC if used where it shouldn't be. I had built in a failsafe if it doesnt detect a parent folder named "Season" or "Series" but I have removed it for the time being in favor of working on automation instead.

## Contribution
All types of contribution is welcome. I am still very new to Go and would be happy to learn from others that want to improve the tool.
