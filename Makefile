all:
	- cd cmd; go build -o ../chatfs

chatfs: 
	- cd cmd; go build -o ../chatfs

clean:
	- rm chatfs

run: chatfs
	- mkdir -p /tmp/chatfs
	- sudo umount /tmp/chatfs
	- ./chatfs
