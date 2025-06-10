# book-formatter
A simple formatter written in Golang that helps converting markdown files to a PDF using pandoc and pdflatex. This was made during the writing process of my book. 


## Features
- It removes frontmatters completely
- Allows for ordered PDF generation through frontmatter tag `chapter_order`
- Custom DIVs with the help of Lua Script

## Usage
1. Install the required dependencies:
   - `pandoc`
   - `pdflatex`
   - `golang`

2. Clone the repository:
   ```bash
    git clone git@github.com:LilithTheDeveloper/book-formatter.git
    ```

3. Navigate to the project directory:
    ```bash
    cd book-formatter
    ```

4. Build the project:
    ```bash
    go build -o book-formatter
    ```

5. (Optional) Install the project globally:
    ```bash
    go install
    ```

6. (Optional) Create a config file: 
    ```bash
    cp config.yaml.example config.yaml
    ```

7. Run the formatter:
    ```bash
    ./book-formatter format --input <path_to_markdown_file> --output <path_to_output_pdf>
    ```    

## Configuration
You can customize the behavior of the formatter by editing the `config.yaml` file. The example configuration file is provided as `config.yaml.example`.

### Example config.yaml
```yaml
# Configuration file for document processing
input_dir: ""

# Input file extensions to process
output_dir: "" 

# Input file extensions to process
output_format: "pdf"

# Output file name
generate_single_file: true

# Generate a single file for each input file
generate_statistics: true 

# The PDF engine to use for LaTeX processing
pandoc_pdf_engine: "xelatex"

# Ignored files
ignored_files: [""]

# Ignored directories
ignored_dirs: [""]

# Config file path 
config_file_path: "./config.yaml"
# Cache
cache_location: "./cache"
# Lua file path
lua_file_path: "./lua/custom_divs.lua"
# LaTeX preamble file path
preamble_file_path: "./latex/preamble.tex"
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or have suggestions for improvements.