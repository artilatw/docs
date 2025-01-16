import yaml
import re

def parse_markdown_to_dict(md_content):
    # 初始化字典
    data_dict = {}
    current_section = None
    current_subsection = None
    current_subsubsection = None
    
    # 需要特殊處理的二級標題
    array_sections = ["Features"]
    text_sections = ["Introduction"]
    
    # 按行處理內容
    lines = md_content.split('\n')
    for line in lines:
        # 去除空白
        line = line.strip()
        if not line:
            continue
            
        # 檢測標題
        if line.startswith('# '):
            current_section = line.replace('# ', '').strip()
            data_dict[current_section] = {}
            current_subsection = None
            current_subsubsection = None
        elif line.startswith('## '):
            current_subsection = line.replace('## ', '').strip()
            if current_section:
                if current_subsection in text_sections:
                    data_dict[current_section][current_subsection] = ""
                elif current_subsection in array_sections:
                    data_dict[current_section][current_subsection] = []
                else:
                    data_dict[current_section][current_subsection] = {}
            current_subsubsection = None
        elif line.startswith('### '):
            current_subsubsection = line.replace('### ', '').strip()
            if current_section and current_subsection and current_subsection not in (text_sections + array_sections):
                data_dict[current_section][current_subsection][current_subsubsection] = []
        # 處理內容
        elif line:
            if current_subsection in text_sections and current_section:
                # 對於 Introduction，直接將內容附加到字串中
                current_content = data_dict[current_section][current_subsection]
                if current_content:
                    data_dict[current_section][current_subsection] = current_content + "\n" + line
                else:
                    data_dict[current_section][current_subsection] = line
            elif current_subsection in array_sections and current_section:
                # 對於 Features，將內容作為數組項
                content = line.lstrip('- ').strip()
                if content:
                    data_dict[current_section][current_subsection].append(content)
            elif current_subsubsection and current_subsection and current_section:
                # 如果是列表項，移除開頭的減號和空格
                content = line.lstrip('- ').strip()
                if content:  # 確保不是空字串
                    data_dict[current_section][current_subsection][current_subsubsection].append(content)
            elif ':' in line and not line.startswith('http'):
                # 處理鍵值對
                key, value = line.split(':', 1)
                key = key.strip()
                value = value.strip()
                
                if current_subsection and current_section and current_subsection not in (text_sections + array_sections):
                    data_dict[current_section][current_subsection][key] = value
                elif current_section:
                    data_dict[current_section][key] = value
    
    return data_dict

def main():
    try:
        # 讀取 markdown 文件
        with open('matrix752/datasheet.md', 'r', encoding='utf-8') as file:
            md_content = file.read()
        
        # 解析為字典
        data_dict = parse_markdown_to_dict(md_content)
        
        # 轉換為 YAML 並保存
        with open('matrix752/datasheet.yaml', 'w', encoding='utf-8') as file:
            yaml.dump(data_dict, file, allow_unicode=True, sort_keys=False)
            
        print("成功將 markdown 轉換為 YAML 格式！")
        
    except FileNotFoundError:
        print("錯誤：找不到 datasheet.md 文件")
    except Exception as e:
        print(f"發生錯誤：{str(e)}")

if __name__ == "__main__":
    main()
