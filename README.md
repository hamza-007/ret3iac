# ret3iac

**ret3iac** is a lightweight Infrastructure as Code (IaC) CLI tool designed to manage cloud resources using YAML configuration files. Inspired by Terraform and Pulumi, it simplifies defining and provisioning infrastructure across multiple cloud providers.

## Features
- **YAML Configuration** for defining infrastructure resources.
- **Cross-Cloud Support** with easy configuration.
- **CLI Interface** for deploying, updating, and deleting resources.
- **Human-Readable Configuration** for ease of use and management.

## Getting Started

### Prerequisites
- **Go** version 1.18 or higher.
- Supported cloud provider accounts (e.g., AWS, Azure).

### Installation
1. Clone the repository:

    ```bash
    git clone https://github.com/hamza-007/ret3iac.git
    cd ret3iac
    ```

2. Build the executable:

    ```bash
    go build -o ret3iac main.go
    ```

3. Verify the installation:

    ```bash
    ./ret3iac --help
    ```

### Usage
Create a configuration file (`config.yml`) with the following structure:

```yaml
resource:
  - type: "aws_instance"
    name: "example_instance"
    config:
      ami: "ami-12345678"
      instance_type: "t2.micro"
      region: "us-east-1"
```

Apply configuration with cli command apply:

```bash
./ret3iac apply config.yml
```

### Directory Structure

This project follows a specific directory structure to maintain organization:

```bash
ret3iac/
├── cmd/               # CLI commands
├── pkg/               # Core library and configuration logic
└── main.go            # Entry point for the CLI application
```

### Configuration

The tool utilizes a configuration file config.yml for defining resources and their settings.

### Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and create a pull request.

### License

This project is licensed under the MIT License.