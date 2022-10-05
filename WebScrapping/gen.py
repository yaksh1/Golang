import time
import math
def gen_prime(n):
  start = time.time()
  n-=1
  x = 2
  y = 2
  pl = []
  sl=[2]
  while n!=0:
    i=2
    prime=False
    while True: 
      if x%i==0:
        break
      if i==math.ceil(x/y):
        prime=True
        pl.append(i)
        y=pl[-1]
        break
      i += 1
    if prime:
      sl.append(x)
      # print(x,end=" ")
      n -= 1
    x += 1
  end = time.time()
  print(sl)
  print("time taken in seconds: ", end-start)
gen_prime(100000)

  
  