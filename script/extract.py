import logging
import time

import pandas as pd
import requests

username = '2006400038'
password = '2006400038'
loginUrl = f'http://127.0.0.1:8080/login/teacher?username={username}&password={password}'
resp = requests.Session().post(loginUrl)
token = resp.json()['token']


creatUrl = 'http://127.0.0.1:8080/unit/create'

df = pd.read_excel('units.xlsx',header=0)
for index, row in df.iterrows():
    id = row['id']
    unit_name = row['unit_name']
    father_id = row['father_id']
    desc = row['unit_name']
    print(f'{id}\t{unit_name}\t{father_id}\t{desc}')
    headers = {
        "Content-Type": "application/json",
        "Authorization": token
    }
    payload = {
        "id": id,
        "unit_name": unit_name,
        "unit_desc": desc,
        "father_unit_id": father_id
    }

    try:
        resp = requests.Session().post(creatUrl, json=payload, headers=headers)

        if resp.ok:
            logging.debug("成功: %s", payload)
        else:
            logging.warning("失败: %s - %s", resp.status_code, resp.text)
    except Exception as exc:
        logging.warning("异常: %s - %s", payload, exc)

    time.sleep(0.2)
