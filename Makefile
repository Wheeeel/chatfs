all:
	- cd cmd; go build -o ../chatfs

chatfs: 
	- cd cmd; go build -o ../chatfs

clean:
	- rm chatfs

run: chatfs
	- sudo umount /tmp/chatfs
	- ./chatfs
