# modsecurity_filter

`modsecurity_filter` is a Go program designed to filter ModSecurity logs by IP prefix, extracting and displaying IDs, messages, and URIs for quick security analysis.

## Features

- Filter ModSecurity logs by IP prefix.
- Extract and display relevant IDs, messages, and URIs.
- Option to specify log file (default: `confluence_error.log`).

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/modsecurity_filter.git
    cd modsecurity_filter
    ```

2. Build the program:
    ```sh
    ./build.sh
    ```

## Usage

```sh
modsecurity_filter [options]

-logfile: Path to the log file. Default is confluence_error.log.
-ip_prefix: IP prefix to filter logs.
-show_msg : Display the msg field (optional, default: false).
-show_uri : Display the uri field (optional, default: false).
-version: Display the version of the program.
```

## Examples
1 - Filter logs with IP prefix 192.168.:
    ```sh
    modsecurity_filter -ip_prefix 192.168.
    ```

2 - Use a different log file:
    ```sh
    modsecurity_filter -logfile /path/to/your/logfile.log -ip_prefix 192.168.
    ```

## Output

Without -show_msg and -show_uri:

IP: 192.168.242.48, ID: 942151
IP: 192.168.242.48, ID: 932370

With -show_msg=true:
IP: 163.116.242.48, ID: 942151, MSG: SQL Injection Attack: SQL function name detected
IP: 163.116.242.48, ID: 932370, MSG: Remote Command Execution: Windows Command Injection


