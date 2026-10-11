import psycopg2
from faker import Faker

fake = Faker()

conn = psycopg2.connect(
    host="localhost", port=5432,
    dbname="postgres", user="postgres", password="postgres"
)
cur = conn.cursor()

cur.execute("""
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100),
        created_at TIMESTAMP
    );
""")

for _ in range(1000):
    cur.execute(
        "INSERT INTO users (name, email, created_at) VALUES (%s, %s, %s)",
        (fake.name(), fake.email(), fake.date_time_this_year())
    )

conn.commit()
cur.close()
conn.close()