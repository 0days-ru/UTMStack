#!/bin/bash
cd /home/utmstack/
rm installer
wget http://github.com/0days-ru/UTMStack/releases/latest/download/installer
chmod +x installer
./installer
