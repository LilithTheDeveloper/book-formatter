package types

type Config struct {
	InputDir          string   `yaml:"input_dir"`
	OutputDir         string   `yaml:"output_dir"`
	OutputFormat      string   `yaml:"output_format"`
	GenerateSingleFile bool     `yaml:"generate_single_file"`
	GenerateStatistics bool     `yaml:"generate_statistics"`
	IgnoredFiles      []string `yaml:"ignored_files"`
	IgnoredDirs       []string `yaml:"ignored_dirs"`
	ConfigFilePath    string   `yaml:"config_file_path"`
	LuaFilePath	  string   `yaml:"lua_file_path"`
	PreambleFilePath string   `yaml:"preamble_file_path"`
}

type MarkdownFile struct {
	FilePath string
	Frontmatter map[string]string
	ChapterOrder int
}

type MarkdownFileStatistics struct {
	Lines int
	Words int
	Characters int
}