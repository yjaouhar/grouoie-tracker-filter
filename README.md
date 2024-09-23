# GROUPIE-TRACKER Project

## Overview

The Groupie project is a Go web application that serves artist and event-related data. It fetches information about artists, locations, dates, and relations from a predefined source and renders it using HTML templates. The project includes functionality for displaying artist details, handling errors, and serving static files.

## Features

- **Fetch Data**: Retrieve artist, location, date, and relation data from an external source.
- **Display Artist Information**: Render artist details on a dedicated page.
- **Error Handling**: Display user-friendly error pages for various HTTP errors.
- **Serve Static Files**: Serve static style files for the frontend.

## Project Structure

- `groupie-tracker/`: Contains the main application code.
  - `Fetch.go`: Functions to fetch data from an external source.
  - `handler.go`: HTTP handlers for serving different routes.
  - `template.go`: Functions for rendering HTML templates.
- `template/`: Directory containing HTML template files.
  - `index.html`: Template for the home page.
  - `index2.html`: Template for displaying artist details.
  - `error.html`: Template for error pages.

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://learn.zone01oujda.ma/git/yjaouhar/groupie-tracker
   cd groupie-tracker
## Usage

- **Home Page**:  Access the home page at http://localhost:8080/.
- **Artist Details**:  View details for a specific artist at http://localhost:8080/artist/{id}, where {id} is the artist's numeric ID (e.g., 1 to 52).

## Error Handling

The application provides custom error pages for various HTTP status codes:
- **404 Not Found**:   Displayed when a page or resource cannot be found.
- **500 Internal Server Error**: Displayed for unexpected server errors.

## Contact
- [Yassien JAOUHARY](https://github.com/yjaouhar)
- [Mohamed BAKHCHA](https://github.com/simonbkh)
- [YOUSSEF BASTA](https://github.com/ybasta)
