# GIN ASSISTANT

This project helps to accelerate the creation of repetitive scripts in a Gin project.

- [Usage](#usage)
- [License](#license)

## Usage


1. **Clone the repository:**

    ```bash
    git clone git@github.com:denimars/gin-assistant.git
    ```

2. **Compile Project"**

    ```bash
    cd gin-assistant
    go build
    ```

3. **Copy the build output into the Gin project**

    ```bash
    mv gin-assistant ../<gin-project-name>
    ```


4. **Create project with gin-assistant**

    ```bash
    ./gin-assistant init
    ```
5. **Create service**

    ```
    ./gin-assistant service <nameService> 
    ```
6. **Create Auth**

    ```
    ./gin-assistant auth
    ```
7. **Run Project**
    ```
    ./gin-assistant run
    ```



## License

Distributed under the MIT License. See `LICENSE` for more information.
