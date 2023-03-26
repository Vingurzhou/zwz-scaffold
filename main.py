#  by zhouwenzhe at 2023/3/26
import shutil

from jinja2 import PackageLoader, Environment, FileSystemLoader

if __name__ == '__main__':
    dest_folder = "/Users/zhouwenzhe/src/zwz-scaffold/gen/"
    service_name = "usercenter"
    host = "http://127.0.0.1"
    port = "80"

    src_folder = "template/src"
    # 使用shutil模块的copytree()函数复制文件夹及其内容
    shutil.copytree(src_folder, dest_folder, dirs_exist_ok=True)

    loader = FileSystemLoader(searchpath=".")
    env = Environment(loader=loader)

    template = env.get_template("template/api/main.tpl")  # 模板文件
    buf = template.render(service_name=service_name)
    with open(dest_folder + "cmd/api/main.go", "w") as file:
        file.write(buf)

    template = env.get_template("template/api/etc.tpl")  # 模板文件
    buf = template.render(service_name=service_name, host=host, port=port)
    with open(dest_folder + "cmd/api/etc/config.yml", "w") as file:
        file.write(buf)
