# This script will enable 127.0.0.1 up to 127.0.0.256 interfaces
for ((i=2;i<256;i++)) do 
  sudo ifconfig lo0 alias 127.0.0.$i up
done 
