
from time import sleep

for s in ["Hi!", "What's up!", "Hello!"]:
	for ch in s:
		print(ch, end="", flush=True)
		sleep(0.2)
	sleep(.5)
	for ch in s:
		print("\b \b", end="", flush=True)
		sleep(0.2)