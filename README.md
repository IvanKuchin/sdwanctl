# sdwanctl
Cisco SDWAN control CLI utility


## vManage URL and credentials

Config file must be located in `~/.sdwan/config`
Example:
```
vmanages:
- name: vmanage1
  server: https://vmanage-lab.viptela.net:8443
- name: vmanage2
  server: https://vmanage.home.lab:8443

users:
- name: user1
  login: login1
  password: password1
- name: user2
  login: login2
  password: secret2

contexts:
- name: home-lab
  vmanage: vmanag1
  user: user1
- name: prod
  vmanage: vmanag2
  user: user2

current-context: home-lab
```

Currently supported API-calls:
- `sdwanctl get devices` - list all devices in a vManage
- `sdwanctl get config {deviceId}` - collect config from a device
- `sdwanctl describe device {deviceId}` - detailed info about a device
- `sdwanctl describe system device {deviceId}` - detailed info about a device
- `sdwanctl describe template input {deviceId}` - collect template inputs for all feature templates
