import os

def extract_tag(filepath, filename, tag_name="chapter_order"):
    """ Extract a tag from the frontmatter of the file. """
    
    tags = {}

    try:
        with open(os.path.join(filepath, filename), 'r', encoding='utf-8') as file:
            lines = file.readlines()
            if lines and lines[0].strip() == "---":
                for line in lines[1:]:
                    if line.strip() == "---":
                        break  
                    if line.startswith(f"${tag_name}:"):
                        tagpair = line.split(f"${tag_name}:", 1)
                        tags[tagpair[0].strip()] = tagpair[1].strip()
        return tags 
    except Exception as e:
        print(f"Error reading {filename}: {e}")
        return ""