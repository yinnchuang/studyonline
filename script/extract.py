import logging
import time

import pandas as pd
import requests

url = 'http://127.0.0.1:8080/unit/create'

df = pd.read_excel('units.xlsx',header=0, skiprows=1)
df = df.iloc[1:, [0, 3, 5]]

father_units = df[df.isnull().any(axis=1)]
son_units = df[~df.isnull().any(axis=1)]

father_units_list = []
for father_unit in father_units['*名称']:
    father_units_list.append(father_unit)

for index, son_unit in son_units.iterrows():
    unit_name = son_unit['*名称']
    unit_desc = son_unit['属性.1']
    father_name = son_unit['*上级知识名称']
    try:
        idx = father_units_list.index(father_name)+1
        print(unit_name, idx)
        headers = {
            "Content-Type": "application/json",
            "Authorization": "bd6fa9c0-7558-44fa-8791-620b5a5c50ed"
        }
        payload = {
            "unit_name": unit_name,
            "unit_desc": unit_desc,
            "father_unit_id": idx
        }

        try:
            resp = requests.Session().post(url, json=payload, headers=headers)

            if resp.ok:
                logging.debug("成功: %s", payload)
            else:
                logging.warning("失败: %s - %s", resp.status_code, resp.text)
        except Exception as exc:
            logging.warning("异常: %s - %s", payload, exc)

        time.sleep(0.2)
    except Exception:
        continue

