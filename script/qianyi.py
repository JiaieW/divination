import mysql.connector

# 数据库配置
config = {
    'host': 'dtnnzjewsurk.mysql.sae.sina.com.cn',  # 数据库服务器地址
    'user': 'root',  # 数据库用户名
    'password': 'rootroot',  # 数据库密码
    'database': 'zhouyi',  # 要连接的数据库名称
    'port': 10045  # 端口号，MySQL默认是3306
}

try:
    # 连接到数据库
    conn = mysql.connector.connect(**config)
    print("成功连接到数据库")
    cursor = conn.cursor()

    # 从gua64表获取数据
    cursor.execute("SELECT id, yao_1, yao_2, yao_3, yao_4, yao_5, yao_6 FROM gua64")
    rows = cursor.fetchall()

    # 检查是否有数据返回
    if rows:
        # 更新yao386表
        for row in rows:
            id, yao_1, yao_2, yao_3, yao_4, yao_5, yao_6 = row
            for yao_pos, yao_ci in enumerate([yao_1, yao_2, yao_3, yao_4, yao_5, yao_6], start=1):
                cursor.execute("UPDATE yao386 SET yao_ci = %s WHERE gua_id = %s AND yao_pos = %s", (yao_ci, id, yao_pos))
        conn.commit()
        print("更新完成")
    else:
        print("未找到数据")

except mysql.connector.Error as e:
    print(f"数据库操作失败: {e}")
    conn.rollback()
finally:
    if conn.is_connected():
        cursor.close()
        conn.close()
        print("数据库连接已关闭")