# Joblinker

[Live](https://nixjoblinker.vercel.app/)
Note: The seacrh may take time as the backend is deployed on a free instance.
Even after a long time there are no results the search API maximum limit might have been exceeded. 

Joblinker is a comprehensive job matching platform designed to provide relevant job opportunities based on user resumes and preferences. This project leverages Golang for the backend, TypeScript for the frontend, and integrates advanced tools such as Gemini and the Google Search API.

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [System Architecture](#system-architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- **Profile Management:** Users can upload resumes, and set preferences.
- **Job Search and Match Display:** Fetches and displays relevant job listings with headings, links, descriptions, and images.
- **Tag Input for Enhanced Search:** Users can enter tags to improve the accuracy and relevance of job matches.

## Technology Stack

- **Backend:** Golang
- **Frontend:** React, TypeScript, Tailwind CSS
- **APIs:** Gemini for resume parsing, Google Search API for job listings
- **Deployment:** Backend on Render, Frontend on Vercel

## System Architecture

- **Backend:** 
  - Built with Golang
  - API Layer, Data Processing, Job Matching Engine
  - Deployed on Render
- **Frontend:**
  - Developed with TypeScript, Tailwind CSS, React
  - Job Search
  - Deployed on Vercel

## Installation

### Prerequisites

- Golang
- Node.js and npm
- Docker (optional for containerized deployment)

### Backend

1. Clone the repository:

   ```bash
   git clone https://github.com/NIXBLACK11/GeminiChallenge.git
   cd GeminiChallenge
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Start the backend server:

   ```bash
   go run ./cmd/main.go
   ```

### Frontend
[Frontend Code](https://github.com/NIXBLACK11/GeminiChallengeFrontend)

1. Clone the repository:

   ```bash
   git clone https://github.com/NIXBLACK11/GeminiChallengeFrontend.git
   cd GeminiChallengeFrontend
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the frontend server:

   ```bash
   npm run dev
   ```

## Usage

1. Open your browser and navigate to `http://localhost:5173/` to access the frontend.
2. Upload your resume and set your job preferences.
3. Search for jobs and view the listings tailored to your profile.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Open a pull request.