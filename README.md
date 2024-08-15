# Go Fiber Project Template

This project serves as a base template for building Go applications using the Fiber web framework. It includes a clean and scalable structure with various essential features to kickstart your development process.
https://github.com/Chintuka-Code/go-fiber-setup/settings

## Features

- **Fiber**: High-performance Go web framework.
- **Global Error Handling**: Centralized error management for better debugging and error reporting.
- **Validation**: Input validation using the `validator` package.
- **Global Validation Handling**: Consistent validation error responses across the application.
- **Custom Validation**: Extend validation logic with custom rules.
- **Clean Code**: Structured project layout promoting maintainability and scalability.
- **Git Hooks Configuration**: Pre-configured Git hooks to enforce coding standards and run tests.
- **Task Automation**: Use `Taskfile.yml` to automate common tasks like building, testing, and linting.
- **Internationalization (i18n) Support**: Easily initialize localization for multiple languages.
- **Swagger Support**: Auto-generate API documentation with Swagger.
- **Routes Grouping**: Organized route management for different modules or features.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. If not, you can download it from [golang.org](https://golang.org/dl/).

### Project Setup

1. **Clone the Repository**

   ```bash
   https://github.com/Chintuka-Code/go-fiber-setup.git
   cd go-fiber-setup
   ```

2. ### Make the `setup.sh` Script Executable

   Before running the setup script, ensure it has the correct permissions:

   ```bash
   chmod +x setup.sh
   ```

3. ### Run the Setup Script

   The `setup.sh` script installs the necessary tools and runs the project setup:

   ```bash
   ./setup.sh
   ```

   **The setup script will:**

   - Install the `Task` command runner globally if not already installed.
   - Run the task setup:app command to initialize the project.
