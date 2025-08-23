import uuid, datetime

n=999931
ts=datetime.datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S")
p="gen.accounts.csv"
 
with open(p,"w") as f:
    for _ in range(n):
        f.write("{},255,14,{},{},{}\n".format(uuid.uuid4(), ts, ts, ts))