<center style="margin-bottom: 20px;">
<img src='./public/logo-solid.svg' />
</center>

SMS Trap is a fake SMS-sending API that provides a dashboard to view all sent SMS messages. It's designed for testing and development purposes, offering code samples for implementation in various languages and framework-specific guides.

## Features

- Fake SMS sending API
- Browser notifications for each SMS
- SMS character count
- SMS encoding detection
- Dashboard to view sent messages
- Code samples for multiple languages
- Integration guides for Laravel and Django

## Usage

### Installation

```sh
sudo curl https://raw.githubusercontent.com/OmarFaruk-0x01/sms-trap/master/setup.sh | sh
```

### CLI Options

SMS Trap accepts the following command-line flags:

- `-db-path`: Path to the SQLite database file (default: /tmp - delete after closing the server)
- `-port`: Port to run the server on (default: 1290)

Example:

```sh
# It will not persist your sms
sms-trap
```

```sh
# To Persist sms, mention the path.
sms-trap -db-path=./sms-trap.db
```

### API Endpoint

Base URL: `http://localhost:1290/api/v1`

#### Send SMS

`GET` `/trap`

Query Parameters:

- `phones[]`: List of phone numbers
- `message`: SMS message content
- `label`: Type of SMS (transactional/promotional)

Example:

```sh
curl http://localhost:1290/api/v1/trap?phones[]=1234567890&phones[]=0987654321&message=Hello%20World&label=transactional
```

### Development Configuration

### Installation

To install SMS Trap, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/sms-trap.git
   cd sms-trap
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   npm install
   ```

### Running the server

To start the SMS Trap server:

```sh
make run
```

This will build the project and start the server on the default port (1290).

### Build Server

To build the project:

```sh
make build
```

This will generate necessary files, build CSS, and compile the Go binary.

### Hot Reloading

For development with hot-reloading:

```sh
make hot
```

The project uses `air` for live-reloading during development. The configuration is in the `.air.toml` file in the project root.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Roadmap

1. SMS Rule Violation Detection
   - Implement rules to detect violations in SMS content.
   - Notify users of any detected violations.
2. SMS Cost Estimation
   - Estimate the cost of sending each SMS based on the message content and SMS providers
   - Provide a detailed cost breakdown in the dashboard for selected providers.
3. SMS Log Capture
   - Capture detailed production logs for each sent SMS.

## Acknowledgements

- This project uses [air](https://github.com/cosmtrek/air) for live-reloading during development.
- CSS is built using [Tailwind CSS](https://tailwindcss.com/).
