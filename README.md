# phpipamsync

Small script synchronizes phpipam hosts with pi-hole resolvers. 

## Use case

Phpipam considered to be a source of truth for DNS resolution. 
* Network admins create hostnames in phpipam. 
* Pihole used as a DNS-server.
* (optional) Cisco router acts as a DHCP-server

To synchronize phpipam hostanmes with pihole DNS-resolver 
```
phpipamsync get pi-hole
```

## phpipamsync config

Location ~/.phpipamsync/config

```
ipam_site_url: https://192.168.169.170:8670/                    # IP and port where phpipam exposes API endpoints
ipam_app_id: phpipamsync                                        # phpipam app id
ipam_app_code: AVyzJl7ceWJ5WrbKZ1QjAIgNdMSggw12                 # phpipam app code
ipam_subnets:                                                   # ipam subnets to synchronize
- 192.168.250.0/24
- 192.168.251.0/24

domain: home                                                    # domain to add to hostnames. For example (home will add .home to all hostnames)
pi_hole: /home/ikuchin/docker/pi.hole/etc-pihole/custom.list    # full path to pi-hole custom.list file 
```

**ATTENTION** pi-hole custom file will be rewritten, not appended.

## phpipam API token

To generate API topken, in phpipam go to Administration -> API -> Create API token. 

* Type - SSL with App code token
* App permission - Read

Important:  HTTPS access should be enabled to phpipam

## Cisco DHCP config

To generate DHCP config snippet for Cisco router configuration 

```
phpipamsync get cisco-dhpc
```

Once snippet generated it can be copy/paste-ed to actual router configuration. There is no integration between phpipamsync and router.

## Examples

### DNS resolution

```
e:\docs\src\go\phpipamsync>nslookup tesla.home
...
Non-authoritative answer:
Name:    tesla.home
Address:  192.168.168.49

e:\docs\src\go\phpipamsync>nslookup 192.168.168.49
...
Name:    tesla.home
Address:  192.168.168.49
```

### Cisco IOS DHCP config

ClientID generated from either of options below:
* MAC address (by adding prefix 01.....)
* Custom field named ClientID ( Administration -> Custom fields -> Custom IP addresses fields )

```
$ phpipamsync get cisco-dhcp

ip dhcp pool fridge
 host 192.168.251.13 255.255.255.0
 client-identifier 0170.2c1f.3ed2.cb
 default-router 192.168.251.1
!
ip dhcp pool tesla
 host 192.168.168.49 255.255.255.0
 client-identifier 014c.fcaa.8ec6.62
 default-router 192.168.251.1
!
```

