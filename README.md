# Spotify_ISRC_Finder

This GoLang-based Music Metadata API allows you to store and retrieve metadata for music tracks in a relational database. It utilizes the Spotify API to fetch additional information about tracks using their ISRC (International Standard Recording Code).

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Dependencies](#dependencies)
- [Database Configuration](#database-confiquration)
- [Spotify Developer Account Configuration](#spotify-developer-account-configuration)
- [API Endpoints](#api-endpoints)
- [Swagger Documentation With Output](#swagger-documentation-with-output)

## Prerequisites

Make sure you have the following installed:

- Go Lang
- PostgreSQL
- Spotify Developer Account

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder.git
  

2. install dependencies:
   ```bash
   go mod init


## Dependencies

The project uses the following Go modules for its dependencies:

- [gin-gonic/gin](https://github.com/gin-gonic/gin) v1.9.1
  - A web framework for Go.

- [lib/pq](https://github.com/lib/pq) v1.10.9  for PostgreSQL driver for Go.

- [swaggo/files](https://github.com/swaggo/files) v1.0.1  for Swaggo static files.

- [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) v1.6.0 for Swagger middleware for Gin.

- [swaggo/swag](https://github.com/swaggo/swag) v1.16.2.  It's A tool to generate Go documentation with Swagger comments.

- [zmb3/spotify](https://github.com/zmb3/spotify) v1.3.0 Go library for the Spotify Web API.

- [oauth2](https://golang.org/x/oauth2) v0.15.0 for OAuth2 client implementation.

- [gorm.io/driver/postgres](https://gorm.io/driver/postgres) v1.5.4 to use PostgreSQL driver for GORM.

- [gorm.io/gorm](https://gorm.io/gorm) v1.25.5 to use ORM library for Golang.

These are the main dependences used for this project. You can see all the depenndencies on go.mod file.

## Database Configuration

### PostgreSQL

1. **Install PostgreSQL:**

   Make sure PostgreSQL is installed on your machine. You can download it from the [official website](https://www.postgresql.org/download/) or use a package manager.

2. **Create a Database:**

   ```bash
   createdb your_database_name

3. **Connection string**
   ```bash
   DATABASE_URL=postgres://your_database_user:your_database_password@localhost:5432/your_database_name?sslmode=disable

replace your_database_user, your_database_password and your_database_name with your actiual PostgreSQL credentials.

## Spotify Developer Account Configuration

To use the Spotify functionality in this project, you need to set up a Spotify Developer account and obtain the required credentials.

### Steps to Obtain Spotify Developer Credentials:

1. **Create a Spotify Developer Account:**
   - Go to the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard/login).
   - Log in with your existing Spotify account or create a new one.

2. **Create a New App:**
   - After logging in, click on the "Create an App" button.
   - Fill in the required information for your new app, and click on "Create."

3. **Get Client ID and Client Secret:**
   - Once the app is created, you'll see the details of your app.
   - Note down the "Client ID" and "Client Secret" values. These will be used for authentication.

Once client id, client secret is obtained replece CLIENT_ID and CLIENT_SECREAT with your actual value in routes.go. 

## API Endpoints

1.**Endpoints:**
The below are the apiendpoins listed in this project.
         
      POST /tracks/create           //For storing tracks to DB for spotify api.
      GET /tracks/{isrc}            //Get the Track with isrc from spotify api.
      GET /tracks/artist/{artist}   //Get the Tracks with artist name form database.
      PUT /tracks/update/{isrc}    //Update the track with isrc code.

2.**Sample Request:**

      
      curl -X POST http://localhost:8080/tracks/your-isrc
      
replace the your-isrc with your actual isrc.

3.**Sample Response:**

- #### Success - 200 OK:
The below is the sample response for tracks stored in database.
    
        {
        "message": "Track created successfully",
        "track": {
       "ISRC": "your-isrc",
       "SpotifyImage": "spotify-image-url",
       "Title": "Track Title",
       "ArtistNames": ["Artist1", "Artist2"],
       "Popularity": 80
        } 

- #### Error - 400 Bad Request:
      
        {
           "error": "Failed to create track / track already exists"
         }
  
## Swagger Documentation With Output
1. **Create docs**
   Run the Swag commit comment to create the swagger docs.

         swag init
2. **Start Application**
   Start the application using the following comment in your root repository.

      go run main.go
   
   you can vist your vist your output in (http://localhost:8080/swagger/index.html) this page.

4. **Sample Output**
Here are the output page for the swagger documentation
<img width="1367" alt="swagger_documentated_output" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/b862229a-4285-4831-be0d-2d74bc5409ac">

Then create a track by passing "ISRC" code in POST method for storing the track details on PostgreSQL.

<img width="855" alt="post_track" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/7b783fa1-b4e5-4f1c-aff8-08b8b1ed2f72">

To get the track dteails by passing "ISRC" in Get track with ISRC method. This fetches the data from the database if not it fetches from the Spotify API.

<img width="1410" alt="get_track_by_isrc" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/1eb27c30-eb02-417d-81c2-716b8bc890ac">


To get the track dteails by passing Artist name in Get track with artist method. This fetches the data from the database if not it returns track not found error.

<img width="1419" alt="get_track_by_artist" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/edd48e64-6386-4ac3-922a-c66927a3899b">

To update the track details by passing ISRC in the PUT method. It updated the updated value in database.

![Update_Track](./assets/outputs/update_track.png)

These are the structs i have used for this project for Swagger documentation.

<img width="1388" alt="models" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/9a13f67c-40a7-41fb-b740-cca53cea7ee8">

The tracks stored in the database are shows below

<img width="1440" alt="database_output" src="https://github.com/KAVINKISHORE13/Spotify_ISRC_Finder/assets/78995411/10d59d77-3e9b-4fdd-b1c6-f26f10fc56f6">

