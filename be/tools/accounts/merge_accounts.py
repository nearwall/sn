#!/usr/bin/env python3

import datetime


def merge_csv_files() ->int:
    # Открываем файлы для чтения с latin-1 кодировкой
    parsed = 0
    ts=datetime.datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S")

    with open('gen.accounts.csv', 'r') as accounts_file, \
         open('people.v2.csv', 'r') as people_file, \
         open('gen.merged_accounts.csv', 'w') as output_file:
        
        # Читаем строки из обоих файлов одновременно
        for accounts_line, people_line in zip(accounts_file, people_file):
            if len(accounts_line) != 0 and len(people_line) != 0:
                parsed += 1
                
                # Берем UUID из первого файла (первое поле до первой запятой)
                uuid_field = accounts_line.split(',')[0]
                
                # Обрабатываем первое поле из people.v2.csv: заменяем пробел на запятую
                name_field = people_line.split(',')[0].replace(' ', ',')
                
                # Получаем остальные поля из people.v2.csv
                people_parts = people_line.strip().split(',')
                remaining_fields = people_parts[1:] if len(people_parts) > 1 else []
                
                # Объединяем: UUID + обработанное имя + остальные поля из people.v2.csv
                merged_line = f"{uuid_field},{name_field}"
                
                if remaining_fields:
                    merged_line += "," + ",".join(remaining_fields[:-1]) + ",,{}".format(remaining_fields[:-1][0]) + f",{ts},{ts}"

                # Записываем объединенную строку
                output_file.write(merged_line + '\n')



    return parsed

if __name__ == "__main__":
    parsed = merge_csv_files()
    print(f"Файлы успешно объединены в merged_accounts.csv c {parsed} строками")
