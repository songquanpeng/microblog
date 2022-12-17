import argparse
import sqlite3
from datetime import datetime


def main(args):
    connection1 = sqlite3.connect(args.v1_db_path)
    cursor1 = connection1.cursor()
    connection2 = sqlite3.connect(args.v2_db_path)
    cursor2 = connection2.cursor()
    # cursor.executemany("INSERT OR IGNORE INTO words (text) VALUES(?)", words)
    # connection.commit()
    res = cursor1.execute("SELECT id, content, time FROM nonsenses")
    posts = res.fetchall()
    for post_id, content, time_str in posts:
        time = datetime.strptime(time_str[2:], '%y-%m-%d %H:%M:%S')
        timestamp = int(time.timestamp())
        cursor2.execute("INSERT INTO posts (id, content, timestamp, status) VALUES(?, ?, ?, ?)",
                        (post_id, content, timestamp, 1))
    connection2.commit()


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--v1_db_path', type=str, default='../microblog-v0.1.db')
    parser.add_argument('--v2_db_path', type=str, default='../microblog.db')
    main(parser.parse_args())
