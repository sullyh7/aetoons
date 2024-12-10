
# API Documentation

This documentation provides details about the REST API endpoints for the **AeToons** application. This application allows users to manage shows and episodes while integrating with Vimeo and the MyAnimeList (MAL) API for extended functionality.

---

## Base URL

```
http://<your-server-url>/
```

---

## Endpoints

### **1. Test Endpoint**
- **URL**: `/`
- **Method**: `GET`
- **Description**: Simple test endpoint to verify the API is running.
- **Response**:
  ```json
  {
    "test": "success"
  }
  ```

---

### **2. Fetch All Shows**
- **URL**: `/shows`
- **Method**: `GET`
- **Description**: Retrieves all shows stored in the database.
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    [
      {
        "id": 1,
        "title": "One Piece",
        "main_picture": {
          "medium": "https://cdn.myanimelist.net/images/anime/1244/138851.jpg",
          "large": "https://cdn.myanimelist.net/images/anime/1244/138851l.jpg"
        },
        "episodes": [
          {
            "id": 1,
            "title": "Episode 1",
            "episode_number": 1,
            "video_url": "https://vimeo.com/example"
          }
        ]
      }
    ]
    ```

---

### **3. Add a Show**
- **URL**: `/add-show`
- **Method**: `POST`
- **Description**: Adds a new show to the database using data from the MAL API.
- **Request Body** (JSON):
  ```json
  {
    "mal_id": 21
  }
  ```
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "id": 21,
      "title": "One Piece",
      "main_picture": {
        "medium": "https://cdn.myanimelist.net/images/anime/1244/138851.jpg",
        "large": "https://cdn.myanimelist.net/images/anime/1244/138851l.jpg"
      }
    }
    ```

---

### **4. Add an Episode (File Upload)**
- **URL**: `/add-episode`
- **Method**: `POST`
- **Description**: Adds a new episode by uploading a video file. The video is transcribed and uploaded to Vimeo.
- **Form Data**:
  - `title` (string, required): The title of the episode.
  - `episode_number` (integer, required): The episode number.
  - `show_id` (integer, required): The ID of the show.
  - `file` (file, required): The video file to upload.
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "id": 1,
      "title": "Episode 1",
      "episode_number": 1,
      "video_url": "https://vimeo.com/example",
      "show_id": 1
    }
    ```

---

### **5. Add an Episode (From URL)**
- **URL**: `/add-episode-from-url`
- **Method**: `POST`
- **Description**: Adds a new episode by providing a video URL. The video is downloaded, transcribed, and uploaded to Vimeo.
- **Request Body** (JSON):
  ```json
  {
    "title": "Episode 1",
    "episode_number": 1,
    "show_id": 1,
    "video_url": "https://example.com/video.mp4"
  }
  ```
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "id": 1,
      "title": "Episode 1",
      "episode_number": 1,
      "video_url": "https://vimeo.com/example",
      "show_id": 1
    }
    ```

---

## Models

### **Show**
```json
{
  "id": 1,
  "title": "One Piece",
  "main_picture": {
    "medium": "https://cdn.myanimelist.net/images/anime/1244/138851.jpg",
    "large": "https://cdn.myanimelist.net/images/anime/1244/138851l.jpg"
  },
  "episodes": [
    {
      "id": 1,
      "title": "Episode 1",
      "episode_number": 1,
      "video_url": "https://vimeo.com/example"
    }
  ]
}
```

### **Episode**
```json
{
  "id": 1,
  "title": "Episode 1",
  "episode_number": 1,
  "video_url": "https://vimeo.com/example",
  "show_id": 1
}
```

---

## Error Responses

- **400 Bad Request**:
  - Missing or invalid fields in the request.
  ```json
  {
    "error": "Missing required fields"
  }
  ```

- **500 Internal Server Error**:
  - Server-side errors, such as issues with file uploads, database operations, or external API calls.
  ```json
  {
    "error": "Failed to save file"
  }
  ```

---

## Notes
- Ensure the MAL Client ID is correctly set in the configuration (`MAL_AUTH_PARAM_NAME`) for adding shows.
- The upload directory for videos is `./uploads/`. Ensure this directory is writable by the server.
- Subtitles are automatically generated during the transcription process.

For additional questions or issues, refer to the AeToons documentation or contact the API support team.
