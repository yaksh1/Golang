d = {'a':1,'b':2,'c':3,'d':4}
kv = {'a':'A','b':'B','c':'C','d':'D'}
for ki in d.keys():
  d[kv[ki]]=d.pop(ki)
print(d)