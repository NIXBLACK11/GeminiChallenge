## To get a Google Search API key and Search Engine ID, you need to follow these steps:

### Step 1: Get a Google API Key

1. Create a Project in Google Cloud Console:
- Go to the [Google Cloud Console](https://console.cloud.google.com/).
- Click on the project drop-down and select or create a new project.

2. Enable the Custom Search API:
- In the Google Cloud Console, go to the [APIs & Services Dashboard](https://console.cloud.google.com/apis).
- Click on "+ ENABLE APIS AND SERVICES".
- Search for "Custom Search API" and enable it.

3. Create API Credentials:
- Go to the [Credentials page](https://console.cloud.google.com/apis/credentials).
- Click on "Create Credentials" and select "API Key".
- Copy the generated API key and keep it safe. This will be your API key for accessing the Google Custom Search API.

### Step 2: Get a Search Engine ID
1. Create a Custom Search Engine:
- Go to the [Google Custom Search Engine](https://developers.google.com/custom-search/v1/introduction).
- Click on "New search engine".
- Fill in the required fields:
- In the "Sites to search" field, you can add any website you want to search, or use "example.com" for a basic setup.
- Click "Create".

2. Retrieve the Search Engine ID:
- After creating your custom search engine, go to the [Custom Search Engine control panel](https://programmablesearchengine.google.com/controlpanel/create).
- Click on the search engine you just created.
- Under "Details", you will find the "Search engine ID". Copy this ID as you will need it to perform searches using the API.
