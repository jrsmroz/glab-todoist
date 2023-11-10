# GitLab to Todoist Integration utility

This Go application automates the process of fetching issues assigned to a specific user from GitLab and creating corresponding tasks in Todoist if they do not already exist.

## Features

- Fetches all issues assigned to a specified GitLab user across all projects.
- Checks if these issues already have corresponding tasks in Todoist.
- Creates Todoist tasks for GitLab issues that are not already tracked in Todoist.

## Prerequisites

To use this application, you will need:

- GitLab personal access token with `read_api` scope.
- Todoist API token.
- Go installed on your system (version 1.19 or higher recommended).

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/jrsmroz/glab-todoist.git
   ```
2. Navigate to the cloned directory:
   ```
   cd glab-todoist
   ```

## Configuration

Set the following environment variables:

- `GITLAB_ACCESS_TOKEN`: Your GitLab personal access token.
- `GITLAB_USER_ID`: Your GitLab user ID.
- `TODOIST_ACCESS_TOKEN`: Your Todoist API token.

## Usage

Run the application with:

```bash
go run main.go
```

The program will fetch issues from GitLab and create tasks in Todoist based on the fetched issues.

## License

This project is licensed under the [MIT License](LICENSE).

---

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your changes.

## Support

If you encounter any issues or have any questions about this tool, please submit an issue in the GitHub repository.

---
