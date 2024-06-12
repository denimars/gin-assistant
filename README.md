# GIN ASSISTANT

This project helps to accelerate the creation of repetitive scripts in a Gin project.

## Table of Contents

- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Usage


1. **Clone the repository:**

    ```bash
    git clone git@github.com:denimars/gin-assistant.git
    ```

2. **Compile Project"**

    ```bash
    go build
    ```

3. **Copy the build output into the Gin project**

    ```bash
    mv ./gin-assistant ../<gin-project-name>
    ```

4. **Create module **

    ```bash
    go mod init <gin-project-name>
    ```
5. **Create project with gin-assistant**

    ```bash
    ./gin-assistant init
    ```
6. **Create service**

    ```
    ./gin-assistant service <nameService> 
    ```


## License

Distributed under the MIT License. See `LICENSE` for more information.
