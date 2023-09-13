package main 

import ( 
	"fmt"
	"flag"
	"os"
)  

func main() { 
	//'videos get' subcommand 
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for 'videos get' command 
	getAll := getCmd.Bool("all" , false , "Get all videos")
	getID := getCmd.String("id", "", "Youtube video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError) 
	addId := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video Title") 
	addUrl := addCmd.String("url", "", "YouTube video URL") 
	addImageUrl := addCmd.String("imageurl", "", "Youtube video Image") 
	addDesc := addCmd.String("desc", "", "Youtube video description") 

	if len(os.Args)< 2 { 
		fmt.Println("expected 'get' or 'add' subcommands") 
		os.Exit(1)
	} 

	//handle get, add and wrong input by the user 
	switch os.Args[1]{ 
	case "get": //if its the get command 
	//handle get here  
		HandleGet(getCmd, getAll ,getID)
	case "add": 
		HandleAdd(addCmd , addID , addTitle , addUrl , addImageUrl , addDesc)
	default: //if we don't understand the input 
	}
}   

func HandleGet(getCmd *flag.FlagSet , all *bool , id *string){ 
	getCmd.Parse(os.Args[2]) 

	if *all == false && *id == "" { 
		fmt.Print("id is require or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	} 

	if *all { 
		//return all videos
		videos := getVideos() 
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n") 
		
		for _, video := range videos { 
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n ", video.Id , video.Title ,video.Url , video.ImageUrl , video.Description) 
		} 
		return 
	} 

	if *id != "" { 
		videos := getVideos() 
		id := *id 
		for _, video := range videos { 
			fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id , video.Title , video.Url , video.ImageUrl , video.Description)
		}
	}
}  

func VallidateVideo(addCmd *flag.FlagSet , id *string , title *string , url *string , imageurl *string, description *string) { 
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description = "" { 
		fmt.Print("all fields are required for adding a video") 
		addCmd.PrintDefaults() //built in in the flag package
		os.Exit(1) 
	}
}

func HandleAdd(addCmd *flag.FlagSet , id *string , title *string , url *string , imageurl *string, description *string) { 
	VallidateVideo(addCmd , id , title, url , imageUrl ,description) 

	video := video { 
		Id : *id , 
		Title : *title, 
		Description: *description, 
		Imageurl : *imageUrl, 
		Url: *url, 
	}  

	videos := getVideos() 
	videos = append(videos,video) 

	saveVideos(videos) 
} 