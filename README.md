# go-architecture-template

A structured and opinionated architecture template for building Go applications. This template is designed to provide a solid foundation for your Go projects, promoting maintainability, scalability, and best practices.

## Features

- Clean project structure based on domain-driven design principles.
- Separation of concerns with clear boundaries between layers.
- Environment using echo.
- Docker support for containerization.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/IanNarsa/go-architecture-template.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-architecture-template
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Build and run the application:

   ```bash
   go run main.go
   ```

## Project Structure

The project follows a structured architecture with clear separation of concerns:

- **cmd:** Application entry points and executable files.
- **config:** Configuration setup and management.
- **internal:** Core application logic.
  - **handler:** HTTP request handlers.
  - **repository:** Data access and storage.
  - **service:** Business logic and domain services.
  - **util:** Utility functions and helpers.
- **internal:** Shared code that can be reused across projects.
- **test:** Test helpers and resources.
- **scripts:** Useful scripts for development and deployment.

## Contributing

Contributions are welcome! Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## License

This project is public.

---

Feel free to customize this template to provide more specific information about your project.