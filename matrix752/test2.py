import yaml

def list_second_level_keys(yaml_file):
    try:
        # 讀取 YAML 文件
        with open(yaml_file, 'r', encoding='utf-8') as file:
            data = yaml.safe_load(file)
        
        # 收集所有二級標題（保持順序）
        second_level_keys = []
        for section in data.values():
            if isinstance(section, dict):
                for key in section.keys():
                    if key not in second_level_keys:
                        second_level_keys.append(key)
        
        # 打印結果
        print("二級標題列表：")
        for key in second_level_keys:
            print(f"- {key}")
            
        return second_level_keys
            
    except FileNotFoundError:
        print("錯誤：找不到 YAML 文件")
    except Exception as e:
        print(f"發生錯誤：{str(e)}")
        
if __name__ == "__main__":
    list_second_level_keys('matrix752/datasheet.yaml')
