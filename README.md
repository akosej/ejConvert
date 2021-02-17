# ejConver
Converter developed in golang, using ffmpeg libraries. 

## Install Linux
```console
wget https://github.com/akosej/ejConvert/raw/main/ejConvert
sudo cp ./ejConver /usr/local/bin/
cudo chmod +x /usr/local/bin/ejConver
```

## Usage
Example of use. If you want to convert the avi files to mp4 that are in the directory.
ejConvert converts all the files that are inside the directory, with the specified extension, in this case it will convert all the .avi to .mp4 

```console
cd /home/user/videos/
ejConvert avi mp4
```

